/**
 *  author: lim
 *  data  : 18-3-24 下午4:35
 */

package server

import (
	"bytes"
	"errors"
	"net"
	"strings"
	"sync"
	"sync/atomic"

	"hack"

	"time"

	"fmt"

	"strconv"

	"github.com/lemonwx/TxMgr/proto"
	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/client"
	"github.com/lemonwx/xsql/config"
	"github.com/lemonwx/xsql/meta"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/xsql/node"
	"github.com/lemonwx/xsql/router"
	"github.com/lemonwx/xsql/sqlparser"
)

var baseConnId uint32 = 1000

type MultiExecSyncer struct {
	sync.Mutex
	sync.WaitGroup
	errs []error
	rets []*mysql.Result
}

func NewMS(size int) *MultiExecSyncer {
	ms := &MultiExecSyncer{
		errs: make([]error, 0, size),
		rets: make([]*mysql.Result, 0, size),
	}
	ms.Add(size)
	return ms
}

func (ms *MultiExecSyncer) appendErr(err error) {
	ms.Lock()
	ms.errs = append(ms.errs, err)
	ms.Unlock()
}

func (ms *MultiExecSyncer) appendRet(ret *mysql.Result) {
	ms.Lock()
	ms.rets = append(ms.rets, ret)
	ms.Unlock()
}

type MidConn struct {
	cli           *client.CliConn
	nodes         []*node.Node
	db            string
	closed        bool
	ConnectionId  uint32
	RemoteAddr    string
	status        uint16
	defaultStatus uint16

	VersionsInUse map[uint64]bool
	NextVersion   uint64

	nodeIdx []int // node that has exec sql in the trx

	baseStmtId uint32
	stmts      map[uint32]*Stmt
	myStmts    map[uint32]myStmt

	pools     map[int]*node.Pool
	execNodes map[int]*node.Node
	svr       *Server
	stat      *Stat
	resp      chan *response
}

func NewMidConn(conn net.Conn, cfg *config.Conf, pools map[int]*node.Pool, s *Server) (*MidConn, error) {

	var err error
	midConn := new(MidConn)

	// conn id
	baseConnId = atomic.AddUint32(&baseConnId, 1)
	midConn.ConnectionId = baseConnId

	cli := client.NewClieConn(conn, midConn.ConnectionId, cfg)
	err = cli.Handshake()
	if err != nil {
		cli.WriteError(err)
		return nil, err
	}

	// cli conn between mysqlCli and xsql, this cli has handshake with mysql cli
	midConn.cli = cli
	midConn.db = cli.Db

	// init and connect to back mysql server
	midConn.nodes = make([]*node.Node, len(meta.GetNodeAddrs()))

	/*
		for idx, nodeCfg := range meta.GetNodeAddrs() {
			tmpNode := node.NewNode(nodeCfg.Host, nodeCfg.Port, nodeCfg.User, nodeCfg.Password,
				cli.Db, midConn.ConnectionId)
			midConn.nodes[idx] = tmpNode
		}

		var wg sync.WaitGroup
		wg.Add(len(midConn.nodes))

		for idx := 0; idx < len(midConn.nodes); idx += 1 {
			go func(tmp int) {
				if err = midConn.nodes[tmp].Connect(); err != nil {
					log.Errorf("connected to backend mysqld %d failed: %v", tmp, err)
				} else {
					log.Debugf("[%d] connect to mysqld [%s] success",
						midConn.ConnectionId, midConn.nodes[tmp])
				}
				wg.Done()
			}(idx)
		}
		wg.Wait()

		if err != nil {
			midConn.cli.WriteError(err)
			return nil, err
		} else {
			// hand shake with cli finish
			log.Debugf("[%d] hand shake with cli and mysqld finish", midConn.ConnectionId)
			if err := midConn.cli.WriteOK(nil); err != nil {
				return nil, err
			}
			midConn.cli.SetPktSeq(0)
		}
	*/
	midConn.pools = pools
	midConn.execNodes = make(map[int]*node.Node)
	midConn.closed = false
	midConn.RemoteAddr = conn.RemoteAddr().String()
	midConn.defaultStatus = mysql.SERVER_STATUS_AUTOCOMMIT
	midConn.status = midConn.defaultStatus

	midConn.VersionsInUse = nil
	midConn.NextVersion = 0

	midConn.stmts = make(map[uint32]*Stmt)
	midConn.myStmts = map[uint32]myStmt{}
	midConn.baseStmtId = 1
	midConn.svr = s
	midConn.stat = newStat()
	midConn.resp = make(chan *response)

	return midConn, nil
}

