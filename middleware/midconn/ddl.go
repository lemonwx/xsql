/**
 *  author: lim
 *  data  : 18-4-9 下午12:33
 */

package midconn

import (
	"github.com/lemonwx/xsql/sqlparser"
	"github.com/lemonwx/xsql/mysql"
)

func (conn *MidConn) handleDDL(stmt *sqlparser.DDL, sql string) error {
	// ddl send to all nodes
	rs, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(sql), nil)
	if err != nil {
		return err
	} else {
		return conn.HandleExecRets(rs)
	}
}
