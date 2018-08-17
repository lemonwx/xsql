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
	"fmt"
	"strings"
	"utils"

	"hack"

	"sync"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/middleware/version"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/xsql/node"
	"github.com/lemonwx/xsql/sqlparser"
)

func (conn *MidConn) handleServerNotServe(data []byte) {
	log.Debugf("[%d] midconn is under not serve status, and recv is %s", conn.ConnectionId, string(data[1:]))
	if utils.StringIn(strings.ToLower(hack.String(data[1:])), "rollback", "commit") {
		err := conn.handleCommit(hack.String(data[1:]))
		if err != nil {
			log.Errorf("[%d] handle commit faild: %v", conn.ConnectionId, err)
			if err := conn.cli.WriteError(err); err != nil {
				log.Errorf("[%d] send err to cli failed: %v", conn.ConnectionId, err)
			}
			conn.cli.SetPktSeq(0)
		} else {
			log.Debugf("[%d] handle commit successed", conn.ConnectionId)
			if err := conn.cli.WriteOK(nil); err != nil {
				log.Errorf("[%d] send ok to cli failed: %v", conn.ConnectionId, err)
			}
			conn.cli.SetPktSeq(0)
		}
	} else {
		if err := conn.cli.WriteError(MUST_ROLLBACK_OR_COMMIT_ERR); err != nil {
			log.Errorf("[%d] send err to cli failed: %v", conn.ConnectionId, err)
		}
		conn.cli.SetPktSeq(0)
	}
}

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
}

func (conn *MidConn) handleCommit(sql string) error {
	log.Debugf("[%d] mid conn's status: %v", conn.ConnectionId, conn.status)

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
		log.Debugf("[%d] need exec: %s", conn.ConnectionId, sql)

		if conn.NextVersion != 0 {
			log.Debugf("[%d] release %v", conn.ConnectionId, conn.NextVersion)
			err := version.ReleaseVersion(conn.NextVersion)
			if err != nil {
				log.Errorf("[%d] release version failed: %v", conn.ConnectionId, err)
				return mysql.NewDefaultError(mysql.MID_ER_RELEASE_VERSION_FAILED)
			}
		}
		conn.NextVersion = 0
		conn.VersionsInUse = nil

		if commitErr := conn.clearExecNodes([]byte(sql)); commitErr != nil {
			return mysql.NewDefaultError(mysql.MID_ER_EXEC_COMMIT_ROLLBACK_FAILED)
		}

		conn.status[0] = conn.defaultStatus
		conn.status[1] = conn.defaultStatus
	}
	return nil
}

func (conn *MidConn) clearExecNodes(sql []byte) error {
	if len(conn.execNodes) == 1 {
		for nodeIdx, back := range conn.execNodes {
			if _, err := back.Execute(mysql.COM_QUERY, sql); err != nil {
				back.Close()
				conn.pools[nodeIdx].PutConn(back)
				delete(conn.execNodes, nodeIdx)
				return err
			}
			conn.pools[nodeIdx].PutConn(back)
			delete(conn.execNodes, nodeIdx)
		}
		return nil
	} else {
		var retErr error
		var wg sync.WaitGroup
		wg.Add(len(conn.execNodes))

		for nodeIdx, back := range conn.execNodes {
			go func(idx int, backNode *node.Node) {
				_, retErr = backNode.Execute(mysql.COM_QUERY, sql)
				if retErr != nil {
					backNode.Close()
				}
				conn.pools[nodeIdx].PutConn(backNode)
				wg.Done()
			}(nodeIdx, back)
			delete(conn.execNodes, nodeIdx)
		}
		wg.Wait()

		return retErr
	}
}

func (conn *MidConn) handleStmtTrx(data []byte) error {
	conn.handleBegin(false)

	err := conn.handleStmtExecute(data)

	if err != nil {
		if conn.status[0] == mysql.SERVER_STATUS_IN_TRANS {
			conn.status[0] = mysql.SERVER_NOT_SERVE
		}

		return err
	}

	return conn.handleCommit("")
}

func (conn *MidConn) handleTrx(stmt sqlparser.Statement, sql string) error {
	conn.handleBegin(false)
	var execErr error
	var handleCommitErr error

	var rets []*mysql.Result
	var isSelect bool

	switch v := stmt.(type) {
	case *sqlparser.Select:
		isSelect = true
		rets, execErr = conn.handleSelect(v)
	case *sqlparser.Insert:
		isSelect = false
		rets, execErr = conn.handleInsert(v, sql)
	case *sqlparser.Update:
		isSelect = false
		rets, execErr = conn.handleUpdate(v, "")
	case *sqlparser.Delete:
		isSelect = false
		rets, execErr = conn.handleDelete(v, sql)
	default:
		isSelect = false
		execErr = mysql.NewDefaultError(mysql.MID_ER_UNEXPECTED)
	}

	handleCommitErr = conn.handleCommit("")

	err := conn.myHandleErr(execErr, handleCommitErr)
	if err != nil {
		return err
	} else {
		if isSelect {
			return conn.HandleSelRets(rets)
		} else {
			return conn.HandleExecRets(rets)
		}
	}
}

func (conn *MidConn) myHandleErr(execErr, handleCommitErr error) error {
	switch {
	case execErr == nil && handleCommitErr == nil:
		return nil
	case execErr == nil && handleCommitErr != nil:
		if conn.status[0] == mysql.SERVER_STATUS_IN_TRANS && conn.status[1] == mysql.SERVER_STATUS_AUTOCOMMIT {
			conn.status[0] = mysql.SERVER_NOT_SERVE
		}
		return handleCommitErr
	case execErr != nil && handleCommitErr == nil:
		return execErr
	case execErr != nil && handleCommitErr != nil:
		if conn.status[0] == mysql.SERVER_STATUS_IN_TRANS && conn.status[1] == mysql.SERVER_STATUS_AUTOCOMMIT {
			conn.status[0] = mysql.SERVER_NOT_SERVE
		}
		return fmt.Errorf("%v -- %v", execErr, handleCommitErr)
	default:
		return mysql.NewDefaultError(mysql.MID_ER_UNEXPECTED)
	}
}
