/**
 *  author: lim
 *  data  : 18-8-28 下午10:43
 */

package server

import "time"

type Stat struct {
	routeT    time.Duration
	sqlparseT time.Duration
	versionT  time.Duration
	chkInuseT time.Duration
	execT     time.Duration
}
