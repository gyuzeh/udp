package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gyuzeh/udp/pkg/client"
)

func main() {
	arguments := os.Args
	if len(arguments) == 2 {
		fmt.Println("Please provide a host and port string")
		return
	}

	host := arguments[1]
	port := arguments[2]

	udpClient, _ := client.CreateUDPClient(host, port, "udp4")

	fmt.Printf("The UDP server is %s\n", udpClient.CurrentServerConnection())

	defer udpClient.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		data := []byte(text + "\n")

		udpClient.Write(data)

		message := udpClient.Read()

		fmt.Printf("Reply: %s\n", message)
	}
}
