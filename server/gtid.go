/**
 *  author: lim
 *  data  : 18-8-30 下午8:55
 */

package server

import (
	"common"
	"net/rpc"
	"quicklock"
	"time"

	"github.com/lemonwx/TxMgr/proto"
	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/config"
)

const (
	SendByDemonTicker uint8 = iota
	SendByFullQueue
)

type req struct {
	cmd uint8
	co  *MidConn
	ts  time.Time
}

type response struct {
	Max    uint64
	Active map[uint64]bool
	Err    error
}

var (
	maxT      time.Duration
	maxQueue  chan *req
	ticker    <-chan time.Time
	pool      *common.Pool
	ql        *quicklock.QuickLock
	cmds      []uint8
	cos       []*MidConn
	gtidTodel []uint64
)

func SetVars(cfg *config.Conf) {
	// max wait time for an version request
	maxT = time.Duration(time.Microsecond * time.Duration(cfg.VWaitBatchTime))
	ticker = time.Tick(maxT)
	// queue of request
	maxQueue = make(chan *req, cfg.VWaitBatchCount)
	// only one receiver consume from maxQueue
	ql = quicklock.NewQL()
	// for every request
	log.Debugf("wait time: %v, wait count: %d", maxT, cfg.VWaitBatchCount)
}

func InitVPool(cfg *config.Conf) error {
	var err error
	pool, err = common.NewPool(cfg.VInitSize, cfg.VIdleSize, cfg.VMaxSize,
		func() (*rpc.Client, error) {
			return rpc.DialHTTP("tcp", cfg.VerSeqAddr)
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func sendall(flag uint8) {

	send := func() {
		reqs := &proto.Request{
			Ts: time.Now(),
		}
		cos := []*MidConn{}
		//hasSleep := false
		b := false
		for {
			select {
			case req := <-maxQueue:
				reqs.Cmds = append(reqs.Cmds, req.cmd)
				cos = append(cos, req.co)
				if req.cmd == proto.D {
					reqs.ToDels = append(reqs.ToDels, req.co.NextVersion)
				}
				req.co.stat.VWaitBatchT.add(int64(time.Since(req.ts)))
			default:
				/*
				if ! hasSleep {
					time.Sleep(maxT)
					hasSleep = true
				} else {
					b = true
				}*/
				b = true
			}

			if b {
				break
			}
		}

		if len(reqs.Cmds) == 0 {
			return
		}

		go func(request *proto.Request, cos []*MidConn) {
			resp := proto.Response{}

			log.Debugf("%d requests merge to send", len(request.Cmds))

			cli, err := pool.Get()
			if err != nil {
				panic(err)
			}

			defer pool.Put(cli)

			err = cli.Call("VSeq.PushReq", request, &resp)
			if err != nil {
				panic(err)
			}

			cos[0].stat.VWaitRespT.add(int64(time.Since(request.Ts)))
			cos[0].stat.BatchReqCount.add(int64(len(request.Cmds)))

			active := make(map[uint64]bool)
			for _, v := range resp.Maxs {
				active[v] = false
			}

			for _, v := range resp.Active {
				active[v] = false
			}

			r := &response{Err: nil}
			for idx, co := range cos {
				switch request.Cmds[idx] {
				case proto.Q:
					r.Active = active
					r.Max = 0
					co.resp <- r
				case proto.C:
					r.Active = nil
					r.Max = resp.Maxs[0]
					resp.Maxs = resp.Maxs[1:]
					co.resp <- r
				case proto.D:
					r.Active = nil
					r.Max = 0
					co.resp <- r
				case proto.C_Q:
					r.Active = active
					r.Max = resp.Maxs[0]
					resp.Maxs = resp.Maxs[1:]
					co.resp <- r
				}
			}

		}(reqs, cos)
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
		log.Debugf("demon ticker")
		sendall(SendByDemonTicker)
	}

}

func Push(cmd uint8, co *MidConn) {
	select {
	case maxQueue <- &req{cmd, co, time.Now()}:
	default:
		log.Debug("send all by full queue")
		sendall(SendByFullQueue)
		maxQueue <- &req{cmd, co, time.Now()}
		return
	}
}
