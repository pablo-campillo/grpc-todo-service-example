package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/pablo-campillo/grpc-todo-service-example/proto/todo/v1"
)

// addTask calls the AddTask unary endpoint with a AddTaskRequest
// generated from description and dueDate.
func addTask(c pb.TodoServiceClient, description string, dueDate time.Time) uint64 {
	req := &pb.AddTaskRequest{
		Description: description,
		DueDate:     timestamppb.New(dueDate),
	}
	res, err := c.AddTask(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("added task: %d\n", res.Id)
	return res.Id
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalln("usage: client [IP_ADDR]")
	}

	addr := args[0]
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.NewClient(addr, opts...)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer func(conn *grpc.ClientConn) {
		if err := conn.Close(); err != nil {
			log.Fatalf("unexpected error: %v", err)
		}
	}(conn)

	c := pb.NewTodoServiceClient(conn)

	fmt.Println("--------ADD--------")
	dueDate := time.Now().Add(5 * time.Second)
	addTask(c, "This is a task", dueDate)
	addTask(c, "This is another task", dueDate)
	addTask(c, "And yet another task", dueDate)
	fmt.Println("-------------------")
}
