/**
 *  author: lim
 *  data  : 18-4-6 下午1:11
 */

package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-yaml/yaml"
)

type Node struct {
	Host     string
	Port     int
	User     string
	Password string
}

type RedisConf struct {
	Host     string
	Port     int
	User     string
	Password string
}

type Auth struct {
	User, Password string
}

type Conf struct {
	Id       string
	Addr     string
	Nodes    []*Node
	LogLevel int       `yaml:"loglevel"`
	RedisCfg RedisConf `yaml:"RedisCfg"`
	Xa       bool      `yaml:"XA"`

	BackInitSize    uint32 `yaml:"BackInitSize"`
	BackMaxIdleSize uint32 `yaml:"BackMaxIdleSize"`
	BackMaxSize     uint32 `yaml:"BackMaxSize"`

	VerSeqAddr string  `yaml:"VerSeqAddr"`
	MaxSize    int     `yaml:"MaxSize"`
	InitSize   int     `yaml:"InitSize"`
	Batch      int     `yaml:"Batch"`
	Auths      []*Auth `yaml:"auth"`
}

func ReadConf() *Conf {

	c := &Conf{}
	yamlF, err := ioutil.ReadFile("/home/lim/space/src/github.com/lemonwx/xsql/etc/c.yaml")
	if err != nil {
		fmt.Println("read ../etc/c.yaml failed: %v", err)
		os.Exit(-1)
	}

	err = yaml.Unmarshal(yamlF, c)
	if err != nil {
		fmt.Println("Unmarshal failed: %v", err)
		os.Exit(-1)
	}
	return c
}
