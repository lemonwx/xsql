/**
 *  author: lim
 *  data  : 18-3-24 下午4:35
 */

package midconn

import (
	"github.com/lemonwx/xsql/client"
	"github.com/lemonwx/xsql/node"
	"fmt"
	"os"
	"net"
)

var (
	node1Addr = "192.168.1.5:5518"
	node2Addr = "192.168.1.5:5520"
)

type MidConn struct {
	cli *client.CliConn
	node []*node.Node
}


func NewMidConn(conn net.Conn) *MidConn {
	// handshake with mysql client
	cli := client.NewClieConn(conn)
	err := cli.Handshake()
	if err != nil {
		cli.WriteError(err)
		os.Exit(-1)
	}

	midConn := new(MidConn)
	// cli conn between mysqlCli and xsql, this cli has handshake with mysql cli
	midConn.cli = cli
	// init and connect to back mysql server
	node1 := node.NewNode(node1Addr)
	node2 := node.NewNode(node2Addr)

	midConn.node = []*node.Node{node1, node2}

	return midConn
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