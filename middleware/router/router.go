/**
 *  author: lim
 *  data  : 18-4-16 下午7:45
 */

package router

import (
	"errors"
	"strconv"

	"github.com/lemonwx/xsql/sqlparser"
	"github.com/lemonwx/xsql/middleware/meta"
)

const (
	EID_NODE = iota
	VALUE_NODE
	LIST_NODE
	OTHER_NODE
)

var UNSUPPORTED_SQL_ROUTER_ERR error = errors.New("UNSUPPORTED SQL ROUTER ERR")
var SHARDDING_KEY_NOT_IN_COL_LIST_ERR error = errors.New("SHARDDING KEY NOT IN COL LIST ERR")



func GetNodeIdxs(stmt sqlparser.Statement) ([]int, error) {

	rule := meta.Meta{
		Db: "db",
		Tb: "tt",
		DisKey: "id",
		DisType: meta.HASH,
		DisNode: meta.FullNodeIdxs,
	}

	switch v := stmt.(type) {

	case *sqlparser.Insert:

		if err := judgeSupport(v.Rows.(sqlparser.Values)); err != nil {
			return nil, err
		}

		for idx, col := range v.Columns {
			// find dis key
			if sqlparser.String(col) == rule.DisKey {
				retIdxs := []int{-1}
				criteria := v.Rows.(sqlparser.Values)

				// for every rows
				for _, val := range criteria {
					key := val.(sqlparser.ValTuple)[idx]
					idx := sharding(&rule, key)
					if retIdxs[0] == -1 {
						retIdxs[0] = idx
					} else if retIdxs[0] != idx {
						return nil, UNSUPPORTED_SQL_ROUTER_ERR
					}
				}
				return retIdxs, nil
			}
		}

		return nil, SHARDDING_KEY_NOT_IN_COL_LIST_ERR

	case *sqlparser.Select:
		if v.Where == nil {
			return rule.DisNode, nil
		}

		return shardByBoolean(v.Where.Expr, &rule)
	case *sqlparser.Update:
		if v.Where == nil {
			return rule.DisNode, nil
		}

		return shardByBoolean(v.Where.Expr, &rule)
	case *sqlparser.Delete:
		if v.Where == nil {
			return rule.DisNode, nil
		}

		return shardByBoolean(v.Where.Expr, &rule)
	}

	return nil, UNSUPPORTED_SQL_ROUTER_ERR
}

func shardByCriteria(criteria sqlparser.SQLNode) ([]int, error) {

	switch criteria := criteria.(type) {
	case sqlparser.Values:
		index := shardByValues(criteria)
		return []int{index}, nil
	case sqlparser.BoolExpr:
		return shardByBoolean(criteria.(sqlparser.BoolExpr), nil)
	default:
		return nil, nil
	}
}

func shardByValues (vals sqlparser.Values) int {
	index := -1
	return index
}

func shardByBoolean (node sqlparser.BoolExpr, rule *meta.Meta) ([]int, error) {
	switch expr := node.(type) {
	case *sqlparser.ComparisonExpr :
		if expr.Operator == "=" {
			col, val := getColVal(expr)

			if sqlparser.String(col) == rule.DisKey {
				idx := sharding(rule, val)
				return []int{idx}, nil
			} else {
				return rule.DisNode, nil
			}
		}
	}

	return nil, UNSUPPORTED_SQL_ROUTER_ERR
}

func getColVal(expr *sqlparser.ComparisonExpr) (sqlparser.ValExpr, sqlparser.ValExpr) {
	l := routingAnalyzeValue(expr.Left)
	r := routingAnalyzeValue(expr.Right)
	var col, val sqlparser.ValExpr

	if l == EID_NODE && r == VALUE_NODE {
		col, val = expr.Left, expr.Right
	} else if l == VALUE_NODE && r == EID_NODE {
		col, val = expr.Right, expr.Left
	}
	return col, val
}

func judgeSupport(vals sqlparser.Values) error {
	for i := 0; i < len(vals); i++ {
		switch tuple := vals[i].(type) {
		case sqlparser.ValTuple:
			result := routingAnalyzeValue(tuple[0])
			if result != VALUE_NODE {

			}
		default:
			return UNSUPPORTED_SQL_ROUTER_ERR
		}
	}

	return nil
}

func routingAnalyzeValue(valExpr sqlparser.ValExpr) int {
	switch node := valExpr.(type) {
	case *sqlparser.ColName:
			return EID_NODE
	case sqlparser.ValTuple:
		for _, n := range node {
			if routingAnalyzeValue(n) != VALUE_NODE {
				return OTHER_NODE
			}
		}
		return LIST_NODE
	case sqlparser.StrVal, sqlparser.NumVal, sqlparser.ValArg:
		return VALUE_NODE
	}
	return OTHER_NODE
}

func sharding(rule *meta.Meta, val sqlparser.ValExpr) int {
	switch rule.DisType {
	case meta.HASH:
		shard := HashShard{ShardNum: len(rule.DisNode)}
		return shard.FindForKey(getBoundValue(val))
	case meta.RANGE:
	case meta.DUP:
	case meta.LIST:
	}
	return -1
}


func getBoundValue(valExpr sqlparser.ValExpr) interface{} {
	switch node := valExpr.(type) {
	case sqlparser.ValTuple:
		if len(node) != 1 {
			return errors.New("tuples not allowed as insert values")
		}
		// TODO: Change parser to create single value tuples into non-tuples.
		return getBoundValue(node[0])
	case sqlparser.StrVal:
		return string(node)
	case sqlparser.NumVal:
		val, err := strconv.ParseInt(string(node), 10, 64)
		if err != nil {
			return err
		}
		return val
		/*
	case sqlparser.ValArg:
		return plan.bindVars[string(node[1:])]
		*/
	}
	return errors.New("Unexpected token")
}