func (conn *MidConn) Serve() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		if err := conn.cli.ReadIntoBuf(); err != nil {
			log.Debugf("[%d] read failed: %v, client link off %s",
				conn.ConnectionId, err, conn.RemoteAddr)
			go func() {
				for {
					time.Sleep(time.Millisecond)
					err := conn.abnormalClose()
					if err == nil {
						break
					}
				}
			}()
		}
		wg.Done()
	}()

	go func() {
		for {
			data, err := conn.cli.ReadFromBuf()
			if err != nil {
				conn.Close()
			}

			if err = conn.dispatch(data); err != nil {
				conn.cli.WriteError(err)
			}
			conn.cli.SetPktSeq(0)
		}
		wg.Done()
	}()

	wg.Wait()
}

func (conn *MidConn) dispatch(sql []byte) error {
	ts := time.Now()
	defer func() {
		conn.stat.Dispatch.add(int64(time.Since(ts)))
	}()
	opt, sql := sql[0], sql[1:]
	log.Debugf("[%d] general: %d:%s", conn.ConnectionId, opt, sql)
	switch opt {
	case mysql.COM_QUERY:
		return conn.handleQuery(string(sql))
	case mysql.COM_QUIT:
		conn.Close()
	case mysql.COM_FIELD_LIST:
		return conn.handleFieldList(sql)
	case mysql.COM_INIT_DB:
		return conn.handleUse(sql)
	case mysql.COM_STMT_PREPARE:
		return conn.handlePrepare(hack.String(sql))
	case mysql.COM_STMT_EXECUTE:
		//return conn.handleExecute(sql)
		//return conn.handleStmtExecute(sql)
		return conn.handleStmtTrx(sql)
	case mysql.COM_STMT_CLOSE:
		return conn.handleStmtClose(sql)
	}

	return nil
}

func (conn *MidConn) handleQuery(sql string) error {

	sql = strings.TrimRight(sql, ";")
	ts := time.Now()
	stmt, err := sqlparser.Parse(sql)
	conn.stat.SqlparseT.add(int64(time.Since(ts)))
	if err != nil {
		log.Errorf("[%d] parse [%s] failed: %v", conn.ConnectionId, sql, err)
		return err
	}

	switch v := stmt.(type) {
	case *sqlparser.Set:
		return conn.handleSet(v, sql)
	case *sqlparser.Begin:
		conn.handleBegin(true)
		return conn.cli.WriteOK(nil)
	case *sqlparser.Commit, *sqlparser.Rollback:
		err = conn.handleTrxFinish(sqlparser.String(v))
		if err != nil {
			return err
		}
		return conn.cli.WriteOK(nil)
	case *sqlparser.DDL:
		return conn.handleDDL(v, sql)
	case *sqlparser.SimpleSelect:
		return conn.handleSimpleSelect(v, sql)
	case *sqlparser.Show:
		return conn.handleShow(v, sql)
	case *sqlparser.Select, *sqlparser.Insert, *sqlparser.Update, *sqlparser.Delete:
		log.Debugf("[%d] sql:[%s] need to execute in trx", conn.ConnectionId, sql)
		return conn.handleTrx(stmt, sql)
	case *sqlparser.Admin:
		return conn.handleAdmin(v, sql)
	case *sqlparser.Kill:
		return conn.handleKill(v)
	default:
		return errors.New("not support this sql")
	}
	return nil
}

func (conn *MidConn) handleKill(kill *sqlparser.Kill) error {
	log.Debugf("[%d] handle %s", conn.ConnectionId, sqlparser.String(kill))
	killId, err := strconv.ParseUint(string(kill.Id), 10, 64)
	if err != nil {
		return err
	}
	ids := conn.svr.GetBackIds(uint32(killId))
	log.Debugf("[%d] %v to be kill", conn.ConnectionId, ids)
	for nodeIdx, id := range ids {
		back, err := conn.getSingleBackConn(nodeIdx)
		if err != nil {
			return err
		}
		killSql := fmt.Sprintf("kill %d", id)
		_, err = conn.execute(back, mysql.COM_QUERY, []byte(killSql))
		if err != nil {
			log.Debugf("[%d] kill %d under node: %s failed: %v", conn.ConnectionId, id, back, err)
			return err
		}
	}
	return conn.cli.WriteOK(nil)
}

