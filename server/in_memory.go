package main

import (
	"time"

	pb "github.com/pablo-campillo/grpc-todo-service-example/proto/todo/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// inMemoryDb is a fake database.
// Its purpose is to let us focus on gRPC
// and to not add any database dependencies such as ORM, ...
type inMemoryDb struct {
	tasks []*pb.Task
}

// New creates a new instance of inMemoryDb
func New() db {
	return &inMemoryDb{}
}

// addTask appends a Task, generated from description and dueDate, to the underlying array.
// It never returns an error but a real database might!
func (d *inMemoryDb) addTask(description string, dueDate time.Time) (uint64, error) {
	nextId := uint64(len(d.tasks) + 1)
	task := &pb.Task{
		Id:          nextId,
		Description: description,
		DueDate:     timestamppb.New(dueDate),
	}

	d.tasks = append(d.tasks, task)
	return nextId, nil
}
