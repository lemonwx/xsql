/**
 *  author: lim
 *  data  : 18-5-30 下午10:33
 */

package node

import (
	"errors"
	"strconv"
	//"math"

	"encoding/binary"
	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/mysql"
)

var UNEXPECTED_NODE_ERR = errors.New("UNEXPECTED NODE ERROR")

func (node *Node) node2cliNullMask(data *[]byte, fieldCount int) ([]byte, error) {

	cliSize := (fieldCount +7 + 2) >> 3
	nodePos := 1 + ((fieldCount + node.ExtraSize + 7 + 2) >> 3)

	if nodePos >= len(*data) {
		log.Errorf("[%d] unexpected node err: node pos: %d >= len(*data): %d", node.ConnectionId, nodePos, len(*data))
		return nil, UNEXPECTED_NODE_ERR
	}

	nodeNullMask := (*data)[1:nodePos]
	cliNullMask := make([]byte, cliSize)

	for idx := 0; idx < fieldCount; idx += 1 {
		nodeidx := idx + node.ExtraSize
		if (nodeidx + 2 ) >> 3 >= len(nodeNullMask){
			log.Errorf("[%d] unexpected node err: idx: %d >= len(node null mask): %d", node.ConnectionId,
				(nodeidx + 2) >> 3 , len(nodeNullMask))
			return nil, UNEXPECTED_NODE_ERR
		}
		if ((nodeNullMask[(nodeidx+2)>>3] >> uint((nodeidx+2)&7)) & 1) == 1 {
			cliNullMask[(idx + 2) >> 3] += 1 << uint((idx + 2) % 8)
		}
	}

	log.Debugf("[%d] 2 cli null mask : %v, len: %d, %d, %d", node.ConnectionId, cliNullMask, cliSize, nodePos, fieldCount)
	return cliNullMask, nil
}

func (node *Node) hideExtraCols(rs *mysql.Result, data *[]byte, vs map[uint64]uint8) error {

	if node.IsStmt {
		pos := 1 + ((len(rs.Fields)+ node.ExtraSize + 7 + 2) >> 3)
		nullMask := (*data)[1:pos]

		for idx := 0; idx < node.ExtraSize; idx += 1 {
			if ((nullMask[(idx+2)>>3] >> uint((idx+2)&7)) & 1) == 1 {
				log.Errorf("[%d] unexpected node err: version parsed from ret is nil", node.ConnectionId)
				return UNEXPECTED_NODE_ERR
			}
			extra := uint64(binary.LittleEndian.Uint64((*data)[pos : pos+8]))
			log.Debugf("[%d] extra col val: %v", node.ConnectionId, extra)
			if _, ok := vs[extra]; ok {
				return errors.New("data in use by another session, pls try again later")
			}
			pos += 8
		}
		mask, err := node.node2cliNullMask(data, len(rs.Fields))
		if err != nil {
			return err
		}
		mask = append((*data)[:1], mask...)
		*data = append(mask, (*data)[pos:]...)
	} else {
		idx := 1 + (*data)[0]
		vStr := string((*data)[1:(*data)[0]+1])
		res, err := strconv.ParseUint(vStr, 10, 64)
		if err != nil {
			log.Errorf("[%d] ParseUint from %v failed: %v", vStr, err)
			return mysql.NewDefaultError(mysql.MID_ER_HIDE_EXTRA_FAILED)
		}

		if _, ok := vs[res]; ok {
			err = mysql.NewDefaultError(mysql.MID_ER_ROWS_IN_USE_BY_OTHER_SESSION)
			log.Errorf("[%d] hide extra col failed: %v", node.ConnectionId, err)
			return err
		}

		for count := 0; count < node.ExtraSize-1; count += 1 {
			s := idx + 1
			e := s + (*data)[idx]

			vStr :=  string((*data)[s:e])
			res, err := strconv.ParseUint(vStr, 10, 64)
			if err != nil {
				log.Errorf("[%d] ParseUint from %v failed: %v", vStr, err)
				return mysql.NewDefaultError(mysql.MID_ER_HIDE_EXTRA_FAILED)
			}
			if _, ok := vs[res]; ok {
				err = mysql.NewDefaultError(mysql.MID_ER_ROWS_IN_USE_BY_OTHER_SESSION)
				log.Errorf("[%d] hide extra col failed: %v", node.ConnectionId, err)
				return err
			}

			idx = (*data)[idx] + idx + 1
		}
		*data = (*data)[idx:]
	}
	return nil
}
