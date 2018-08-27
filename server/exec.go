/**
 *  author: lim
 *  data  : 18-4-9 下午12:45
 */

package server

import (
	"fmt"
	"strconv"

	"sync"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/errors"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/xsql/node"
	"github.com/lemonwx/xsql/server/version"
	"github.com/lemonwx/xsql/sqlparser"
)

var String = sqlparser.String

func (conn *MidConn) handleDelete(stmt *sqlparser.Delete, sql string) ([]*mysql.Result, error) {
	var err error

	if conn.nodeIdx, err = conn.getShardList(stmt); err != nil {
		log.Errorf("[%d] get shard list failed:%v", conn.ConnectionId, err)
		return nil, conn.NewMySQLErr(ERR_UNSUPPORTED_SHARD)
	}

	if len(conn.nodeIdx) == 0 {
		return nil, nil
	}

	table := sqlparser.String(stmt.Table)
	where := sqlparser.String(stmt.Where)
	beContinue, err := conn.chkAndLockRows(table, where)
	if err != nil {
		return nil, err
	}

	if !beContinue {
		return nil, nil
	}

	if err := conn.getNextVersion(); err != nil {
		return nil, err
	}

	updateSql := fmt.Sprintf("update %s set version = %d %s", table, conn.NextVersion, where)
	log.Debugf("[%d] update version sql: %s", conn.ConnectionId, updateSql)

	if _, err := conn.ExecuteOnNodePool([]byte(updateSql), conn.nodeIdx); err != nil {
		return nil, err
	}

	newSql := sqlparser.String(stmt)
	return conn.ExecuteOnNodePool([]byte(newSql), conn.nodeIdx)
}

func (conn *MidConn) handleInsert(stmt *sqlparser.Insert, sql string) ([]*mysql.Result, error) {
	var err error
	var shardList []int

	if shardList, err = conn.getShardList(stmt); err != nil {
		log.Errorf("[%d] get shard list failed:%v", conn.ConnectionId, err)
		return nil, conn.NewMySQLErr(ERR_UNSUPPORTED_SHARD)
	}

	if len(shardList) != 1 {
		log.Errorf("[%d] insert stmt must route to 1 node, but recv: %d", conn.ConnectionId, len(conn.nodeIdx))
		return nil, fmt.Errorf("insert must route to 1 node")
	}

	if err := conn.getNextVersion(); err != nil {
		return nil, err
	}

	// add extra col for every rows
	vals := make(sqlparser.Values, len(stmt.Rows.(sqlparser.Values)))
	for idx, row := range stmt.Rows.(sqlparser.Values) {
		t := row.(sqlparser.ValTuple)
		t[0] = sqlparser.NumVal(strconv.FormatUint(conn.NextVersion, 10))
		vals[idx] = t
	}
	stmt.Rows = vals

	newSql := sqlparser.String(stmt)
	log.Debugf("[%d]: after convert sql: %s", conn.ConnectionId, newSql)

	// exec
	nodeIdx := shardList[0]
	back, err := conn.getSingleBackConn(nodeIdx)
	if err != nil {
		return nil, err
	}

	ret, err := back.Execute(mysql.COM_QUERY, []byte(newSql))
	if err != nil {
		return nil, err
	}

	return []*mysql.Result{ret}, nil
}

func (conn *MidConn) execSingleUpdate(stmt *sqlparser.Update, idx int) ([]*mysql.Result, error) {
	back, err := conn.getSingleBackConn(idx)
	if err != nil {
		return nil, err
	}

	var vInUse map[uint64]uint8
	var vErr error
	var chkRet *mysql.Result
	var chkErr error
	var cvtSql string

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		vInUse, vErr = conn.getCurVInUse(UPDATE_OR_DELETE)
		stmt.Exprs[0].Expr = sqlparser.NumVal(strconv.FormatUint(conn.NextVersion, 10))
		cvtSql = sqlparser.String(stmt)
		wg.Done()
	}()

	go func() {
		selSql := fmt.Sprintf("select version from %s %s for update", String(stmt.Table), String(stmt.Where))
		chkRet, chkErr = back.Execute(mysql.COM_QUERY, []byte(selSql))
		wg.Done()
	}()
	wg.Wait()

	if vErr != nil {
		return nil, vErr
	}
	if chkErr != nil {
		return nil, chkErr
	}

	if len(chkRet.RowDatas) == 0 {
		return nil, nil
	}

	if err := conn.chkInUse(&[]*mysql.Result{chkRet}, 1, vInUse); err != nil {
		return nil, err
	}

	ret, err := back.Execute(mysql.COM_QUERY, []byte(cvtSql))
	if err != nil {
		return nil, err
	}
	return []*mysql.Result{ret}, nil
}

