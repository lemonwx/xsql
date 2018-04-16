/**
 *  author: lim
 *  data  : 18-4-12 下午12:29
 */

// almost all real sql is in transaction
// 	so let:
// 		backend node conn, default is not autocommit,
// 			[[[ --- for decrease exchange with backend node --- ]]]
// 		proxy default is autocommit,
// 			[[[ --- for mysqld used to --- ]]]

package midconn

import (
	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/middleware/version"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/xsql/sqlparser"
)

func (conn *MidConn) handleBegin() {

	if conn.status[0] == conn.defaultStatus {
		conn.status[0] = mysql.SERVER_STATUS_IN_TRANS
	}

	log.Debug(conn.status)
}

func (conn *MidConn) handleCommit(nodeIdx []int, sql string) error {

	commit := false

	switch {
	case conn.status[0] == mysql.SERVER_STATUS_IN_TRANS &&
		conn.status[1] == mysql.SERVER_STATUS_AUTOCOMMIT:
		commit = true
	case sql == "commit":
		commit = true
	default:
		commit = false
	}

	if commit {
		log.Debugf("[%d] need commit", conn.ConnectionId)

		_, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte("commit"), nil)
		if err != nil {
			return err
		}
		conn.status[0] = conn.defaultStatus
		conn.status[1] = conn.defaultStatus

		if conn.NextVersion != nil {
			version.ReleaseVersion(conn.NextVersion)
		}
		conn.NextVersion = nil
		conn.VersionsInUse = nil
	}
	return nil

}

func (conn *MidConn) handleTrx(stmt sqlparser.Statement, sql string) error {
	conn.handleBegin()
	var err error

	switch v := stmt.(type) {
	case *sqlparser.Select:
		err = conn.handleSelect(v, sql)
	case *sqlparser.Insert:
		err = conn.handleInsert(v, sql)
	case *sqlparser.Update:
		err = conn.handleUpdate(v, sql)
	case *sqlparser.Delete:
		err = conn.handleDelete(v, sql)
	default:
		err = UNEXPECT_MIDDLE_WARE_ERR
	}

	if err != nil {
		log.Debugf("exec errr %v", err)
		return err
	}
	err = conn.handleCommit(nil, "")
	return err
}
