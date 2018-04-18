/**
 *  author: lim
 *  data  : 18-4-9 下午12:33
 */

package midconn

import (
	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/xsql/sqlparser"
	"strings"
)

func (conn *MidConn) handleDDL(stmt *sqlparser.DDL, sql string) error {
	// ddl send to all nodes

	log.Debugf("[%d]: recv ddl sql: %s", conn.ConnectionId, sql)
	sql = conn.addExtraCol(sql)
	rs, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(sql), nil)
	if err != nil {
		return err
	} else {
		return conn.HandleExecRets(rs)
	}
}

func (conn *MidConn) addExtraCol(sql string) string {
	if strings.Contains(strings.ToLower(sql), "table") {
		log.Debugf("[%d] create sql need to add extra col", conn.ConnectionId)
		idx := strings.Index(sql, "(")
		if idx != -1 {
			sql = sql[:idx+1] + extraColDef + sql[idx+1:]
		}
	}
	return sql
}
