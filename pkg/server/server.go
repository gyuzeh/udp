package server

import (
	"fmt"
	"net"
	"strconv"
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

// Run waits for comunications to the Server and responses
func (server *Server) Run() {

	defer server.close()

	buffer := make([]byte, 1024)

	for {
		n, addr, err := server.connection.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Print("-> ", string(buffer[0:n-1]))

		data := []byte(strconv.Itoa(n))
		fmt.Printf("data: %s\n", string(data))

		server.write(data, addr)
	}
}

func (server *Server) close() {
	server.connection.Close()
}

func (server *Server) write(data []byte, addr *net.UDPAddr) {
	_, err := server.connection.WriteToUDP(data, addr)

	if err != nil {
		fmt.Println(err)
		server.close()
	}
}
