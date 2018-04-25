/**
 *  author: lim
 *  data  : 18-4-9 下午12:45
 */

package midconn

import (
	"strconv"
	"fmt"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/xsql/sqlparser"
	"github.com/lemonwx/xsql/middleware/version"
	"github.com/lemonwx/xsql/middleware/meta"
)

func (conn *MidConn) handleDelete(stmt *sqlparser.Delete, sql string) error {
	var err error
	if err = conn.getNodeIdxs(stmt); err != nil {
		return err
	} else if conn.nodeIdx == nil {
		return conn.cli.WriteOK(nil)
	}

	var tb string = sqlparser.String(stmt.Table)
	var where string = sqlparser.String(stmt.Where)

	if err = conn.handleSelectForUpdate(tb, where); err != nil {
		return err
	}

	if err = conn.getNextVersion(); err != nil {
		return err
	}

	updateSql := fmt.Sprintf("update %s set version = %s %s", tb, conn.NextVersion, where)
	if _, err = conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(updateSql), conn.nodeIdx); err != nil {
		if err != nil {
			log.Errorf("[%d] execute in multi node failed: %v", conn.ConnectionId, err)
			return err
		}
		log.Debugf("[%d] exec update in multi node finish", conn.ConnectionId)
	}

	log.Debugf("[%d] after convert sql: %s", conn.ConnectionId, sql)
	rs, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(sql), conn.nodeIdx)
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
	} else if conn.nodeIdx == nil {
		return conn.cli.WriteOK(nil)
	}

	// get next version
	if err = conn.getNextVersion(); err != nil {
		return err
	}

	// add extra col
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
	rs, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(newSql), conn.nodeIdx)
	if err != nil {
		return err
	} else {
		return conn.HandleExecRets(rs)
	}
}

func (conn *MidConn) handleUpdate(stmt *sqlparser.Update, sql string) error {

	var err error

	if err = conn.getNodeIdxs(stmt); err != nil {
		return err
	} else if conn.nodeIdx == nil {
		return conn.cli.WriteOK(nil)
	}

	if err = conn.handleSelectForUpdate(
		sqlparser.String(stmt.Table), sqlparser.String(stmt.Where)); err != nil {
		return err
	}

	if err = conn.getNextVersion(); err != nil {
		return err
	}

	// add extra col, get new sql
	stmt.Exprs[0].Expr = sqlparser.NumVal(strconv.FormatUint(conn.NextVersion, 10))
	newSql := sqlparser.String(stmt)
	log.Debugf("[%d] sql convert to: %s", conn.ConnectionId, newSql)
	log.Debugf("generallog--[%d] 3:%s", conn.ConnectionId, newSql)

	if rs, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(newSql), conn.nodeIdx); err != nil {
		return err
	} else {
		return conn.HandleExecRets(rs)
	}
}

func (conn *MidConn) handleSelectForUpdate(table, where string) error {
	var err error

	selSql := fmt.Sprintf("select version from %s %s for update", table, where)

	if err = conn.getVInUse(); err != nil {
		return err
	} else if conn.nodeIdx == nil {
		return conn.cli.WriteOK(nil)
	}

	conn.setupNodeStatus(conn.VersionsInUse, true)
	defer conn.setupNodeStatus(nil, false)

	_, err = conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(selSql), conn.nodeIdx)
	if err != nil {
		log.Debugf("[%d] row data in use by another session, update failed: %v", conn.ConnectionId, err)
		return err
	}
	log.Debugf("[%d] select for update success", conn.ConnectionId)
	return nil
}

func (conn *MidConn) needGetNextV(nodeIdxs []int) bool {
	// judge if this sql need to get next version or not
	need := true

	if conn.status[0] == mysql.SERVER_STATUS_IN_TRANS &&
		conn.status[1] == mysql.SERVER_STATUS_AUTOCOMMIT &&
		len(nodeIdxs) == 1 {
		need = false
	}
	return need
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

func (conn *MidConn) getNodeIdxs(stmt sqlparser.Statement) error {
	var err error
	if conn.db == "" {
		return mysql.NewDefaultError(mysql.ER_NO_DB_ERROR)
	}
	conn.nodeIdx, err = sqlparser.GetStmtShardListIndex(stmt, meta.GetRouter(conn.db), nil)
	//conn.nodeIdx, err = router.GetNodeIdxs(stmt)

	if err != nil {
		log.Debugf("[%d] get node idxs failed: %v", conn.ConnectionId, err)
		return err
	}
	log.Debugf("[%d] get node idxs: %v", conn.ConnectionId, conn.nodeIdx)
	return nil
}


func (conn *MidConn) makeBindVars(args []interface{}) map[string]interface{} {
	bindVars := make(map[string]interface{}, len(args))

	for i, v := range args {
		bindVars[fmt.Sprintf("v%d", i+1)] = v
	}

	return bindVars
}