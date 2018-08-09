/**
 *  author: lim
 *  data  : 18-8-9 下午9:25
 */

package node

import "github.com/lemonwx/xsql/config"

const (
	MaxInitFailedSize = 5
)

type Pool struct {
	adminConn   *Node
	idleChan    chan *Node
	maxConnSize uint32

	host     string
	port     int
	user     string
	password string
}

func (p *Pool) NewConn() (*Node, error) {
	conn := NewNode(p.host, p.port, p.user, p.password, "", 0)
	if err := conn.Connect(); err != nil {
		return nil, err
	}
	return conn, nil
}

func NewNodePool(initSize, idleSize, maxConnSize uint32, cfg *config.Node) (*Pool, error) {
	p := &Pool{
		maxConnSize: maxConnSize,
		host:        cfg.Host,
		port:        cfg.Port,
		user:        cfg.User,
		password:    cfg.Password,
	}

	if conn, err := p.NewConn(); err != nil {
		return nil, err
	} else {
		p.adminConn = conn
	}

	failedSize := 0
	p.idleChan = make(chan *Node, idleSize)
	for count := uint32(0); count < initSize; {
		if conn, err := p.NewConn(); err != nil {
			failedSize++
		} else {
			p.idleChan <- conn
			count++
			failedSize++
		}
	}

	return p, nil
}

func (pool *Pool) GetConn() *Node {
	conn := <-pool.idleChan
	return conn
}
