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
	"github.com/lemonwx/xsql/middleware/meta"
	"github.com/lemonwx/xsql/middleware/version"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/xsql/sqlparser"
)

func (conn *MidConn) handleUpdateForDelete(stmt *sqlparser.Delete) error {

	updateSql := fmt.Sprintf("update %s set version = %d %s", sqlparser.String(stmt.Table), conn.NextVersion, sqlparser.String(stmt.Where))
	updateStmt, err := sqlparser.Parse(updateSql)

	if err != nil {
		return err
	}

	_, err = conn.handleUpdate(updateStmt.(*sqlparser.Update), updateSql)
	if err != nil {
		return err
	}

	return nil
}
func (conn *MidConn) handleDelete(stmt *sqlparser.Delete, sql string) ([]*mysql.Result, error) {
	if err := conn.getNodeIdxs(stmt, nil); err != nil {
		return nil, err
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

func (conn *MidConn) handleDelete1(stmt *sqlparser.Delete, sql string) ([]*mysql.Result, error) {
	var err error
	if err = conn.getNodeIdxs(stmt, nil); err != nil {
		return nil, err
	} else if conn.nodeIdx == nil {
		return nil, UNEXPECT_MIDDLE_WARE_ERR
	}

	/*
		var tb string = sqlparser.String(stmt.Table)
		var where string = sqlparser.String(stmt.Where)

		if err = conn.handleSelectForUpdate(tb, where); err != nil {
			return nil, err
		}

		if err = conn.getNextVersion(); err != nil {
			return nil, err
		}

		updateSql := fmt.Sprintf("update %s set version = %d %s", tb, conn.NextVersion, where)
		if _, err = conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(updateSql), conn.nodeIdx); err != nil {
			if err != nil {
				log.Errorf("[%d] execute in multi node failed: %v", conn.ConnectionId, err)
				return nil, err
			}
			log.Debugf("[%d] exec update in multi node finish", conn.ConnectionId)
		}
	*/
	if err = conn.handleUpdateForDelete(stmt); err != nil {
		return nil, err
	}

	log.Debugf("[%d] after convert sql: %s", conn.ConnectionId, sql)
	return conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(sql), conn.nodeIdx)
}

func (conn *MidConn) handleInsert(stmt *sqlparser.Insert, sql string) ([]*mysql.Result, error) {
	var err error

	// router
	if err = conn.getNodeIdxs(stmt, nil); err != nil {
		return nil, err
	} else if conn.nodeIdx == nil {
		return nil, UNEXPECT_MIDDLE_WARE_ERR
	}

	if len(conn.nodeIdx) != 1 {
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
	nodeIdx := conn.nodeIdx[0]
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
	if err := conn.getNodeIdxs(stmt, nil); err != nil {
		return nil, err
	}

	if err := conn.getNextVersion(); err != nil {
		return nil, err
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

	stmt.Exprs[0].Expr = sqlparser.NumVal(strconv.FormatUint(conn.NextVersion, 10))
	newSql := sqlparser.String(stmt)
	log.Debugf("[%d] sql after convert: %s", conn.ConnectionId, newSql)

	return conn.ExecuteOnNodePool([]byte(newSql), conn.nodeIdx)
}

func (conn *MidConn) chkAndLockRows(table, where string) (bool, error) {
	selSql := fmt.Sprintf("select version from %s %s for update", table, where)
	rets, err := conn.executeSelect(selSql, 1)
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

/*

func (conn *MidConn) change2caseWhen(stmt *sqlparser.Update) error {

	maxV := []byte("10000000000000")
	for idx, expr := range stmt.Exprs {

		// case when's cond
		var cond sqlparser.BoolExpr

		rCond := &sqlparser.ComparisonExpr{
			Operator: "<",
			Left: &sqlparser.ColName{
				Name: []byte(extraColName),
			},
			Right: sqlparser.NumVal(maxV),
		}

		if len(conn.VersionsInUse) != 0 {
			vs := make([]sqlparser.ValExpr, 0, len(conn.VersionsInUse))
			for v, _ := range conn.VersionsInUse {
				vs = append(vs, sqlparser.NumVal([]byte(fmt.Sprintf("%d", v))))
			}
			cond = &sqlparser.AndExpr{
				Left: &sqlparser.ComparisonExpr{
					Operator: "not in",
					Left: &sqlparser.ColName{
						Name: []byte(extraColName),
					},
					Right: sqlparser.ValTuple(vs),
				},
				Right: rCond,
			}
		} else {
			cond = rCond
		}

		stmt.Exprs[idx].Expr = &sqlparser.CaseExpr{
			Expr: nil,
			Whens: []*sqlparser.When{
				&sqlparser.When{
					Cond: cond,
					Val:  expr.Expr,
				},
			},
			Else: expr.Name,
		}
	}

	return nil
}

func (conn *MidConn) handleSelectForUpdate1(table, where string) error {
	var err error

	selSql := fmt.Sprintf("select version from %s %s for update", table, where)

	if err = conn.getVInUse(); err != nil {
		return err
	} else if conn.nodeIdx == nil {
		return conn.cli.WriteOK(nil)
	}

	conn.setupNodeStatus(conn.VersionsInUse, true, false, 1)
	defer conn.setupNodeStatus(nil, false, false, 0)

	_, err = conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(selSql), conn.nodeIdx)
	if err != nil {
		log.Debugf("[%d] row data in use by another session, update failed: %v", conn.ConnectionId, err)
		return err
	}
	log.Debugf("[%d] select for update success", conn.ConnectionId)
	return nil
}
*/
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

func (conn *MidConn) getCurVInUse() (map[uint64]uint8, error) {
	vs, err := version.VersionsInUse()
	if err != nil {
		return nil, err
	}

	if _, ok := vs[conn.NextVersion]; ok {
		delete(vs, conn.NextVersion)
	}

	return vs, nil
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

	r, err := meta.GetRouter(conn.db)
	if err != nil {
		log.Errorf("[%d] get router failed: %v", err)
		return err
	}

	conn.nodeIdx, err = sqlparser.GetStmtShardListIndex(stmt, r, bindVars)

	if err != nil {
		log.Debugf("[%d] get node idxs failed: %v", conn.ConnectionId, err)
		return err
	}
	log.Debugf("[%d] get node idxs: %v", conn.ConnectionId, conn.nodeIdx)

	return nil
}
