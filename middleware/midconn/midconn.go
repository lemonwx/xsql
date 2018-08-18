/**
 *  author: lim
 *  data  : 18-3-24 下午4:35
 */

package midconn

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"
	"sync/atomic"

	"hack"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/client"
	"github.com/lemonwx/xsql/config"
	"github.com/lemonwx/xsql/middleware/meta"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/xsql/node"
	"github.com/lemonwx/xsql/sqlparser"
)

var baseConnId uint32 = 1000

type MidConn struct {
	cli           *client.CliConn
	nodes         []*node.Node
	db            string
	closed        bool
	ConnectionId  uint32
	RemoteAddr    net.Addr
	status        []uint16 // 0:trx status, 1:defaultStatus at trx begin
	defaultStatus uint16

	VersionsInUse map[uint64]uint8
	NextVersion   uint64

	nodeIdx []int // node that has exec sql in the trx

	stmts map[uint32]*Stmt

	pools     map[int]*node.Pool
	execNodes map[int]*node.Node
}

func NewMidConn(conn net.Conn, cfg *config.Conf, pools map[int]*node.Pool) (*MidConn, error) {

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
	midConn.RemoteAddr = conn.RemoteAddr()
	midConn.defaultStatus = mysql.SERVER_STATUS_AUTOCOMMIT
	midConn.status = []uint16{midConn.defaultStatus, midConn.defaultStatus}

	midConn.VersionsInUse = nil
	midConn.NextVersion = 0

	midConn.stmts = make(map[uint32]*Stmt)

	return midConn, nil
}

func (conn *MidConn) Serve() {
	for {
		conn.cli.SetPktSeq(0)
		data, err := conn.cli.ReadPacket()
		if err != nil {
			log.Errorf("[%d] cli conn read packet failed: %v", conn.ConnectionId, err)
			break
		}
		if conn.status[0] == mysql.SERVER_NOT_SERVE {
			conn.handleServerNotServe(data)
		} else if err = conn.dispatch(data); err != nil {
			conn.cli.WriteError(err)
			conn.cli.SetPktSeq(0)
		}
	}
}

func (conn *MidConn) dispatch(sql []byte) error {
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
	}

	return nil
}

func (conn *MidConn) handleQuery(sql string) error {

	sql = strings.TrimRight(sql, ";")
	stmt, err := sqlparser.Parse(sql)
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
		err = conn.handleCommit(sqlparser.String(v))
		if err != nil {
			if conn.status[0] == mysql.SERVER_STATUS_IN_TRANS {
				conn.status[0] = mysql.SERVER_NOT_SERVE
			}
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

	default:
		return errors.New("not support this sql")
	}
	return nil
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
		return conn.cli.WriteFieldList(conn.status[0], fs)
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
			if err.Code == ROWS_IN_USE_BY_OTHER_SESSION {
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

	return conn.cli.WriteResultsets(conn.status[0], rs)
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

func (conn *MidConn) Close() {
	conn.closed = true
	conn.cli.Close()

	conn.clearExecNodes([]byte("rollback"))
}

func (c *MidConn) newEmptyResultset(stmt *sqlparser.Select) *mysql.Resultset {
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

func (conn *MidConn) getShardList(stmt sqlparser.Statement) ([]int, error) {
	var err error
	if conn.db == "" {
		err := mysql.NewDefaultError(mysql.ER_NO_DB_ERROR)
		log.Errorf("[%d] get conn.db failed: %v", conn.ConnectionId, err)
		return nil, err
	}

	r, err := meta.GetRouter(conn.db)
	if err != nil {
		log.Errorf("[%d] get router failed: %v", conn.ConnectionId, err)
		return nil, err
	}

	switch s := stmt.(type) {
	case *sqlparser.Select:
		p, err := sqlparser.GeneralPlanForSelect(r, s)
		if err != nil {
			return nil, err
		}
		return p.ShardList, nil
	default:
		return nil, fmt.Errorf("[%d] unsupported for this type of sql")
	}
}

func (conn *MidConn) getMultiBackConn(idxs []int) error {
	for _, idx := range idxs {
		if _, ok := conn.execNodes[idx]; ok {
			continue
		}

		back, err := conn.pools[idx].GetConn(conn.db)
		if err != nil {
			return err
		} else {
			conn.execNodes[idx] = back
		}
	}
	return nil
}

func (conn *MidConn) getSingleBackConn(idx int) (*node.Node, error) {
	back, ok := conn.execNodes[idx]
	if ok {
		return back, nil
	}
	back, err := conn.pools[idx].GetConn(conn.db)
	if err != nil {
		return nil, err
	}

	conn.execNodes[idx] = back
	return back, nil
}
