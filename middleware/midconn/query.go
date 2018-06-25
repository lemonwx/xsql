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
			if stmt.Limit.Offset != nil {
				log.Errorf("[%d] offset : %v not nil, not support this sql now", conn.ConnectionId, stmt.Limit.Offset)
				return nil, mysql.NewDefaultError(mysql.MID_ER_UNSUPPORTED_SQL)
			}
			log.Debugf("[%d] offset: %v, rows count: %d", conn.ConnectionId, stmt.Limit.Offset, stmt.Limit.Rowcount)

			count, err := strconv.ParseUint(string(stmt.Limit.Rowcount.(sqlparser.NumVal)), 10, 64)

			if err != nil {
				log.Errorf("[%d] parse limit count failed: %v", conn.ConnectionId, err)
				return nil, err
			}

			allCount := 0

			r := &mysql.Result{
				Status:       0,
				AffectedRows: 0,
				Resultset:    &mysql.Resultset{},
			}

			r.Status = rets[0].Status
			r.InsertId = rets[0].InsertId
			r.AffectedRows = rets[0].AffectedRows
			r.Fields = rets[0].Fields
			r.FieldNames = rets[0].FieldNames
			r.RowDatas = make([]mysql.RowData, count)

			s := 0
			for _, ret := range rets {
				tmp := len(ret.RowDatas)
				if allCount + tmp > int(count) {
					tmp = int(count) - allCount
					copy(r.RowDatas[s:], ret.RowDatas[:tmp])
					return []*mysql.Result{r}, nil
				}
				copy(r.RowDatas[s:], ret.RowDatas[:tmp])
				s += tmp
				allCount += tmp
			}
		}
	}
	return rets, err
}

func (conn *MidConn) setupNodeStatus(vInUse map[uint64]byte, hide bool, isStmt bool, extraSize int) {
	for idx, _ := range conn.nodes {
		conn.nodes[idx].VersionsInUse = vInUse
		conn.nodes[idx].NeedHide = hide
		conn.nodes[idx].IsStmt = isStmt
		conn.nodes[idx].ExtraSize = extraSize
	}
}
