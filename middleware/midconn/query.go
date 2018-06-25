/**
 *  author: lim
 *  data  : 18-4-6 下午5:15
 */

package midconn

import (
	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/middleware/meta"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/xsql/sqlparser"
	"strconv"
)

func (conn *MidConn) handleShow(stmt *sqlparser.Show, sql string) error {
	// show only send to one node
	rets, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(sql), []int{0})
	if err != nil {
		log.Errorf("execute in multi node failed: %v", err)
		return err
	}

	return conn.HandleSelRets(rets)

}

func (conn *MidConn) handleSimpleSelect(stmt *sqlparser.SimpleSelect, sql string) error {
	log.Debugf("[%d] handle simple select", conn.ConnectionId)
	rets, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(sql), meta.GetFullNodeIdxs())
	if err != nil {
		log.Errorf("execute in multi node failed: %v", err)
		return err
	}

	return conn.HandleSelRets(rets)
}

func (conn *MidConn) handleSelect(stmt *sqlparser.Select, sql string) ([]*mysql.Result, error) {

	var err error

	if p, err := conn.getPlan(stmt); err != nil {
		return nil, err
	} else {
		log.Debugf("[%d] get shard list's: %v", conn.ConnectionId, p.ShardList)
		if len(p.ShardList) == 0 {
			r := conn.newEmptyResultset(stmt)
			return []*mysql.Result{&mysql.Result{Resultset: r}}, nil
		}
		conn.nodeIdx = p.ShardList
	}

	if err = conn.getVInUse(); err != nil {
		return nil, err
	}

	conn.setupNodeStatus(conn.VersionsInUse, true, false, len(stmt.ExtraCols))
	defer conn.setupNodeStatus(nil, false, false, 0)

	newSql := sqlparser.String(stmt)
	rets, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(newSql), conn.nodeIdx)

	if err == nil {
		if stmt.Limit != nil {
			rets, err = conn.handleLimit(rets, stmt.Limit)
		}

		// group by [having] , order by, distinct
	}

	return rets, err
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
			if allCount + tmp >= limitCount {

				rets[idx].RowDatas = rets[idx].RowDatas[:limitCount - allCount]
				return rets[:idx + 1], nil
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
