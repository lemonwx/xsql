/**
 *  author: lim
 *  data  : 18-4-6 下午6:27
 */

package server

import (
	"errors"
)

var extraColName string = "version"
var extraColDef string = "version bigint unsigned not null default 0, "

type extraColType uint64

var ROW_DATA_IN_USE_ERR error = errors.New("this row data inuse by another session, pls try again later")
var UNEXPECT_MIDDLE_WARE_ERR error = errors.New("UNEXPECT MIDDLE WARE ERROR")
var UNEXPECT_COMMIT_ERR error = errors.New("UNEXPECT COMMIT ERROR")
var MUST_ROLLBACK_OR_COMMIT_ERR error = errors.New("this midconn is under NOT_SERVE status, please rollback or commit")
var NONE_DB_ERR error = errors.New("no db selected")

// for all back node conn, autocommit = 0
// mid conn status:
// liucheng
const (
	AUTOCOMMIT = iota
	IN_TRANSACTION
)

const (
	updateOrDelete = iota
	Select
)

const (
	errUnsupportedShard    = 10001
	errUnsupportedMultiSet = 10002
	errInternal            = 10003
	errUnexpected          = 10004
	errUnsupportedSql      = 10005

	errShardPanic
	rowsInUseByOtherSession
	sqlOnlySuccessInPartlyOfNode
	unexpectedMiddleErr
)

var MySQLErrName = map[uint16]string{
	errUnsupportedShard:    "unsupported shard for this sql",
	errUnsupportedMultiSet: "unsupported multi set",
	errInternal:            "internal error: %v",
	errUnexpected:          "unexpected midconn error",
	errUnsupportedSql:      "unsupported for this sql",
}
