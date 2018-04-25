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
	"encoding/binary"
	"strconv"
)


var paramFieldData []byte = (&mysql.Field{}).Dump()
var columnFieldData []byte = (&mysql.Field{}).Dump()

type Stmt struct {
	id uint32

	params  int
	columns uint16

	args []interface{}

	s sqlparser.Statement

	sql string
}


func (s *Stmt) ResetParams() {
	s.args = make([]interface{}, s.params)
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
			data = append(data, paramFieldData...)

			if err := conn.cli.WritePacket(data); err != nil {
				return err
			}
		}

		if err := conn.cli.WriteEOF(conn.status[0]); err != nil {
			return err
		}
	}

	if s.columns > 0 {
		for i := uint16(0); i < s.columns; i++ {
			data = data[0:4]
			data = append(data, columnFieldData...)

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

func (conn *MidConn) handlePrepare(sql string) error {
	log.Debugf("[%d] handle prepare %s", conn.ConnectionId, sql)

	var err error

	if conn.db == "" {
		return mysql.NewDefaultError(mysql.ER_NO_DB_ERROR)
	}

	s := new(Stmt)

	sql = strings.TrimRight(sql, ";")
	s.s, err = sqlparser.Parse(sql)
	if err != nil {
		return fmt.Errorf(`parse sql "%s" error`, sql)
	}

	s.sql = sql

	// send prepare to node[0]
	if err = conn.prepare(s, 0); err != nil {
		return err
	}

	// send prepare result to mysql cli
	if err = conn.writePrepare(s); err != nil {
		return err
	}

	s.ResetParams()
	return nil
}

func (conn *MidConn) prepare(stmt *Stmt, idx int) error {
	return conn.nodes[idx].ExecutePrepare([]byte(stmt.sql), &stmt.id, &stmt.columns, &stmt.params)
}

/*
func (conn *MidConn) handleStmtExecute(data []byte) error {
	if len(data) < 9 {
		return ErrMalformPacket
	}

	pos := 0
	id := binary.LittleEndian.Uint32(data[0:4])
	pos += 4

	s, ok := c.stmts[id]
	if !ok {
		return NewDefaultError(ER_UNKNOWN_STMT_HANDLER,
			strconv.FormatUint(uint64(id), 10), "stmt_execute")
	}

	flag := data[pos]
	pos++
	//now we only support CURSOR_TYPE_NO_CURSOR flag
	if flag != 0 {
		return NewError(ER_UNKNOWN_ERROR, fmt.Sprintf("unsupported flag %d", flag))
	}

	//skip iteration-count, always 1
	pos += 4

	var nullBitmaps []byte
	var paramTypes []byte
	var paramValues []byte

	paramNum := s.params

	if paramNum > 0 {
		nullBitmapLen := (s.params + 7) >> 3
		if len(data) < (pos + nullBitmapLen + 1) {
			return ErrMalformPacket
		}
		nullBitmaps = data[pos : pos+nullBitmapLen]
		pos += nullBitmapLen

		//new param bound flag
		if data[pos] == 1 {
			pos++
			if len(data) < (pos + (paramNum << 1)) {
				return ErrMalformPacket
			}

			paramTypes = data[pos : pos+(paramNum<<1)]
			pos += (paramNum << 1)

			paramValues = data[pos:]
		}

		if err := c.bindStmtArgs(s, nullBitmaps, paramTypes, paramValues); err != nil {
			return err
		}
	}

	var err error

	switch stmt := s.s.(type) {
	case *sqlparser.Select:
		err = c.handleSelect(stmt, s.sql, s.args)
	case *sqlparser.Insert:
		err = c.handleExec(s.s, s.sql, s.args)
	case *sqlparser.Update:
		err = c.handleExec(s.s, s.sql, s.args)
	case *sqlparser.Delete:
		err = c.handleExec(s.s, s.sql, s.args)
	case *sqlparser.Replace:
		err = c.handleExec(s.s, s.sql, s.args)
	default:
		err = fmt.Errorf("command %T not supported now", stmt)
	}

	s.ResetParams()

	return err
	return nil
}

/*
func (conn *MidConn) handleStmtExecute(data []byte) error {
	if len(data) < 9 {
		return UNEXPECT_MIDDLE_WARE_ERR
	}

	pos := 0
	id := binary.LittleEndian.Uint32(data[0:4])
	pos += 4

	s, ok := conn.stmts[id]
	if !ok {
		return mysql.NewDefaultError(mysql.ER_UNKNOWN_STMT_HANDLER,
			strconv.FormatUint(uint64(id), 10), "stmt_execute")
	}

	flag := data[pos]
	pos++
	//now we only support CURSOR_TYPE_NO_CURSOR flag
	if flag != 0 {
		return mysql.NewError(mysql.ER_UNKNOWN_ERROR, fmt.Sprintf("unsupported flag %d", flag))
	}

	//skip iteration-count, always 1
	pos += 4

	var nullBitmaps []byte
	var paramTypes []byte
	var paramValues []byte

	paramNum := s.params

	if paramNum > 0 {
		nullBitmapLen := (s.params + 7) >> 3
		if len(data) < (pos + nullBitmapLen + 1) {
			return mysql.ErrMalformPacket
		}
		nullBitmaps = data[pos : pos+nullBitmapLen]
		pos += nullBitmapLen

		//new param bound flag
		if data[pos] == 1 {
			pos++
			if len(data) < (pos + (paramNum << 1)) {
				return mysql.ErrMalformPacket
			}

			paramTypes = data[pos : pos+(paramNum<<1)]
			pos += (paramNum << 1)

			paramValues = data[pos:]
		}

		if err := conn.bindStmtArgs(s, nullBitmaps, paramTypes, paramValues); err != nil {
			return err
		}
	}

	var err error

	switch stmt := s.s.(type) {
	case *sqlparser.Select:
		err = conn.handleSelect(stmt, s.sql, s.args)
	case *sqlparser.Insert:
		err = conn.handleExec(s.s, s.sql, s.args)
	case *sqlparser.Update:
		err = conn.handleExec(s.s, s.sql, s.args)
	case *sqlparser.Delete:
		err = conn.handleExec(s.s, s.sql, s.args)
	case *sqlparser.Replace:
		err = conn.handleExec(s.s, s.sql, s.args)
	default:
		err = fmt.Errorf("command %T not supported now", stmt)
	}

	s.ResetParams()

	return err
}
*/