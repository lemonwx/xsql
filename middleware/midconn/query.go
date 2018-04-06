/**
 *  author: lim
 *  data  : 18-4-6 下午5:15
 */

package midconn

import (
	"github.com/lemonwx/xsql/sqlparser"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/log"
)

func handleShow() {

}

func (conn *MidConn)handleSimpleSelect(stmt *sqlparser.SimpleSelect, sql string) error {
	rets, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(sql), nil)
	if err != nil {
		log.Errorf("execute in multi node failed: %v", err)
		return err
	}

	return conn.HandleSelRets(rets)
}

func (conn *MidConn)handleSelect(stmt *sqlparser.Select, sql string) error {
	rets, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(sql), nil)
	if err != nil {
		return err
	}

	return conn.HandleSelRets(rets)
}