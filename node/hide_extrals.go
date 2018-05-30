/**
 *  author: lim
 *  data  : 18-5-30 下午10:33
 */

package node

import (
	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/log"
)

func (node *Node) hideExtraCols(rs *mysql.Result, data *[]byte, vs map[uint64]uint8) {
	log.Debugf("[%d] hide extra cols: %v", node.ConnectionId, data)

	idx := 1 + (*data)[0]
	for count := 0; count < 3 - 1; count += 1 {
		idx = (*data)[idx] + idx + 1
	}
	log.Debug(idx, (*data)[idx:])
}
