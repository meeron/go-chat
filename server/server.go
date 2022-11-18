package server

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/meeron/go-chat/shared"
	"net"
)

var (
	connections = make(map[uuid.UUID]*connection, 0)
)

func Run(address string, status chan<- *shared.ServerStatus) *net.Listener {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	go listen(&listener, status)

	return &listener
}

func listen(listener *net.Listener, status chan<- *shared.ServerStatus) {
	fmt.Println("Listening on ", (*listener).Addr())
	for {
		netConn, err := (*listener).Accept()
		if err != nil {
			break
		}

		id := uuid.New()
		connections[id] = &connection{
			id:      id,
			netConn: &netConn,
		}
		go handle(id)
	}

	fmt.Println("Closing connections...")

	for _, conn := range connections {
		(*conn).close()
	}

	fmt.Println("Server closed.")
	status <- &shared.ServerStatus{
		IsClosed: true,
	}
}

func handle(connId uuid.UUID) {
	conn := *connections[connId]
	fmt.Printf("Connection %s from %v\n", connId, conn.remoteAddr())

	if err := conn.sendId(); err != nil {
		fmt.Println(err)
	}
}
