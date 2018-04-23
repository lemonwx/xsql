/**
 *  author: lim
 *  data  : 18-3-24 下午3:18
 */

package server

import (
	"fmt"
	"net"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/middleware/midconn"
	"github.com/lemonwx/xsql/middleware/meta"
	"github.com/lemonwx/xsql/config"
	"github.com/lemonwx/xsql/middleware/router"
)

type Server struct {
	lis  net.Listener
	addr string
}

func NewServer(cfg *config.Conf) (*Server, error) {
	s := new(Server)
	s.addr = cfg.Addr
	s.parseSchemas(cfg)

	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return nil, err
	}
	s.lis = lis
	return s, nil
}

func (s *Server) Run() error {

	for {
		conn, err := s.lis.Accept()
		if err != nil {
			fmt.Println("xsql server accept failed")
		}

		go s.ServeConn(conn)
	}
	return nil
}

// serve for mysql client conn(get by lis.Accept)
func (s *Server) ServeConn(conn net.Conn) {
	// init and connect with back mysql server
	if midConn, err := midconn.NewMidConn(conn); err != nil {
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

	fullNodeIdx := make([]int, 0, len(cfg.Nodes))
	nodeAddrs := make([]*config.Node, 0, len(cfg.Nodes))
	for idx, node := range cfg.Nodes {
		fullNodeIdx = append(fullNodeIdx, idx)
		nodeAddrs = append(nodeAddrs, node)
	}

	rs := make(map[string]*router.Router)
	rs["db"] = &router.Router{
		DB         :"db",
		Rules: map[string]*router.Rule{
			"tb":&router.Rule{
				DB:    "db",
				Table: "tb",
				Key:   "id",
				Type:  "hash",
				Nodes: []string{"1", "2"},
				Shard:  &router.HashShard{2},
			},
			"tt":&router.Rule{
				DB:    "db",
				Table: "tt",
				Key:   "id",
				Type:  "hash",
				Nodes: []string{"1", "2"},
				Shard:  &router.HashShard{2},
			},

		},
		DefaultRule :router.NewDefaultRule("db", ""),
	}

	meta.SetMetas(&meta.Meta{
		NodeAddrs: nodeAddrs,
		FullNodeIdxs:fullNodeIdx,
		Routers:rs,
	})

	/*

	for _, schemaCfg := range s.cfg.Schemas {
		if _, ok := s.schemas[schemaCfg.DB]; ok {
			return fmt.Errorf("duplicate schema [%s].", schemaCfg.DB)
		}
		if len(schemaCfg.Nodes) == 0 {
			return fmt.Errorf("schema [%s] must have a node.", schemaCfg.DB)
		}

		nodes := make(map[string]*Node)
		for _, n := range schemaCfg.Nodes {
			if s.getNode(n) == nil {
				return fmt.Errorf("schema [%s] node [%s] config is not exists.", schemaCfg.DB, n)
			}

			if _, ok := nodes[n]; ok {
				return fmt.Errorf("schema [%s] node [%s] duplicate.", schemaCfg.DB, n)
			}

			nodes[n] = s.getNode(n)
		}

		rule, err := router.NewRouter(&schemaCfg)
		if err != nil {
			return err
		}

		s.schemas[schemaCfg.DB] = &Schema{
			db:    schemaCfg.DB,
			nodes: nodes,
			rule:  rule,
		}
	}
	*/

	return nil
}