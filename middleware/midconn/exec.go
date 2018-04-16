/**
 *  author: lim
 *  data  : 18-4-9 下午12:45
 */

package midconn

import (
	"fmt"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/middleware/version"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/xsql/sqlparser"
	"github.com/lemonwx/xsql/middleware/router"
)

func (conn *MidConn) handleDelete(stmt *sqlparser.Delete, sql string) error {
	var err error

	tb := sqlparser.String(stmt.Table)
	where := sqlparser.String(stmt.Where)
	err = conn.handleSelectForUpdate(tb, where, nil)
	if err != nil {
		log.Warnf("[%d] select for update failed: %v", conn.ConnectionId, err)
		return err
	}

	if conn.NextVersion == nil {
		conn.NextVersion, err = version.NextVersion()
		if err != nil {
			log.Errorf("[%d] request for next version failed: %v", conn.ConnectionId, err)
			return err
		}
	}
	log.Debugf("[%d] get next version: %v", conn.ConnectionId, conn.NextVersion)

	updateSql := fmt.Sprintf("update %s set version = %s %s", tb, conn.NextVersion, where)
	log.Debugf("[%d] exec update sql : %s", conn.ConnectionId, updateSql)

	_, err = conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(updateSql), nil)
	if err != nil {
		log.Errorf("[%d] execute in multi node failed: %v", conn.ConnectionId, err)
		return err
	}
	log.Debugf("[%d] exec update in multi node finish", conn.ConnectionId)
	log.Debugf("[%d] after convert sql: %s", conn.ConnectionId, sql)

	rs, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(sql), nil)
	if err != nil {
		return err
	} else {
		return conn.HandleExecRets(rs)
	}
}

func (conn *MidConn) handleInsert(stmt *sqlparser.Insert, sql string) error {
	var err error

	// router
	if err = conn.getNodeIdxs(stmt); err != nil {
		return err
	}

	// get next version
	if err = conn.getNextVersion(); err != nil {
		return err
	}


	// add extra col
	vals := make(sqlparser.Values, len(stmt.Rows.(sqlparser.Values)))
	for idx, row := range stmt.Rows.(sqlparser.Values) {
		t := row.(sqlparser.ValTuple)
		t[0] = sqlparser.NumVal(conn.NextVersion)
		vals[idx] = t
	}
	stmt.Rows = vals

	newSql := sqlparser.String(stmt)
	log.Debugf("[%d]: after convert sql: %s", conn.ConnectionId, newSql)

	// exec
	rs, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(newSql), conn.nodeIdx)
	if err != nil {
		return err
	} else {
		return conn.HandleExecRets(rs)
	}
}

func (conn *MidConn) handleUpdate(stmt *sqlparser.Update, sql string) error {
	var err error

	if err = conn.handleSelectForUpdate(
		sqlparser.String(stmt.Table), sqlparser.String(stmt.Where), nil); err != nil {
		log.Debugf("[%d] row data in use by another session, update failed",
			conn.ConnectionId)
		return err
	}

	if conn.NextVersion == nil {
		conn.NextVersion, err = version.NextVersion()
		if err != nil {
			log.Errorf("[%d] get nextversion failed: %v", conn.ConnectionId, err)
			return err
		}
	}
	log.Debugf("[%d] get nextversion is: %d", conn.ConnectionId, conn.NextVersion)

	expr := &sqlparser.UpdateExpr{
		Name: &sqlparser.ColName{
			Name: []byte(extraColName),
		},
		Expr: sqlparser.NumVal(conn.NextVersion),
	}
	stmt.Exprs = append(stmt.Exprs, expr)
	newSql := sqlparser.String(stmt)
	log.Debugf("[%d] sql convert to: %s", conn.ConnectionId, newSql)

	if rs, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(newSql), nil); err != nil {
		return err
	} else {
		return conn.HandleExecRets(rs)
	}
}

func (conn *MidConn) handleSelectForUpdate(table, where string, nodeIdx []int) error {
	selSql := fmt.Sprintf("select version from %s %s for update", table, where)
	log.Debugf("[%d] select for update sql: %s",
		conn.ConnectionId, selSql)

	var err error
	if conn.VersionsInUse == nil {
		conn.VersionsInUse, err = version.VersionsInUse()
		if err != nil {
			log.Debugf("[%d] get v in user failed %v", conn.ConnectionId, err)
			return err
		}
	}
	log.Debugf("[%d] get vInuse: %v", conn.ConnectionId, conn.VersionsInUse)

	conn.setupNodeStatus(conn.VersionsInUse, true)
	defer conn.setupNodeStatus(nil, false)

	_, err = conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(selSql), nodeIdx)
	if err != nil {
		return err
	}
	log.Debugf("[%d] select for update success", conn.ConnectionId)
	return nil
}

func (conn *MidConn)needGetNextV(nodeIdxs []int) bool {
	// judge if this sql need to get next version or not
	need := true

	if conn.status[0] == mysql.SERVER_STATUS_IN_TRANS &&
		conn.status[1] == mysql.SERVER_STATUS_AUTOCOMMIT &&
			len(nodeIdxs) == 1 {
		need = false
	}
	return  need
}

func (conn *MidConn) getNextVersion() error {
	// get next version
	var err error
	if conn.NextVersion == nil {
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

func (conn *MidConn) getNodeIdxs(stmt sqlparser.Statement) error {
	var err error
	conn.nodeIdx, err = router.GetNodeIdxs(stmt)
	if err != nil {
		return err
	}
	log.Debugf("[%d] get node idxs: %v", conn.ConnectionId,  conn.nodeIdx)
	return nil
}