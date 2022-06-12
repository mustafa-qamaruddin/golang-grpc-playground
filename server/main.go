package main

import (
	"google.golang.org/grpc"
	"log"
	todos "mustafa-qamaruddin/golang-grpc-playground/blob/main/proto/todos.proto"
	"net"
)

type TodoService struct {
}

func main() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("unable to listen on port 3000 %v", err)
	}
	server := grpc.NewServer()
	todos.RegisterTodosServer(server, &TodoService{})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("unable to serve %v", err)
	}
}
