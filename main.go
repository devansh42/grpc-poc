package main

import (
	"log"
	"time"
)

const (
	port = 8000
)

func main() {
	go startServer(port)
	time.Sleep(time.Second * 2)
	startClient(port)
	log.Printf("Shutting Down")
}
