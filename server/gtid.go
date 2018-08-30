/**
 *  author: lim
 *  data  : 18-8-30 下午8:55
 */

package server

import (
	"time"

	"common"

	"net"

	"github.com/lemonwx/log"
)

const (
	C uint8 = iota
	D
	Q
	Q_C
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

func encode(cmds []uint8) []byte {
	pktLen := len(cmds)
	pkt := make([]uint8, 0, pktLen+4)
	pkt = append(pkt, byte(pktLen), byte(pktLen>>8), byte(pktLen>>16), byte(pktLen>>24))
	pkt = append(pkt, cmds...)
	return pkt
}

func ret(cmds []byte) map[int]*response {
	resps := make(map[int]*response)
	size := len(cmds)
	hasQ := false
	for idx, cmd := range cmds {
		switch cmd {
		case Q_C:
			resps[idx] = &response{Max: 1001}
			hasQ = true
		case C:
			resps[idx] = &response{Max: 1001}
		case D:
			resps[idx] = &response{}
		case Q:
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
	log.Debugf("%d request merge to send", size)
	for i := 0; i < size-1; i++ {
		req := <-maxQueue
		cmds = append(cmds, req.cmd)
		cos = append(cos, req.co)
	}
	cmds = append(cmds, exReq.cmd)
	cos = append(cos, exReq.co)

	pkt := encode(cmds)
	log.Debug(cmds, pkt)
	resps := ret(cmds)

	log.Debugf("response: ")
	for idx, co := range cos {
		if resp, ok := resps[idx]; ok {
			if cmds[idx] == Q_C {
				resp.Active = resps[size].Active
			}
			//log.Debugf("%d, %d, %v, %v", idx, co.ConnectionId, cmds[idx], resp)
			co.resp <- resp
		} else {
			co.resp <- resps[size]
			//log.Debugf("%d, %d, %v, %v", idx, co.ConnectionId, cmds[idx], resps[mergeSize])
		}
	}

	conn, err := pool.Get()
	if err != nil {
		panic(err)
	}
	conn.Write(pkt)
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
	select {
	case <-ticker:
		// send all and cur
		log.Debug("send all by ticker")
		sendall(&req{cmd, co})
		return
	default:
	}

	select {
	case maxQueue <- &req{cmd, co}:
	default:
		// send all and cur
		log.Debug("send all by full queue")
		sendall(&req{cmd, co})
		return
	}
}
