/**
 *  author: lim
 *  data  : 18-8-30 下午8:55
 */

package server

import (
	"common"
	"time"

	"net/rpc"

	"quicklock"

	"github.com/lemonwx/TxMgr/proto"
	"github.com/lemonwx/log"
)

type req struct {
	cmd uint8
	co  *MidConn
}

type response struct {
	Max    uint64
	Active map[uint64]bool
	Err    error
}

var (
	maxT     = time.Duration(time.Millisecond * 300)
	maxQueue = make(chan *req, 1024)
	ticker   = time.Tick(maxT)
	pool     *common.Pool
	ql       = quicklock.NewQL()
)

func InitGtidPool(addr string) {
	log.Debug(addr)
	var err error
	pool, err = common.NewPool(10, 10, 10,
		func() (*rpc.Client, error) {
			return rpc.DialHTTP("tcp", addr)
		},
	)
	if err != nil {
		panic(err)
	}
}

func sendall(exReq *req) {

	send := func() {
		size := len(maxQueue)
		cmds := make([]uint8, 0, size)
		cos := make([]*MidConn, 0, size)
		gtidTodel := make([]uint64, 0, size)
		for {
			var req *req
			select {
			case req = <-maxQueue:
				cmds = append(cmds, req.cmd)
				cos = append(cos, req.co)
				if req.cmd == proto.D {
					gtidTodel = append(gtidTodel, req.co.NextVersion)
				}
			default:
				req = nil
			}

			if req == nil {
				break
			}
		}

		cmds = append(cmds, exReq.cmd)
		cos = append(cos, exReq.co)
		if exReq.cmd == proto.D {
			gtidTodel = append(gtidTodel, exReq.co.NextVersion)
		}

		req := proto.Request{
			Cmds:   cmds,
			ToDels: gtidTodel,
			Ts:     time.Now(),
		}

		log.Debugf("%d request merge to send", len(req.Cmds))

		cli, err := pool.Get()
		if err != nil {
			panic(err)
		}

		resp := &proto.Response{}
		err = cli.Call("VSeq.PushReq", req, &resp)
		if err != nil {
			panic(err)
		}

		pool.Put(cli)

		log.Debug(resp.Maxs, resp.Active)

		log.Debugf("response: ")
		for idx, co := range cos {
			switch cmds[idx] {
			case proto.Q:
				co.resp <- &response{Active: resp.Active}
			case proto.C:
				max := resp.Maxs[0]
				resp.Maxs = resp.Maxs[1:]
				co.resp <- &response{Max: max}
			case proto.D:
				co.resp <- &response{Err: nil}
			case proto.C_Q:
				max := resp.Maxs[0]
				resp.Maxs = resp.Maxs[1:]
				co.resp <- &response{Max: max}
			}
		}
	}

	if ql.Lock() {
		send()
		ql.UnLock()
	} else {
		log.Debugf("another send all executing, return")
	}
}

func RequestSender() {
	for {
		<-ticker
		log.Debugf("send all by demon ticker")
		req := <-maxQueue
		sendall(req)
	}

}

func Push(cmd uint8, co *MidConn) {
	select {
	case maxQueue <- &req{cmd, co}:
	default:
		log.Debug("send all by full queue")
		sendall(&req{cmd, co})
		return
	}
}
