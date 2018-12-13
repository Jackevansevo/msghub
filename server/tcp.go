package server

import (
	"bufio"
	"log"
	"net"
	"strings"

	"github.com/Jackevansevo/msghub/client"
)

// TCPServer defines a message server
type TCPServer struct{ hub *Hub }

// NewTCPServer returns a new server
func NewTCPServer(hub *Hub) *TCPServer {
	return &TCPServer{hub: hub}
}

// Start initializes the server
func (s TCPServer) Start() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		client := client.NewTCPClient(conn)
		go client.Listen()
		go s.serveClient(client)
	}
}

func (s *TCPServer) serveClient(tcpClient *client.TCPClient) {

	scanner := bufio.NewScanner(tcpClient.Conn)
	for scanner.Scan() {

		cmd := scanner.Text()

		s.hub.log.Println("Recieved:", cmd)
		tokens := strings.Split(cmd, " ")

		// [TODO] Abstract out message parsing logic
		switch tokens[0] {
		case "pub":
			message := client.Message{Author: tcpClient.Client, Topic: tokens[1], Body: tokens[2]}
			s.hub.Publish(message)
		case "sub":
			topic := tokens[1]
			s.hub.Subscribe(topic, &tcpClient.Client)
		case "unsub":
		default:
			s.hub.log.Println("Invalid cmd: ", tokens[0])

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
