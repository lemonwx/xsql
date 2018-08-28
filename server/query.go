/**
 *  author: lim
 *  data  : 18-4-6 下午5:15
 */

package server

import (
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

func (conn *MidConn) executeSelect(sql string, extraSz int, flag uint8) ([]*mysql.Result, error) {
	if len(conn.nodeIdx) == 0 {
		return nil, errors.New2("can't execute sql under empty node idxs")
	}

	var wg sync.WaitGroup
	wg.Add(2)

	var vErr error
	var vInUse map[uint64]uint8
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
		return nil, vErr
	}

	if exeErr != nil {
		return nil, exeErr
	}

	if err := conn.chkInUse(&rets, extraSz, vInUse); err != nil {
		return nil, err
	}

	return rets, nil
}

func (conn *MidConn) chkInUse(rets *[]*mysql.Result, extraSz int, vInUse map[uint64]uint8) error {
	ts := time.Now()
	defer func() {
		conn.stat.chkInuseT += time.Since(ts)
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

func (conn *MidConn) hideExtraCols(data *mysql.RowData, size int, vs map[uint64]uint8) error {
	idx := uint8(0)
	for count := 0; count < size; count += 1 {
		s := idx + 1
		e := s + (*data)[idx]

		vStr := string((*data)[s:e])
		res, err := strconv.ParseUint(vStr, 10, 64)
		if err != nil {
			log.Errorf("[%d] ParseUint from %v failed: %v", vStr, err)
			return mysql.NewDefaultError(mysql.MID_ER_HIDE_EXTRA_FAILED)
		}
		if _, ok := vs[res]; ok {
			err = mysql.NewDefaultError(mysql.MID_ER_ROWS_IN_USE_BY_OTHER_SESSION)
			log.Errorf("[%d] hide extra col failed: %v", conn.ConnectionId, err)
			return err
		}
		idx = (*data)[idx] + idx + 1
	}
	(*data) = (*data)[idx:]
	return nil
}

func (conn *MidConn) handleSelect(stmt *sqlparser.Select) ([]*mysql.Result, error) {
	var err error
	if conn.nodeIdx, err = conn.getShardList(stmt); err != nil {
		log.Errorf("[%d] get shard list failed:%v", conn.ConnectionId, err)
		return nil, conn.NewMySQLErr(ERR_UNSUPPORTED_SHARD)
	}

	if len(conn.nodeIdx) == 0 {
		log.Debugf("[%d] get empty shard list, nothing to execute, pls chk your sql", conn.ConnectionId)
		rs := conn.newEmptyResultset(stmt)
		ret := &mysql.Result{Resultset: rs}
		return []*mysql.Result{ret}, nil
	}

	rets, err := conn.executeSelect(sqlparser.String(stmt), len(stmt.ExtraCols), SELECT)
	return rets, err
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

func (conn *MidConn) handleLimit(rets []*mysql.Result, limit *sqlparser.Limit) ([]*mysql.Result, error) {

	if len(rets) == 0 {
		log.Errorf("[%d] handle limit rets's len == 0, unexpected err", conn.ConnectionId)
		return nil, mysql.NewDefaultError(mysql.MID_ER_UNEXPECTED)
	}

	if limit != nil {
		if limit.Offset != nil {
			log.Errorf("[%d] offset : %v not nil, not support this sql now", conn.ConnectionId, limit.Offset)
			return nil, mysql.NewDefaultError(mysql.MID_ER_UNSUPPORTED_SQL)
		}
		log.Debugf("[%d] offset: %v, rows count: %d", conn.ConnectionId, limit.Offset, limit.Rowcount)

		limitCount, err := strconv.ParseUint(string(limit.Rowcount.(sqlparser.NumVal)), 10, 64)
		if err != nil {
			log.Errorf("[%d] parse limit count failed: %v", conn.ConnectionId, err)
			return nil, err
		}

		allCount := uint64(0)
		for idx, ret := range rets {
			tmp := uint64(len(ret.RowDatas))
			if allCount+tmp >= limitCount {

				rets[idx].RowDatas = rets[idx].RowDatas[:limitCount-allCount]
				return rets[:idx+1], nil
			}
			allCount += tmp
		}
	}

	return rets, nil
}

func (conn *MidConn) setupNodeStatus(vInUse map[uint64]byte, hide bool, isStmt bool, extraSize int) {
	for idx, _ := range conn.nodes {
		conn.nodes[idx].VersionsInUse = vInUse
		conn.nodes[idx].NeedHide = hide
		conn.nodes[idx].IsStmt = isStmt
		conn.nodes[idx].ExtraSize = extraSize
	}
}
