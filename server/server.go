package main

import (
	pb "github.com/pablo-campillo/grpc-todo-service-example/proto/todo/v1"
)

type server struct {
	d db

	pb.UnimplementedTodoServiceServer
}
