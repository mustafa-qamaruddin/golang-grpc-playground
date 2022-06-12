package main

import (
	"context"
	todos "github.com/mustafa-qamaruddin/golang-grpc-playground/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type TodosServer struct {
	todos.UnimplementedTodosServer
}

func (t TodosServer) CreateTodo(ctx context.Context, todo *todos.Todo) (*todos.Todo, error) {
	//TODO implement me
	return todo, nil
}

func main() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("unable to listen on port 3000 %v", err)
	}
	server := grpc.NewServer()
	todos.RegisterTodosServer(server, &TodosServer{})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("unable to serve %v", err)
	}
}