func (conn *MidConn) handleFieldList(data []byte) error {
	index := bytes.IndexByte(data, 0x00)
	table := string(data[0:index])
	wildcard := string(data[index+1:])

	if conn.db == "" {
		return mysql.NewDefaultError(mysql.ER_NO_DB_ERROR)
	}

	back, err := conn.pools[0].GetConn(conn.db)
	if err != nil {
		return err
	}

	defer conn.pools[0].PutConn(back)

	if fs, err := back.FieldList(table, wildcard); err != nil {
		log.Errorf("node 0 execute fieldList failed: %v", err)
		return err
	} else {
		return conn.cli.WriteFieldList(conn.status, fs)
	}
}

func (conn *MidConn) handleUse(db []byte) error {
	tmp := string(db)
	conn.db = tmp
	conn.cli.Db = tmp

	/*
		// rets, errs := conn

		rets, err := conn.ExecuteMultiNode(mysql.COM_INIT_DB, db, meta.GetFullNodeIdxs())
		if err != nil {
			return err
		}
		return conn.HandleExecRets(rets)
	*/
	return conn.cli.WriteOK(nil)
}

func (conn *MidConn) writeResultset(status uint16, r *mysql.Resultset) error {
	return conn.cli.WriteResultset(status, r)
}

func (conn *MidConn) ExecuteMultiNode(opt uint8, sql []byte, nodeIdxs []int) ([]*mysql.Result, error) {

	nodeSize := len(nodeIdxs)
	wg := sync.WaitGroup{}
	wg.Add(nodeSize)

	rets := make([]*mysql.Result, 0, nodeSize)
	errs := make([]error, 0, nodeSize)

	for idx := 0; idx < nodeSize; idx += 1 {
		go func(tmp int) {
			if rs, err := conn.nodes[nodeIdxs[tmp]].Execute(opt, sql); err != nil {
				errs = append(errs, err)
			} else {
				rets = append(rets, rs)
			}
			wg.Done()
		}(idx)
	}
	wg.Wait()

	// handle returns from multi nodes
	switch {
	case len(errs) == 0 && len(rets) == nodeSize:
		// 所有节点返回的都是执行成功
		log.Debugf("[%d] all %d nodes return success", conn.ConnectionId, nodeSize)
		return rets, nil

	case len(rets) == 0 && len(errs) == nodeSize:
		// 所有节点都执行出错
		desc := errs[0].Error()
		for _, err := range errs[1:] {
			// 错误内容不一致，预期外的情况
			if err.Error() != desc {
				log.Debugf("[%d] all %d nodes return error, but err's desc not equal", conn.ConnectionId, nodeSize)
				return nil, UNEXPECT_MIDDLE_WARE_ERR
			}
		}

		// 所有节点返回的 错误内容一致
		if err, ok := errs[0].(*mysql.SqlError); ok {
			if err.Code == errRowsInuseByOthers {
				log.Debugf("[%d] all %d nodes return error: %v", conn.ConnectionId, nodeSize, err.Message)
				return nil, errs[0]
			}
		}

		log.Debugf("[%d] all %d nodes return error: %v", conn.ConnectionId, errs[0].Error())
		return nil, errs[0]

	default:
		// 既有错误的节点，也有成功的节点， 分布式事务可能不一致, 必须手动 rollback
		err := mysql.NewDefaultError(mysql.MID_ER_SQL_ONLY_SUCCESS_IN_PARTLY_OF_NODE)
		log.Errorf("[%d] %v", err)
		return nil, err
	}
}

func (conn *MidConn) ExecuteMultiNodePrepare(args []interface{}, stmtMeta map[int]uint32, nodeIdxs []int) ([]*mysql.Result, error) {
	rets := make([]interface{}, len(nodeIdxs))
	wg := sync.WaitGroup{}
	wg.Add(len(nodeIdxs))

	for idx := 0; idx < len(nodeIdxs); idx += 1 {
		go func(tmp int) {
			nodeStmtId := stmtMeta[nodeIdxs[tmp]]
			execData := conn.makePkt(args, nodeStmtId)
			if rs, err := conn.nodes[nodeIdxs[tmp]].Execute(mysql.COM_STMT_EXECUTE, execData); err != nil {
				rets[tmp] = err
			} else {
				rets[tmp] = rs
			}
			wg.Done()
		}(idx)
	}
	wg.Wait()

	rs := make([]*mysql.Result, 0, len(nodeIdxs))
	for _, ret := range rets {
		if err, ok := ret.(error); ok {
			return nil, err
		}
		rs = append(rs, ret.(*mysql.Result))
	}
	return rs, nil
}

