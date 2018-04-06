/**
 *  author: lim
 *  data  : 18-4-6 下午1:11
 */

package config

import (
	"io/ioutil"
	"fmt"
	"os"

	"github.com/go-yaml/yaml"
)

type Node struct {
	Host string
	Port int
	User string
	Password string
}

type RedisConf struct {
	Host string
	Port int
	User string
	Password string
}

type Conf struct {
	Id string
	Addr string
	Nodes []*Node
	LogLevel int `yaml:"loglevel"`
	RedisCfg RedisConf `yaml:"RedisCfg"`
	Xa bool `yaml:"XA"`
}

func ReadConf() (*Conf) {

	c := &Conf{}
	yamlF, err := ioutil.ReadFile("../etc/c.yaml")
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