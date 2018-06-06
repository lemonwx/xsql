package router

import (
	"fmt"
	"strings"
)

type R interface {
	ISR()
	GetKey() string
	GetKeyType() string
	GetShardType() string
	GetTB() string

	// 判断两个 rule 是否等价
	Equal(r R) bool
	GetRule() *Rule
	KeyEqual(col string) bool
	SetAs(as string)
	GetShard() *Shard
}

type Rule struct {
	DB      string
	Table   string
	As      string
	Key     string
	KeyType string

	Type string

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

func (r *JoinRule) KeyEqual(col string) bool {
	if r.GetKey() == col {
		return true
	}
	if strings.Contains(col, ".") {
		if fmt.Sprintf("%s.%s", r.GetAs(), r.GetKey()) == col {
			return true
		}
	}
	return false
}

func (r *JoinRule) GetShard() *Shard {
	return r.Lr.GetShard()
}

func (r *JoinRule) GetTB() string {
	return r.Lr.GetTB()
}

func (r *JoinRule) GetAs() string {
	return r.As
}

func (r *JoinRule) SetAs(as string) {
	r.As = as
}

func (r *JoinRule) GetRule() *Rule {
	return r.Lr.GetRule()
}

func (r *JoinRule) GetKey() string {
	return r.Lr.GetKey()
}

func (r *JoinRule) GetKeyType() string {
	return r.Lr.GetKeyType()
}

func (r *JoinRule) GetShardType() string {
	return r.Lr.GetShardType()
}

func (r *Rule) ISR() {}

type JoinRule struct {
	Lr, Rr    R
	As        string
	ShardList []int
}

func (r *JoinRule) ISR() {}

func (r *JoinRule) Equal(r1 R) bool {
	if r.GetKeyType() == r.GetKeyType() && // 分发键的类型一致
		ShardEqual(r.GetShard(), r1.GetShard()) { // 分发类型一致
		return true
	}
	return false
}

func (r *Rule) Equal(r1 R) bool {
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
	return fmt.Sprintf("%s.%s?key=%v&shard=%s&nodes=%s",
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

func (r *Router) GetRule(table string) *Rule {
	idx := strings.Index(table, ".")
	if idx != -1 {
		table = table[idx+1:]
	}
	rule, ok := r.Rules[table]
	if !ok {
		panic(fmt.Errorf("can't find Rule for this tb: %v", table))
	}

	if rule == nil {
		return r.DefaultRule
	} else {
		return rule
	}
}

type Router struct {
	DB          string
	Rules       map[string]*Rule //key is <table name>
	DefaultRule *Rule
	nodes       []string //just for human saw
}

/*
func NewRouter(schemaConfig *config.SchemaConfig) (*Router, error) {

	if !includeNode(schemaConfig.Nodes, schemaConfig.RulesConifg.Default) {
		return nil, fmt.Errorf("default node[%s] not in the nodes list.",
			schemaConfig.RulesConifg.Default)
	}

	rt := new(Router)
	rt.DB = schemaConfig.DB
	rt.nodes = schemaConfig.Nodes
	rt.Rules = make(map[string]*Rule, len(schemaConfig.RulesConifg.ShardRule))
	rt.DefaultRule = NewDefaultRule(rt.DB, schemaConfig.RulesConifg.Default)

	for _, shard := range schemaConfig.RulesConifg.ShardRule {
		rc := &RuleConfig{shard}
		for _, node := range shard.Nodes {
			if !includeNode(rt.nodes, node) {
				return nil, fmt.Errorf("shard table[%s] node[%s] not in the schema.nodes list:[%s].",
					shard.Table, node, strings.Join(shard.Nodes, ","))
			}
		}
		rule, err := rc.ParseRule(rt.DB)
		if err != nil {
			return nil, err
		}

		if rule.Type == DefaultRuleType {
			return nil, fmt.Errorf("[default-rule] duplicate, must only one.")
		} else {
			if _, ok := rt.Rules[rule.Table]; ok {
				return nil, fmt.Errorf("table %s rule in %s duplicate", rule.Table, rule.DB)
			}
			rt.Rules[rule.Table] = rule
		}
	}
	return rt, nil
}

*/
func includeNode(nodes []string, node string) bool {
	for _, n := range nodes {
		if n == node {
			return true
		}
	}
	return false
}
