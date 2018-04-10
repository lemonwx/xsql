/**
 *  author: lim
 *  data  : 18-4-6 下午5:15
 */

package midconn

import (
	"github.com/lemonwx/xsql/sqlparser"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/middleware/version"
	"time"
)

func (conn *MidConn) handleShow(stmt *sqlparser.Show, sql string) error {
	// show only send to one node
	rets, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(sql), []int{0,})
	if err != nil {
		log.Errorf("execute in multi node failed: %v", err)
		return err
	}

	return conn.HandleSelRets(rets)

}

func (conn *MidConn) handleSimpleSelect(stmt *sqlparser.SimpleSelect, sql string) error {
	rets, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(sql), nil)
	if err != nil {
		log.Errorf("execute in multi node failed: %v", err)
		return err
	}

	return conn.HandleSelRets(rets)
}

func (conn *MidConn) handleSelect(stmt *sqlparser.Select, sql string) error {

	ts := time.Now()
	var err error
	conn.VersionsInUse, err = version.VersionsInUse()
	if err != nil {
		log.Errorf("[%d] get xa.VersionsInUse failed: %v", err)
		return err
	}
	log.Debugf("[%d] get xa.VersionsInUse: %v", conn.ConnectionId, conn.VersionsInUse)

	conn.setupNodeStatus(conn.VersionsInUse, true)
	defer conn.setupNodeStatus(nil, false)

	rets, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(sql), nil)
	if err != nil {
		return err
	}

	err = conn.HandleSelRets(rets)
	log.Debugf("[%d] handle select cost: %v", conn.ConnectionId, time.Since(ts))

	return err
}

func (conn *MidConn) setupNodeStatus(vInUse [][]byte, hide bool) {
	for _, node := range conn.nodes {
		node.VersionsInUse = vInUse
		node.NeedHide = hide
	}
}