func (conn *MidConn) HandleExecRets(rets []*mysql.Result) error {
	if rs, err := conn.mergeExecResult(rets); err != nil {
		return conn.cli.WriteError(err)
	} else if rs != nil {
		return conn.cli.WriteOK(rs)
	} else {
		return conn.cli.WriteOK(nil)
	}

}

func (conn *MidConn) HandleSelRets(rets []*mysql.Result) error {

	rs := make([]*mysql.Resultset, len(rets))
	for idx, ret := range rets {
		rs[idx] = ret.Resultset
	}

	return conn.cli.WriteResultsets(conn.status, rs)
}

func (conn *MidConn) mergeExecResult(rets []*mysql.Result) (*mysql.Result, error) {
	ret := new(mysql.Result)

	for _, r := range rets {
		ret.Status |= r.Status
		ret.AffectedRows += r.AffectedRows
	}
	return ret, nil
}

func (conn *MidConn) mergeSelResult(rets []*mysql.Result) (*mysql.Result, error) {
	if len(rets) == 0 {
		return nil, UNEXPECT_MIDDLE_WARE_ERR
	}
	if rets[0] == nil {
		return nil, UNEXPECT_MIDDLE_WARE_ERR
	}

	rs := rets[0]

	// merge data row numbers
	finalLen := 0
	for _, r := range rets {
		finalLen += len(r.RowDatas)
	}

	tgtRs := &mysql.Resultset{
		Fields:   rs.Fields,
		RowDatas: make([]mysql.RowData, finalLen),
	}

	copy(tgtRs.RowDatas, rs.RowDatas)
	from := len(rs.RowDatas)
	for _, rs := range rets[1:] {
		copy(tgtRs.RowDatas[from:], rs.RowDatas)
		from += len(rs.RowDatas)
	}

	return &mysql.Result{
		Status:    0,
		Resultset: tgtRs,
	}, nil

}

func (conn *MidConn) abnormalClose() error {
	if conn.closed {
		return nil
	}

	log.Debugf("[%d] client may abnormal exit, kill back conn link with it", conn.ConnectionId)
	for idx, node := range conn.execNodes {
		back, err := conn.svr.pools[idx].GetConn("")
		if err != nil {
			return err
		}
		defer conn.svr.pools[idx].PutConn(back)
		conn.execute(back, mysql.COM_QUERY, []byte(fmt.Sprintf("kill %d", node.BackCoId)))
		node.Close()
	}

	conn.Close()
	return nil
}

func (conn *MidConn) Close() {
	conn.closed = true
	conn.cli.Close()
	conn.clearExecNodes([]byte("rollback"))
}

func (conn *MidConn) newEmptyResultset(stmt *sqlparser.Select) *mysql.Resultset {
	r := new(mysql.Resultset)
	r.Fields = make([]*mysql.Field, len(stmt.SelectExprs))

	for i, expr := range stmt.SelectExprs {
		r.Fields[i] = &mysql.Field{}
		switch e := expr.(type) {
		case *sqlparser.StarExpr:
			r.Fields[i].Name = []byte("*")
		case *sqlparser.NonStarExpr:
			if e.As != nil {
				r.Fields[i].Name = e.As
				r.Fields[i].OrgName = hack.Slice(sqlparser.String(e.Expr))
			} else {
				r.Fields[i].Name = hack.Slice(sqlparser.String(e.Expr))
			}
		default:
			r.Fields[i].Name = hack.Slice(sqlparser.String(e))
		}
	}

	r.Values = make([][]interface{}, 0)
	r.RowDatas = make([]mysql.RowData, 0)

	return r
}

func (conn *MidConn) getShardList(stmt sqlparser.Statement, args map[int]interface{}) ([]int, error) {
	ts := time.Now()
	defer func() {
		conn.stat.RouteT.add(int64(time.Since(ts)))
	}()

	if conn.db == "" {
		err := mysql.NewDefaultError(mysql.ER_NO_DB_ERROR)
		log.Errorf("[%d] get conn.db failed: %v", conn.ConnectionId, err)
		return nil, err
	}

	if r, err := meta.GetRouter(conn.db); err != nil {
		log.Errorf("[%d] get router failed: %v", conn.ConnectionId, err)
		return nil, err
	} else {
		return router.GeneralShardList(r, stmt, args)
	}
}

