/**
 *  author: lim
 *  data  : 18-5-6 下午8:56
 */

package server

import (
	"fmt"

	"encoding/binary"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/xsql/sqlparser"
)

type myStmt interface {
	prepare(idx int) error
	execute(data []byte) error
	response() error
	close() error
}

type baseStmt struct {
	stmtId uint32

	cliFieldCount uint16
	svrFieldCount uint16

	cliArgCount int
	svrArgCount int

	svrStmtIds map[int]uint32 // key: nodeIdx, val: svr resp stmt id

	mid *MidConn
	s   sqlparser.Statement
	sql string

	args     map[int]interface{}
	argTypes []byte
	argFlags []byte
}

func (bs *baseStmt) reset() {
	bs.args = map[int]interface{}{}
	bs.argTypes = bs.argTypes[:0]
	bs.argFlags = bs.argFlags[:0]
}

func (bs *baseStmt) prepare(idx int) error {
	// get back conn of node[idx]
	back, err := bs.mid.getSingleBackConn(idx)
	if err != nil {
		return err
	}
	defer bs.mid.putConn(idx, back)

	// send prepare cmd to node[idx]'s svr
	var id uint32
	var fieldCount uint16
	var argCount int
	err = back.ExecutePrepare([]byte(bs.sql), &id, &fieldCount, &argCount)
	if err != nil {
		return err
	}

	// chk and assign
	if _, ok := bs.svrStmtIds[idx]; ok {
		log.Errorf("[%d] had send prepare cmd for this node and sql, but receive again", bs.mid.ConnectionId)
		return newMySQLErr(errRepeatPrepare)
	} else {
		bs.svrStmtIds[idx] = id
	}

	if bs.svrFieldCount == 0 {
		bs.svrFieldCount = fieldCount
	} else if bs.svrFieldCount != fieldCount {
		return newMySQLErr(errMultiPrepareNotEqual)
	}

	if bs.svrArgCount == 0 {
		bs.svrArgCount = argCount
	} else if bs.svrArgCount != argCount {
		return newMySQLErr(errMultiPrepareNotEqual)
	}

	return nil
}

func (bs *baseStmt) close() error {
	for idx, stmtId := range bs.svrStmtIds {
		back, err := bs.mid.getSingleBackConn(idx)
		if err != nil {
			return err
		}
		err = back.WriteCmd(mysql.COM_STMT_CLOSE, mysql.Uint32ToBytes(stmtId))
		if err != nil {
			return err
		}
	}

	return nil
}

func (bs *baseStmt) stmtChk(shardList []int) error {
	for _, nodeIdx := range shardList {
		if _, ok := bs.svrStmtIds[nodeIdx]; !ok {
			if err := bs.prepare(nodeIdx); err != nil {
				return err
			}
		}
	}
	return nil
}

func (bs *baseStmt) parseArgs(data []byte) error {
	if bs.cliArgCount == 0 {
		return nil
	}
	pos := 0
	bitMapzie := (bs.cliArgCount + 7) / 8
	nullBitMap := data[pos : pos+bitMapzie]
	pos += bitMapzie
	readFieldType := data[pos]
	pos += 1
	if readFieldType != 1 {
		return newMySQLErr(errUnsupportedStmtExecWithoutFieldType)
	}
	for idx := 0; idx < bs.cliArgCount; idx += 1 {
		bs.argTypes = append(bs.argTypes, data[pos])
		pos += 1
		bs.argFlags = append(bs.argFlags, data[pos])
		pos += 1
	}

	for idx := 0; idx < bs.cliArgCount; idx += 1 {
		if nullBitMap[idx>>3]&(1<<(uint(idx)%8)) > 0 {
			bs.args[idx] = nil
			continue
		}

		tp := bs.argTypes[idx]
		isUnsigned := bs.argFlags[idx]&mysql.UNSIGNED_FLAG > 0
		switch tp {
		case mysql.MYSQL_TYPE_NULL:
			bs.args[idx] = nil
		case mysql.MYSQL_TYPE_LONGLONG:
			if isUnsigned {
				bs.args[idx] = binary.LittleEndian.Uint64(data[pos : pos+8])
			} else {
				bs.args[idx] = int64(binary.LittleEndian.Uint64(data[pos : pos+8]))
			}
			pos += 8
		default:
			return newDefaultMySQLError(errUnsupportedStmtFieldType, tp)
		}
	}

	return nil
}

func (bs *baseStmt) execute(data []byte) error {
	return nil
}

func (bs *baseStmt) response() error {
	data := make([]byte, 4, 128)

	//status ok
	data = append(data, 0)
	//stmt id
	data = append(data, mysql.Uint32ToBytes(bs.stmtId)...)
	//number columns
	data = append(data, mysql.Uint16ToBytes(uint16(bs.cliFieldCount))...)
	//number params
	data = append(data, mysql.Uint16ToBytes(uint16(bs.cliArgCount))...)
	//filter [00]
	data = append(data, 0)
	//warning count
	data = append(data, 0, 0)

	if err := bs.mid.cli.WritePacket(data); err != nil {
		return err
	}

	if bs.cliArgCount > 0 {
		for i := 0; i < bs.cliArgCount; i++ {
			data = data[0:4]
			data = append(data, paramFieldData...)

			if err := bs.mid.cli.WritePacket(data); err != nil {
				return err
			}
		}

		if err := bs.mid.cli.WriteEOF(bs.mid.status); err != nil {
			return err
		}
	}

	if bs.cliFieldCount > 0 {
		for i := uint16(0); i < bs.cliFieldCount; i++ {
			data = data[0:4]
			data = append(data, columnFieldData...)

			if err := bs.mid.cli.WritePacket(data); err != nil {
				return err
			}
		}

		if err := bs.mid.cli.WriteEOF(bs.mid.status); err != nil {
			return err
		}

	}
	return nil
}

