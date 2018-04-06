/**
 *  author: lim
 *  data  : 18-3-24 下午3:18
 */

package server

import (
	"fmt"
	"net"

	"github.com/lemonwx/xsql/middleware/midconn"
	"github.com/lemonwx/log"
)

type Server struct {
	lis  net.Listener
	addr string
}

func NewServer(addr string) (*Server, error) {
	s := new(Server)
	s.addr = addr
	lis, err := net.Listen("tcp", addr)
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
		log.Debugf("[%s] connected, midConn [%d] serve for it",
			midConn.RemoteAddr, midConn.COnnectionId)
		midConn.Serve()
		midConn.Close()
		log.Errorf("conn [%s] colesed, midconn [%d]'s goroutine will exit",
			conn.RemoteAddr(), midConn.COnnectionId)
	}
}
