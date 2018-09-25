/**
 *  author: lim
 *  data  : 18-5-22 下午10:40
 */

package router

import (
	"fmt"
	"hack"
	"sort"
	"strconv"

	"utils"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/errors"
	"github.com/lemonwx/xsql/sqlparser"
)

var (
	String                = sqlparser.String
	UNSUPPORTED_SHARD_ERR = fmt.Errorf("unsupported shard for this sql")
)

const (
	EID_NODE = iota
	VALUE_NODE
	LIST_NODE
	OTHER_NODE
)

type Plan struct {
	rule      *Rule
	fullList  []int
	ShardList []int

	//for insert
	shardKeyIdx int
	stmtArgs    map[int]interface{}
}

func SliceEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	sort.Ints(s1)
	sort.Ints(s2)

	for idx, item := range s1 {
		if item != s2[idx] {
			return false
		}
	}

	return true
}

func SliceIn(s1, s2 []int) bool {
	if s2 == nil {
		return false
	}

	if len(s2) == 0 {
		return false
	}

	mp := make(map[int]uint8)

	for _, item := range s2 {
		mp[item] = 0
	}

	for _, item := range s1 {
		if _, ok := mp[item]; !ok {
			return false
		}
	}

	return true
}

func makeList(start, end int) []int {
	list := make([]int, end-start)
	for i := start; i < end; i++ {
		list[i-start] = i
	}
	return list
}

// l1 & l2
func interList(l1 []int, l2 []int) []int {
	if len(l1) == 0 || len(l2) == 0 {
		return []int{}
	}

	l3 := make([]int, 0, len(l1)+len(l2))
	var i = 0
	var j = 0
	for i < len(l1) && j < len(l2) {
		if l1[i] == l2[j] {
			l3 = append(l3, l1[i])
			i++
			j++
		} else if l1[i] < l2[j] {
			i++
		} else {
			j++
		}
	}

	return l3
}

// l1 | l2
func unionList(l1 []int, l2 []int) []int {
	if len(l1) == 0 {
		return l2
	} else if len(l2) == 0 {
		return l1
	}

	l3 := make([]int, 0, len(l1)+len(l2))

	var i = 0
	var j = 0
	for i < len(l1) && j < len(l2) {
		if l1[i] < l2[j] {
			l3 = append(l3, l1[i])
			i++
		} else if l1[i] > l2[j] {
			l3 = append(l3, l2[j])
			j++
		} else {
			l3 = append(l3, l1[i])
			i++
			j++
		}
	}

	if i != len(l1) {
		l3 = append(l3, l1[i:]...)
	} else if j != len(l2) {
		l3 = append(l3, l2[j:]...)
	}

	return l3
}

// l1 - l2
func differentList(l1 []int, l2 []int) []int {
	if len(l1) == 0 {
		return []int{}
	} else if len(l2) == 0 {
		return l1
	}

	l3 := make([]int, 0, len(l1))

	var i = 0
	var j = 0
	for i < len(l1) && j < len(l2) {
		if l1[i] < l2[j] {
			l3 = append(l3, l1[i])
			i++
		} else if l1[i] > l2[j] {
			j++
		} else {
			i++
			j++
		}
	}

	if i != len(l1) {
		l3 = append(l3, l1[i:]...)
	}

	return l3
}

