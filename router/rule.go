package router

import (
	"fmt"
	"strings"
)

type Rule struct {
	DB      string
	Table   string
	As      string
	Key     string `json:"Key"`
	KeyType string

	Type string `json:"Type"`

	Nodes     []string
	Shard     Shard
	ShardList []int
}

func (r *Rule) GetKey() string {
	return r.Key
}

func (r *Rule) GetKeyType() string {
	return r.KeyType
}

func (r *Rule) KeyEqual(col string) bool {
	if r.Key == col {
		return true
	}
	if strings.Contains(col, ".") {
		if fmt.Sprintf("%s.%s", r.Table, r.Key) == col {
			return true
		}

		if fmt.Sprintf("%s.%s", r.As, r.Key) == col {
			return true
		}
	}
	return false
}

func (r *Rule) GetRule() *Rule {
	return r
}

func (r *Rule) GetShard() *Shard {
	return &r.Shard
}

func (r *Rule) GetTB() string {
	return r.Table
}

func (r *Rule) SetAs(as string) {
	r.As = as
}

func (r *Rule) GetShardType() string {
	return r.Type
}

func (r *Rule) ISR() {}

func (r *Rule) Equal(r1 *Rule) bool {
	if r.KeyType == r1.GetKeyType() &&
		ShardEqual(r.GetShard(), r1.GetShard()) {
		return true
	}
	return false
}

func (r *Rule) FindNode(key interface{}) string {
	i := r.Shard.FindForKey(key)
	return r.Nodes[i]
}

func (r *Rule) FindNodeIndex(key interface{}) int {
	return r.Shard.FindForKey(key)
}

func (r *Rule) String() string {
	return fmt.Sprintf("%s.%s?key=%v&type=%s&nodes=%s",
		r.DB, r.Table, r.Key, r.Type, strings.Join(r.Nodes, ", "))
}

func NewDefaultRule(db string, node string) *Rule {
	var r *Rule = &Rule{
		DB:    db,
		Type:  DefaultRuleType,
		Nodes: []string{node},
		Shard: new(DefaultShard),
	}
	return r
}
