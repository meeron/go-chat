package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/meeron/go-chat/server"
	"github.com/meeron/go-chat/shared"
	"net"
	"os"
	"strings"
)

func main() {
	var listner *net.Listener
	serverStatus := make(chan *shared.ServerStatus, 0)

	listen := flag.String("listen", "", "Listen address. Usage: -listen :8080")
	connect := flag.String("connect", "", "Connect to server. Usage -connect localhost:8080")
	flag.Parse()

	if *listen == "" && *connect == "" {
		flag.PrintDefaults()
		return
	}

	if *listen != "" && *connect != "" {
		fmt.Println("Choose either '-connect' or '-listen' flag")
		return
	}

	if *listen != "" {
		listner = server.Run(*listen, serverStatus)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Command: ")
		cmd, _ := reader.ReadString('\n')

		cmd = strings.TrimSpace(cmd)

		if cmd == "quit" {
			(*listner).Close()

			status := <-serverStatus

			if (*status).IsClosed {
				return
			}
		}

		fmt.Println("Invalid command ", cmd)
	}
}
