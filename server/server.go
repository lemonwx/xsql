/**
 *  author: lim
 *  data  : 18-3-24 下午3:18
 */

package  server

import (
	"net"
	"fmt"

	"github.com/lemonwx/xsql/middleware/midconn"
)

type Server struct {
	lis net.Listener
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

func (s *Server) Run()  error {

	for  {
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
	midConn := midconn.NewMidConn(conn)
	midConn.Serve()
	fmt.Println("this goroutine will exit", midConn)
}