func (p *Plan) ShardForFrom(r *Router, preWhere *sqlparser.Where, froms ...sqlparser.TableExpr) (*Plan, error) {

	if len(froms) == 1 {
		switch v := froms[0].(type) {
		case *sqlparser.AliasedTableExpr:
			switch node := v.Expr.(type) {
			case *sqlparser.TableName:
				rule := r.GetRule(hack.String(node.Name))
				rule.As = hack.String(v.As)
				p.rule = rule
				p.fullList = makeList(0, len(rule.Nodes))

				if preWhere == nil {
					p.ShardList = p.fullList
				} else {
					p.ShardList = p.routingAnalyzeBoolean(preWhere.Expr)
				}
			case *sqlparser.Subquery:
				switch sel := node.Select.(type) {
				case *sqlparser.SimpleSelect:
					panic(UNSUPPORTED_SHARD_ERR)
				case *sqlparser.Select:
					plan, err := GeneralPlanForSelect(r, sel, p.stmtArgs)
					if err != nil {
						panic(err)
					}

					plan.rule.As = hack.String(v.As)
					p.rule = plan.rule
					p.ShardList = plan.ShardList
					p.fullList = plan.fullList
				case *sqlparser.Union:
					panic(UNSUPPORTED_SHARD_ERR)
				}
			}
		case *sqlparser.JoinTableExpr:
			pl := &Plan{}
			pr := &Plan{}

			pl.ShardForFrom(r, nil, v.LeftExpr)
			pr.ShardForFrom(r, nil, v.RightExpr)

			lr := pl.rule
			rr := pr.rule

			if !lr.Equal(rr) {
				panic(errors.New(UNSUPPORTED_SHARD_ERR))
			}

			log.Debugf("lr equal rr: %v == %v", lr, rr)

			if b, ok := v.On.(*sqlparser.ComparisonExpr); ok {
				if b.Operator != "=" {
					panic(UNSUPPORTED_SHARD_ERR)
				}

				l := p.AnalyzeValue(b.Left)
				r := p.AnalyzeValue(b.Right)

				if l == r && l == EID_NODE {
					// only support left and rigth all col
					// if l.colname == key and

					if lr.KeyEqual(String(b.Left)) && rr.KeyEqual(String(b.Right)) ||
						lr.KeyEqual(String(b.Right)) && rr.KeyEqual(String(b.Left)) {

						if len(pl.ShardList) == len(pr.ShardList) {
							if SliceEqual(pl.ShardList, pr.ShardList) {
								p.ShardList = pr.ShardList
							} else {
								panic(UNSUPPORTED_SHARD_ERR)
							}
						} else if len(pl.ShardList) < len(pr.ShardList) {
							p.ShardList = pl.ShardList
						} else {
							p.ShardList = pr.ShardList
						}
						p.rule = pl.rule
						p.fullList = pl.fullList
					} else {
						panic(UNSUPPORTED_SHARD_ERR)
					}
				} else {
					panic(UNSUPPORTED_SHARD_ERR)
				}
			} else {
				log.Debugf("join table expr's on expr: [%v] not compare(\"=\")", v.On)
			}

			if preWhere == nil {
				if p.rule == nil {
					log.Debugf("join table expr's where: [%v], and can't shard by on", preWhere)
					panic(UNSUPPORTED_SHARD_ERR)
				} else {
					log.Debugf("select from join, and can by on")
				}
			} else {
				if p.rule == nil {
					compare, ok := preWhere.Expr.(*sqlparser.ComparisonExpr)
					if ok && compare.Operator == "=" {
						l := p.AnalyzeValue(compare.Left)
						r := p.AnalyzeValue(compare.Right)
						if l == EID_NODE && r == EID_NODE {
							if lr.KeyEqual(String(compare.Left)) && rr.KeyEqual(String(compare.Right)) {
								p.rule = lr
								p.ShardList = makeList(0, len(lr.Nodes))
								p.fullList = p.ShardList
							} else if lr.KeyEqual(String(compare.Right)) && rr.KeyEqual(String(compare.Left)) {
								p.rule = lr
								p.ShardList = makeList(0, len(lr.Nodes))
								p.fullList = p.ShardList
							} else {
								log.Debug("join table expr's where: [%v] is compare(\"=\") but not [ keys equal to rule.key ], and can't by on", preWhere)
								panic(UNSUPPORTED_SHARD_ERR)
							}
						} else {
							log.Debugf("join table expr's where: [%s] is compare(\"=\") but not [ l == r == EID_NODE ], and can't shard by on",
								String(preWhere))
							panic(errors.New(UNSUPPORTED_SHARD_ERR))
						}
					} else {
						log.Debug("join table expr's where: [%v] is not compare(\"=\"), and can't by on", preWhere)
						panic(UNSUPPORTED_SHARD_ERR)
					}
				} else {
					log.Debug("join table expr can be shard by on, now use where to reduce ShardList")
					p.ShardList = p.routingAnalyzeBoolean(preWhere.Expr)
				}
			}
		case *sqlparser.ParenTableExpr:
			panic(UNSUPPORTED_SHARD_ERR)
		}
	} else {
		var rules []*Rule
		var rule *Rule
		for _, tb := range froms {
			switch v := tb.(type) {
			case *sqlparser.AliasedTableExpr:
				switch node := v.Expr.(type) {
				case *sqlparser.TableName:
					rule = r.GetRule(hack.String(node.Name))
					rule.As = hack.String(v.As)
				case *sqlparser.Subquery:
					switch sel := node.Select.(type) {
					case *sqlparser.SimpleSelect:
						panic(UNSUPPORTED_SHARD_ERR)
					case *sqlparser.Select:
						plan, err := GeneralPlanForSelect(r, sel, p.stmtArgs)
						if err != nil {
							panic(err)
						}
						plan.rule.As = hack.String(v.As)
						rule = plan.rule
					case *sqlparser.Union:
						panic(UNSUPPORTED_SHARD_ERR)
					}
				default:
					panic(UNSUPPORTED_SHARD_ERR)
				}
			}

			if len(rules) != 0 {
				if !rules[0].Equal(rule) {
					panic(UNSUPPORTED_SHARD_ERR)
				}
			}

			rules = append(rules, rule)
		}

		if preWhere == nil {
			panic(UNSUPPORTED_SHARD_ERR)
		} else {
			switch b := preWhere.Expr.(type) {
			case *sqlparser.AndExpr:
				panic(UNSUPPORTED_SHARD_ERR)
			case *sqlparser.ComparisonExpr:
				if len(froms) != 2 {
					panic(UNSUPPORTED_SHARD_ERR)
				}

				if b.Operator != "=" {
					panic(UNSUPPORTED_SHARD_ERR)
				}

				l := p.AnalyzeValue(b.Left)
				r := p.AnalyzeValue(b.Right)

				if !(l == EID_NODE && r == EID_NODE) {
					panic(UNSUPPORTED_SHARD_ERR)
				}

				if rules[0].KeyEqual(String(b.Left)) && rules[1].KeyEqual(String(b.Right)) {
					p.rule = rules[0]
					p.ShardList = makeList(0, len(rules[0].Nodes))
					p.fullList = p.ShardList
				} else if rules[1].KeyEqual(String(b.Left)) && rules[0].KeyEqual(String(b.Right)) {
					p.rule = rules[0]
					p.ShardList = makeList(0, len(rules[0].Nodes))
					p.fullList = p.ShardList
				} else {
					log.Debug("join table expr's where: [%v] is compare(\"=\") but not [ keys equal to rule.key ], and can't by on", preWhere)
					panic(UNSUPPORTED_SHARD_ERR)
				}

			default:
				panic(UNSUPPORTED_SHARD_ERR)
			}
			log.Debug(preWhere.Type)

			p.ShardList = p.routingAnalyzeBoolean(preWhere.Expr)
		}
	}

	return nil, nil
}

