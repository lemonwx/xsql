/**
 *  author: lim
 *  data  : 18-4-11 下午11:06
 */

package server

import (
	"strings"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/errors"
	"github.com/lemonwx/xsql/meta"
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/xsql/sqlparser"
)

func (conn *MidConn) handleSet(stmt *sqlparser.Set, sql string) error {
	log.Debugf("[%d] handle set: %v", conn.ConnectionId, stmt.Exprs)

	// default
	if len(stmt.Exprs) != 2 {
		return newMySQLErr(errUnsupportedMultiSet)
	}

	var cvtExprs sqlparser.UpdateExprs
	for _, expr := range stmt.Exprs {
		field := string(expr.Name.Name)
		if string(field) == "autocommit" {
			if val, ok := expr.Expr.(sqlparser.NumVal); ok {
				if string(val) == "1" {
					conn.status = mysql.SERVER_STATUS_AUTOCOMMIT
				} else if string(val) == "0" {
					conn.status = mysql.SERVER_STATUS_IN_TRANS
				} else {
					return errors.New2("unsupported autocommit value")
				}
			} else {
				return errors.New2("unsupported autocommit value")
			}
		} else {
			cvtExprs = append(cvtExprs, expr)
		}
	}

	if !strings.Contains(strings.ToLower(sql), "autocommit") {

		rets, err := conn.ExecuteOnNodePool([]byte(sql), meta.GetFullNodeIdxs())
		if err != nil {
			return err
		}
		return conn.HandleExecRets(rets)
	}

	/*
		expr := stmt.Exprs[0]

		if v, ok := expr.Expr.(sqlparser.NumVal); ok {
			log.Debugf("[%d], set num %v", conn.ConnectionId, v)
		}

		if v, ok := expr.Expr.(sqlparser.StrVal); ok {
			log.Debugf("[%d], set str d g%v", conn.ConnectionId, v)
		}
	*/

	/*
		if on :
			default = on
		if off
			default = off
	*/

	return conn.cli.WriteOK(nil)

}
