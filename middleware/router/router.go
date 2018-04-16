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

var UNSUPPORT_SQL_ROUTER_ERR error = errors.New("UNSUPPORT SQL ROUTER ERR")
var SHARDDING_KEY_NOT_IN_COL_LIST_ERR error = errors.New("SHARDDING KEY NOT IN COL LIST ERR")



func GetNodeIdxs(stmt sqlparser.Statement) ([]int, error) {

	switch v := stmt.(type) {

	case *sqlparser.Insert:

		if err := judgeSupport(v.Rows.(sqlparser.Values)); err != nil {
			return nil, err
		}

		rule := meta.Meta{
			Db: "db",
			Tb: "tb",
			DisKey: "id",
			DisType: meta.HASH,
			DisNode: meta.FullNodeIdxs,
		}

		for idx, col := range v.Columns {
			// find dis key
			if sqlparser.String(col) == rule.DisKey {
				retIdxs := make([]int, 0, len(rule.DisNode))
				criteria := v.Rows.(sqlparser.Values)

				// for every rows
				for _, val := range criteria {
					key := val.(sqlparser.ValTuple)[idx]
					idx := sharding(&rule, key)
					retIdxs = append(retIdxs, idx)
				}
				return retIdxs, nil
			}
		}
		return nil, SHARDDING_KEY_NOT_IN_COL_LIST_ERR
	}
	return nil, nil
}

func shardByCriteria(criteria sqlparser.SQLNode) ([]int, error) {

	switch criteria := criteria.(type) {
	case sqlparser.Values:
		index := shardByValues(criteria)
		return []int{index}, nil
	case sqlparser.BoolExpr:
		return shardByBoolean(criteria), nil
	default:
		return nil, nil
	}
}

func shardByValues (vals sqlparser.Values) int {
	index := -1
	return index
}

func shardByBoolean (expr sqlparser.BoolExpr) []int {
	return nil
}

func judgeSupport(vals sqlparser.Values) error {
	for i := 0; i < len(vals); i++ {
		switch tuple := vals[i].(type) {
		case sqlparser.ValTuple:
			result := routingAnalyzeValue(tuple[0])
			if result != VALUE_NODE {

			}
		default:
			return UNSUPPORT_SQL_ROUTER_ERR
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
	default:
		 return -1
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