func (plan *Plan) routingAnalyzeBoolean(where sqlparser.BoolExpr) []int {
	switch node := where.(type) {
	case *sqlparser.AndExpr:
		left := plan.routingAnalyzeBoolean(node.Left)
		right := plan.routingAnalyzeBoolean(node.Right)
		return interList(left, right)
	case *sqlparser.OrExpr:
		left := plan.routingAnalyzeBoolean(node.Left)
		right := plan.routingAnalyzeBoolean(node.Right)
		return unionList(left, right)
	case *sqlparser.ParenBoolExpr:
		return plan.routingAnalyzeBoolean(node.Expr)
	case *sqlparser.ComparisonExpr:
		switch {
		case utils.StringIn(node.Operator, "=", "<", ">", "<=", ">=", "<=>"):
			left := plan.routingAnalyzeValue(node.Left)
			right := plan.routingAnalyzeValue(node.Right)
			if (left == EID_NODE && right == VALUE_NODE) || (left == VALUE_NODE && right == EID_NODE) {
				return plan.findConditionShard(node)
			}
		case utils.StringIn(node.Operator, "in", "not in"):
			//judge node.Left is col and dis key
			left := plan.routingAnalyzeValue(node.Left)
			right := plan.routingAnalyzeValue(node.Right)
			if left == EID_NODE && right == LIST_NODE {
				if node.Operator == "in" {
					return plan.findShardList(node.Right)
				} else {
					return plan.fullList
				}
			}
		}
	case *sqlparser.RangeCond:
		left := plan.routingAnalyzeValue(node.Left)
		from := plan.routingAnalyzeValue(node.From)
		to := plan.routingAnalyzeValue(node.To)
		if left == EID_NODE && from == VALUE_NODE && to == VALUE_NODE {
			return plan.findConditionShard(node)
		}
	case *sqlparser.ExistsExpr:
		switch node.Subquery.Select.(type) {
		case *sqlparser.SimpleSelect:
			return plan.fullList
		default:
			panic(fmt.Errorf("unsupported sharding type"))
		}
	}

	return plan.fullList

}

