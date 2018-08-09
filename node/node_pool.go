/**
 *  author: lim
 *  data  : 18-8-9 下午9:25
 */

package node

import (
	"fmt"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/config"
)

const (
	MaxInitFailedSize = 5
)

type Pool struct {
	adminConn *Node
	idleConns chan *Node
	freeConns chan *Node

	maxConnSize uint32

	host     string
	port     int
	user     string
	password string
}

func (p *Pool) NewConn() *Node {
	return NewNode(p.host, p.port, p.user, p.password, "", 0)
}

func (p *Pool) NewAndConnect() (*Node, error) {
	conn := NewNode(p.host, p.port, p.user, p.password, "", 0)
	if err := conn.Connect(); err != nil {
		return nil, err
	}
	return conn, nil
}

func NewNodePool(initSize, idleSize, maxConnSize uint32, cfg *config.Node) (*Pool, error) {
	if initSize > idleSize {
		return nil, fmt.Errorf("pool's init size must < idle size")
	}
	if idleSize > maxConnSize {
		return nil, fmt.Errorf("pool's idle size must < max size")
	}

	p := &Pool{
		maxConnSize: maxConnSize,
		host:        cfg.Host,
		port:        cfg.Port,
		user:        cfg.User,
		password:    cfg.Password,
	}

	if conn, err := p.NewAndConnect(); err != nil {
		return nil, err
	} else {
		p.adminConn = conn
	}

	failedSize := 0
	p.idleConns = make(chan *Node, idleSize)
	p.freeConns = make(chan *Node, maxConnSize-idleSize)

	count := uint32(0)
	for count < initSize {
		if conn, err := p.NewAndConnect(); err != nil {
			failedSize++
			if failedSize > MaxInitFailedSize {
				return nil, fmt.Errorf("too many errors when connect to backend")
			}
		} else {
			p.idleConns <- conn
			count++
			failedSize = 0
		}
	}

	for count < maxConnSize {
		if count < idleSize {
			p.idleConns <- p.NewConn()
		} else {
			p.freeConns <- p.NewConn()
		}
		count++
	}
	return p, nil
}

func (p *Pool) GetConn() (*Node, error) {
	var conn *Node
	select {
	case conn = <-p.idleConns:
	case conn = <-p.freeConns:
	}

	if conn.conn != nil {
		return conn, nil
	}

	if err := conn.Connect(); err != nil {
		p.freeConns <- conn
		return nil, err
	}
	return conn, nil
}

func (p *Pool) freeConn(node *Node) {
	node.Close()
	select {
	case p.freeConns <- node:
		return
	default:
		log.Errorf("unexpected both full of idle and free node list")
		return
	}
}

func (p *Pool) PutConn(node *Node) {
	select {
	case p.idleConns <- node:
		return
	default:
		p.freeConn(node)
	}
}

func (p *Pool) DumpInfo() {
	log.Debug(len(p.idleConns))
	log.Debug(len(p.freeConns))
}
