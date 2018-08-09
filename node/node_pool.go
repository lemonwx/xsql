/**
 *  author: lim
 *  data  : 18-8-9 下午9:25
 */

package node

import (
	"fmt"

	"github.com/lemonwx/xsql/config"
	"github.com/lemonwx/xsql/errors"
)

const (
	MaxInitFailedSize = 5
)

type Pool struct {
	adminConn   *Node
	idleChan    chan *Node
	maxConnSize uint32
}

func NewNodePool(initSize, idleSize, maxConnSize uint32, cfg *config.Node) (*Pool, error) {
	m := &Pool{}
	m.idleChan = make(chan *Node, idleSize)
	m.maxConnSize = maxConnSize

	node := NewNode(cfg.Host, cfg.Port, cfg.User, cfg.Password, "", 0)
	if err := node.Connect(); err != nil {
		return nil, err
	} else {
		m.adminConn = node
	}

	failedSize := 0
	for count := uint32(0); count < initSize; {
		node := NewNode(cfg.Host, cfg.Port, cfg.User, cfg.Password, "", 0)
		if err := node.Connect(); err != nil {
			failedSize++
			if failedSize >= MaxInitFailedSize {
				err := errors.New(fmt.Errorf("init conn failed too many times"))
				return nil, err
			}
		} else {
			failedSize = 0
			m.idleChan <- node
			count++
		}
	}

	return m, nil
}

func (pool *Pool) GetConn() *Node {
	conn := <-pool.idleChan
	return conn
}
