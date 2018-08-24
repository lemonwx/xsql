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

package server

import (
	"fmt"
	"strings"
	"utils"

	"hack"

	"sync"

	"time"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/xsql/node"
	"github.com/lemonwx/xsql/server/version"
	"github.com/lemonwx/xsql/sqlparser"
)

func (conn *MidConn) handleServerNotServe(data []byte) {
	log.Debugf("[%d] midconn is under not serve status, and recv is %s", conn.ConnectionId, string(data[1:]))
	if utils.StringIn(strings.ToLower(hack.String(data[1:])), "rollback", "commit") {
		err := conn.handleTrxFinish(hack.String(data[1:]))
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
		if conn.status == mysql.SERVER_STATUS_IN_TRANS {
			// current is in trx, so need rollback first
			conn.handleTrxFinish("commit")
		}
		conn.status = mysql.SERVER_STATUS_IN_TRANS
	} else {
		conn.status = mysql.SERVER_STATUS_IN_TRANS
	}
}

func (conn *MidConn) handleTrxFinish(sql string) error {
	log.Debugf("[%d] mid conn's status: %v", conn.ConnectionId, conn.status)

	commit := false

	switch {
	case conn.status == mysql.SERVER_STATUS_IN_TRANS:
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
		reset := func() {
			conn.NextVersion = 0
			conn.status = conn.defaultStatus
		}

		log.Debugf("[%d] need exec: %s", conn.ConnectionId, sql)
		if err := conn.clearExecNodes([]byte(sql)); err != nil {
			/*
				roll back with binlog, release conn.NextVersion when finish
			*/
			log.Errorf("[%d] clear exec nodes failed:%v", conn.ConnectionId, err)
			reset()
			return err
		}

		if conn.NextVersion != 0 {
			log.Debugf("[%d] release %v", conn.ConnectionId, conn.NextVersion)
			for {
				// retry until release success, then response to client
				if err := version.ReleaseVersion(conn.NextVersion); err == nil {
					break
				}
				time.Sleep(time.Second * 3)
			}
		}

		reset()
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
		retErrs := make([]error, len(conn.execNodes))
		var wg sync.WaitGroup
		wg.Add(len(conn.execNodes))

		for nodeIdx, back := range conn.execNodes {
			go func(idx int, backNode *node.Node) {
				_, retErr := backNode.Execute(mysql.COM_QUERY, sql)
				if retErr != nil {
					backNode.Close()
				}
				retErrs[idx] = retErr
				conn.pools[idx].PutConn(backNode)
				wg.Done()
			}(nodeIdx, back)
			delete(conn.execNodes, nodeIdx)
		}
		wg.Wait()

		for _, retErr := range retErrs {
			if err, ok := retErr.(error); ok {
				return err
			}
		}
		return nil
	}
}

func (conn *MidConn) handleStmtTrx(data []byte) error {
	conn.handleBegin(false)

	err := conn.handleStmtExecute(data)

	if err != nil {
		return err
	}

	return conn.handleTrxFinish("")
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

	if execErr != nil {
		// exec error, rollback then response
		conn.handleTrxFinish("rollback")
		conn.status = conn.defaultStatus
		return execErr
	}

	handleCommitErr = conn.handleTrxFinish("")
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
		return handleCommitErr
	case execErr != nil && handleCommitErr == nil:
		return execErr
	case execErr != nil && handleCommitErr != nil:
		return fmt.Errorf("%v -- %v", execErr, handleCommitErr)
	default:
		return mysql.NewDefaultError(mysql.MID_ER_UNEXPECTED)
	}
}
