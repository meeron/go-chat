package server

import (
	"bytes"
	"encoding/binary"
	"github.com/google/uuid"
	"github.com/meeron/go-chat/shared"
	"net"
)

type connection struct {
	id      uuid.UUID
	netConn *net.Conn
}

func (conn connection) sendId() error {
	buffer := bytes.Buffer{}

	buffer.WriteByte(byte(shared.CmdOk))

	connIdBytes, err := conn.id.MarshalBinary()
	if err != nil {
		return err
	}

	err = binary.Write(&buffer, binary.BigEndian, int32(len(connIdBytes)))
	if err != nil {
		return err
	}

	buffer.Write(connIdBytes)

	_, sentErr := (*conn.netConn).Write(buffer.Bytes())

	return sentErr
}

func (conn connection) remoteAddr() net.Addr {
	return (*conn.netConn).RemoteAddr()
}

func (conn connection) close() error {
	return (*conn.netConn).Close()
}
