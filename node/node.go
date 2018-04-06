/**
 *  author: lim
 *  data  : 18-3-24 下午4:36
 */

package node

import (
	"fmt"
	"net"
	"encoding/binary"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/mysql"
)

type Node struct {
	conn net.Conn
	pkt *mysql.PacketIO


	addr     string
	user     string
	password string
	Db string

	capcbility uint32
	status uint16
}

func NewNode(host string, port int, user, password, db string) *Node {

	node := &Node {
		addr : fmt.Sprintf("%s:%d", host, port),
		user: user,
		password: password,
		Db: db,
	}

	return node
}

func (node *Node) Connect() error {
	return nil
}

func (node *Node) Execute(opt uint8, data []byte) (*mysql.Result, error) {
	if err := node.executeSql(opt, data); err != nil {
		return nil, err
	} else {
		return node.parseResult()
	}

}

func (node *Node) executeSql(opt uint8, data []byte) error{
	node.pkt.Sequence = 0
	length := len(data) + 1
	send := make([]byte, length + 4)
	send[4] = opt
	copy(send[5:], data)

	log.Debugf("send  [%s] to node [%s]", data, node.addr)
	return node.writePacket(send)
}

func (node *Node) parseResult() (*mysql.Result, error) {
	data, err := node.readPacket()
	if err != nil {
		log.Errorf("parse result from %v failed: %v", node.addr, err)
	}
	log.Debugf("recv [%d]-[%s] from node %v", data[0], data[1:], node.addr)

	switch data[0] {
	case mysql.OK_HEADER:
		return node.parseOKPkt(data)
	case mysql.ERR_HEADER:
		return nil, node.parseErrPkt(data)
	}
	return node.readResultset(data, false)
}

func (node *Node) readResultset(data []byte, binary bool) (*mysql.Result, error) {
	ret := &mysql.Result{
		Status: 0,
		AffectedRows: 0,
		Resultset: &mysql.Resultset{},
	}

	count, _, n := mysql.LengthEncodedInt(data)
	if n != len(data) {
		return nil, mysql.ErrMalformPacket
	}

	ret.Fields = make([]*mysql.Field, count)
	ret.FieldNames = make(map[string]int, count)
	if err := node.readResultColumns(ret); err != nil {
		return nil, err
	}
	if err := node.ReadResultRows(ret, binary); err != nil {
		return nil, err
	}
	return ret, nil


}

func (node *Node) readResultColumns(result *mysql.Result) error {
	var i int = 0
	var data []byte
	var err error

	for  {
		data, err = node.pkt.ReadPacket()
		if err != nil {
			return err
		}

		if node.isEOFPacket(data) {
			if node.capcbility & mysql.CLIENT_PROTOCOL_41 > 0 {
				result.Status = binary.LittleEndian.Uint16(data[3:])
				node.status = result.Status
			}

			if i != len(result.Fields) {
				err = mysql.ErrMalformPacket
			}

			return err
		}

		result.Fields[i], err = mysql.FieldData(data).Parse()
		if err != nil {
			return err
		}

		result.FieldNames[string(result.Fields[i].Name)] = i

		i += 1
	}
}

func (node *Node) ReadResultRows(result *mysql.Result, isBinary bool) error {
	var err error
	var data []byte
	for {
		data, err = node.pkt.ReadPacket()

		if err != nil {
			return err
		}

		if node.isEOFPacket(data) {
			if node.capcbility & mysql.CLIENT_PROTOCOL_41 > 0 {
				result.Status = binary.LittleEndian.Uint16(data[3:])
				node.status = result.Status
			}

			break
		}
		result.RowDatas = append(result.RowDatas, data)
	}

	result.Values = make([][]interface{}, len(result.RowDatas))
	for i := range result.Values {
		result.Values[i], err = result.RowDatas[i].Parse(result.Fields, isBinary)

		if err != nil {
			return err
		}
	}
	return nil
}

func (node *Node) isEOFPacket(data []byte) bool {
	return data[0] == mysql.EOF_HEADER && len(data) <= 5
}

func (node *Node) parseOKPkt(data []byte) (*mysql.Result, error){
	var n int
	var pos int = 1

	r := new(mysql.Result)
	r.AffectedRows, _,n = mysql.LengthEncodedInt(data[pos:])
	pos += n
	r.InsertId, _, n = mysql.LengthEncodedInt(data[pos:])
	pos += n

	if node.capcbility & mysql.CLIENT_PROTOCOL_41 > 0 {
		r.Status = binary.LittleEndian.Uint16(data[pos:])
		node.status = r.Status
		pos += 2
	} else if node.capcbility & mysql.CLIENT_TRANSACTIONS > 0 {
		r.Status = binary.LittleEndian.Uint16(data[pos:])
		node.status = r.Status
		pos += 2
	}
	return r, nil
}

func (node *Node) parseErrPkt(data []byte) error {
	e := new(mysql.SqlError)
	pos := 1
	e.Code = binary.LittleEndian.Uint16(data[pos:])
	pos += 2

	if node.capcbility & mysql.CLIENT_PROTOCOL_41 > 0 {
		pos += 1
		e.State = string(data[pos : pos + 5])
		pos += 5
	}
	e.Message = string(data[pos:])
	return e
}


func (node *Node) writePacket(data []byte) error {
	return node.pkt.WritePacket(data)
}

func (node *Node) readPacket()([]byte, error) {
	return node.pkt.ReadPacket()
}