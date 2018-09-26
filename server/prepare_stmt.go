/**
 *  author: lim
 *  data  : 18-5-6 下午8:56
 */

package server

import (
	"fmt"

	"encoding/binary"

	"sync"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/xsql/node"
	"github.com/lemonwx/xsql/sqlparser"
)

type myStmt interface {
	prepare(idx int) error
	execute(data []byte) ([]*mysql.Result, error)
	response() error
	close() error
	reset()
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

	lockStartIdx int
}

func (bs *baseStmt) reset() {
	bs.args = map[int]interface{}{}
	bs.argTypes = bs.argTypes[:0]
	bs.argFlags = bs.argFlags[:0]
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
	if len(bs.args) != 0 {
		return nil
	}
	log.Debug(bs.cliArgCount)
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

func (bs *baseStmt) prepare(idx int) error {
	// get back conn of node[idx]
	back, err := bs.mid.getSingleBackConn(idx)
	if err != nil {
		return err
	}

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

func (bs *baseStmt) execute(data []byte, fun func(map[int]interface{}, uint32) ([]byte, error)) (
	[]*mysql.Result, error) {
	var shardList []int
	var err error
	if err = bs.parseArgs(data); err != nil {
		return nil, err
	}

	log.Debug(bs.args, data, sqlparser.String(bs.s))
	if shardList, err = bs.mid.getShardList(bs.s, bs.args); err != nil {
		return nil, err
	}

	if len(shardList) == 0 {
		// todo: return based on sql type
		switch v := bs.s.(type) {
		case *sqlparser.Select:
			rs := bs.mid.newEmptyResultset(v)
			ret := &mysql.Result{Resultset: rs}
			return []*mysql.Result{ret}, nil
		default:
			return nil, nil
		}
	}

	if err = bs.stmtChk(shardList); err != nil {
		return nil, err
	}

	mes := MultiExecSyncer{}
	mes.Add(len(shardList))
	for _, nodeIdx := range shardList {
		stmtId, ok := bs.svrStmtIds[nodeIdx]
		if !ok {
			return nil, newDefaultMySQLError(errInternal, "svr stmt id not exists")
		}

		go func(nodeIdx int, stmtId uint32) {
			var back *node.Node
			var svrData []byte
			var err error
			if back, err = bs.mid.getSingleBackConn(nodeIdx); err != nil {
				mes.appendErr(err)
				return
			}

			if svrData, err = fun(bs.args, stmtId); err != nil {
				mes.appendErr(err)
				return
			}

			if ret, err := bs.mid.execute(back, mysql.COM_STMT_EXECUTE, svrData); err != nil {
				mes.appendErr(err)
			} else {
				mes.appendRet(ret)
			}

			mes.Done()
		}(nodeIdx, stmtId)
	}
	mes.Wait()
	switch {
	case len(mes.errs) == len(shardList):
		return nil, mes.errs[0]
	case len(mes.rets) == len(shardList):
		return mes.rets, nil
	default:
		return nil, newMySQLErr(errMultiStmtExecNotEqual)
	}
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
		delete(bs.svrStmtIds, idx)
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

func (sel *selStmt) execute(data []byte) ([]*mysql.Result, error) {
	var wg sync.WaitGroup
	wg.Add(2)
	var vErr, eErr error
	var vInuse map[uint64]bool
	var rets []*mysql.Result
	go func() {
		vInuse, vErr = sel.mid.getCurVInUse(Select)
		wg.Done()
	}()
	go func() {
		f := func(args map[int]interface{}, stmtId uint32) ([]byte, error) {
			svrData := make([]byte, 0, len(data)+4+1+4)
			svrData = append(svrData, mysql.Uint32ToBytes(stmtId)...) //int<4> statement id
			svrData = append(svrData, 0)                              //int<1> flags:
			svrData = append(svrData, 1, 0, 0, 0)                     //int<4> Iteration count (always 1)
			svrData = append(svrData, data...)
			return svrData, nil
		}

		rets, eErr = sel.baseStmt.execute(data, f)
		wg.Done()
	}()
	wg.Wait()

	if vErr != nil {
		return nil, vErr
	}

	if eErr != nil {
		return nil, eErr
	}

	extraSize := len(sel.s.(*sqlparser.Select).ExtraCols)
	if err := sel.mid.chkInUse(&rets, extraSize, vInuse, true); err != nil {
		return nil, err
	}

	return rets, nil
}

type istStmt struct {
	*baseStmt
}

func (ist *istStmt) prepare(idx int) error {
	if err := ist.baseStmt.prepare(idx); err != nil {
		return err
	}
	ist.cliArgCount = ist.svrArgCount - 1
	ist.cliFieldCount = ist.svrFieldCount
	ist.mid.myStmts[ist.stmtId] = ist
	ist.args = map[int]interface{}{}
	ist.argTypes = make([]byte, 0, ist.cliArgCount)
	ist.argFlags = make([]byte, 0, ist.cliArgCount)
	return nil
}

func (ist *istStmt) execute(data []byte) ([]*mysql.Result, error) {
	if err := ist.mid.getNextVersion(); err != nil {
		return nil, err
	}

	f := func(args map[int]interface{}, stmtId uint32) ([]byte, error) {
		svrArgs := make([]interface{}, 1, len(args)+1)
		svrArgs[0] = int64(ist.mid.NextVersion)
		for _, v := range args {
			svrArgs = append(svrArgs, v)
		}
		ret := ist.mid.makePkt(svrArgs, stmtId)
		return ret, nil
	}

	return ist.baseStmt.execute(data, f)
}

type updStmt struct {
	*baseStmt
	lockStmt *selStmt
}

func (upd *updStmt) reset() {
	upd.baseStmt.reset()
	upd.lockStmt.reset()
}

func (upd *updStmt) prepare(idx int) error {
	if err := upd.baseStmt.prepare(idx); err != nil {
		return err
	}
	upd.cliArgCount = upd.svrArgCount - 1
	upd.cliFieldCount = upd.svrFieldCount
	upd.mid.myStmts[upd.stmtId] = upd
	upd.args = map[int]interface{}{}
	upd.argTypes = make([]byte, 0, upd.cliArgCount)
	upd.argFlags = make([]byte, 0, upd.cliArgCount)

	u, ok := upd.s.(*sqlparser.Update)
	if !ok {
		return newDefaultMySQLError(errInternal, "mid update stmt, but back stmt not update")
	}
	selList := []sqlparser.SelectExpr{&sqlparser.NonStarExpr{Expr: &sqlparser.ColName{Name: []byte("version")}}}
	lockS := &sqlparser.Select{
		Comments:    nil,
		Distinct:    "",
		SelectExprs: selList,
		From:        []sqlparser.TableExpr{&sqlparser.AliasedTableExpr{Expr: u.Table}},
		Where:       u.Where,
		GroupBy:     nil,
		Having:      nil,
		OrderBy:     u.OrderBy,
		Limit:       u.Limit,
		Lock:        " for update",
		ExtraCols:   selList,
	}

	if s, err := newMyStmt(lockS, upd.mid); err != nil {
		return err
	} else {
		upd.lockStmt = s.(*selStmt)
		if err := upd.lockStmt.prepare(idx); err != nil {
			return err
		}
	}

	return nil
}

func (upd *updStmt) execute(data []byte) ([]*mysql.Result, error) {
	ch := make(chan map[uint64]bool)
	var vErr error

	// get v inUse and next v
	go func() {
		defer close(ch)
		var vInUse map[uint64]bool
		vInUse, vErr = upd.mid.getCurVInUse(updateOrDelete)
		if vErr != nil {
			return
		}
		ch <- vInUse
	}()

	if err := upd.parseArgs(data); err != nil {
		return nil, err
	}

	// calc update exprs size
	upd.lockStartIdx = 0
	if s, ok := upd.s.(*sqlparser.Update); ok {
		// update stmt will add extra cols, first expr must be ValArg
		for _, expr := range s.Exprs[1:] {
			if _, ok := expr.Expr.(sqlparser.ValArg); ok {
				upd.lockStartIdx += 1
			}
		}
	}

	for idx := upd.lockStartIdx; idx < len(upd.args); idx += 1 {
		upd.lockStmt.args[idx] = upd.args[idx]
	}

	log.Debug(upd.args, upd.lockStmt.args)

	f := func(args map[int]interface{}, stmtId uint32) ([]byte, error) {
		svrArgs := make([]interface{}, 0, len(args)-upd.lockStartIdx)
		for idx := 0; idx < len(args); idx += 1 {
			svrArgs = append(svrArgs, args[idx+upd.lockStartIdx])
		}

		log.Debug(upd.lockStartIdx, args, svrArgs)
		ret := upd.mid.makePkt(svrArgs, stmtId)
		return ret, nil
	}

	ret, err := upd.lockStmt.baseStmt.execute(data, f)
	if err != nil {
		return nil, err
	}

	vInUse, ok := <-ch
	if !ok {
		return nil, newDefaultMySQLError(errGetVersionFailed, vErr)
	}
	extraColSize := len(upd.lockStmt.s.(*sqlparser.Select).ExtraCols)
	if err := upd.mid.chkInUse(&ret, extraColSize, vInUse, true); err != nil {
		return nil, err
	}

	f1 := func(args map[int]interface{}, stmtId uint32) ([]byte, error) {
		svrArgs := make([]interface{}, 1, len(args)+1)
		svrArgs[0] = int64(upd.mid.NextVersion)
		for idx := 0; idx < len(args); idx += 1 {
			svrArgs = append(svrArgs, args[idx])
		}

		log.Debug(svrArgs)
		ret := upd.mid.makePkt(svrArgs, stmtId)
		return ret, nil
	}

	return upd.baseStmt.execute(data, f1)
}

type delStmt struct {
	*baseStmt
	upd *updStmt
}

func (del *delStmt) execute(data []byte) ([]*mysql.Result, error) {
	return nil, nil
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

	switch v := s.(type) {
	case *sqlparser.Select:
		return &selStmt{baseStmt: stmt}, nil
	case *sqlparser.Insert:
		return &istStmt{baseStmt: stmt}, nil
	case *sqlparser.Delete:
		return &delStmt{baseStmt: stmt}, nil
	case *sqlparser.Update:
		v.Exprs[0].Expr = sqlparser.ValArg("?")
		stmt.sql = sqlparser.String(v)
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
