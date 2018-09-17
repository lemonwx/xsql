/**
 *  author: lim
 *  data  : 18-4-6 下午5:15
 */

package server

import (
	"bytes"
	"hack"
	"sort"
	"strconv"

	"sync"

	"fmt"

	"time"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/errors"
	"github.com/lemonwx/xsql/meta"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/xsql/sqlparser"
)

func (conn *MidConn) handleShow(stmt *sqlparser.Show, sql string) error {
	defer conn.clearExecNodes([]byte("rollback"))

	if err := conn.getMultiBackConn(meta.GetFullNodeIdxs()); err != nil {
		return err
	}

	var rets []*mysql.Result
	for _, back := range conn.execNodes {
		ret, err := conn.execute(back, mysql.COM_QUERY, []byte(sql))
		if err != nil {
			return err
		}

		ret.FieldNames["node"] = 2
		ret.Fields = append(ret.Fields, &mysql.Field{Name: []byte("node")})
		val := back.String()
		col := make([]byte, len(val)+1)
		col[0] = byte(len(val))
		copy(col[1:], val)

		for idx, _ := range ret.RowDatas {
			ret.RowDatas[idx] = append(ret.RowDatas[idx], col...)
		}

		rets = append(rets, ret)
	}

	return conn.HandleSelRets(rets)
}

func (conn *MidConn) handleSimpleSelect(stmt *sqlparser.SimpleSelect, sql string) error {
	log.Debugf("[%d] handle simple select", conn.ConnectionId)

	back, err := conn.pools[0].GetConn(conn.db)
	if err != nil {
		return err
	}
	defer conn.pools[0].PutConn(back)

	ret, err := conn.execute(back, mysql.COM_QUERY, []byte(sql))
	if err != nil {
		return err
	}
	return conn.HandleSelRets([]*mysql.Result{ret})
}

func (conn *MidConn) executeSelect(sql string, flag uint8) ([]*mysql.Result, map[uint64]bool, error) {
	if len(conn.nodeIdx) == 0 {
		return nil, nil, errors.New2("can't execute sql under empty node idxs")
	}

	var wg sync.WaitGroup
	wg.Add(2)

	var vErr error
	var vInUse map[uint64]bool
	go func() {
		vInUse, vErr = conn.getCurVInUse(flag)
		wg.Done()
	}()

	var rets []*mysql.Result
	var exeErr error
	go func() {
		rets, exeErr = conn.ExecuteOnNodePool([]byte(sql), conn.nodeIdx)
		wg.Done()
	}()

	wg.Wait()

	if vErr != nil {
		log.Errorf("[%d] get version failed: %v", conn.ConnectionId, vErr)
		return nil, nil, vErr
	}

	if exeErr != nil {
		log.Errorf("[%d] execute sql failed: %v", conn.ConnectionId, exeErr)
		return nil, nil, exeErr
	}

	return rets, vInUse, nil
}

func (conn *MidConn) chkInUse(rets *[]*mysql.Result, extraSz int, vInUse map[uint64]bool) error {
	ts := time.Now()
	defer func() {
		conn.stat.ChkInuseT.add(int64(time.Since(ts)))
	}()

	for idx, ret := range *rets {
		ret.Fields = ret.Fields[extraSz:]
		for rowIdx, _ := range ret.RowDatas {
			if err := conn.hideExtraCols(&ret.RowDatas[rowIdx], extraSz, vInUse); err != nil {
				return err
			}
		}
		(*rets)[idx] = ret
	}
	return nil
}

func (conn *MidConn) hideExtraCols(data *mysql.RowData, size int, vs map[uint64]bool) error {
	idx := uint8(0)
	for count := 0; count < size; count += 1 {
		s := idx + 1
		e := s + (*data)[idx]
		idx = (*data)[idx] + idx + 1

		vStr := string((*data)[s:e])
		res, err := strconv.ParseUint(vStr, 10, 64)
		if err != nil {
			log.Errorf("[%d] ParseUint from %v failed: %v", vStr, err)
			return mysql.NewDefaultError(mysql.MID_ER_HIDE_EXTRA_FAILED)
		}

		if res == conn.NextVersion {
			continue
		}

		if _, ok := vs[res]; ok {
			err = mysql.NewDefaultError(mysql.MID_ER_ROWS_IN_USE_BY_OTHER_SESSION)
			log.Errorf("[%d] hide extra col failed: %v", conn.ConnectionId, err)
			return err
		}
	}
	(*data) = (*data)[idx:]
	return nil
}

func (conn *MidConn) handleSelect(stmt *sqlparser.Select) ([]*mysql.Result, error) {
	var err error
	if conn.nodeIdx, err = conn.getShardList(stmt); err != nil {
		log.Errorf("[%d] get shard list failed:%v", conn.ConnectionId, err)
		return nil, conn.NewMySQLErr(errUnsupportedShard)
	}

	if len(conn.nodeIdx) == 0 {
		log.Debugf("[%d] get empty shard list, nothing to execute, pls chk your sql", conn.ConnectionId)
		rs := conn.newEmptyResultset(stmt)
		ret := &mysql.Result{Resultset: rs}
		return []*mysql.Result{ret}, nil
	}

	rets, vInUse, err := conn.executeSelect(sqlparser.String(stmt), Select)
	if err != nil {
		return nil, mysql.NewDefaultError(errUnexpected)
	}

	if err := conn.handleLimit(&rets, stmt.Limit); err != nil {
		log.Debug(err)
		return nil, err
	}

	if err := conn.chkInUse(&rets, len(stmt.ExtraCols), vInUse); err != nil {
		return nil, err
	}

	if len(stmt.OrderBy) != 0 {
		if rs, err := conn.handleOrderBy(rets, stmt); err != nil {
			return nil, err
		} else {
			return []*mysql.Result{rs}, nil
		}
	}

	return rets, nil
}

