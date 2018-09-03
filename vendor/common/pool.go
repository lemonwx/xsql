/**
 *  author: lim
 *  data  : 18-8-30 下午10:30
 */

package common

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/lemonwx/log"
)

const (
	MaxInitFailedSize = 5
)

type Conn interface {
	net.Conn
}

type Pool struct {
	idleConns chan *rpc.Client
	freeConns chan *rpc.Client

	New func() (*rpc.Client, error)
}

func NewPool(initSize, idleSize, maxSize uint32, f func() (*rpc.Client, error)) (*Pool, error) {
	if initSize > idleSize {
		return nil, fmt.Errorf("pool's init size must < idle size")
	}
	if idleSize > maxSize {
		return nil, fmt.Errorf("pool's idle size must < max size")
	}

	p := &Pool{New: f}
	p.idleConns = make(chan *rpc.Client, idleSize)
	p.freeConns = make(chan *rpc.Client, maxSize-idleSize)

	failedSize := 0
	count := uint32(0)
	for count < initSize {
		if conn, err := p.New(); err != nil {
			log.Errorf("conn to version server failed: %v", err)
			failedSize++
			if failedSize > MaxInitFailedSize {
				return nil, fmt.Errorf("too many errors when connect to version server")
			}
		} else {
			p.idleConns <- conn
			count++
			failedSize = 0
		}
	}

	for count < maxSize {
		if count < idleSize {
			p.idleConns <- nil
		} else {
			p.freeConns <- nil
		}
		count++
	}
	return p, nil
}

func (p *Pool) tryReuse(conn *rpc.Client) (*rpc.Client, error) {
	log.Debug(conn)
	if conn != nil {
		log.Debugf("return conn")
		return conn, nil
	}
	return p.New()
}

func (p *Pool) GetConnFromIdle() (*rpc.Client, error) {
	var conn *rpc.Client
	select {
	case conn = <-p.idleConns:
		return p.tryReuse(conn)
	default:
		return nil, fmt.Errorf("idle list empty, try get from free list")
	}
}

func (p *Pool) Get() (*rpc.Client, error) {
	// first try get conn from idle
	if conn, err := p.GetConnFromIdle(); err == nil {
		return conn, nil
	}

	// during this time,
	//      1. may some conn put back to idle list, we expect all conn get from idle list
	//   or 2. may free empty and idle not empty
	// so try get conn from both idle and free
	var conn *rpc.Client
	select {
	case conn = <-p.idleConns:
	case conn = <-p.freeConns:
	}
	return p.tryReuse(conn)
}

func (p *Pool) freeConn(conn *rpc.Client) {
	conn.Close()
	conn = nil
	select {
	case p.freeConns <- conn:
		return
	default:
		log.Errorf("unexpected both full of idle and free node list")
		return
	}
}

func (p *Pool) Put(conn *rpc.Client) {
	log.Debug(len(p.idleConns), len(p.freeConns))
	select {
	case p.idleConns <- conn:
		return
	default:
		p.freeConn(conn)
	}
}
