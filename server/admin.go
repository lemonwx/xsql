/**
 *  author: lim
 *  data  : 18-8-24 下午8:28
 */

package server

import (
	"fmt"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/xsql/sqlparser"
)

const (
	UseVersions = "versions"
)

func (conn *MidConn) handleAdmin(stmt *sqlparser.Admin, sql string) error {
	log.Debugf("[%d] handle admin command: %s", conn.ConnectionId, sql)
	log.Debug(stmt.Name, stmt.Values)

	switch string(stmt.Name) {
	case UseVersions:
		rs := &mysql.Resultset{
			Fields:     []*mysql.Field{{Name: []byte("node")}, {Name: []byte("cur version in use")}},
			FieldNames: map[string]int{"node": 0, "cur version in use": 1},
		}
		rows := make([]mysql.RowData, 0, 16)
		conn.svr.cos.RLock()
		log.Debug(conn.svr.cos.midConns)
		for remote, conn := range conn.svr.cos.midConns {
			v := []byte(fmt.Sprintf("%d", conn.NextVersion))
			log.Debug(remote, v)
			log.Debug(conn)

			row := make([]byte, 0, len(remote)+1+len(v)+1)
			row = append(row, byte(len(remote)))
			row = append(row, remote...)

			row = append(row, byte(len(v)))
			row = append(row, v...)

			rows = append(rows, row)
		}
		conn.svr.cos.RUnlock()

		log.Debug(rows)
		rs.RowDatas = rows

		return conn.writeResultset(conn.status, rs)
	default:
		return fmt.Errorf("unsupported this is admin sql")
	}

}
