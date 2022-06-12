package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

type TodoService struct {
}

func main() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("unable to listen on port 3000", err)
	}
	server := grpc.NewServer()

}
