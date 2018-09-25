/**
 *  author: lim
 *  data  : 18-9-18 下午9:34
 */

package server

import (
	"bytes"
	"fmt"
	"hack"
	"sort"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/errors"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/xsql/sqlparser"
)

type order struct {
	key  uint64
	dire string
}

type OrderedResult struct {
	*mysql.Resultset
	orderKeys []*order
	vals      [][]interface{}
}

func NewOrderedResult(rets []*mysql.Result, orderBy sqlparser.OrderBy, extraSize int) (OrderedResult, error) {
	ret := OrderedResult{}
	if len(rets) == 0 {
		return ret, errors.New2("unexpected middleware error")
	}
	if rets[0] == nil {
		return ret, errors.New2("unexpected middleware error")
	}

	rs := rets[0]

	// merge data row numbers
	finalLen := 0
	for _, r := range rets {
		finalLen += len(r.RowDatas)
	}

	tgtRs := &mysql.Resultset{
		Fields:     rs.Fields,
		FieldNames: rs.FieldNames,
		RowDatas:   make([]mysql.RowData, finalLen),
	}

	copy(tgtRs.RowDatas, rs.RowDatas)
	from := len(rs.RowDatas)
	for _, rs := range rets[1:] {
		copy(tgtRs.RowDatas[from:], rs.RowDatas)
		from += len(rs.RowDatas)
	}

	ret.Resultset = tgtRs
	ret.orderKeys = make([]*order, 0, len(orderBy))

	for _, orderby := range orderBy {
		switch node := orderby.Expr.(type) {
		case sqlparser.NumVal:
			keyIdx, err := node.GetLocalVal()
			if err != nil {
				return ret, err
			}

			if keyIdx > uint64(len(ret.Fields)) {
				return ret, newMySQLErr(errOrderByIdxOutOfRange)
			}

			ret.orderKeys = append(ret.orderKeys, &order{keyIdx - uint64(extraSize), orderby.Direction})
		case *sqlparser.ColName:
			log.Debugf("cols: %v, want: %s", ret.FieldNames, node.Name)
			if colIdx, ok := ret.FieldNames[string(node.Name)]; ok {
				ret.orderKeys = append(ret.orderKeys, &order{uint64(colIdx - extraSize), orderby.Direction})
			} else {
				return ret, newMySQLErr(errOrderByColMustInSelectList)
			}
		default:
			return ret, errors.New2("unsupported order key type")
		}
	}
	if err := ret.getLocalVal(); err != nil {
		return ret, err
	}
	return ret, nil
}

func (oRet *OrderedResult) getLocalVal() error {
	oRet.vals = make([][]interface{}, 0, len(oRet.RowDatas))
	for _, row := range oRet.RowDatas {
		val, err := row.Parse(oRet.Fields, false)
		if err != nil {
			return err
		}
		oRet.vals = append(oRet.vals, val)
	}
	return nil
}

func (oRet OrderedResult) Len() int {
	return len(oRet.RowDatas)
}

//compare value using asc
func cmpValue(v1 interface{}, v2 interface{}) int {
	if v1 == nil && v2 == nil {
		return 0
	} else if v1 == nil {
		return -1
	} else if v2 == nil {
		return 1
	}

	switch v := v1.(type) {
	case string:
		s := v2.(string)
		return bytes.Compare(hack.Slice(v), hack.Slice(s))
	case []byte:
		s := v2.([]byte)
		return bytes.Compare(v, s)
	case int64:
		s := v2.(int64)
		if v < s {
			return -1
		} else if v > s {
			return 1
		} else {
			return 0
		}
	case uint64:
		s := v2.(uint64)
		if v < s {
			return -1
		} else if v > s {
			return 1
		} else {
			return 0
		}
	case float64:
		s := v2.(float64)
		if v < s {
			return -1
		} else if v > s {
			return 1
		} else {
			return 0
		}
	default:
		//can not go here
		panic(fmt.Sprintf("invalid type %T", v))
	}
}

func (oRet OrderedResult) Less(i, j int) bool {
	row1 := oRet.vals[i]
	row2 := oRet.vals[j]

	for _, orderKey := range oRet.orderKeys {
		f1 := row1[orderKey.key]
		f2 := row2[orderKey.key]

		if orderKey.dire == "desc" {
			f1, f2 = f2, f1
		}

		comRet := cmpValue(f1, f2)
		if comRet < 0 {
			return true
		} else if comRet > 0 {
			return false
		}
	}
	return false
}

func (oRet OrderedResult) Swap(i, j int) {
	oRet.RowDatas[i], oRet.RowDatas[j] = oRet.RowDatas[j], oRet.RowDatas[i]
	oRet.vals[i], oRet.vals[j] = oRet.vals[j], oRet.vals[i]
}

func (conn *MidConn) handleOrderBy(rets []*mysql.Result, sel *sqlparser.Select) (*mysql.Result, error) {
	oRet, err := NewOrderedResult(rets, sel.OrderBy, len(sel.ExtraCols))
	if err != nil {
		return nil, err
	}
	sort.Sort(oRet)
	return &mysql.Result{Resultset: oRet.Resultset}, nil
}
