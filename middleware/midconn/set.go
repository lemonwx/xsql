/**
 *  author: lim
 *  data  : 18-4-11 下午11:06
 */

package midconn

import (
	"github.com/lemonwx/xsql/sqlparser"
	"github.com/lemonwx/log"
)

func (conn *MidConn)handleSet(stmt *sqlparser.Set, sql string) error {
	log.Debugf("[%d] handle set", conn.ConnectionId)
	if len(stmt.Exprs) != 1 {
		return UNEXPECT_MIDDLE_WARE_ERR
	}

	expr := stmt.Exprs[0]

	if v, ok := expr.Expr.(sqlparser.NumVal); ok {
		log.Debugf("[%d], set num %v", conn.ConnectionId, v)
	}

	if v, ok := expr.Expr.(sqlparser.StrVal); ok {
		log.Debugf("[%d], set str d g%v", conn.ConnectionId, v)
	}



	return conn.cli.WriteOK(nil)

}
