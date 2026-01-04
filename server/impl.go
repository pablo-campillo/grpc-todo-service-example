package main

import (
	"context"

	pb "github.com/pablo-campillo/grpc-todo-service-example/proto/todo/v1"
)

// AddTask adds a Task to the database.
// It returns the id of the newly inserted Task or an error.
func (s *server) AddTask(_ context.Context, in *pb.AddTaskRequest) (*pb.AddTaskResponse, error) {
	id, _ := s.d.addTask(in.Description, in.DueDate.AsTime())

	return &pb.AddTaskResponse{Id: id}, nil
}
