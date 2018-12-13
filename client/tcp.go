package client

import (
	"fmt"
	"net"
)

// TCPClient represents a connection to a server
type TCPClient struct {
	Client
	net.Conn
}

// NewTCPClient returns a new Client
func NewTCPClient(conn net.Conn) *TCPClient {
	return &TCPClient{NewClient(), conn}
}

func (c TCPClient) String() string {
	return fmt.Sprintf("<TCPClient(uuid: %s)>", c.UUID)
}

// Listen will block infinite, attempting to read of the clients message channel
func (c TCPClient) Listen() {
	for {
		select {
		case msg := <-c.Messages:
			content := fmt.Sprintf("Recieved: %s\n", msg.Body)
			c.Conn.Write([]byte(content))
		}
	}
}