func (conn *MidConn) getMultiBackConn(idxs []int) error {
	ts := time.Now()
	defer func() {
		conn.stat.GetConn.add(int64(time.Since(ts)))
	}()

	for _, idx := range idxs {
		if _, ok := conn.execNodes[idx]; ok {
			continue
		}

		back, err := conn.pools[idx].GetConn(conn.db)
		if err != nil {
			return err
		} else {
			back.ConnectionId = conn.ConnectionId
			conn.svr.StoreMidSession(conn.ConnectionId, back.BackCoId, idx)
			conn.execNodes[idx] = back
		}
	}
	return nil
}

func (conn *MidConn) getSingleBackConn(idx int) (*node.Node, error) {
	ts := time.Now()
	defer func() {
		conn.stat.GetConn.add(int64(time.Since(ts)))
	}()

	back, ok := conn.execNodes[idx]
	if ok {
		conn.svr.StoreMidSession(conn.ConnectionId, back.BackCoId, idx)
		return back, nil
	}
	back, err := conn.pools[idx].GetConn(conn.db)
	if err != nil {
		return nil, err
	}

	back.ConnectionId = conn.ConnectionId
	conn.execNodes[idx] = back
	conn.svr.StoreMidSession(conn.ConnectionId, back.BackCoId, idx)
	return back, nil
}

func (conn *MidConn) putConn(idx int, back *node.Node) {
	ts := time.Now()
	defer func() {
		conn.stat.PutConn.add(int64(time.Since(ts)))
	}()

	conn.svr.RmMidSession(conn.ConnectionId, idx)
	conn.pools[idx].PutConn(back)
}

func (conn *MidConn) getNextVersion() error {
	ts := time.Now()
	defer func() {
		conn.stat.VersionT.add(int64(time.Since(ts)))
	}()

	if conn.NextVersion == 0 {
		Push(proto.C, conn)
		r := <-conn.resp
		if r.Err != nil {
			log.Debugf("[%d] get from async version failed: %v", conn.ConnectionId, r.Err)
		}
		log.Debugf("[%d] get from async gtid: %v", conn.ConnectionId, r.Max)
		conn.NextVersion = r.Max
	}
	return nil
}

func (conn *MidConn) getCurVInUse(flag uint8) (map[uint64]bool, error) {
	ts := time.Now()
	defer func() {
		conn.stat.VersionT.add(int64(time.Since(ts)))
	}()

	var cmd uint8
	if flag == updateOrDelete && conn.NextVersion == 0 {
		cmd = proto.C_Q
	} else {
		cmd = proto.Q
	}

	Push(cmd, conn)
	r := <-conn.resp
	log.Debug(r)
	if r.Err != nil {
		return nil, r.Err
	}

	if cmd == proto.C_Q {
		conn.NextVersion = r.Max
	}

	return r.Active, nil
}

func (conn *MidConn) getVInUse() error {
	Push(proto.Q, conn)
	r := <-conn.resp
	conn.VersionsInUse = r.Active
	return nil
}

func (conn *MidConn) getNodeIdxs(stmt sqlparser.Statement, bindVars map[string]interface{}) error {
	var err error
	if conn.db == "" {
		err = mysql.NewDefaultError(mysql.ER_NO_DB_ERROR)
		return err
	}

	conn.nodeIdx, err = conn.getShardList(stmt, nil)

	if err != nil {
		log.Debugf("[%d] get node idxs failed: %v", conn.ConnectionId, err)
		return err
	}
	log.Debugf("[%d] get node idxs: %v", conn.ConnectionId, conn.nodeIdx)

	return nil
}

func (conn *MidConn) execute(back *node.Node, opt uint8, data []byte) (*mysql.Result, error) {
	ts := time.Now()
	defer func() {
		conn.stat.ExecT.add(int64(time.Since(ts)))
	}()
	return back.Execute(opt, data)
}

func newMySQLErr(errCode uint16) *mysql.SqlError {
	return mysql.NewError(errCode, MySQLErrName[errCode])
}

func newDefaultMySQLError(errCode uint16, args ...interface{}) *mysql.SqlError {
	return mysql.NewDefaultError(errCode, args)
}
