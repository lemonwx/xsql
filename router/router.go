/**
 *  author: lim
 *  data  : 18-8-22 下午8:48
 */

package router

import (
	"fmt"
	"strings"
)

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
	RuleMeta    map[string]map[string]*Rule //key is schema.table
	nodes       []string                    //just for human saw
	DefaultRule *Rule

	DB          string
	defRuleMeta map[string]*Rule //key is table, if not specify key, will use this to shard
	Rules       map[string]*Rule //key is table
}