func (conn *MidConn) execMultiUpdate(stmt *sqlparser.Update) ([]*mysql.Result, error) {
	if err := conn.getMultiBackConn(conn.nodeIdx); err != nil {
		return nil, err
	}

	var newSql string
	shardSize := len(conn.nodeIdx)
	ch := make(chan map[uint64]uint8, shardSize)

	go func() {
		// assume the row want to update not used by other, so get next version as the same time
		vInUse, err := conn.getCurVInUse(UPDATE_OR_DELETE)
		if err != nil {
			// notify all the executor get failed
			close(ch)
		} else {
			// send to executor, setup cvt sql
			stmt.Exprs[0].Expr = sqlparser.NumVal(strconv.FormatUint(conn.NextVersion, 10))
			newSql = sqlparser.String(stmt)
			for _ = range conn.execNodes {
				ch <- vInUse
			}
			log.Debugf("[%d] cvt sql and send to executor done.", conn.ConnectionId)
		}
	}()

	selSql := fmt.Sprintf("select version from %s %s for update", String(stmt.Table), String(stmt.Where))
	ms := NewMS(shardSize)

	for _, back := range conn.execNodes {
		go func(back *node.Node) {
			defer ms.Done()
			// first exec chk in use sql and lock the rows
			chkRet, err := back.Execute(mysql.COM_QUERY, []byte(selSql))
			if err != nil {
				ms.appendErr(err)
				return
			}

			// wait for get active versions, if ch closed, may get failed, response to client
			vInUse, ok := <-ch
			if !ok {
				ms.appendErr(errors.New2("get version in use failed"))
				return
			}

			// if result empty, direct response to client, affect rows: 0
			if len(chkRet.RowDatas) == 0 {
				ms.appendRet(&mysql.Result{AffectedRows: 0})
				return
			}

			// if chk in use failed, response to client
			if err := conn.chkInUse(&[]*mysql.Result{chkRet}, 1, vInUse); err != nil {
				ms.appendErr(err)
				return
			}

			// final exec cvt sql
			if execRet, err := back.Execute(mysql.COM_QUERY, []byte(newSql)); err != nil {
				ms.appendErr(err)
				return
			} else {
				ms.appendRet(execRet)
			}
		}(back)
	}
	ms.Wait()

	switch {
	case len(ms.rets) == shardSize:
		return ms.rets, nil
	case len(ms.errs) == shardSize:
		return nil, ms.errs[0]
	default:
		return nil, errors.New2("unexpected multi node response not equal")
	}
}

func (conn *MidConn) handleUpdate(stmt *sqlparser.Update, sql string) ([]*mysql.Result, error) {
	var err error

	if conn.nodeIdx, err = conn.getShardList(stmt); err != nil {
		log.Errorf("[%d] get shard list failed:%v", conn.ConnectionId, err)
		return nil, conn.NewMySQLErr(ERR_UNSUPPORTED_SHARD)
	}

	switch len(conn.nodeIdx) {
	case 0:
		return nil, nil
	case 1:
		return conn.execSingleUpdate(stmt, conn.nodeIdx[0])
	default:
		return conn.execMultiUpdate(stmt)
	}
}

func (conn *MidConn) chkAndLockRows(table, where string) (bool, error) {
	selSql := fmt.Sprintf("select version from %s %s for update", table, where)
	rets, err := conn.executeSelect(selSql, 1, UPDATE_OR_DELETE)
	if err != nil {
		return false, err
	}

	for _, ret := range rets {
		if len(ret.RowDatas) != 0 {
			return true, nil
		}
	}
	return false, nil
}

func (conn *MidConn) getNextVersion() error {
	// get next version
	var err error

	if conn.NextVersion == 0 {
		conn.NextVersion, err = version.NextVersion()
		if err != nil {
			log.Debugf("[%d] conn next version is nil, but get failed %v", conn.ConnectionId, conn.NextVersion)
			return err
		}
		log.Debugf("[%d] conn next version is nil, get one: %v", conn.ConnectionId, conn.NextVersion)
	} else {
		log.Debugf("[%d] use next version get from pre sql in this trx: %v", conn.ConnectionId, conn.NextVersion)
	}
	return nil
}

func (conn *MidConn) getCurVInUse(flag uint8) (map[uint64]uint8, error) {
	var err error
	var ret map[uint64]uint8

	if flag == UPDATE_OR_DELETE && conn.NextVersion == 0 {
		log.Debugf("[%d] chk v in use for update, get next version at the same time", conn.ConnectionId)
		base, err := version.InUseAndNext()
		if err != nil {
			return nil, err
		}
		conn.NextVersion = base.Next
		ret = conn.VersionsInUse
	} else {
		ret, err = version.VersionsInUse()
		if err != nil {
			return nil, err
		}
	}

	if _, ok := ret[conn.NextVersion]; ok {
		delete(ret, conn.NextVersion)
	}

	return ret, nil
}

func (conn *MidConn) getVInUse() error {
	// get v in use by other session
	var err error

	if conn.VersionsInUse == nil {
		conn.VersionsInUse, err = version.VersionsInUse()
		if err != nil {
			log.Debugf("[%d] conn's vInuse in use is nil, but get v in user failed %v", conn.ConnectionId, err)
			return err
		}

		if _, ok := conn.VersionsInUse[conn.NextVersion]; ok {
			delete(conn.VersionsInUse, conn.NextVersion)
			log.Debugf("[%d] delete pre sql's next version %s in the same trx", conn.ConnectionId, conn.NextVersion)
		}
		log.Debugf("[%d] get vInuse: %v", conn.ConnectionId, conn.VersionsInUse)
	} else {
		log.Debugf("[%d] use vInuse get from pre sel sql in this trx: %v", conn.ConnectionId, conn.NextVersion)
	}
	return nil
}

func (conn *MidConn) getNodeIdxs(stmt sqlparser.Statement, bindVars map[string]interface{}) error {
	var err error
	if conn.db == "" {
		err = errors.New(mysql.NewDefaultError(mysql.ER_NO_DB_ERROR))
		return err
	}

	conn.nodeIdx, err = conn.getShardList(stmt)

	if err != nil {
		log.Debugf("[%d] get node idxs failed: %v", conn.ConnectionId, err)
		return err
	}
	log.Debugf("[%d] get node idxs: %v", conn.ConnectionId, conn.nodeIdx)

	return nil
}
