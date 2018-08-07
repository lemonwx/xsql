/**
 *  author: lim
 *  data  : 18-4-6 下午6:27
 */

package midconn

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
	ROWS_IN_USE_BY_OTHER_SESSION = 10001
	SQL_ONLY_SUCCESS_IN_PARTLY_OF_NODE
	UNEXPECTED_MIDDLE_ERR
)
