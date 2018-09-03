/**
 *  author: lim
 *  data  : 18-8-24 下午8:28
 */

package server

import (
	"fmt"

	"reflect"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/xsql/sqlparser"
)

const (
	UseVersions = "versions"
	TimeStat    = "time"
	Clear       = "clear"
)

func (conn *MidConn) handleAdmin(stmt *sqlparser.Admin, sql string) error {
	log.Debugf("[%d] handle admin command: %s", conn.ConnectionId, sql)

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
			Fields: []*mysql.Field{
				{Name: []byte("phase")},
				{Name: []byte("cost")},
				{Name: []byte("count")},
				{Name: []byte("avg")},
			},
			FieldNames: map[string]int{
				"phase": 0, "all": 1, "counts": 2, "avg": 3},
		}

		ret := newStat()
		t := reflect.TypeOf(*ret)
		v := reflect.ValueOf(*ret)

		for _, co := range conn.svr.cos.midConns {
			sVal := reflect.ValueOf(*co.stat)
			for i := 0; i < t.NumField(); i++ {
				sField := sVal.Field(i).Interface().(field)
				tField := v.Field(i).Interface().(field)
				tField.plus(sField)
			}
		}

		for _, stat := range conn.svr.stats {
			sVal := reflect.ValueOf(*stat)
			for i := 0; i < t.NumField(); i++ {
				sField := sVal.Field(i).Interface().(field)
				tField := v.Field(i).Interface().(field)
				tField.plus(sField)
			}
		}

		rs.RowDatas = make([]mysql.RowData, 0, t.NumField()+1)
		for i := 0; i < t.NumField(); i++ {
			phase := t.Field(i).Name
			field := v.Field(i).Interface().(field)
			row := make([]byte, 0, len(phase)*4)

			row = append(row, byte(len(phase)))
			row = append(row, phase...)

			row = append(row, field.fmt()...)
			rs.RowDatas = append(rs.RowDatas, row)
		}

		row := []byte{}
		row = append(row, byte(len("theory")))
		row = append(row, "theory"...)
		row = append(row, 1)
		row = append(row, 32)
		row = append(row, 1)
		row = append(row, 32)
		theory := ret.getTheoryAvg().String()
		row = append(row, byte(len(theory)))
		row = append(row, theory...)
		rs.RowDatas = append(rs.RowDatas, row)
		return conn.writeResultset(conn.status, rs)
	case Clear:
		conn.svr.stats = []*Stat{}
		for _, co := range conn.svr.cos.midConns {
			co.stat.clear()
		}
		return conn.cli.WriteOK(nil)
	default:
		return fmt.Errorf("unsupported this is admin sql")
	}

}
