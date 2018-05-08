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

func (conn *MidConn) handleBegin(isBegin bool) {
	if isBegin {
		// 显式 begin / start transaction
		if conn.status[0] == mysql.SERVER_STATUS_IN_TRANS {
			// current is in trx, so need commit first
			conn.handleCommit("commit")
		}
		conn.status = []uint16{mysql.SERVER_STATUS_IN_TRANS, ^mysql.SERVER_STATUS_AUTOCOMMIT}
	} else {
		if conn.status[0] == conn.defaultStatus {
			conn.status[0] = mysql.SERVER_STATUS_IN_TRANS
		}
	}

	conn.executedIdx = make(map[int]uint8)
}

func (conn *MidConn) getExecutedNodeIdx() []int {
	ret := make([]int, 0, len(conn.executedIdx))
	for k, _ := range conn.executedIdx {
		ret = append(ret, k)
	}

	log.Debug(conn.executedIdx, ret)

	return ret
}

func (conn *MidConn) handleCommit(sql string) error {

	commit := false

	switch {
	case conn.status[0] == mysql.SERVER_STATUS_IN_TRANS &&
		conn.status[1] == mysql.SERVER_STATUS_AUTOCOMMIT:
		commit = true
		sql = "commit"
	case sql == "commit":
		commit = true
	case sql == "rollback":
		commit = true
	default:
		commit = false
	}

	if commit {
		log.Debugf("[%d] need commit", conn.ConnectionId)
		conn.status[0] = conn.defaultStatus
		conn.status[1] = conn.defaultStatus

		if conn.NextVersion != 0 {
			log.Debugf("[%d] release %v", conn.ConnectionId, conn.NextVersion)
			err := version.ReleaseVersion(conn.NextVersion)
			if err != nil {
				return err
			}
		}
		conn.NextVersion = 0
		conn.VersionsInUse = nil

		_, err := conn.ExecuteMultiNode(mysql.COM_QUERY, []byte(sql), conn.getExecutedNodeIdx())
		if err != nil {
			return err
		}
	}
	return nil

}

func (conn *MidConn) handleStmtTrx(data []byte) error {
	conn.handleBegin(false)


	err := conn.handleStmtExecute(data)

	if err != nil {
		if conn.status[0] == mysql.SERVER_STATUS_IN_TRANS &&
			conn.status[1] == mysql.SERVER_STATUS_AUTOCOMMIT {
				conn.status[0] = mysql.SERVER_NOT_SERVE
		}

		return err
	}

	return conn.handleCommit("")
}

func (conn *MidConn) handleTrx(stmt sqlparser.Statement, sql string) error {
	conn.handleBegin(false)
	var err error

	switch v := stmt.(type) {
	case *sqlparser.Select:
		err = conn.handleSelect(v, sql)
	case *sqlparser.Insert:
		err = conn.handleInsert(v, sql)
	case *sqlparser.Update:
		err = conn.handleUpdate(v, "")
	case *sqlparser.Delete:
		err = conn.handleDelete(v, sql)
	default:
		err = UNEXPECT_MIDDLE_WARE_ERR
	}

	if err != nil {
		log.Debugf("exec err %v, this trx is in uncommited status", err)
		if conn.status[0] == mysql.SERVER_STATUS_IN_TRANS &&
			conn.status[1] == mysql.SERVER_STATUS_AUTOCOMMIT {
			conn.status[0] = mysql.SERVER_NOT_SERVE
		}
		return err
	}

	return conn.handleCommit("")
}
