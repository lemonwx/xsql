package meta

import (
	"fmt"
	"github.com/lemonwx/xsql/config"
	"github.com/lemonwx/xsql/middleware/router"
	"github.com/lemonwx/xsql/node"
)

var m *Meta

func SetMetas(meta *Meta) {
	m = meta
}

func GetRouter(db string) (*router.Router, error) {
	r, ok := m.Routers[db]
	if !ok {
		return nil, fmt.Errorf("can't find router for this db: %v", db)
	}
	return r, nil
}

type Meta struct {
	NodeAddrs    []*config.Node
	FullNodeIdxs []int
	nodes        map[string]*node.Node
	Routers      map[string]*router.Router
}

func GetNodeAddrs() []*config.Node {
	return m.NodeAddrs
}

func GetFullNodeIdxs() []int {
	return m.FullNodeIdxs
}
