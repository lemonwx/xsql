/**
 *  author: lim
 *  data  : 18-4-10 下午10:16
 */

package version

import (
	"net/rpc"
	"sync"
)

const(
	DIALERR = iota
	USEING
	FREE
	CLOSED

)

var lock sync.Mutex

type Pool struct {
	clis []*rpc.Client
	used []uint8
	mu   sync.RWMutex
}

var pool Pool
func NewRpcPool(size int, addr string) {
	pool = Pool{}
	pool.clis = make([]*rpc.Client, size)
	pool.used = make([]uint8, size)

	for idx := 0; idx < size; idx += 1 {
		cli, err := rpc.DialHTTP("tcp", addr)
		pool.clis[idx] = cli
		if err == nil {
			pool.used[idx] = DIALERR
		} else {
			pool.used[idx] = FREE
		}
	}
}

func (p *Pool) getConn() *rpc.Client {
	p.mu.Lock()
	defer p.mu.Unlock()
	for idx := 0; idx < len(p.used); idx += 1 {
		if p.used[idx] == FREE {
			p.used[idx] = USEING
			return p.clis[idx]
		}
	}
	return nil
}

func (p *Pool) Close() {
	for idx, cli := range p.clis {
		cli.Close()
		p.used[idx] = CLOSED
	}
}