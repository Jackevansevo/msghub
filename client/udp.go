package client

import (
	"fmt"
	"net"
)

// UDPClient represents a connection to a server
type UDPClient struct {
	Client
	net.Conn
}

// NewUDPClient returns a new Client
func NewUDPClient(conn net.Conn) *UDPClient {
	return &UDPClient{NewClient(), conn}
}

func (c UDPClient) String() string {
	return fmt.Sprintf("<UDPClient(uuid: %s)>", c.UUID)
}

// Listen will block infinite, attempting to read of the clients message channel
func (c UDPClient) Listen() {
	for {
		select {
		case msg := <-c.Messages:
			content := fmt.Sprintf("Recieved: %s\n", msg.Body)
			c.Conn.Write([]byte(content))
		}
	}
}
