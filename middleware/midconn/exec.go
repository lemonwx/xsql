/**
 *  author: lim
 *  data  : 18-4-9 下午12:45
 */

package midconn

import (
	"github.com/lemonwx/xsql/sqlparser"
	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/xsql/middleware/version"
	"fmt"
)

func (conn *MidConn) handleInsert(stmt *sqlparser.Insert, sql string) error {

	nextV, err := version.NextVersion()
	if err != nil {
		return err
	}

	log.Debugf("[%d] get next version : %v", conn.ConnectionId, nextV)
	// add extra col
	extraCol := &sqlparser.NonStarExpr{
		Expr: &sqlparser.ColName{Name: []byte(extraColName)},
	}
	stmt.Columns = append(stmt.Columns, extraCol)
	vals := make(sqlparser.Values, len(stmt.Rows.(sqlparser.Values)))

	for idx, row := range stmt.Rows.(sqlparser.Values) {
		t := row.(sqlparser.ValTuple)
		t = append(t, sqlparser.NumVal(nextV))
		vals[idx] = t
	}
	stmt.Rows = vals

	newSql := sqlparser.String(stmt)
	log.Debugf("[%d]: after convert sql: %s", conn.ConnectionId, newSql)

	rs, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(newSql), nil)
	if err != nil {
		return err
	} else {
		err = conn.HandleExecRets(rs)
		if err != nil {
			return err
		}
		return version.ReleaseVersion(nextV)
	}
}

func (conn *MidConn) handleUpdate(stmt *sqlparser.Update, sql string) error {

	if err := conn.handleSelectForUpdate(stmt, nil); err != nil {
		log.Debugf("[%d] row data in use by another session, update failed",
			conn.ConnectionId)
		return err
	}

	nextVersion, err := version.NextVersion()
	if err != nil {
		log.Errorf("[%d] get nextversion failed: %v", conn.ConnectionId, err)
		return err
	}
	log.Debugf("[%d] get nextversion is: %d", conn.ConnectionId, nextVersion)

	expr := &sqlparser.UpdateExpr{
		Name: &sqlparser.ColName{
			Name: []byte(extraColName),
		},
		Expr:sqlparser.NumVal(nextVersion),
	}
	stmt.Exprs = append(stmt.Exprs, expr)
	newSql := sqlparser.String(stmt)
	log.Debugf("[%d] sql convert to: %s", conn.ConnectionId, newSql)

	if rs, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(newSql), nil);err != nil {
			return err
	} else {
		err = conn.HandleExecRets(rs)
		if err != nil {
			return err
		} else {
			return version.ReleaseVersion(nextVersion)
		}
	}
}

func (conn *MidConn) handleSelectForUpdate(uStmt *sqlparser.Update, nodeIdx []int) error {
	selSql := fmt.Sprintf("select version from %s %s for update",
		sqlparser.String(uStmt.Table), sqlparser.String(uStmt.Where))
	log.Debugf("[%d] select for update sql: %s",
		conn.ConnectionId, selSql)

	vInuse, err := version.VersionsInUse()
	if err != nil {
		return err
	}
	log.Debugf("[%d] get vInuse: %v", conn.ConnectionId, vInuse)

	conn.setupNodeStatus(vInuse, true)
	defer conn.setupNodeStatus(vInuse, false)

	_, err = conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(selSql), nodeIdx)
	if err != nil {
		return err
	}
	log.Debugf("[%d] select for update success", conn.ConnectionId)
	return nil
}
