package client

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/google/uuid"
	"github.com/meeron/go-chat/shared"
	"net"
)

const (
	BufferLen = 1024
)

func Connect(address string) {
	fmt.Println("Connecting...")

	connection, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}

	defer connection.Close()

	buffer := make([]byte, BufferLen)

	_, err = connection.Read(buffer)
	if err != nil {
		panic(err)
	}

	cmd := shared.Cmd(buffer[0])
	var count int32

	b := bytes.NewBuffer(buffer[1:5])

	if err = binary.Read(b, binary.BigEndian, &count); err != nil {
		panic(err)
	}

	if cmd == shared.CmdOk {
		connId, err := uuid.FromBytes(buffer[5 : 5+count])
		if err != nil {
			panic(err)
		}

		fmt.Printf("Connected. Connection id=%v\n", connId)
	}
}
