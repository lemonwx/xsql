/**
 *  author: lim
 *  data  : 18-3-24 下午3:22
 */

package client

import (
	"bytes"
	"encoding/binary"
	"net"

	"github.com/lemonwx/xsql/mysql"
	"github.com/lemonwx/log"
)

var DEFAULT_CAPABILITY uint32 = mysql.CLIENT_LONG_PASSWORD | mysql.CLIENT_LONG_FLAG |
	mysql.CLIENT_CONNECT_WITH_DB | mysql.CLIENT_PROTOCOL_41 |
	mysql.CLIENT_TRANSACTIONS | mysql.CLIENT_SECURE_CONNECTION

var baseConnId int32 = 1000

type CliConn struct {
	conn         net.Conn
	pkt          *mysql.PacketIO
	connectionId uint32
	salt         []byte
	capability   uint32

	status    uint16
	collation mysql.CollationId
	charset   string

	user string
	Db   string

	defaultUser   string
	defaultPasswd string
}

func NewClieConn(conn net.Conn, connid uint32) *CliConn {

	tcpConn := conn.(*net.TCPConn)
	tcpConn.SetNoDelay(false)

	cli := &CliConn{
		conn: tcpConn,
		pkt: mysql.NewPacketIO(tcpConn),
	}
	cli.pkt.Sequence = 0
	cli.status = mysql.SERVER_STATUS_AUTOCOMMIT
	cli.salt, _ = mysql.RandomBuf(20)
	cli.charset = mysql.DEFAULT_CHARSET
	cli.collation = mysql.DEFAULT_COLLATION_ID
	cli.connectionId = connid

	return cli
}

func (c *CliConn) Handshake() error {
	if err := c.writeInitialHandshake(); err != nil {
		return err
	}

	if err := c.readHandshakeResponse(); err != nil {
		return err
	}

	/*
	if err := c.WriteOK(nil); err != nil {
		return err
	}
	c.pkt.Sequence = 0
	*/

	return nil
}

func (c *CliConn) writeInitialHandshake() error {
	data := make([]byte, 4, 128)

	//min version 10
	data = append(data, 10)

	//server version[00]
	data = append(data, mysql.ServerVersion...)
	data = append(data, 0)

	//connection id
	data = append(data, byte(c.connectionId), byte(c.connectionId>>8), byte(c.connectionId>>16), byte(c.connectionId>>24))

	//auth-plugin-data-part-1
	data = append(data, c.salt[0:8]...)

	//filter [00]
	data = append(data, 0)

	//capability flag lower 2 bytes, using default capability here
	data = append(data, byte(DEFAULT_CAPABILITY), byte(DEFAULT_CAPABILITY>>8))

	//charset, utf-8 default
	data = append(data, uint8(mysql.DEFAULT_COLLATION_ID))

	//status
	data = append(data, byte(c.status), byte(c.status>>8))

	//below 13 byte may not be used
	//capability flag upper 2 bytes, using default capability here
	data = append(data, byte(DEFAULT_CAPABILITY>>16), byte(DEFAULT_CAPABILITY>>24))

	//filter [0x15], for wireshark dump, value is 0x15
	data = append(data, 0x15)

	//reserved 10 [00]
	data = append(data, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)

	//auth-plugin-data-part-2
	data = append(data, c.salt[8:]...)

	//filter [00]
	data = append(data, 0)
	return c.writePacket(data)
}

func (c *CliConn) readHandshakeResponse() error {
	data, err := c.ReadPacket()

	if err != nil {
		return err
	}

	pos := 0

	//capability
	c.capability = binary.LittleEndian.Uint32(data[:4])
	pos += 4

	//skip max packet size
	pos += 4

	//charset, skip, if you want to use another charset, use set names
	//c.collation = CollationId(data[pos])
	pos++

	//skip reserved 23[00]
	pos += 23

	//user name
	c.user = string(data[pos : pos+bytes.IndexByte(data[pos:], 0)])

	pos += len(c.user) + 1

	//auth length and auth
	authLen := int(data[pos])
	pos++
	auth := data[pos : pos+authLen]

	checkAuth := mysql.CalcPassword(c.salt, []byte(c.defaultPasswd))
	if c.user != c.defaultUser || !bytes.Equal(auth, checkAuth) {

	}
	if false {
		return mysql.NewDefaultError(mysql.ER_ACCESS_DENIED_ERROR, c.user, c.conn.RemoteAddr().String(), "Yes")
	}

	pos += authLen

	var db string
	if c.capability&mysql.CLIENT_CONNECT_WITH_DB > 0 {
		if len(data[pos:]) == 0 {
			return nil
		}

		db = string(data[pos : pos+bytes.IndexByte(data[pos:], 0)])
		pos += len(c.Db) + 1

	}
	c.Db = db

	return nil
}


func (c *CliConn) SetPktSeq(sz uint8) {
	c.pkt.Sequence = sz
}


func (c *CliConn) ReadPacket() ([]byte, error) {
	return c.pkt.ReadPacket()
}

