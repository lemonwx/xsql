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
	"math"
	"io"
	"github.com/lemonwx/xsql/middleware/meta"
	"utils"
)


var paramFieldData []byte = (&mysql.Field{}).Dump()
var columnFieldData []byte = (&mysql.Field{}).Dump()

type Stmt struct {
	id uint32
	ids []uint32

	params  int
	columns uint16

	args []interface{}

	s sqlparser.Statement

	sql string

	nodeIdx []int
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

	stmt := new(Stmt)

	sql = strings.TrimRight(sql, ";")
	stmt.s, err = sqlparser.Parse(sql)
	if err != nil {
		return fmt.Errorf(`parse sql "%s" error`, sql)
	}

	stmt.sql = sqlparser.String(stmt.s)

	// send prepare to node[0]
	if err = conn.prepare(stmt, 0); err != nil {
		return err
	}

	stmt.nodeIdx = []int{0}
	stmt.ids = []uint32{stmt.id}
	// send prepare result to mysql cli
	if err = conn.writePrepare(stmt); err != nil {
		return err
	}


	conn.stmts[stmt.id] = stmt

	stmt.ResetParams()
	return nil
}

func (conn *MidConn) handleStmtExecute(data []byte) error {
	log.Debugf("[%d] handle stmt execute %v", conn.ConnectionId, data)

	var err error

	if len(data) < 9 {
		return mysql.ErrMalformPacket
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

	if conn.nodeIdx, err = sqlparser.GetStmtShardListIndex(
		s.s, meta.GetRouter(conn.db), conn.makeBindVars(s.args)); err != nil {
			log.Debugf("[%d] get nodeidx failed: %v", conn.ConnectionId, err)
			return err
	}
	log.Debugf("[%d] get nodeidx %v", conn.ConnectionId, conn.nodeIdx)

	if err = conn.chkPrepare(s); err != nil {
		return err
	}

	log.Debugf("[%d] prepare stmt: %v, exec: %v", conn.ConnectionId, s, data)

	switch s.s.(type) {
	case *sqlparser.Select:
		return conn.ExecuteSelect(data)
	default:
		return UNEXPECT_MIDDLE_WARE_ERR
	}
}

func (conn *MidConn) ExecuteSelect(data []byte) error {
	var err error
	if err = conn.getVInUse(); err != nil {
		return err
	}

	conn.setupNodeStatus(conn.VersionsInUse, true, true)
	defer conn.setupNodeStatus(nil, false, false)

	if rets, err := conn.ExecuteMultiNode(mysql.COM_STMT_EXECUTE, data, conn.nodeIdx); err != nil {
		return err
	} else {
		return conn.HandleSelRets(rets)
	}
}


func (conn *MidConn) prepare(stmt *Stmt, idx int) error {
	return conn.nodes[idx].ExecutePrepare([]byte(stmt.sql), &stmt.id, &stmt.columns, &stmt.params)
}

func (conn *MidConn) chkPrepare(stmt *Stmt) error {

	if utils.CompareIntSlice(conn.nodeIdx, stmt.nodeIdx) {
		return nil
	}

	for _, idx := range conn.nodeIdx {
		if ! utils.ContainsIntSlice(stmt.nodeIdx, idx) {
			log.Debugf("[%d] node :%d need to prepare", conn.ConnectionId, idx )
			tmpStmt := new(Stmt)
			tmpStmt.s = stmt.s
			tmpStmt.sql = stmt.sql
			if err := conn.prepare(tmpStmt, idx); err != nil {
				return err
			} else {
				if tmpStmt.columns == stmt.columns && tmpStmt.params == tmpStmt.params {
					stmt.ids = append(stmt.ids, tmpStmt.id)
					stmt.nodeIdx = append(stmt.nodeIdx, idx)
				} else {
					return UNEXPECT_MIDDLE_WARE_ERR
				}
			}
		}
	}
	return nil
}

func (conn *MidConn) bindStmtArgs(s *Stmt, nullBitmap, paramTypes, paramValues []byte) error {
	args := s.args

	pos := 0

	var v []byte
	var n int = 0
	var isNull bool
	var err error

	for i := 0; i < s.params; i++ {
		if nullBitmap[i>>3]&(1<<(uint(i)%8)) > 0 {
			args[i] = nil
			continue
		}

		tp := paramTypes[i<<1]
		isUnsigned := (paramTypes[(i<<1)+1] & 0x80) > 0

		switch tp {
		case mysql.MYSQL_TYPE_NULL:
			args[i] = nil
			continue

		case mysql.MYSQL_TYPE_TINY:
			if len(paramValues) < (pos + 1) {
				return mysql.ErrMalformPacket
			}

			if isUnsigned {
				args[i] = uint8(paramValues[pos])
			} else {
				args[i] = int8(paramValues[pos])
			}

			pos++
			continue

		case mysql.MYSQL_TYPE_SHORT, mysql.MYSQL_TYPE_YEAR:
			if len(paramValues) < (pos + 2) {
				return mysql.ErrMalformPacket
			}

			if isUnsigned {
				args[i] = uint16(binary.LittleEndian.Uint16(paramValues[pos : pos+2]))
			} else {
				args[i] = int16((binary.LittleEndian.Uint16(paramValues[pos : pos+2])))
			}
			pos += 2
			continue

		case mysql.MYSQL_TYPE_INT24, mysql.MYSQL_TYPE_LONG:
			if len(paramValues) < (pos + 4) {
				return mysql.ErrMalformPacket
			}

			if isUnsigned {
				args[i] = uint32(binary.LittleEndian.Uint32(paramValues[pos : pos+4]))
			} else {
				args[i] = int32(binary.LittleEndian.Uint32(paramValues[pos : pos+4]))
			}
			pos += 4
			continue

		case mysql.MYSQL_TYPE_LONGLONG:
			if len(paramValues) < (pos + 8) {
				return mysql.ErrMalformPacket
			}

			if isUnsigned {
				args[i] = binary.LittleEndian.Uint64(paramValues[pos : pos+8])
			} else {
				args[i] = int64(binary.LittleEndian.Uint64(paramValues[pos : pos+8]))
			}
			pos += 8
			continue

		case mysql.MYSQL_TYPE_FLOAT:
			if len(paramValues) < (pos + 4) {
				return mysql.ErrMalformPacket
			}

			args[i] = float32(math.Float32frombits(binary.LittleEndian.Uint32(paramValues[pos : pos+4])))
			pos += 4
			continue

		case mysql.MYSQL_TYPE_DOUBLE:
			if len(paramValues) < (pos + 8) {
				return mysql.ErrMalformPacket
			}

			args[i] = math.Float64frombits(binary.LittleEndian.Uint64(paramValues[pos : pos+8]))
			pos += 8
			continue

		case mysql.MYSQL_TYPE_DECIMAL, mysql.MYSQL_TYPE_NEWDECIMAL, mysql.MYSQL_TYPE_VARCHAR,
			mysql.MYSQL_TYPE_BIT, mysql.MYSQL_TYPE_ENUM, mysql.MYSQL_TYPE_SET, mysql.MYSQL_TYPE_TINY_BLOB,
			mysql.MYSQL_TYPE_MEDIUM_BLOB, mysql.MYSQL_TYPE_LONG_BLOB, mysql.MYSQL_TYPE_BLOB,
			mysql.MYSQL_TYPE_VAR_STRING, mysql.MYSQL_TYPE_STRING, mysql.MYSQL_TYPE_GEOMETRY,
			mysql.MYSQL_TYPE_DATE, mysql.MYSQL_TYPE_NEWDATE,
			mysql.MYSQL_TYPE_TIMESTAMP, mysql.MYSQL_TYPE_DATETIME, mysql.MYSQL_TYPE_TIME:
			if len(paramValues) < (pos + 1) {
				return mysql.ErrMalformPacket
			}

			v, isNull, n, err = LengthEnodedString(paramValues[pos:])
			pos += n
			if err != nil {
				return err
			}

			if !isNull {
				args[i] = v
				continue
			} else {
				args[i] = nil
				continue
			}
		default:
			return fmt.Errorf("Stmt Unknown FieldType %d", tp)
		}
	}
	return nil
}

func LengthEnodedString(b []byte) ([]byte, bool, int, error) {
	// Get length
	num, isNull, n := LengthEncodedInt(b)
	if num < 1 {
		return nil, isNull, n, nil
	}

	n += int(num)

	// Check data length
	if len(b) >= n {
		return b[n-int(num) : n], false, n, nil
	}
	return nil, false, n, io.EOF
}

func LengthEncodedInt(b []byte) (num uint64, isNull bool, n int) {
	switch b[0] {

	// 251: NULL
	case 0xfb:
		n = 1
		isNull = true
		return

		// 252: value of following 2
	case 0xfc:
		num = uint64(b[1]) | uint64(b[2])<<8
		n = 3
		return

		// 253: value of following 3
	case 0xfd:
		num = uint64(b[1]) | uint64(b[2])<<8 | uint64(b[3])<<16
		n = 4
		return

		// 254: value of following 8
	case 0xfe:
		num = uint64(b[1]) | uint64(b[2])<<8 | uint64(b[3])<<16 |
			uint64(b[4])<<24 | uint64(b[5])<<32 | uint64(b[6])<<40 |
			uint64(b[7])<<48 | uint64(b[8])<<56
		n = 9
		return
	}

	// 0-250: value of first byte
	num = uint64(b[0])
	n = 1
	return
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



func (conn *MidConn) makeBindVars(args []interface{}) map[string]interface{} {
	bindVars := make(map[string]interface{}, len(args))

	for i, v := range args {
		bindVars[fmt.Sprintf("v%d", i)] = v
	}

	return bindVars
}