/**
 *  author: lim
 *  data  : 18-4-9 下午12:45
 */

package midconn

import (
	"github.com/lemonwx/xsql/sqlparser"
	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/mysql"
)

func (conn *MidConn) handleInsert(stmt *sqlparser.Insert, sql string) error {

	// add extra col


	ex := &sqlparser.UpdateExpr{
		Name: &sqlparser.ColName{Name: []byte(extraColName)},
		Expr: sqlparser.NumVal("123456"),
	}



	extraCol := &sqlparser.NonStarExpr{
		Expr: ex.Name,
	}
	stmt.Columns = append(stmt.Columns, extraCol)
	vals := make(sqlparser.Values, len(stmt.Rows.(sqlparser.Values)))

	for idx, row := range stmt.Rows.(sqlparser.Values) {
		t := row.(sqlparser.ValTuple)
		t = append(t, ex.Expr)
		vals[idx] = t
	}
	stmt.Rows = vals

	newSql := sqlparser.String(stmt)
	log.Debugf("[%d]: after convert sql: %s", conn.ConnectionId, newSql)

	rs, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(newSql), nil)
	if err != nil {
		return err
	} else {
		return conn.HandleExecRets(rs)
	}
}

