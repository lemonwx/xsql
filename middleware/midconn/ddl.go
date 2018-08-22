/**
 *  author: lim
 *  data  : 18-4-9 下午12:33
 */

package midconn

import (
	"strings"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/meta"
	"github.com/lemonwx/xsql/sqlparser"
)

func (conn *MidConn) handleDDL(stmt *sqlparser.DDL, sql string) error {
	// ddl send to all nodes
	log.Debugf("[%d]: recv ddl sql: %s: %v", conn.ConnectionId, sql, stmt)
	sql = conn.addExtraCol(sql)

	defer func() {
		log.Debug(conn.execNodes)
		for idx, back := range conn.execNodes {
			conn.pools[idx].PutConn(back)
			delete(conn.execNodes, idx)
		}
		log.Debug(conn.execNodes)
	}()

	rs, err := conn.ExecuteOnNodePool([]byte(sql), meta.GetFullNodeIdxs())
	if err != nil {
		return err
	}

	if stmt.Action == sqlparser.AST_CREATE {
		if stmt.Type == sqlparser.DATABASE {
			meta.AddDatabase(string(stmt.NewName))
		}

		if stmt.Type == sqlparser.TABLE {
			db := conn.db
			tb := string(stmt.TableName.Name)

			if stmt.TableName.Qualifier != nil {
				db = string(stmt.TableName.Qualifier)
			}
			meta.AddTable(db, tb)
		}
	}

	return conn.HandleExecRets(rs)
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
