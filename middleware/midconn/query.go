/**
 *  author: lim
 *  data  : 18-4-6 下午5:15
 */

package midconn

import (
	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/xsql/sqlparser"
	"github.com/lemonwx/xsql/middleware/meta"
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
	rets, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(sql), meta.GetFullNodeIdxs())
	if err != nil {
		log.Errorf("execute in multi node failed: %v", err)
		return err
	}

	return conn.HandleSelRets(rets)
}

func (conn *MidConn) handleSelect(stmt *sqlparser.Select, sql string) error {

	var hide bool = true
	var err error

	if err = conn.getNodeIdxs(stmt); err != nil {
		return err
	} else if conn.nodeIdx == nil {
		return conn.writeResultset(conn.status[0], conn.newEmptyResultset(stmt))
	}

	if err = conn.getVInUse(); err != nil {
		return err
	}

	// judge extra col hide or not
	if _, ok := stmt.SelectExprs[0].(*sqlparser.StarExpr); ok {
		log.Debugf("[%d] select * not need to convert", conn.ConnectionId)
	} else if expr, ok := stmt.SelectExprs[0].(*sqlparser.NonStarExpr); ok {
		colName := sqlparser.String(expr)
		log.Debugf("[%d] select %s, expr, add extra col add first", conn.ConnectionId, colName)
		if colName != extraColName {
			tmp := make(sqlparser.SelectExprs, len(stmt.SelectExprs)+1)
			copy(tmp[1:], stmt.SelectExprs[:])
			tmp[0] = &sqlparser.NonStarExpr{
				Expr: &sqlparser.ColName{Name: []byte(extraColName)},
			}
			stmt.SelectExprs = tmp
		} else {
			log.Debugf("[%d] select contains extra col, set hide = false", conn.ConnectionId)
			hide = false
		}
	}

	if hide {
		conn.setupNodeStatus(conn.VersionsInUse, true)
		defer conn.setupNodeStatus(nil, false)
	}

	newSql := sqlparser.String(stmt)
	rets, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(newSql), conn.nodeIdx)
	if err != nil {
		return err
	}

	return conn.HandleSelRets(rets)
}

func (conn *MidConn) setupNodeStatus(vInUse map[uint64]byte, hide bool) {
	for _, node := range conn.nodes {
		node.VersionsInUse = vInUse
		node.NeedHide = hide
	}
}
