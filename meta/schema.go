package meta

import (
	"fmt"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/config"
	"github.com/lemonwx/xsql/router"
)

type Meta struct {
	NodeAddrs    []*config.Node
	FullNodeIdxs []int
	Routers      map[string]*router.Router
}

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

func GetNodeAddrs() []*config.Node {
	return m.NodeAddrs
}

func GetFullNodeIdxs() []int {
	return m.FullNodeIdxs
}

func AddDatabase(db string) {
	log.Debug(m.Routers[db])
	m.Routers[db] = &router.Router{
		DB: db,
	}

	log.Debug(m.Routers[db])
}

func AddTable(db, tb string) error {
	log.Debug(db)
	log.Debug(tb)
	r, err := GetRouter(db)
	if err != nil {
		return err
	}

	r.Rules[tb] = &router.Rule{
		DB:    db,
		Table: tb,
	}
	return nil
}
