/**
 *  author: lim
 *  data  : 18-3-24 下午4:35
 */

package midconn

import (
	"sync"
	"net"
	"sync/atomic"

	"github.com/lemonwx/xsql/client"
	"github.com/lemonwx/xsql/node"
	"github.com/lemonwx/xsql/middleware/meta"
	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/mysql"
)

var baseConnId uint32 = 1000

type MidConn struct {
	cli  *client.CliConn
	nodes []*node.Node
	db string
	closed bool
	COnnectionId uint32
	RemoteAddr net.Addr
	status uint16
}

func NewMidConn(conn net.Conn) (*MidConn, error) {
	// handshake with mysql client
	var err error
	cli := client.NewClieConn(conn)
	err = cli.Handshake()
	if err != nil {
		cli.WriteError(err)
		return nil, err
	}

	midConn := new(MidConn)
	// cli conn between mysqlCli and xsql, this cli has handshake with mysql cli
	midConn.cli = cli
	// init and connect to back mysql server
	midConn.nodes = make([]*node.Node, len(meta.NodeAddrs))

	for idx, nodeCfg := range meta.NodeAddrs {
		tmpNode := node.NewNode(nodeCfg.Host, nodeCfg.Port, nodeCfg.User, nodeCfg.Password, cli.Db)
		midConn.nodes[idx] = tmpNode
	}

	var wg sync.WaitGroup
	wg.Add(len(midConn.nodes))

	for idx := 0; idx < len(midConn.nodes); idx += 1{
		go func(tmp int) {
			if err = midConn.nodes[tmp].Connect(); err != nil {
				log.Errorf("connected to backend mysqld %d failed: %v", tmp, err)
			} else {
				log.Debugf("connect to mysqld [%v] success", midConn.nodes[tmp])
			}
			wg.Done()
		}(idx)
	}
	wg.Wait()

	if err != nil {
		midConn.cli.WriteError(err)
	} else {
		// hand shake with cli finish
		log.Debug("hand shake with cli and mysqld finish")
		if err := midConn.cli.WriteOK(nil); err != nil {
			return nil, err
		}
		midConn.cli.SetPktSeq(0)
	}
	midConn.closed = false
	baseConnId = atomic.AddUint32(&baseConnId, 1)
	midConn.COnnectionId = baseConnId
	midConn.RemoteAddr = conn.RemoteAddr()
	midConn.status = mysql.SERVER_STATUS_AUTOCOMMIT
	return midConn, nil
}

func (conn *MidConn) Serve() {
	for {
		conn.cli.SetPktSeq(0)
		data, err := conn.cli.ReadPacket()
		if err != nil {
			log.Errorf("cli conn read packet failed: %v", err)
			break
		}
		if err = conn.dispatch(data); err != nil {
			conn.cli.WriteError(err)
			conn.cli.SetPktSeq(0)
		}
		conn.cli.WriteOK(nil)
		conn.cli.SetPktSeq(0)
	}
}

func (conn *MidConn) dispatch(sql []byte) error {
	opt, sql := sql[0], sql[1:]
	log.Debugf("recv [%s] from cli", sql)
	switch opt {
	case mysql.COM_QUERY:
	case mysql.COM_QUIT:
	case mysql.COM_FIELD_LIST:
	case mysql.COM_INIT_DB:
	}

	return nil
}

func (conn *MidConn) handleUse(db []byte) error {
	tmp := string(db)
	conn.db = tmp
	conn.cli.Db = tmp
	// rets, errs := conn
	return nil
}

func (conn *MidConn) ExecuteMultiNode2(opt uint8, sql []byte, nodeIdxs []int)(
	[]*mysql.Result, error) {

	if nodeIdxs == nil {
		log.Debug("nodeIdxs is nil. use meta.FullNodeIdxs to execute")
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
		}(idx)
	}

	return nil, nil
}

func (conn *MidConn) Close() {
	conn.closed = true
	conn.cli.Close()
	for _, node := range conn.nodes {
		node.Close()
	}
}