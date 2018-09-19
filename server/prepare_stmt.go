/**
 *  author: lim
 *  data  : 18-5-6 下午8:56
 */

package server

import (
	"fmt"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/sqlparser"
)

type myStmt interface {
	prepare(idx int) error
	execute(args ...interface{}) error
	response() error
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

func (bs *baseStmt) execute(args ...interface{}) error {
	return nil
}

func (bs *baseStmt) response() error {
	return bs.mid.cli.WriteOK(nil)
}

type selStmt struct {
	*baseStmt
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
