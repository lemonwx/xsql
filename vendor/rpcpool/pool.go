package rpcpool

import (
	"errors"
	"net/rpc"
	"sync"
)

var DIALERR = errors.New("dial rpc server err")
var POOL_HAS_CLOSED_ERR = errors.New("pool has closed err")
var POOL_HAS_FULL_ERR = errors.New("pool has full err")

type Factory func() (*rpc.Client, error)

type ConnPool struct {
	mu sync.Mutex

	conns chan *Conn

	factory Factory

	closed bool
}

func NewConnPool(factory Factory, initSize, maxSize int) (*ConnPool, error) {

	cp := &ConnPool{
		mu:      sync.Mutex{},
		conns:   make(chan *Conn, maxSize),
		factory: factory,
		closed:  false,
	}

	for i := 0; i < initSize; i++ {

		cli, err := cp.factory()
		if err != nil {
			cp.Close()
			return nil, err
		}
		cp.conns <- &Conn{cli: cli}
	}

	return cp, nil
}

func (cp *ConnPool) Get() (*Conn, error) {
	if cp.closed {
		return nil, POOL_HAS_CLOSED_ERR
	}

	select {
	case conn, ok := <-cp.conns:
		{
			if !ok {
				return nil, POOL_HAS_CLOSED_ERR
			}

			return conn, nil
		}
	default:
		{
			cli, err := cp.factory()
			if err != nil {
				return nil, err
			}
			return &Conn{cli: cli}, nil
		}
	}

}

func (cp *ConnPool) Put(conn *Conn) error {
	if cp.closed {
		return POOL_HAS_CLOSED_ERR
	}

	select {
	case cp.conns <- conn:
		{
			return nil
		}
	default:
		{
			conn.Close()
			return POOL_HAS_FULL_ERR
		}
	}
}

func (cp *ConnPool) Close() {
	if cp.closed {
		return
	}
	cp.mu.Lock()
	cp.closed = true
	close(cp.conns)

	for conn := range cp.conns {
		conn.Close()
	}
	cp.mu.Unlock()
}

func (cp *ConnPool) len() int {
	return len(cp.conns)
}
