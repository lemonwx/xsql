/**
 *  author: lim
 *  data  : 18-3-24 下午4:36
 */

package node

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"strconv"

	"github.com/lemonwx/log"
	"github.com/lemonwx/xsql/mysql"
)

type Node struct {
	conn net.Conn
	pkt  *mysql.PacketIO

	addr     string
	user     string
	password string
	Db       string

	ConnectionId uint32
	capability   uint32
	status       uint16
	collation    mysql.CollationId
	charset      string
	salt         []byte

	VersionsInUse map[uint64]uint8
	NextVersion   uint64
	NeedHide      bool
	IsStmt        bool
}

func (node *Node) String() string {
	return node.addr
}

func NewNode(host string, port int, user, password, db string, connid uint32) *Node {

	node := &Node{
		addr:         fmt.Sprintf("%s:%d", host, port),
		user:         user,
		password:     password,
		Db:           db,
		ConnectionId: connid,
	}

	return node
}

func (node *Node) Connect() error {
	conn, err := net.Dial("tcp", node.addr)
	if err != nil {
		log.Errorf("dial to backend mysqld [%v] failed %v", node, err)
	}

	tcpConn := conn.(*net.TCPConn)
	tcpConn.SetNoDelay(false)
	node.conn = tcpConn
	node.pkt = mysql.NewPacketIO(tcpConn)

	if err := node.readInitialHandshake(); err != nil {
		log.Errorf("read init handshake from mysqld [%v] failed: %v", node.addr, err)
		node.Close()
		return err
	}

	if err := node.writeAuthHandshake(); err != nil {
		log.Errorf("write auth handshake from mysqld [%v] failed: %v", node.addr, err)
		node.Close()
		return err
	}

	if _, err := node.readOK(); err != nil {
		log.Errorf("hand shake with mysqld [%v], read ok failed: %v", node.addr, err)
		return err
	}

	_, err = node.Execute(mysql.COM_QUERY, []byte("set autocommit = 0"))
	if err != nil {
		log.Errorf("[%d] execute set autocommit = 0 failed: %v", node.ConnectionId, err)
		return err
	}

	return nil
}

func (node *Node) readInitialHandshake() error {
	data, err := node.readPacket()
	if err != nil {
		return err
	}

	if data[0] == mysql.ERR_HEADER {
		return errors.New("read initial handshake error")
	}

	if data[0] < mysql.MinProtocolVersion {
		return fmt.Errorf("invalid protocol version %d, must >= 10", data[0])
	}

	//skip mysql version and connection id
	//mysql version end with 0x00
	//connection id length is 4
	pos := 1 + bytes.IndexByte(data[1:], 0x00) + 1 + 4

	node.salt = append(node.salt, data[pos:pos+8]...)

	//skip filter
	pos += 8 + 1

	//capability lower 2 bytes
	node.capability = uint32(binary.LittleEndian.Uint16(data[pos : pos+2]))

	pos += 2

	if len(data) > pos {
		//skip server charset
		//c.charset = data[pos]
		pos += 1

		node.status = binary.LittleEndian.Uint16(data[pos : pos+2])
		pos += 2

		node.capability = uint32(binary.LittleEndian.Uint16(data[pos:pos+2]))<<16 | node.capability

		pos += 2

		//skip auth data len or [00]
		//skip reserved (all [00])
		pos += 10 + 1

		// The documentation is ambiguous about the length.
		// The official Python library uses the fixed length 12
		// mysql-proxy also use 12
		// which is not documented but seems to work.
		node.salt = append(node.salt, data[pos:pos+12]...)
	}

	return nil
}

func (node *Node) writeAuthHandshake() error {
	// Adjust client capability flags based on server support
	capability := mysql.CLIENT_PROTOCOL_41 | mysql.CLIENT_SECURE_CONNECTION |
		mysql.CLIENT_LONG_PASSWORD | mysql.CLIENT_TRANSACTIONS | mysql.CLIENT_LONG_FLAG

	capability &= node.capability

	//packet length
	//capbility 4
	//max-packet size 4
	//charset 1
	//reserved all[0] 23
	length := 4 + 4 + 1 + 23

	//username
	length += len(node.user) + 1

	//we only support secure connection
	auth := mysql.CalcPassword(node.salt, []byte(node.password))

	length += 1 + len(auth)

	if len(node.Db) > 0 {
		capability |= mysql.CLIENT_CONNECT_WITH_DB

		length += len(node.Db) + 1
	}

	node.capability = capability

	data := make([]byte, length+4)

	//capability [32 bit]
	data[4] = byte(capability)
	data[5] = byte(capability >> 8)
	data[6] = byte(capability >> 16)
	data[7] = byte(capability >> 24)

	//MaxPacketSize [32 bit] (none)
	//data[8] = 0x00
	//data[9] = 0x00
	//data[10] = 0x00
	//data[11] = 0x00

	//Charset [1 byte]
	data[12] = byte(node.collation)

	//Filler [23 bytes] (all 0x00)
	pos := 13 + 23

	//User [null terminated string]
	if len(node.user) > 0 {
		pos += copy(data[pos:], node.user)
	}
	//data[pos] = 0x00
	pos++

	// auth [length encoded integer]
	data[pos] = byte(len(auth))
	pos += 1 + copy(data[pos+1:], auth)

	// db [null terminated string]
	if len(node.Db) > 0 {
		pos += copy(data[pos:], node.Db)
		//data[pos] = 0x00
	}

	return node.writePacket(data)
}