type selStmt struct {
	*baseStmt
}

func (sel *selStmt) prepare(idx int) error {
	if err := sel.baseStmt.prepare(idx); err != nil {
		return err
	}

	sel.cliArgCount = sel.svrArgCount
	sel.cliFieldCount = sel.svrFieldCount - 1
	sel.mid.myStmts[sel.stmtId] = sel
	sel.args = map[int]interface{}{}
	sel.argTypes = make([]byte, 0, sel.cliArgCount)
	sel.argFlags = make([]byte, 0, sel.cliArgCount)
	return nil
}

func (sel *selStmt) execute(data []byte) error {
	if err := sel.parseArgs(data); err != nil {
		return err
	}

	shardList, err := sel.mid.getShardList(sel.s, sel.args)
	if err != nil {
		return err
	}

	if err := sel.stmtChk(shardList); err != nil {
		return err
	}

	mes := MultiExecSyncer{}
	mes.Add(len(shardList))
	for _, nodeIdx := range shardList {
		stmtId, ok := sel.svrStmtIds[nodeIdx]
		if !ok {
			return newDefaultMySQLError(errInternal, "svr stmt id not exists")
		}
		go func(node int, id uint32) {
			back, err := sel.mid.getSingleBackConn(node)
			if err != nil {
				mes.appendErr(err)
				return
			}

			svrData := make([]byte, 0, len(data)+4+1+4)
			svrData = append(svrData, mysql.Uint32ToBytes(id)...) //int<4> statement id
			svrData = append(svrData, 0)                          //int<1> flags:
			svrData = append(svrData, 1, 0, 0, 0)                 //int<4> Iteration count (always 1)
			svrData = append(svrData, data...)
			if ret, err := back.Execute(mysql.COM_STMT_EXECUTE, svrData); err != nil {
				mes.appendErr(err)
			} else {
				mes.appendRet(ret)
			}
			mes.Done()
		}(nodeIdx, stmtId)
	}
	mes.Wait()

	switch {
	case len(mes.errs) == len(sel.svrStmtIds):
	case len(mes.rets) == len(sel.svrStmtIds):
	default:
		return newMySQLErr(errMultiStmtExecNotEqual)
	}

	log.Debug(sel.args)
	log.Debug(sel.argTypes)
	log.Debug(sel.argFlags)

	return nil
}

type istStmt struct {
	*baseStmt
}

type updStmt struct {
	*baseStmt
	lockStmt *selStmt
}

type delStmt struct {
	*baseStmt
	upd *updStmt
}

func newMyStmt(s sqlparser.Statement, co *MidConn) (myStmt, error) {
	co.baseStmtId += 1
	stmt := &baseStmt{
		s:          s,
		mid:        co,
		sql:        sqlparser.String(s),
		svrStmtIds: map[int]uint32{},
		stmtId:     co.baseStmtId,
	}

	switch s.(type) {
	case *sqlparser.Select:
		return &selStmt{baseStmt: stmt}, nil
	case *sqlparser.Insert:
		return &istStmt{baseStmt: stmt}, nil
	case *sqlparser.Delete:
		return &delStmt{baseStmt: stmt}, nil
	case *sqlparser.Update:
		return &updStmt{baseStmt: stmt}, nil
	default:
		log.Errorf("[%d] unsupported prepare for this sql", co.ConnectionId)
		return nil, newMySQLErr(errUnsupportedSql)
	}

}

type Stmt struct {
	id uint32

	cliParams  int
	nodeParams int

	cliColumns  uint16
	nodeColumns uint16

	cliArgs  []interface{}
	nodeArgs []interface{}

	s sqlparser.Statement

	sql       string
	originSql string

	stmtIdMeta map[int]uint32

	//forUpdateStmts []*Stmt
	forUpdateStmts  map[int]*Stmt
	forUpStmtIdMeta map[int]uint32
	forUpdateSql    string

	updateStmts      map[int]*Stmt
	updateStmtIdMeta map[int]uint32
	updateSql        string

	firstPrepare bool
}

func NewStmt() *Stmt {
	stmt := new(Stmt)
	stmt.stmtIdMeta = make(map[int]uint32)

	stmt.forUpdateStmts = make(map[int]*Stmt)
	stmt.forUpStmtIdMeta = make(map[int]uint32)

	stmt.updateStmtIdMeta = make(map[int]uint32)
	stmt.updateStmts = make(map[int]*Stmt)

	stmt.firstPrepare = true

	return stmt
}

func (s *Stmt) InitParams() {
	s.cliArgs = make([]interface{}, s.cliParams)
	s.nodeArgs = make([]interface{}, s.nodeParams)
}

func (s *Stmt) ResetParams(size int) {
	for idx := 0; idx < s.cliParams; idx += 1 {
		s.cliArgs[idx] = nil
	}
	for idx := 0; idx < s.nodeParams; idx += 1 {
		s.nodeArgs[idx] = nil
	}
}

func (s *Stmt) ChkEqual(params int, columns uint16) error {
	if s.firstPrepare {
		return fmt.Errorf("should not call ChkEqual at first prepare")
	} else {
		if s.nodeParams == params && s.nodeColumns == columns {
			return nil
		}
	}

	return fmt.Errorf("exec prepare between multi nodes ret not equal")
}
