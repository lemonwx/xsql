package main

import (
	"fmt"
	"os"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/config"
	"github.com/lemonwx/xsql/middleware/version"
	"github.com/lemonwx/xsql/server"
)

var cfg *config.Conf

func setupConfig() {
	cfg = config.ReadConf()
	/*
	for idx, node := range cfg.Nodes {
		meta.NodeAddrs = append(meta.NodeAddrs, node)
		meta.FullNodeIdxs = append(meta.FullNodeIdxs, idx)
	}
	*/
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
	version.NewRpcPool(cfg.InitSize, cfg.MaxSize, cfg.VerSeqAddr)

	s, err := server.NewServer(cfg)
	if err != nil {
		log.Fatalf("new server failed: %v", err)
	}

	log.Infof("server run under %v", s)
	s.Run()
}
