/**
 *  author: lim
 *  data  : 18-5-30 下午10:33
 */

package node

import (
	"strconv"
	"errors"

	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/log"
)

func (node *Node) hideExtraCols(rs *mysql.Result, data *[]byte, vs map[uint64]uint8) error {
	log.Debugf("[%d] hide extra cols: %v", node.ConnectionId, data)

	idx := 1 + (*data)[0]
	res, err := strconv.ParseUint(string((*data)[1:(*data)[0]+1]), 10, 64)
	if err != nil {
		return err
	}

	if _, ok := vs[res]; ok {
		return errors.New("data in use by another session, pls try again later")
	}

	for count := 0; count < node.ExtraSize - 1; count += 1 {
		s := idx + 1
		e := s + (*data)[idx]

		res, err := strconv.ParseUint(string((*data)[s:e]), 10, 64)
		if err != nil {
			return err
		}
		if _, ok := vs[res]; ok {
			return errors.New("data in use by another session, pls try again later")
		}

		idx = (*data)[idx] + idx + 1
	}
	*data = (*data)[idx:]
	log.Debug(idx, *data)
	return nil
}