func (node *Node) readOK() (*mysql.Result, error) {
	data, err := node.readPacket()
	if err != nil {
		return nil, err
	}

	if data[0] == mysql.OK_HEADER {
		return node.parseOKPkt(data)
	} else if data[0] == mysql.ERR_HEADER {
		return nil, node.parseErrPkt(data)
	} else {
		return nil, errors.New("invalid ok packet")
	}
}

func (node *Node) readPrepareResultPacket(id *uint32, columnCount *uint16, paramCount *int) error {
	data, err := node.pkt.ReadPacket()
	if err == nil {
		// packet indicator [1 byte]

		if data[0] == mysql.ERR_HEADER {
			return node.parseErrPkt(data)
		} else if data[0] != mysql.OK_HEADER {
			return mysql.ErrMalformPacket
		}

		if data[0] != mysql.OK_HEADER {
			return err
		}

		// statement id [4 bytes]
		*id = binary.LittleEndian.Uint32(data[1:5])

		// Column count [16 bit uint]
		*columnCount = binary.LittleEndian.Uint16(data[5:7])

		// Param count [16 bit uint]
		*paramCount = int(binary.LittleEndian.Uint16(data[7:9]))

		if *paramCount > 0 {
			if err := node.pkt.ReadUntilEOF(); err != nil {
				return err
			}
		}

		if *columnCount > 0 {
			if err := node.pkt.ReadUntilEOF(); err != nil {
				return err
			}
		}

		// Reserved [8 bit]

		// Warning count [16 bit uint]

		return nil
	}
	return err
}

func (node *Node) ExecutePrepare(data []byte, id *uint32, columntCount *uint16, paramCount *int) error {
	if err := node.executeSql(mysql.COM_STMT_PREPARE, data); err != nil {
		return err
	} else {
		return node.readPrepareResultPacket(id, columntCount, paramCount)
	}
}

func (node *Node) Execute(opt uint8, data []byte) (*mysql.Result, error) {
	if err := node.executeSql(opt, data); err != nil {
		return nil, err
	} else {
		return node.parseResult()
	}

}

func (node *Node) ExecuteSql(opt uint8, data []byte) error {
	return node.executeSql(opt, data)
}

func (node *Node) executeSql(opt uint8, data []byte) error {
	node.pkt.Sequence = 0
	length := len(data) + 1
	send := make([]byte, length+4)
	send[4] = opt
	copy(send[5:], data)

	log.Debugf("[%d] send [%d:%s] to node [%s]", node.ConnectionId, send[4], send[5:], node.addr)
	return node.writePacket(send)
}

func (node *Node) writeCommandStrStr(command byte, arg1 string, arg2 string) error {
	node.pkt.Sequence = 0

	data := make([]byte, 4, 6+len(arg1)+len(arg2))

	data = append(data, command)
	data = append(data, arg1...)
	data = append(data, 0)
	data = append(data, arg2...)

	return node.writePacket(data)
}

func (node *Node) ParseResult() (*mysql.Result, error) {
	return node.parseResult()
}

