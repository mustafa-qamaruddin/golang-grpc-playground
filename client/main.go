package main

import (
	"context"
	todos "github.com/mustafa-qamaruddin/golang-grpc-playground/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	opts := grpc.WithInsecure()
	conn, err := grpc.Dial("localhost:3000", opts)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer conn.Close()
	c := todos.NewTodosClient(conn)
	todo, err := c.CreateTodo(context.Background(), &todos.Todo{
		Id:          "123",
		Title:       "Implement gRPC",
		Description: "Implemented gRPC server/client",
	})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("[%s]: %s, %s, %s", time.Now(),
		todo.Id, todo.Title, todo.Description)
}
