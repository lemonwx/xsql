/**
 *  author: lim
 *  data  : 18-3-24 下午4:35
 */

package midconn

import (
	"bytes"
	"errors"
	"net"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/client"
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

	VersionsInUse map[string]uint8
	NextVersion   []byte

	NodeIdxs []int // node that has exec sql in the trx
}

func NewMidConn(conn net.Conn) (*MidConn, error) {

	var err error
	midConn := new(MidConn)

	// conn id
	baseConnId = atomic.AddUint32(&baseConnId, 1)
	midConn.ConnectionId = baseConnId

	cli := client.NewClieConn(conn, midConn.ConnectionId)
	err = cli.Handshake()
	if err != nil {
		cli.WriteError(err)
		return nil, err
	}

	// cli conn between mysqlCli and xsql, this cli has handshake with mysql cli
	midConn.cli = cli

	// init and connect to back mysql server
	midConn.nodes = make([]*node.Node, len(meta.NodeAddrs))

	for idx, nodeCfg := range meta.NodeAddrs {
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
				log.Debugf("[%d] connect to mysqld [%v] success",
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
	midConn.closed = false
	midConn.RemoteAddr = conn.RemoteAddr()
	midConn.defaultStatus = mysql.SERVER_STATUS_AUTOCOMMIT
	midConn.status = []uint16{midConn.defaultStatus, midConn.defaultStatus}

	midConn.VersionsInUse = nil
	midConn.NextVersion = nil

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
		if err = conn.dispatch(data); err != nil {
			conn.cli.WriteError(err)
			conn.cli.SetPktSeq(0)
		}
	}
}

func (conn *MidConn) dispatch(sql []byte) error {
	opt, sql := sql[0], sql[1:]
	log.Debugf("[%d] recv [%d:%s] from cli", conn.ConnectionId, opt, sql)
	switch opt {
	case mysql.COM_QUERY:
		return conn.handleQuery(string(sql))
	case mysql.COM_QUIT:
	case mysql.COM_FIELD_LIST:
		return conn.handleFieldList(sql)
	case mysql.COM_INIT_DB:
		return conn.handleUse(sql)
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
		/*
			1. get next version
			2. get versions in use
		*/
		conn.status = []uint16{mysql.SERVER_STATUS_IN_TRANS, ^mysql.SERVER_STATUS_AUTOCOMMIT}
		return conn.cli.WriteOK(nil)
	case *sqlparser.Commit, *sqlparser.Rollback:
		err = conn.handleCommit(nil, sql)
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
		log.Debugf("[%d] sql need to execute in trx", conn.ConnectionId)
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

	if fs, err := conn.nodes[0].FieldList(table, wildcard); err != nil {
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
	// rets, errs := conn

	rets, err := conn.ExecuteMultiNode(mysql.COM_INIT_DB, db, nil)
	if err != nil {
		return err
	}
	return conn.HandleExecRets(rets)
}

func (conn *MidConn) writeResultset(status uint16, r *mysql.Resultset) error {
	return conn.cli.WriteResultset(status, r)
}

func (conn *MidConn) ExecuteMultiNode(opt uint8, sql []byte, nodeIdxs []int) (
	[]*mysql.Result, error) {

	if nodeIdxs == nil {
		log.Debugf("[%d] nodeIdxs is nil. use meta.FullNodeIdxs to execute",
			conn.ConnectionId)
		nodeIdxs = meta.FullNodeIdxs
	}

	rets := make([]interface{}, len(nodeIdxs))

	wg := sync.WaitGroup{}
	wg.Add(len(nodeIdxs))

	for idx := 0; idx < len(nodeIdxs); idx += 1 {
		go func(tmp int) {
			if rs, err := conn.nodes[nodeIdxs[tmp]].Execute(opt, sql); err != nil {
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

	if len(rets) == 0 {
		return UNEXPECT_MIDDLE_WARE_ERR
	}

	rs := make([]*mysql.Resultset, len(rets))
	for idx, ret := range rets {
		rs[idx] = ret.Resultset
	}
	log.Debugf("----%v-----", conn.status[0])
	return conn.cli.WriteResultsets(conn.status[0], rs)

	/*
		if rs, err := conn.mergeSelResult(rets); err != nil {
			log.Errorf("merge select result failed: %v", err)
			return conn.cli.WriteError(err)
		} else if rs != nil {
			return conn.cli.WriteResultset(conn.status, rs.Resultset)
		} else {
			return UNEXPECT_MIDDLE_WARE_ERR
		}
	*/

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
	for _, node := range conn.nodes {
		node.Close()
	}
}