func (node *Node) parseResult() (*mysql.Result, error) {
	data, err := node.readPacket()
	if err != nil {
		log.Errorf("[%d] parse result from %v failed: %v", node.ConnectionId, node.addr, err)
		return nil, err
	}
	log.Debugf("[%d] recv [%d]-[%s] from node %v", node.ConnectionId, data[0], data[1:], node.addr)

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
		Status:       0,
		AffectedRows: 0,
		Resultset:    &mysql.Resultset{},
	}

	count, _, n := mysql.LengthEncodedInt(data)
	if node.NeedHide {
		log.Debugf("[%d] node [%v] read result need to hide extra col", node.ConnectionId, node.addr)
		count -= 1
	}

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

	for idx := 0; ; idx += 1 {
		data, err = node.pkt.ReadPacket()
		if err != nil {
			return err
		}

		if node.isEOFPacket(data) {
			if node.capability&mysql.CLIENT_PROTOCOL_41 > 0 {
				result.Status = binary.LittleEndian.Uint16(data[3:])
				node.status = result.Status
			}

			if i != len(result.Fields) {
				err = mysql.ErrMalformPacket
			}

			return err
		}

		if node.NeedHide && idx == 0 {
			continue
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
	var retErr error
	var err error
	var data []byte
	// pre row's version value
	//var preRowV uint64 = 0

	for {
		data, err = node.pkt.ReadPacket()
		if err != nil {
			return err
		}

		if node.isEOFPacket(data) {
			if node.capability&mysql.CLIENT_PROTOCOL_41 > 0 {
				result.Status = binary.LittleEndian.Uint16(data[3:])
				node.status = result.Status
			}
			break
		}
		if node.NeedHide {
			node.hideExtraCols(result, &data, node.VersionsInUse)
			version, err := node.calcVersion(result, &data)
			if err != nil {
				retErr = err
			} else if _, ok := node.VersionsInUse[version]; ok {
				retErr = errors.New("data in use by another session, pls try again later")
			}
		}
		result.RowDatas = append(result.RowDatas, data)
	}
	/*
		no affect to send resultset to mysql cli
		result.Values = make([][]interface{}, len(result.RowDatas))
		for i := range result.Values {
			result.Values[i], err = result.RowDatas[i].Parse(result.Fields, isBinary)

			if err != nil {
				return err
			}
		}
	*/
	return retErr
}

func (node *Node) calcVersion(rs *mysql.Result, data *[]byte) (uint64, error) {
	if node.IsStmt {
		pos := 1 + (len(rs.Fields)+1+7+2)>>3
		nullMask := (*data)[1:pos]
		if ((nullMask[(0+2)>>3] >> uint((0+2)&7)) & 1) == 1 {
			return 0, errors.New("UNEXPECT VERSION IS NULL")
		}
		*data = append((*data)[0:pos], (*data)[pos+8:]...)
		return uint64(binary.LittleEndian.Uint64((*data)[pos : pos+8])), nil
	} else {
		res, err := strconv.ParseUint(string((*data)[1:(*data)[0]+1]), 10, 64)
		*data = (*data)[(*data)[0]+1:]
		return res, err

	}

}

func (node *Node) isEOFPacket(data []byte) bool {
	return node.pkt.IsEOFPacket(data)
}

func (node *Node) parseOKPkt(data []byte) (*mysql.Result, error) {
	var n int
	var pos int = 1

	r := new(mysql.Result)
	r.AffectedRows, _, n = mysql.LengthEncodedInt(data[pos:])
	pos += n
	r.InsertId, _, n = mysql.LengthEncodedInt(data[pos:])
	pos += n

	if node.capability&mysql.CLIENT_PROTOCOL_41 > 0 {
		r.Status = binary.LittleEndian.Uint16(data[pos:])
		node.status = r.Status
		pos += 2
	} else if node.capability&mysql.CLIENT_TRANSACTIONS > 0 {
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

	if node.capability&mysql.CLIENT_PROTOCOL_41 > 0 {
		pos += 1
		e.State = string(data[pos : pos+5])
		pos += 5
	}
	e.Message = string(data[pos:])
	return e
}

func (node *Node) FieldList(table string, wildcard string) ([]*mysql.Field, error) {
	if err := node.writeCommandStrStr(mysql.COM_FIELD_LIST, table, wildcard); err != nil {
		return nil, err
	}

	data, err := node.readPacket()
	if err != nil {
		return nil, err
	}

	fs := make([]*mysql.Field, 0, 4)
	var f *mysql.Field
	if data[0] == mysql.ERR_HEADER {
		return nil, node.parseErrPkt(data)
	} else {
		for {
			if data, err = node.readPacket(); err != nil {
				return nil, err
			}

			// EOF Packet
			if node.isEOFPacket(data) {
				return fs, nil
			}

			if f, err = mysql.FieldData(data).Parse(); err != nil {
				return nil, err
			}
			fs = append(fs, f)
		}
	}
	return nil, fmt.Errorf("field list error")
}

func (node *Node) WritePacket(data []byte) error {
	return node.pkt.WritePacket(data)
}

func (node *Node) writePacket(data []byte) error {
	return node.pkt.WritePacket(data)
}

func (node *Node) readPacket() ([]byte, error) {
	return node.pkt.ReadPacket()
}

func (node *Node) Close() {
	if node.conn != nil {
		node.conn.Close()
		node.conn = nil
	}
}

func (node *Node) SetPktSeq(sz uint8) {
	node.pkt.Sequence = sz
}