func (plan *Plan) routingAnalyzeValue(valExpr sqlparser.ValExpr) int {
	switch node := valExpr.(type) {
	case *sqlparser.ColName:
		if plan.rule.KeyEqual(String(node)) {
			return EID_NODE
		}
	case sqlparser.ValTuple:
		if plan.shardKeyIdx != 0 {
			if plan.routingAnalyzeValue(node[plan.shardKeyIdx]) != VALUE_NODE {
				return OTHER_NODE
			}
		}
		for _, n := range node {
			if plan.routingAnalyzeValue(n) != VALUE_NODE {
				return OTHER_NODE
			}
		}
		return LIST_NODE
	case sqlparser.StrVal, sqlparser.NumVal, sqlparser.ValArg:
		return VALUE_NODE
	}
	return OTHER_NODE
}

func (plan *Plan) routingAnalyzeValues(vals sqlparser.Values) sqlparser.Values {
	// Analyze first value of every item in the list
	log.Debug(vals)
	for i := 0; i < len(vals); i++ {
		switch tuple := vals[i].(type) {
		case sqlparser.ValTuple:
			result := plan.routingAnalyzeValue(tuple[plan.shardKeyIdx])
			if result != VALUE_NODE {
				panic(NewParserError("insert is too complex"))
			}
		default:
			panic(NewParserError("insert is too complex"))
		}
	}
	return vals
}

