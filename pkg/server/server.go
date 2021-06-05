package server

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"
)

// Server represents the Struct for a UDP Server
type Server struct {
	connection *net.UDPConn
}

// StartUDPServer initiates a UdpServer with a port and network type ("udp4", "udp6")
func StartUDPServer(port, network string) (Server, error) {
	updAddress, err := net.ResolveUDPAddr(network, ":"+port)

	if err != nil {
		fmt.Println(err)
		return Server{}, err
	}

	connection, err := net.ListenUDP(updAddress.Network(), updAddress)

	if err != nil {
		fmt.Println(err)
		return Server{}, err
	}

	fmt.Printf("Stating server on IP: %s and Port: %d\n", updAddress.IP.String(), updAddress.Port)

	return Server{connection: connection}, nil
}

// Listen waits for comunications to the Server and responses
func (server *Server) Listen() {

	defer server.connection.Close()

	buffer := make([]byte, 1024)

	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := server.connection.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Print("-> ", string(buffer[0:n-1]))

		r := rand.New(rand.NewSource(1001))
		data := []byte(strconv.Itoa(int(r.Uint32())))

		fmt.Printf("data: %s\n", string(data))

		_, err = server.connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
