/**
 *  author: lim
 *  data  : 18-4-24 下午9:57
 */

package midconn

import (
	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/sqlparser"
	"github.com/lemonwx/xsql/mysql"
	"strings"
	"fmt"
)


var paramFieldData []byte
var columnFieldData []byte

type Stmt struct {
	id uint32

	params  int
	columns int

	args []interface{}

	s sqlparser.Statement

	sql string
}

func (conn *MidConn) writePrepare(s *Stmt) error {
	data := make([]byte, 4, 128)

	//status ok
	data = append(data, 0)
	//stmt id
	data = append(data, mysql.Uint32ToBytes(s.id)...)
	//number columns
	data = append(data, mysql.Uint16ToBytes(uint16(s.columns))...)
	//number params
	data = append(data, mysql.Uint16ToBytes(uint16(s.params))...)
	//filter [00]
	data = append(data, 0)
	//warning count
	data = append(data, 0, 0)

	if err := conn.cli.WritePacket(data); err != nil {
		return err
	}

	if s.params > 0 {
		for i := 0; i < s.params; i++ {
			data = data[0:4]
			data = append(data, []byte(paramFieldData)...)

			if err := conn.cli.WritePacket(data); err != nil {
				return err
			}
		}

		if err := conn.cli.WriteEOF(conn.status[0]); err != nil {
			return err
		}
	}

	if s.columns > 0 {
		for i := 0; i < s.columns; i++ {
			data = data[0:4]
			data = append(data, []byte(columnFieldData)...)

			if err := conn.cli.WritePacket(data); err != nil {
				return err
			}
		}

		if err := conn.cli.WriteEOF(conn.status[0]); err != nil {
			return err
		}

	}
	return nil
}

func (conn *MidConn) handlePrepare(sql []byte) error {
	log.Debugf("[%d] handle prepare %s", conn.ConnectionId, sql)
	newSql := string(sql)

	s := new(Stmt)

	newSql = strings.TrimRight(newSql, ";")

	var err error
	s.s, err = sqlparser.Parse(newSql)
	if err != nil {
		return fmt.Errorf(`parse sql "%s" error`, sql)
	}

	s.sql = newSql
	s.id = 1001

	conn.writePrepare(s)


	return nil
}
