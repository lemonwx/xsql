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
	"encoding/binary"
)

func (node *Node) hideExtraCols(rs *mysql.Result, data *[]byte, vs map[uint64]uint8) error {
	log.Debugf("[%d] hide extra cols: %v", node.ConnectionId, data)


	if node.IsStmt {
		pos := 1 + (len(rs.Fields)+1+7+2)>>3
		nullMask := (*data)[1:pos]
		if ((nullMask[(0+2)>>3] >> uint((0+2)&7)) & 1) == 1 {
			return errors.New("UNEXPECT VERSION IS NULL")
		}
		start_pos := pos

		for idx := 0; idx < node.ExtraSize; idx += 1 {
			extra := uint64(binary.LittleEndian.Uint64((*data)[pos: pos+8]))
			log.Debugf("[%d] extra col val: %v", node.ConnectionId, extra)
			if _, ok := vs[extra]; ok {
				return errors.New("data in use by another session, pls try again later")
			}
			pos += 8
		}
		*data = append((*data)[0:start_pos], (*data)[start_pos+8*node.ExtraSize:]...)
	} else {

		idx := 1 + (*data)[0]
		res, err := strconv.ParseUint(string((*data)[1:(*data)[0]+1]), 10, 64)
		if err != nil {
			return err
		}

		if _, ok := vs[res]; ok {
			return errors.New("data in use by another session, pls try again later")
		}

		for count := 0; count < node.ExtraSize-1; count += 1 {
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
	}
	return nil
}
