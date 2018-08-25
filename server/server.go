/**
 *  author: lim
 *  data  : 18-3-24 下午3:18
 */

package server

import (
	"fmt"
	"net"

	"time"

	"encoding/json"
	"io/ioutil"

	"sync"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/config"
	"github.com/lemonwx/xsql/errors"
	"github.com/lemonwx/xsql/meta"
	"github.com/lemonwx/xsql/node"
	"github.com/lemonwx/xsql/router"
)

type conns struct {
	midConns map[string]*MidConn // key is client's addr
	sync.RWMutex
}

type Server struct {
	lis   net.Listener
	addr  string
	cfg   *config.Conf
	pools map[int]*node.Pool
	cos   *conns
}

func NewServer(cfg *config.Conf) (*Server, error) {
	var err error

	s := new(Server)
	s.cfg = cfg
	s.addr = cfg.Addr
	s.cos = &conns{
		midConns: map[string]*MidConn{},
	}

	if err = s.parseSchemas(cfg); err != nil {
		return nil, err
	}

	if err = s.newBackendPool(cfg); err != nil {
		return nil, err
	}

	if s.lis, err = net.Listen("tcp", s.addr); err != nil {
		return nil, err
	}

	//go s.dumpPoolsInfo()

	return s, nil
}

func (s *Server) Run() error {

	for {
		conn, err := s.lis.Accept()
		if err != nil {
			fmt.Println("xsql server accept failed")
		}

		//go s.ServeConn(conn)

		midConn, err := NewMidConn(conn, s.cfg, s.pools, s)
		if err != nil {
			log.Errorf("new mid conn failed: %v", err)
		} else {
			go func() {
				midConn.Serve()
				midConn.Close()
				s.cos.Lock()
				delete(s.cos.midConns, midConn.RemoteAddr)
				s.cos.Unlock()

			}()
			s.cos.midConns[midConn.RemoteAddr] = midConn
		}
	}
	return nil
}

// serve for mysql client conn(get by lis.Accept)
func (s *Server) ServeConn(conn net.Conn) {
	// init and connect with back mysql server
	if midConn, err := NewMidConn(conn, s.cfg, s.pools, s); err != nil {
		log.Errorf("new mid conn failed: %v", err)
		return
	} else {
		log.Debugf("[%d] [%s] connected, midConn [%d] serve for it",
			midConn.ConnectionId, midConn.RemoteAddr, midConn.ConnectionId)
		midConn.Serve()
		midConn.Close()
		log.Errorf("[%d] conn [%s] colesed, midconn [%d]'s goroutine will exit",
			midConn.ConnectionId, conn.RemoteAddr(), midConn.ConnectionId)
	}
}

func (s *Server) parseSchemas(cfg *config.Conf) error {
	data, err := ioutil.ReadFile(cfg.Meta)
	if err != nil {
		return err
	}

	m := &meta.Meta{}
	if err = json.Unmarshal(data, &m.Routers); err != nil {
		return err
	}

	for db, r := range m.Routers {
		log.Debug(r.Rules)
		for tb, rule := range r.Rules {
			rule.DB = db
			rule.Table = tb

			switch rule.Type {
			case "hash":
				rule.Shard = &router.HashShard{len(rule.Nodes)}
			default:
				return errors.New2("unsupported shard type: " + rule.Type)
			}

			r.Rules[tb] = rule
		}

		m.Routers[db] = r
	}

	m.FullNodeIdxs = make([]int, len(cfg.Nodes))
	for idx, _ := range cfg.Nodes {
		m.FullNodeIdxs[idx] = idx
	}
	m.FullNodeIdxs = []int{0, 1}
	meta.SetMetas(m)
	return nil
}

func (s *Server) newBackendPool(cfg *config.Conf) error {
	if len(cfg.Nodes) == 0 {
		return fmt.Errorf("length of backend nodes can not be 0")
	}

	s.pools = make(map[int]*node.Pool)

	log.Debug(cfg.BackInitSize, cfg.BackMaxIdleSize, cfg.BackMaxSize)
	for idx, nodeCfg := range cfg.Nodes {
		pool, err := node.NewNodePool(cfg.BackInitSize, cfg.BackMaxIdleSize, cfg.BackMaxSize, nodeCfg)
		if err != nil {
			return err
		} else {
			s.pools[idx] = pool
		}
	}

	return nil
}

func (s *Server) dumpPoolsInfo() {
	for {
		for _, p := range s.pools {
			p.DumpInfo()
		}
		time.Sleep(time.Second * 5)
	}
}