func (c *CliConn) WriteResultsets(status uint16, rs []*mysql.Resultset) error {
	log.Debugf("[%d] send select rets [%v] to cli", c.connectionId, rs[0].FieldNames)

	total := make([]byte, 0, 4096)
	data := make([]byte, 4, 512)
	var err error

	columnLen := mysql.PutLengthEncodedInt(uint64(len(rs[0].Fields)))

	data = append(data, columnLen...)
	total, err = c.writePacketBatch(total, data, false)
	if err != nil {
		return err
	}

	for _, v := range rs[0].Fields {


		data = data[0:4]
		data = append(data, v.Dump()...)

		total, err = c.writePacketBatch(total, data, false)
		if err != nil {
			return err
		}
	}

	total, err = c.writeEOFBatch(total, status, false)
	if err != nil {
		return err
	}

	for _, r := range rs {
		for _, v := range r.RowDatas {

			data = data[0:4]
			data = append(data, v...)

			total, err = c.writePacketBatch(total, data, false)
			if err != nil {
				return err
			}
		}
	}

	total, err = c.writeEOFBatch(total, status, true)
	total = nil
	if err != nil {
		return err
	}

	return nil
}

func (c *CliConn) WriteResultset(status uint16, r *mysql.Resultset) error {
	log.Debugf("[%d] send select rets [%v] to cli", c.connectionId, r.FieldNames)

	total := make([]byte, 0, 4096)
	data := make([]byte, 4, 512)
	var err error

	columnLen := mysql.PutLengthEncodedInt(uint64(len(r.Fields)))

	data = append(data, columnLen...)
	total, err = c.writePacketBatch(total, data, false)
	if err != nil {
		return err
	}

	for _, v := range r.Fields {
		data = data[0:4]
		data = append(data, v.Dump()...)
		total, err = c.writePacketBatch(total, data, false)
		if err != nil {
			return err
		}
	}

	total, err = c.writeEOFBatch(total, status, false)
	if err != nil {
		return err
	}

	for _, v := range r.RowDatas {
		data = data[0:4]
		data = append(data, v...)
		total, err = c.writePacketBatch(total, data, false)
		if err != nil {
			return err
		}
	}

	total, err = c.writeEOFBatch(total, status, true)
	total = nil
	if err != nil {
		return err
	}

	return nil
}

func (c *CliConn) WriteFieldList(status uint16, fs []*mysql.Field) error {

	data := make([]byte, 4, 1024)

	for _, v := range fs {
		data = data[0:4]
		data = append(data, v.Dump()...)
		if err := c.writePacket(data); err != nil {
			return err
		}
	}

	if err := c.writeEOF(status); err != nil {
		return err
	}
	return nil
}


func (c *CliConn) WriteOK(r *mysql.Result) error {
	log.Debugf("[%d] send exec ok to cli: %v", c.connectionId, r )
	if r == nil {
		r = &mysql.Result{Status: c.status}
	}
	data := make([]byte, 4, 32)

	data = append(data, mysql.OK_HEADER)

	data = append(data, mysql.PutLengthEncodedInt(r.AffectedRows)...)
	data = append(data, mysql.PutLengthEncodedInt(r.InsertId)...)

	if c.capability&mysql.CLIENT_PROTOCOL_41 > 0 {
		data = append(data, byte(r.Status), byte(r.Status>>8))
		data = append(data, 0, 0)
	}

	return c.writePacket(data)
}

func (c *CliConn) WriteError(e error) error {
	var m *mysql.SqlError
	var ok bool
	if m, ok = e.(*mysql.SqlError); !ok {
		m = mysql.NewError(mysql.ER_UNKNOWN_ERROR, e.Error())
	}

	data := make([]byte, 4, 16+len(m.Message))

	data = append(data, mysql.ERR_HEADER)
	data = append(data, byte(m.Code), byte(m.Code>>8))

	if c.capability&mysql.CLIENT_PROTOCOL_41 > 0 {
		data = append(data, '#')
		data = append(data, m.State...)
	}

	data = append(data, m.Message...)

	log.Errorf("[%d] send err to cli: %v", c.connectionId, e)
	return c.writePacket(data)
}


func (c *CliConn) writePacket(data []byte) error {
	return c.pkt.WritePacket(data)
}

func (c *CliConn) writePacketBatch(total, data []byte, direct bool) ([]byte, error) {
	return c.pkt.WritePacketBatch(total, data, direct)
}

func (c *CliConn) writeEOF(status uint16) error {
	data := make([]byte, 4, 9)

	data = append(data, mysql.EOF_HEADER)
	if c.capability&mysql.CLIENT_PROTOCOL_41 > 0 {
		data = append(data, 0, 0)
		data = append(data, byte(status), byte(status>>8))
	}

	return c.writePacket(data)
}

func (c *CliConn) writeEOFBatch(total []byte, status uint16, direct bool) ([]byte, error) {
	data := make([]byte, 4, 9)

	data = append(data, mysql.EOF_HEADER)
	if c.capability&mysql.CLIENT_PROTOCOL_41 > 0 {
		data = append(data, 0, 0)
		data = append(data, byte(status), byte(status>>8))
	}

	return c.writePacketBatch(total, data, direct)
}


func (c  *CliConn) Close() {
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}
}
