/**
 *  author: lim
 *  data  : 18-4-9 下午12:45
 */

package midconn

import (
	"fmt"
	"strconv"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/errors"
	"github.com/lemonwx/xsql/middleware/version"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/xsql/sqlparser"
)

func (conn *MidConn) handleDelete(stmt *sqlparser.Delete, sql string) ([]*mysql.Result, error) {
	var err error

	if conn.nodeIdx, err = conn.getShardList(stmt); err != nil {
		return nil, err
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
		return nil, err
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

func (conn *MidConn) handleUpdate(stmt *sqlparser.Update, sql string) ([]*mysql.Result, error) {
	var err error

	if conn.nodeIdx, err = conn.getShardList(stmt); err != nil {
		return nil, err
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

	stmt.Exprs[0].Expr = sqlparser.NumVal(strconv.FormatUint(conn.NextVersion, 10))
	newSql := sqlparser.String(stmt)
	log.Debugf("[%d] sql after convert: %s", conn.ConnectionId, newSql)

	return conn.ExecuteOnNodePool([]byte(newSql), conn.nodeIdx)
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
