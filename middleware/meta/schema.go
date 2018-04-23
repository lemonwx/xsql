package meta

import (
	"github.com/lemonwx/xsql/middleware/router"
	"github.com/lemonwx/xsql/node"
	"github.com/lemonwx/xsql/config"
)

var m *Meta

func SetMetas(meta *Meta) {
	m = meta
}

func GetRouter(db string) *router.Router {
	return m.Routers[db]
}

type Meta struct {
	NodeAddrs []*config.Node
	FullNodeIdxs []int
	nodes map[string]*node.Node
	Routers map[string]*router.Router
}

func GetNodeAddrs() []*config.Node {
	return m.NodeAddrs
}

func GetFullNodeIdxs() []int {
	return m.FullNodeIdxs
}