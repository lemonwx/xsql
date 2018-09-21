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

	"encoding/binary"

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

func (conn *MidConn) chkInUse(rets *[]*mysql.Result, extraSz int, vInUse map[uint64]bool, isBin bool) error {
	ts := time.Now()
	defer func() {
		conn.stat.ChkInuseT.add(int64(time.Since(ts)))
	}()

	for idx, ret := range *rets {
		ret.Fields = ret.Fields[extraSz:]
		for rowIdx, _ := range ret.RowDatas {
			if err := conn.hideExtraCols(&ret.RowDatas[rowIdx], extraSz, vInUse, len(ret.Fields), isBin); err != nil {
				return err
			}
		}
		(*rets)[idx] = ret
	}
	return nil
}

func (conn *MidConn) node2cliNullMask(data []byte, extraSize, fieldCount int) ([]byte, error) {

	cliSize := (fieldCount + 7 + 2) >> 3
	nodePos := 1 + ((fieldCount + extraSize + 7 + 2) >> 3)

	if nodePos >= len(data) {
		err := fmt.Errorf("[%d] unexpected node err: node pos: %d >= len(*data): %d", conn.ConnectionId, nodePos, len(data))
		log.Error(err)
		return nil, newDefaultMySQLError(errInternal, err)
	}

	nodeNullMask := (data)[1:nodePos]
	cliNullMask := make([]byte, cliSize)

	for idx := 0; idx < fieldCount; idx += 1 {
		nodeidx := idx + extraSize
		if (nodeidx+2)>>3 >= len(nodeNullMask) {
			err := fmt.Errorf("[%d] unexpected node err: idx: %d >= len(node null mask): %d", conn.ConnectionId,
				(nodeidx+2)>>3, len(nodeNullMask))
			log.Error(err)
			return nil, newDefaultMySQLError(errInternal, err)
		}
		if ((nodeNullMask[(nodeidx+2)>>3] >> uint((nodeidx+2)&7)) & 1) == 1 {
			cliNullMask[(idx+2)>>3] += 1 << uint((idx+2)%8)
		}
	}

	log.Debugf("[%d] 2 cli null mask : %v, len: %d, %d, %d", conn.ConnectionId, cliNullMask, cliSize, nodePos, fieldCount)
	return cliNullMask, nil
}

func (conn *MidConn) hideExtraCols(data *mysql.RowData, size int, vs map[uint64]bool, fieldSize int, isBin bool) error {
	if isBin {
		pos := 1 + ((fieldSize + size + 7 + 2) >> 3)
		nullMask := (*data)[1:pos]

		for idx := 0; idx < size; idx += 1 {
			if ((nullMask[(idx+2)>>3] >> uint((idx+2)&7)) & 1) == 1 {
				log.Errorf("[%d] unexpected node err: version parsed from ret is nil", conn.ConnectionId)
				return newDefaultMySQLError(errInternal, "decoded extra col is nil")
			}
			extra := uint64(binary.LittleEndian.Uint64((*data)[pos : pos+8]))
			log.Debugf("[%d] extra col val: %v", conn.ConnectionId, extra)
			if _, ok := vs[extra]; ok {
				return newMySQLErr(errRowsInuseByOthers)
			}
			pos += 8
		}
		mask, err := conn.node2cliNullMask(*data, size, fieldSize)
		if err != nil {
			return err
		}
		mask = append((*data)[:1], mask...)
		*data = append(mask, (*data)[pos:]...)
	} else {
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
	}
	return nil
}

func (conn *MidConn) handleSelect(stmt *sqlparser.Select) ([]*mysql.Result, error) {
	var err error
	if conn.nodeIdx, err = conn.getShardList(stmt, nil); err != nil {
		log.Errorf("[%d] get shard list failed:%v", conn.ConnectionId, err)
		return nil, newMySQLErr(errUnsupportedShard)
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

	if err := conn.chkInUse(&rets, len(stmt.ExtraCols), vInUse, false); err != nil {
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
		return newMySQLErr(errUnexpected)
	}

	if limit != nil {
		if limit.Offset != nil {
			log.Errorf("[%d] offset : %v not nil, not support this sql now", conn.ConnectionId, limit.Offset)
			return newMySQLErr(errUnsupportedSql)
		}

		limitCount, err := strconv.ParseUint(string(limit.Rowcount.(sqlparser.NumVal)), 10, 64)
		if err != nil {
			log.Errorf("[%d] parse limit count failed: %v", conn.ConnectionId, err)
			return newMySQLErr(errUnsupportedSql)
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

func (conn *MidConn) setupNodeStatus(vInUse map[uint64]bool, hide bool, isStmt bool, extraSize int) {
}
