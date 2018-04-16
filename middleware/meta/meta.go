/**
 *  author: lim
 *  data  : 18-4-6 下午2:29
 */

package meta

import (
	"github.com/lemonwx/xsql/config"
)

var (
	NodeAddrs = []*config.Node{}
)

var FullNodeIdxs = []int{}

type Meta struct {
	Db      string
	Tb      string
	DisKey  string
	DisType string
	DisNode []int
}

func NewMeta(sql string) *Meta {
	return &Meta{
	}
}

const (
	HASH  string = "hash"
	RANGE string = "range"
	LIST  string = "list"
	DUP   string = "dup"
)

func Hash(val int64) int64 {
	return val % int64(len(FullNodeIdxs))
}