func (conn *MidConn) ExecuteOnSinglePool(sql []byte, nodeIdxs []int) ([]*mysql.Result, error) {
	if len(nodeIdxs) != 1 {
		return nil, fmt.Errorf("len of nodeIdxs must be 1")
	}

	idx := nodeIdxs[0]
	back, err := conn.getSingleBackConn(idx)
	if err != nil {
		return nil, err
	}

	ret, err := conn.execute(back, mysql.COM_QUERY, sql)
	if err != nil {
		return nil, err
	}
	return []*mysql.Result{ret}, nil
}

func (conn *MidConn) ExecuteOnMultiPool(sql []byte, nodeIdxs []int) ([]*mysql.Result, error) {
	if err := conn.getMultiBackConn(nodeIdxs); err != nil {
		return nil, err
	}

	shardSize := len(nodeIdxs)
	rets := make([]*mysql.Result, 0, shardSize)
	errs := make([]error, 0, shardSize)

	var wg sync.WaitGroup
	wg.Add(len(nodeIdxs))
	lock := sync.Mutex{}

	for _, idx := range nodeIdxs {
		go func(idx int) {
			back, ok := conn.execNodes[idx]
			if !ok {
				lock.Lock()
				errs = append(errs, fmt.Errorf("unexpected error, idx should in conn.execNodes"))
				lock.Unlock()
				wg.Done()
				return
			}
			ret, err := conn.execute(back, mysql.COM_QUERY, sql)
			lock.Lock()
			if err != nil {
				errs = append(errs, err)
			} else {
				rets = append(rets, ret)
			}
			lock.Unlock()
			wg.Done()
		}(idx)
	}
	wg.Wait()

	switch {
	case len(errs) == shardSize:
		return nil, errs[0]
	case len(rets) == shardSize:
		return rets, nil
	default:
		return nil, fmt.Errorf("unexpected multi node response not equal")
	}
}

func (conn *MidConn) ExecuteOnNodePool(sql []byte, nodeIdxs []int) ([]*mysql.Result, error) {
	if len(nodeIdxs) == 1 {
		return conn.ExecuteOnSinglePool(sql, nodeIdxs)
	} else {
		return conn.ExecuteOnMultiPool(sql, nodeIdxs)
	}
}

func (conn *MidConn) handleLimit(rets *[]*mysql.Result, limit *sqlparser.Limit) error {
	if len(*rets) == 0 {
		log.Errorf("[%d] handle limit rets's len == 0, unexpected err", conn.ConnectionId)
		return conn.NewMySQLErr(errUnexpected)
	}

	if limit != nil {
		if limit.Offset != nil {
			log.Errorf("[%d] offset : %v not nil, not support this sql now", conn.ConnectionId, limit.Offset)
			return conn.NewMySQLErr(errUnsupportedSql)
		}

		limitCount, err := strconv.ParseUint(string(limit.Rowcount.(sqlparser.NumVal)), 10, 64)
		if err != nil {
			log.Errorf("[%d] parse limit count failed: %v", conn.ConnectionId, err)
			return conn.NewMySQLErr(errUnsupportedSql)
		}

		log.Debugf("[%d] offset: %v, rows count: %d, %d", conn.ConnectionId, limit.Offset, limit.Rowcount, limitCount)

		allCount := uint64(0)
		// finally, should response to cli size
		finIdx := 0
		for idx, ret := range *rets {
			finIdx = idx
			curSize := uint64(len(ret.RowDatas))
			if allCount+curSize >= limitCount {
				(*rets)[idx].RowDatas = (*rets)[idx].RowDatas[:limitCount-allCount]
				break
			}
			allCount += curSize
		}

		log.Debug(finIdx)
		*rets = (*rets)[:finIdx+1]

	}
	return nil
}

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
			ret.orderKeys = append(ret.orderKeys, &order{keyIdx - uint64(extraSize), orderby.Direction})
		case *sqlparser.ColName:
			log.Debugf("cols: %v, want: %s", ret.FieldNames, node.Name)
			if colIdx, ok := ret.FieldNames[string(node.Name)]; ok {
				ret.orderKeys = append(ret.orderKeys, &order{uint64(colIdx - extraSize), orderby.Direction})
			} else {
				return ret, errors.New2("unknown colname")
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

func (conn *MidConn) setupNodeStatus(vInUse map[uint64]bool, hide bool, isStmt bool, extraSize int) {
	for idx, _ := range conn.nodes {
		conn.nodes[idx].VersionsInUse = vInUse
		conn.nodes[idx].NeedHide = hide
		conn.nodes[idx].IsStmt = isStmt
		conn.nodes[idx].ExtraSize = extraSize
	}
}
