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
	maxT     time.Duration
	maxQueue chan *req
	ticker   <-chan time.Time
	pool     *common.Pool
	ql       *quicklock.QuickLock
	cmds     []uint8
	cos      []*MidConn
	toDels   []uint64
)

func SetVars(cfg *config.Conf) {
	// max wait time for an version request
	maxT = time.Duration(time.Microsecond * time.Duration(cfg.VWaitBatchTime))
	ticker = time.Tick(maxT)
	// queue of request
	maxQueue = make(chan *req, cfg.VWaitBatchCount)
	// only one receiver consume from maxQueue
	ql = quicklock.NewQL()
	//
	cmds = make([]uint8, 0, cfg.VWaitBatchCount)
	cos = make([]*MidConn, 0, cfg.VWaitBatchCount)
	toDels = make([]uint64, 0, cfg.VWaitBatchCount)
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

func responseToMidCo(request *proto.Request, resp *proto.Response, cos []*MidConn, err error) {
	if err != nil {
		r := &response{Err: err}
		for _, co := range cos {
			co.resp <- r
		}
		return
	}

	active := make(map[uint64]bool)
	for _, v := range resp.Maxs {
		active[v] = false
	}

	for _, v := range resp.Active {
		active[v] = false
	}

	r := &response{Active: active, Err: nil}
	for idx, co := range cos {
		switch request.Cmds[idx] {
		case proto.Q:
			co.resp <- r
		case proto.C:
			max := resp.Maxs[0]
			resp.Maxs = resp.Maxs[1:]
			co.resp <- &response{Max: max, Err: nil}
		case proto.D:
			co.resp <- r
		case proto.C_Q:
			max := resp.Maxs[0]
			resp.Maxs = resp.Maxs[1:]
			co.resp <- &response{Max: max, Active: active, Err: nil}
		}
	}
}

func batchSend(request *proto.Request, cos []*MidConn) {
	var err error
	var cli *rpc.Client
	resp := proto.Response{}

	log.Debugf("%d requests merge to send", len(request.Cmds))

	cli, err = pool.Get()
	if err != nil {
		log.Debugf("get conn failed: %v, send to midconn", err)
		responseToMidCo(nil, nil, cos, err)
		return
	}

	err = cli.Call("VSeq.PushReq", request, &resp)
	if err != nil {
		log.Debugf("call version failed: %v", err)
		responseToMidCo(nil, nil, cos, err)
		pool.Put(cli, err)
		return
	}

	cos[0].stat.VWaitRespT.add(int64(time.Since(request.Ts)))
	cos[0].stat.BatchReqCount.add(int64(len(request.Cmds)))

	responseToMidCo(request, &resp, cos, nil)
	pool.Put(cli, nil)
}

func send(flag uint8) {
	bContinue := true
	cmds = cmds[:0]
	cos = cos[:0]
	toDels = toDels[:0]

	for bContinue {
		var req *req
		select {
		case req = <-maxQueue:
			cmds = append(cmds, req.cmd)
			cos = append(cos, req.co)
			if req.cmd == proto.D {
				toDels = append(toDels, req.co.NextVersion)
			}
		default:
			bContinue = false
		}
	}

	if len(cmds) == 0 {
		return
	}

	reqs := &proto.Request{
		Cmds:   make([]uint8, len(cmds)),
		ToDels: make([]uint64, len(toDels)),
		Ts:     time.Now(),
	}
	midCos := make([]*MidConn, len(cos))

	copy(reqs.Cmds, cmds)
	copy(reqs.ToDels, toDels)
	copy(midCos, cos)

	cos[0].stat.BlockRequestCount.add(int64(len(maxQueue)))
	if flag == SendByDemonTicker {
		cos[0].stat.TickerReqCount.add(1)
	} else {
		cos[0].stat.FullReqCount.add(1)
	}

	go batchSend(reqs, midCos)
}

func sendall(flag uint8) {
	if ql.Lock() {
		send(flag)
		ql.UnLock()
	}
}

func RequestSender(stat *Stat) {
	for {
		<-ticker
		ts := time.Now()
		sendall(SendByDemonTicker)
		cost := time.Since(ts)
		if cost > time.Microsecond*100 {
			stat.SendT.add(int64(cost))
		}
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
