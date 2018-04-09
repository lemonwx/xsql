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