func (plan *Plan) findConditionShard(expr sqlparser.BoolExpr) (shardList []int) {
	var index int
	switch criteria := expr.(type) {
	case *sqlparser.ComparisonExpr:
		switch criteria.Operator {
		case "=", "<=>":
			var col, val sqlparser.ValExpr
			if plan.routingAnalyzeValue(criteria.Left) == EID_NODE {
				col = criteria.Left
				val = criteria.Right
			} else {
				col = criteria.Right
				val = criteria.Left
			}
			if plan.rule.KeyEqual(String(col)) {
				index = plan.findShard(val)
				return []int{index}
			} else {
				return makeList(index, len(plan.rule.Nodes))
			}
		case "<", "<=":
			if plan.rule.Type == HashRuleType {
				return plan.fullList
			}

			if plan.routingAnalyzeValue(criteria.Left) == EID_NODE {
				index = plan.findShard(criteria.Right)
				if criteria.Operator == "<" {
					index = plan.adjustShardIndex(criteria.Right, index)
				}

				return makeList(0, index+1)
			} else {
				index = plan.findShard(criteria.Left)
				return makeList(index, len(plan.rule.Nodes))
			}
		case ">", ">=":
			if plan.rule.Type == HashRuleType {
				return plan.fullList
			}

			if plan.routingAnalyzeValue(criteria.Left) == EID_NODE {
				index = plan.findShard(criteria.Right)
				return makeList(index, len(plan.rule.Nodes))
			} else {
				index = plan.findShard(criteria.Left)

				if criteria.Operator == ">" {
					index = plan.adjustShardIndex(criteria.Left, index)
				}
				return makeList(0, index+1)
			}
			/*
				case "in":
					return plan.findShardList(criteria.Right)
				case "not in":
					if plan.rule.Type == router.RangeRuleType {
						return plan.fullList
					}

					l := plan.findShardList(criteria.Right)
					return plan.notList(l)
			*/
		}
	case *sqlparser.RangeCond:
		if plan.rule.Type == HashRuleType {
			return plan.fullList
		}

		start := plan.findShard(criteria.From)
		last := plan.findShard(criteria.To)

		if criteria.Operator == "between" {
			if last < start {
				start, last = last, start
			}
			l := makeList(start, last+1)
			return l
		} else {
			if last < start {
				start, last = last, start
				start = plan.adjustShardIndex(criteria.To, start)
			} else {
				start = plan.adjustShardIndex(criteria.From, start)
			}

			l1 := makeList(0, start+1)
			l2 := makeList(last, len(plan.rule.Nodes))
			return unionList(l1, l2)
		}
	default:
		return plan.fullList
	}

	return plan.fullList
}

func (plan *Plan) findShard(valExpr sqlparser.ValExpr) int {
	value := plan.getBoundValue(valExpr)
	return plan.rule.FindNodeIndex(value)
}

func (plan *Plan) getBoundValue(valExpr sqlparser.ValExpr) interface{} {
	switch node := valExpr.(type) {
	case sqlparser.ValTuple:
		if len(node) != 1 {
			panic(NewParserError("tuples not allowed as insert values"))
		}
		// TODO: Change parser to create single value tuples into non-tuples.
		return plan.getBoundValue(node[0])
	case sqlparser.StrVal:
		return string(node)
	case sqlparser.NumVal:
		val, err := strconv.ParseInt(string(node), 10, 64)
		if err != nil {
			panic(NewParserError("%s", err.Error()))
		}
		return val
	case sqlparser.ValArg:
		return plan.stmtArgs[int(node[2]-48)]
		/*
			key := fmt.Sprintf("v%d", plan.disKeyIdx)
			return plan.bindVars[key]
		*/
	}
	panic("Unexpected token")
}

func (plan *Plan) adjustShardIndex(valExpr sqlparser.ValExpr, index int) int {
	value := plan.getBoundValue(valExpr)

	s, ok := plan.rule.Shard.(RangeShard)
	if !ok {
		return index
	}

	if s.EqualStart(value, index) {
		index--
		if index < 0 {
			panic(NewParserError("invalid range sharding"))
		}
	}
	return index
}

func (plan *Plan) findShardList(valExpr sqlparser.ValExpr) []int {
	shardset := make(map[int]bool)
	switch node := valExpr.(type) {
	case sqlparser.ValTuple:
		for _, n := range node {
			index := plan.findShard(n)
			shardset[index] = true
		}
	}
	shardlist := make([]int, len(shardset))
	index := 0
	for k := range shardset {
		shardlist[index] = k
		index++
	}

	sort.Ints(shardlist)
	return shardlist
}

func (plan *Plan) AnalyzeValue(valExpr sqlparser.ValExpr) int {
	switch valExpr.(type) {
	case *sqlparser.ColName:
		return EID_NODE
	default:
		return ^EID_NODE
	}

}

func (plan *Plan) ShardForWhere(where *sqlparser.Where) {
	if where == nil {
		plan.ShardList = plan.fullList
		return
	}
	plan.ShardList = plan.routingAnalyzeBoolean(where.Expr)
}

