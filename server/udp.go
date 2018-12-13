package server

import (
	"bytes"
	"fmt"
	"log"
	"net"
)

// UDPServer defines a message server
type UDPServer struct{ hub *Hub }

// NewUDPServer returns a new server
func NewUDPServer(hub *Hub) *UDPServer {
	return &UDPServer{hub: hub}
}

// Start initializes the server
func (s UDPServer) Start() {
	addr, _ := net.ResolveUDPAddr("udp", ":7070")
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for {
		inputBytes := make([]byte, 2048)
		length, _, err := conn.ReadFromUDP(inputBytes)
		if err != nil {
			log.Fatal(err)
		}
		buffer := bytes.NewBuffer(inputBytes[:length])
		fmt.Println(buffer)
	}
}
