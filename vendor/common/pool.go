/**
 *  author: lim
 *  data  : 18-8-30 下午10:30
 */

package common

import (
	"fmt"
	"net"

	"github.com/lemonwx/log"
)

const (
	MaxInitFailedSize = 5
)

type Conn interface {
	net.Conn
}

type Pool struct {
	idleConns chan Conn
	freeConns chan Conn

	New func() (Conn, error)
}

func NewPool(initSize, idleSize, maxSize uint32, f func() (Conn, error)) (*Pool, error) {
	if initSize > idleSize {
		return nil, fmt.Errorf("pool's init size must < idle size")
	}
	if idleSize > maxSize {
		return nil, fmt.Errorf("pool's idle size must < max size")
	}

	p := &Pool{New: f}
	p.idleConns = make(chan Conn, idleSize)
	p.freeConns = make(chan Conn, maxSize-idleSize)

	failedSize := 0
	count := uint32(0)
	for count < initSize {
		if conn, err := p.New(); err != nil {
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

func (p *Pool) tryReuse(conn Conn) (Conn, error) {
	if conn != nil {
		return conn, nil
	}
	return p.New()
}

func (p *Pool) GetConnFromIdle() (Conn, error) {
	var conn Conn
	select {
	case conn = <-p.idleConns:
		return p.tryReuse(conn)
	default:
		return nil, fmt.Errorf("idle list empty, try get from free list")
	}
}

func (p *Pool) Get() (Conn, error) {
	// first try get conn from idle
	if conn, err := p.GetConnFromIdle(); err == nil {
		return conn, nil
	}

	// during this time,
	//      1. may some conn put back to idle list, we expect all conn get from idle list
	//   or 2. may free empty and idle not empty
	// so try get conn from both idle and free
	var conn Conn
	select {
	case conn = <-p.idleConns:
	case conn = <-p.freeConns:
	}
	return p.tryReuse(conn)
}

func (p *Pool) freeConn(conn Conn) {
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

func (p *Pool) Put(conn Conn) {
	select {
	case p.idleConns <- conn:
		return
	default:
		p.freeConn(conn)
	}
}
