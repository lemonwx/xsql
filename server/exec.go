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
	"github.com/lemonwx/xsql/sqlparser"
)

var String = sqlparser.String

func (conn *MidConn) execSingle(stmt *sqlparser.Update, idx int, delSql string) ([]*mysql.Result, error) {
	back, err := conn.getSingleBackConn(idx)
	if err != nil {
		return nil, err
	}

	var vInUse map[uint64]bool
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
		chkRet, chkErr = conn.execute(back, mysql.COM_QUERY, []byte(selSql))
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

	if delSql != "" {
		// for delete
		if _, err := conn.execute(back, mysql.COM_QUERY, []byte(cvtSql)); err != nil {
			return nil, err
		}

		if ret, err := conn.execute(back, mysql.COM_QUERY, []byte(delSql)); err != nil {
			return nil, err
		} else {
			return []*mysql.Result{ret}, nil
		}
	} else {
		// for update
		if ret, err := conn.execute(back, mysql.COM_QUERY, []byte(cvtSql)); err != nil {
			return nil, err
		} else {
			return []*mysql.Result{ret}, nil
		}
	}
}

func (conn *MidConn) execMulti(stmt *sqlparser.Update, delsql string) ([]*mysql.Result, error) {
	if err := conn.getMultiBackConn(conn.nodeIdx); err != nil {
		return nil, err
	}

	var updateSql string
	shardSize := len(conn.nodeIdx)
	ch := make(chan map[uint64]bool, shardSize)

	go func() {
		// assume the row want to update not used by other, so get next version as the same time
		vInUse, err := conn.getCurVInUse(UPDATE_OR_DELETE)
		if err != nil { // notify all the executor get failed
			close(ch)
		} else { // send to executor, setup cvt sql
			stmt.Exprs[0].Expr = sqlparser.NumVal(strconv.FormatUint(conn.NextVersion, 10))
			updateSql = sqlparser.String(stmt)
			for _ = range conn.execNodes {
				ch <- vInUse
			}
			close(ch)
		}
	}()

	selSql := fmt.Sprintf("select version from %s %s for update", String(stmt.Table), String(stmt.Where))
	ms := NewMS(shardSize)

	for _, back := range conn.execNodes {
		go func(back *node.Node) {
			defer ms.Done()
			// first exec chk in use sql and lock the rows
			chkRet, err := conn.execute(back, mysql.COM_QUERY, []byte(selSql))
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

			// exec is from handle delete
			if delsql != "" {
				// for delete, update rows with cur version, for tx_iso
				if _, err := conn.execute(back, mysql.COM_QUERY, []byte(updateSql)); err != nil {
					ms.appendErr(err)
					return
				}
				if execRet, err := conn.execute(back, mysql.COM_QUERY, []byte(delsql)); err != nil {
					ms.appendErr(err)
					return
				} else {
					ms.appendRet(execRet)
				}
			} else {
				if execRet, err := conn.execute(back, mysql.COM_QUERY, []byte(updateSql)); err != nil {
					ms.appendErr(err)
					return
				} else {
					ms.appendRet(execRet)
				}
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

	shardSize := len(conn.nodeIdx)
	if shardSize == 0 {
		return nil, nil
	} else if shardSize == 1 {
		return conn.execSingle(stmt, conn.nodeIdx[0], "")
	} else {
		return conn.execMulti(stmt, "")
	}
}

func (conn *MidConn) handleDelete(stmt *sqlparser.Delete, sql string) ([]*mysql.Result, error) {
	var err error
	if conn.nodeIdx, err = conn.getShardList(stmt); err != nil {
		log.Errorf("[%d] get shard list failed:%v", conn.ConnectionId, err)
		return nil, conn.NewMySQLErr(ERR_UNSUPPORTED_SHARD)
	}

	shardSize := len(conn.nodeIdx)
	if shardSize == 0 {
		return nil, nil
	}

	update := &sqlparser.Update{
		Table: stmt.Table,
		Exprs: sqlparser.UpdateExprs{
			&sqlparser.UpdateExpr{
				Name: &sqlparser.ColName{Name: []byte("version")},
				Expr: sqlparser.NumVal{}}},
		Where:   stmt.Where,
		OrderBy: stmt.OrderBy,
		Limit:   stmt.Limit,
	}

	if shardSize == 1 {
		return conn.execSingle(update, conn.nodeIdx[0], sql)
	} else {
		return conn.execMulti(update, sql)
	}
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

	ret, err := conn.execute(back, mysql.COM_QUERY, []byte(newSql))
	if err != nil {
		return nil, err
	}

	return []*mysql.Result{ret}, nil
}
