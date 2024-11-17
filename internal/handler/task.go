package handler

import (
	"fmt"
	"github.com/lcok/vault-sync/internal/logger"
)

type Task struct {
	log *logger.Logger
}

func NewTask(log *logger.Logger) *Task {
	// TODO conf from flags
	return &Task{log: log}
}

func (c *Task) Start() {
	// TODO impl this
	fmt.Println("Task run once")
}
