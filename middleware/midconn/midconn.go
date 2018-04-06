/**
 *  author: lim
 *  data  : 18-3-24 下午4:35
 */

package midconn

import (
	"fmt"
	"github.com/lemonwx/xsql/client"
	"github.com/lemonwx/xsql/node"
	"net"
	"github.com/lemonwx/xsql/middleware/meta"
	"sync"
	"github.com/lemonwx/log"
	"sync/atomic"
	"github.com/lemonwx/xsql/middleware"
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
			}
			wg.Done()
		}(idx)
	}
	wg.Wait()

	if err != nil {
		midConn.cli.WriteError(err)
	} else {
		// hand shake with cli finish
		midConn.cli.WriteOK(nil)
		midConn.cli.SetPktSeq(0)
	}
	midConn.closed = false
	baseConnId = atomic.AddUint32(&baseConnId, 1)
	midConn.COnnectionId = baseConnId
	midConn.RemoteAddr = conn.RemoteAddr()
	midConn.status = middleware.SERVER_STATUS_AUTOCOMMIT
	return midConn, nil
}

func (conn *MidConn) Serve() {
	for {
		conn.cli.SetPktSeq(0)
		data, err := conn.cli.ReadPacket()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(data)
		conn.cli.WriteOK(nil)
		conn.cli.SetPktSeq(0)
	}
}
