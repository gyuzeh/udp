package main

import (
	"fmt"
	"os"

	"github.com/gyuzeh/udp/pkg/server"
)

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	port := arguments[1]

	udpServer, _ := server.StartUDPServer(port, "udp4")
	udpServer.Run()
}
