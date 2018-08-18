/**
 *  author: lim
 *  data  : 18-5-22 下午10:40
 */

package sqlparser

import (
	"fmt"
	"hack"
	"sort"
	"strconv"

	"utils"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/errors"
	"github.com/lemonwx/xsql/middleware/router"
)

var UNSUPPORTED_SHARD_ERR = errors.New(fmt.Errorf("unsupported shard for this sql"))

type Plan struct {
	rule      *router.Rule
	fullList  []int
	ShardList []int

	//for insert
	shardKeyIdx int
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

func (p *Plan) ShardForFrom(r *router.Router, preWhere *Where, froms ...TableExpr) (*Plan, error) {

	if len(froms) == 1 {
		switch v := froms[0].(type) {
		case *AliasedTableExpr:
			switch node := v.Expr.(type) {
			case *TableName:
				rule := r.GetRule(hack.String(node.Name))
				rule.As = hack.String(v.As)
				p.rule = rule
				p.fullList = makeList(0, len(rule.Nodes))

				if preWhere == nil {
					p.ShardList = p.fullList
				} else {
					p.ShardList = p.routingAnalyzeBoolean(preWhere.Expr)
				}
			case *Subquery:
				switch sel := node.Select.(type) {
				case *SimpleSelect:
					panic(UNSUPPORTED_SHARD_ERR)
				case *Select:
					plan, err := GeneralPlanForSelect(r, sel)
					if err != nil {
						panic(err)
					}

					plan.rule.As = hack.String(v.As)
					p.rule = plan.rule
					p.ShardList = plan.ShardList
					p.fullList = plan.fullList
				case *Union:
					panic(UNSUPPORTED_SHARD_ERR)
				}
			}
		case *JoinTableExpr:
			pl := &Plan{}
			pr := &Plan{}

			pl.ShardForFrom(r, nil, v.LeftExpr)
			pr.ShardForFrom(r, nil, v.RightExpr)

			lr := pl.rule
			rr := pr.rule

			if !lr.Equal(rr) {
				panic(UNSUPPORTED_SHARD_ERR)
			}

			log.Debugf("lr equal rr: %v == %v", lr, rr)

			if b, ok := v.On.(*ComparisonExpr); ok {
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
					compare, ok := preWhere.Expr.(*ComparisonExpr)
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
							log.Debug("join table expr's where: [%v] is compare(\"=\") but not [ l == r == EID_NODE ], and can't by on", preWhere)
							panic(UNSUPPORTED_SHARD_ERR)
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
		case *ParenTableExpr:
			panic(UNSUPPORTED_SHARD_ERR)
		}
	} else {
		var rules []*router.Rule
		var rule *router.Rule
		for _, tb := range froms {
			switch v := tb.(type) {
			case *AliasedTableExpr:
				switch node := v.Expr.(type) {
				case *TableName:
					rule = r.GetRule(hack.String(node.Name))
					rule.As = hack.String(v.As)
				case *Subquery:
					switch sel := node.Select.(type) {
					case *SimpleSelect:
						panic(UNSUPPORTED_SHARD_ERR)
					case *Select:
						plan, err := GeneralPlanForSelect(r, sel)
						if err != nil {
							panic(err)
						}
						plan.rule.As = hack.String(v.As)
						rule = plan.rule
					case *Union:
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
			case *AndExpr:
				panic(UNSUPPORTED_SHARD_ERR)
			case *ComparisonExpr:
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

func (plan *Plan) routingAnalyzeBoolean(where BoolExpr) []int {
	switch node := where.(type) {
	case *AndExpr:
		left := plan.routingAnalyzeBoolean(node.Left)
		right := plan.routingAnalyzeBoolean(node.Right)
		return interList(left, right)
	case *OrExpr:
		left := plan.routingAnalyzeBoolean(node.Left)
		right := plan.routingAnalyzeBoolean(node.Right)
		return unionList(left, right)
	case *ParenBoolExpr:
		return plan.routingAnalyzeBoolean(node.Expr)
	case *ComparisonExpr:
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
	case *RangeCond:
		left := plan.routingAnalyzeValue(node.Left)
		from := plan.routingAnalyzeValue(node.From)
		to := plan.routingAnalyzeValue(node.To)
		if left == EID_NODE && from == VALUE_NODE && to == VALUE_NODE {
			return plan.findConditionShard(node)
		}
	case *ExistsExpr:
		switch node.Subquery.Select.(type) {
		case *SimpleSelect:
			return plan.fullList
		default:
			panic(fmt.Errorf("unsupported sharding type"))
		}
	}

	return plan.fullList

}

func (plan *Plan) routingAnalyzeValue(valExpr ValExpr) int {
	switch node := valExpr.(type) {
	case *ColName:
		if plan.rule.KeyEqual(String(node)) {
			return EID_NODE
		}
	case ValTuple:
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
	case StrVal, NumVal, ValArg:
		return VALUE_NODE
	}
	return OTHER_NODE
}

func (plan *Plan) routingAnalyzeValues(vals Values) Values {
	// Analyze first value of every item in the list
	log.Debug(vals)
	for i := 0; i < len(vals); i++ {
		switch tuple := vals[i].(type) {
		case ValTuple:
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

func (plan *Plan) findConditionShard(expr BoolExpr) (shardList []int) {
	var index int
	switch criteria := expr.(type) {
	case *ComparisonExpr:
		switch criteria.Operator {
		case "=", "<=>":
			var col, val ValExpr
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
			if plan.rule.Type == router.HashRuleType {
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
			if plan.rule.Type == router.HashRuleType {
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
	case *RangeCond:
		if plan.rule.Type == router.HashRuleType {
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

func (plan *Plan) findShard(valExpr ValExpr) int {
	value := plan.getBoundValue(valExpr)
	return plan.rule.FindNodeIndex(value)
}

func (plan *Plan) getBoundValue(valExpr ValExpr) interface{} {
	switch node := valExpr.(type) {
	case ValTuple:
		if len(node) != 1 {
			panic(NewParserError("tuples not allowed as insert values"))
		}
		// TODO: Change parser to create single value tuples into non-tuples.
		return plan.getBoundValue(node[0])
	case StrVal:
		return string(node)
	case NumVal:
		val, err := strconv.ParseInt(string(node), 10, 64)
		if err != nil {
			panic(NewParserError("%s", err.Error()))
		}
		return val
	case ValArg:
		/*
			key := fmt.Sprintf("v%d", plan.disKeyIdx)
			return plan.bindVars[key]
		*/
	}
	panic("Unexpected token")
}

func (plan *Plan) adjustShardIndex(valExpr ValExpr, index int) int {
	value := plan.getBoundValue(valExpr)

	s, ok := plan.rule.Shard.(router.RangeShard)
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

func (plan *Plan) findShardList(valExpr ValExpr) []int {
	shardset := make(map[int]bool)
	switch node := valExpr.(type) {
	case ValTuple:
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

func (plan *Plan) AnalyzeValue(valExpr ValExpr) int {
	switch valExpr.(type) {
	case *ColName:
		return EID_NODE
	default:
		return ^EID_NODE
	}

}

func (plan *Plan) findInsertShard(val Tuple) int {
	row, ok := val.(ValTuple)
	if !ok {
		panic(errors.New2("can't shard for this type of values"))
	}

	return plan.findShard(row[plan.shardKeyIdx])
}

func GeneralPlanForSelect(r *router.Router, stmt *Select) (plan *Plan, err error) {

	defer handleError(&err)
	plan = &Plan{}

	for _, expr := range stmt.SelectExprs {
		switch ee := expr.(type) {
		case *StarExpr:
			panic("unsupported select *")
		case *NonStarExpr:
			if _, ok := ee.Expr.(*Subquery); ok {
				panic(UNSUPPORTED_SHARD_ERR)
			}
		}
	}

	plan.ShardForFrom(r, stmt.Where, stmt.From...)
	return
}

func GeneralPlanForInsert(r *router.Router, ist *Insert) (plan *Plan, err error) {
	defer handleError(&err)

	if r == nil {
		panic(errors.New(fmt.Errorf("cant't shard use nil router")))
	}

	var ok bool
	var vals Values
	plan = &Plan{}

	// get rule for this table
	if plan.rule, ok = r.Rules[string(ist.Table.Name)]; !ok {
		panic(errors.New(fmt.Errorf("can't find shard rule for this table: %s", ist.Table.Name)))
	}

	// only shard for values, insert select or any others not support
	vals, ok = ist.Rows.(Values)
	if !ok {
		panic(errors.New(fmt.Errorf("can't shard for this kind of insert rows")))
	}

	// can not shard for insert multi vals
	if len(vals) != 1 {
		panic(fmt.Errorf("can't shard for insert multi values"))
	}

	// find shard key idx
	for idx, col := range ist.Columns {
		c := col.(*NonStarExpr).Expr.(*ColName)
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

func GeneralShardList(r *router.Router, stmt Statement) ([]int, error) {
	var plan *Plan
	var err error

	switch s := stmt.(type) {
	case *Select:
		plan, err = GeneralPlanForSelect(r, s)
	case *Insert:
		plan, err = GeneralPlanForInsert(r, s)
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
			return plan.ShardList, nil
		}
	} else {
		return nil, errors.New2("un expected plan is nil")
	}
}
