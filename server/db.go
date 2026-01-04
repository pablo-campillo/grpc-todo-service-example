package main

import "time"

// db is the interface representing the actions we want to take on any database.
type db interface {
	addTask(description string, dueDate time.Time) (uint64, error)
}
