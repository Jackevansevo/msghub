package main

import (
	"log"
	"os"

	"github.com/Jackevansevo/msghub/server"
)

func main() {
	log := log.New(os.Stdout, "", log.LstdFlags)
	hub := server.NewHub(log)
	tcp := server.NewTCPServer(hub)
	udp := server.NewUDPServer(hub)
	go udp.Start()
	go tcp.Start()
	select {}
}
