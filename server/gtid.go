/**
 *  author: lim
 *  data  : 18-8-30 下午8:55
 */

package server

import (
	"common"
	"net"
	"time"

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
}

var (
	maxT     = time.Duration(time.Second * 3)
	maxQueue = make(chan *req, 1024)
	ticker   = time.Tick(maxT)
	pool     *common.Pool
)

func InitGtidPool(addr string) {
	log.Debug(addr)
	var err error
	pool, err = common.NewPool(10, 10, 10,
		func() (common.Conn, error) {
			return net.Dial("tcp", addr)
		},
	)
	if err != nil {
		panic(err)
	}
}

func encode(cmds []uint8, gtids []uint64) []byte {
	log.Debug(gtids)
	size := len(cmds)
	data := make([]uint8, 0, size+4+size*8)
	data = append(data, byte(size), byte(size>>8), byte(size>>16), byte(size>>24))
	data = append(data, cmds...)

	for _, gtid := range gtids {
		data = append(data,
			byte(gtid), byte(gtid>>8), byte(gtid>>16), byte(gtid>>24),
			byte(gtid>>32), byte(gtid>>40), byte(gtid>>48), byte(gtid>>56),
		)
	}

	pktLen := len(data) + 4
	pkt := make([]uint8, 0, pktLen)
	pkt = append(pkt, byte(pktLen), byte(pktLen>>8), byte(pktLen>>16), byte(pktLen>>24))
	pkt = append(pkt, data...)
	return pkt
}

func ret(cmds []byte) map[int]*response {
	resps := make(map[int]*response)
	size := len(cmds)
	hasQ := false
	for idx, cmd := range cmds {
		switch cmd {
		case proto.C_Q:
			resps[idx] = &response{Max: 1001}
			hasQ = true
		case proto.C:
			resps[idx] = &response{Max: 1001}
		case proto.D:
			resps[idx] = &response{}
		case proto.Q:
			hasQ = true
		}
	}

	if hasQ {
		resps[size] = &response{Active: map[uint64]bool{1001: false, 1002: false}}
	}
	return resps
}

func sendall(exReq *req) {
	size := len(maxQueue) + 1

	cmds := make([]byte, 0, size)
	cos := make([]*MidConn, 0, size)
	gtidTodel := make([]uint64, 0, size)

	log.Debugf("%d request merge to send", size)
	for i := 0; i < size-1; i++ {
		req := <-maxQueue
		cmds = append(cmds, req.cmd)
		cos = append(cos, req.co)
		if req.cmd == proto.D {
			gtidTodel = append(gtidTodel, req.co.NextVersion)
		}
	}
	log.Debug(cmds, gtidTodel)

	cmds = append(cmds, exReq.cmd)
	cos = append(cos, exReq.co)
	if exReq.cmd == proto.D {
		gtidTodel = append(gtidTodel, exReq.co.NextVersion)
	}

	pkt := encode(cmds, gtidTodel)
	log.Debug(cmds, pkt)
	resps := ret(cmds)

	log.Debugf("response: ")
	for idx, co := range cos {
		if resp, ok := resps[idx]; ok {
			if cmds[idx] == proto.C_Q {
				resp.Active = resps[size].Active
			}
			log.Debugf("%d, %d, %v, %v", idx, co.ConnectionId, cmds[idx], resp)
			co.resp <- resp
		} else {
			log.Debugf("%d, %d, %v, %v", idx, co.ConnectionId, cmds[idx], resps[size])
			log.Debug(co.resp)
			co.resp <- resps[size]
		}
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
	log.Debug("push", len(maxQueue))
	defer func() {
		log.Debug("push", len(maxQueue))
	}()

	/*
		select {
		case <-ticker:
			// send all and cur
			log.Debug("send all by ticker")
			sendall(&req{cmd, co})
			return
		default:
		}*/

	select {
	case maxQueue <- &req{cmd, co}:
	default:
		// send all and cur
		log.Debug("send all by full queue")
		sendall(&req{cmd, co})
		return
	}
}