func (plan *Plan) findInsertShard(val sqlparser.Tuple) int {
	row, ok := val.(sqlparser.ValTuple)
	if !ok {
		panic(errors.New2("can't shard for this type of values"))
	}

	return plan.findShard(row[plan.shardKeyIdx])
}

func GeneralPlanForSelect(r *Router, stmt *sqlparser.Select, args map[int]interface{}) (plan *Plan, err error) {

	defer handleError(&err)
	plan = &Plan{}
	plan.stmtArgs = args

	for _, expr := range stmt.SelectExprs {
		switch ee := expr.(type) {
		case *sqlparser.StarExpr:
			panic("unsupported select *")
		case *sqlparser.NonStarExpr:
			if _, ok := ee.Expr.(*sqlparser.Subquery); ok {
				panic(UNSUPPORTED_SHARD_ERR)
			}
		}
	}

	plan.ShardForFrom(r, stmt.Where, stmt.From...)
	return
}

func GeneralPlanForInsert(r *Router, ist *sqlparser.Insert, args map[int]interface{}) (plan *Plan, err error) {
	defer handleError(&err)

	if r == nil {
		panic(errors.New(fmt.Errorf("cant't shard use nil router")))
	}

	var ok bool
	var vals sqlparser.Values
	plan = &Plan{stmtArgs: args}

	if plan.rule, ok = r.Rules[string(ist.Table.Name)]; !ok {
		panic(errors.New2("can't shard for this table: " + String(ist.Table)))
	}

	// only shard for values, insert select or any others not support
	vals, ok = ist.Rows.(sqlparser.Values)
	if !ok {
		panic(errors.New(fmt.Errorf("can't shard for this kind of insert rows")))
	}

	// can not shard for insert multi vals
	if len(vals) != 1 {
		panic(fmt.Errorf("can't shard for insert multi values"))
	}

	// find shard key idx
	for idx, col := range ist.Columns {
		c := col.(*sqlparser.NonStarExpr).Expr.(*sqlparser.ColName)
		if string(c.Name) == plan.rule.Key {
			plan.shardKeyIdx = idx
			break
		}
	}

	if plan.shardKeyIdx == 0 {
		panic(errors.New(fmt.Errorf("can't find shard key in insert cols")))
	}

	idx := plan.findInsertShard(vals[0])
	plan.ShardList = []int{idx}

	return
}

func GeneralPlanForWhere(r *Router, table string, where *sqlparser.Where) (plan *Plan, err error) {
	defer handleError(&err)
	plan = &Plan{}

	var ok bool
	if plan.rule, ok = r.Rules[table]; !ok {
		panic(errors.New(fmt.Errorf("can't find shard rule for this table: %s", table)))
	}
	plan.fullList = makeList(0, len(plan.rule.Nodes))
	plan.ShardForWhere(where)
	return
}

func GeneralShardList(r *Router, stmt sqlparser.Statement, args map[int]interface{}) ([]int, error) {
	var plan *Plan
	var err error

	switch s := stmt.(type) {
	case *sqlparser.Select:
		plan, err = GeneralPlanForSelect(r, s, args)
	case *sqlparser.Insert:
		plan, err = GeneralPlanForInsert(r, s, args)
	case *sqlparser.Update:
		plan, err = GeneralPlanForWhere(r, string(s.Table.Name), s.Where)
	case *sqlparser.Delete:
		plan, err = GeneralPlanForWhere(r, string(s.Table.Name), s.Where)
	default:
		return nil, errors.New2("can't shard for this type of sql")
	}

	if err != nil {
		return nil, err
	}

	if plan != nil {
		if plan.ShardList == nil {
			return nil, errors.New2("un expected plan's shard list is nil")
		} else {
			log.Debugf("get shard list: %v", plan.ShardList)
			return plan.ShardList, nil
		}
	} else {
		return nil, errors.New2("un expected plan is nil")
	}
}
