package main

import (
	"os"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/server"
	"github.com/lemonwx/xsql/config"
	"github.com/lemonwx/xsql/middleware/xa"
	"github.com/lemonwx/xsql/middleware/meta"
	"fmt"
)

var cfg *config.Conf

func setupConfig() {
	cfg = config.ReadConf()
	for idx, node := range cfg.Nodes {
		meta.NodeAddrs = append(meta.NodeAddrs, node)
		meta.FullNodeIdxs = append(meta.FullNodeIdxs, idx)
	}
}

func setupLogger() {
	f, err := os.Create("xsql.log")
	if err != nil {
		fmt.Println("touch log file xsql.log failed: %v", err)
	}
	log.NewDefaultLogger(f)
	log.SetLevel(cfg.LogLevel)
	log.Debug("this is xsql's log")
}


func main() {
	setupConfig()
	fmt.Println("init cfg done.")
	setupLogger()
	fmt.Println("init log done.")
	xa.InitPool()

	s, err := server.NewServer(cfg.Addr)
	if err != nil {
		log.Fatalf("new server failed: %v", err)
	}

	log.Infof("server run under %v", s)
	s.Run()
}
