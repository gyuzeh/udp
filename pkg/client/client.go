package client

import (
	"fmt"
	"net"
)

// Client for UDP
type Client struct {
	connection *net.UDPConn
}

// CreateUDPClient creates a udp client
func CreateUDPClient(ip, port, network string) (Client, error) {
	host := fmt.Sprintf("%s:%s", ip, port)
	updAddress, err := net.ResolveUDPAddr(network, host)

	if err != nil {
		fmt.Println(err)
		return Client{}, err
	}

	connection, err := net.DialUDP(network, nil, updAddress)

	if err != nil {
		fmt.Println(err)
		return Client{}, err
	}

	return Client{connection: connection}, nil
}

func (client *Client) Close() {
	client.connection.Close()
}

func (client *Client) CurrentServerConnection() string {
	return client.connection.RemoteAddr().String()
}

func (client *Client) Write(data []byte) {
	_, err := client.connection.Write(data)

	if err != nil {
		fmt.Println(err)
		client.Close()
	}
}

func (client *Client) Read() string {
	buffer := make([]byte, 1024)
	n, _, err := client.connection.ReadFromUDP(buffer)

	if err != nil {
		fmt.Println(err)
		client.Close()
	}

	return string(buffer[0:n])
}
