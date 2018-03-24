package main

import (
	"fmt"

	"github.com/lemonwx/xsql/server"

	"os"
)

const addr string = "192.168.1.5:1234"

func main() {
	fmt.Println("vim-go")
	s, err := server.NewServer(addr)
	if err != nil {
		os.Exit(1)
	}

	s.Run()
}
