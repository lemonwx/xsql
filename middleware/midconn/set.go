/**
 *  author: lim
 *  data  : 18-4-11 下午11:06
 */

package midconn

import (
	"strings"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/meta"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/xsql/sqlparser"
)

func (conn *MidConn) handleSet(stmt *sqlparser.Set, sql string) error {
	log.Debugf("[%d] handle set", conn.ConnectionId, stmt.Exprs)
	if len(stmt.Exprs) != 2 {
		return UNEXPECT_MIDDLE_WARE_ERR
	}

	if !strings.Contains(strings.ToLower(sql), "autocommit") {

		rets, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(sql), meta.GetFullNodeIdxs())
		if err != nil {
			return err
		}
		return conn.HandleExecRets(rets)
	} else {
		if sqlparser.String(stmt.Exprs[1].Expr) == "0" {
			conn.defaultStatus = ^mysql.SERVER_STATUS_AUTOCOMMIT
			conn.status[1] = conn.defaultStatus
			conn.handleBegin(true)
		}
	}

	/*
		expr := stmt.Exprs[0]

		if v, ok := expr.Expr.(sqlparser.NumVal); ok {
			log.Debugf("[%d], set num %v", conn.ConnectionId, v)
		}

		if v, ok := expr.Expr.(sqlparser.StrVal); ok {
			log.Debugf("[%d], set str d g%v", conn.ConnectionId, v)
		}
	*/

	/*
		if on :
			default = on
		if off
			default = off
	*/

	return conn.cli.WriteOK(nil)

}
