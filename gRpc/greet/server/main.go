package main

import (
	"log"
	"net"
)

var addr = "0.0.0.0:50051"

type Server struct {
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)
}
