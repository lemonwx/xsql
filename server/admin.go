/**
 *  author: lim
 *  data  : 18-8-24 下午8:28
 */

package server

import (
	"fmt"

	"time"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/xsql/sqlparser"
)

const (
	UseVersions = "versions"
	TimeStat    = "time"
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
	case TimeStat:
		rs := &mysql.Resultset{
			Fields:     []*mysql.Field{{Name: []byte("phase")}, {Name: []byte("cost")}},
			FieldNames: map[string]int{"phase": 0, "cost": 1},
		}

		phases := map[string]time.Duration{"routeT": 0, "sqlparseT": 0, "versionT": 0, "chkInuseT": 0, "execT": 0}
		rs.RowDatas = make([]mysql.RowData, 0, len(phases))

		for _, conn := range conn.svr.cos.midConns {
			phases["routeT"] += conn.stat.routeT
			phases["sqlparseT"] += conn.stat.sqlparseT
			phases["versionT"] += conn.stat.versionT
			phases["chkInuseT"] += conn.stat.chkInuseT
			phases["execT"] += conn.stat.execT
		}

		for _, phase := range []string{"execT", "versionT", "sqlparseT", "routeT", "chkInuseT"} {
			t := phases[phase]
			lens := len(phase)
			row := make([]byte, 0, lens*2)
			row = append(row, byte(lens))
			row = append(row, []byte(phase)...)
			cost := t.String()
			row = append(row, byte(len(cost)))
			row = append(row, []byte(cost)...)

			rs.RowDatas = append(rs.RowDatas, row)
		}
		return conn.writeResultset(conn.status, rs)
	default:
		return fmt.Errorf("unsupported this is admin sql")
	}

}
