package data

import (
	"my-task-app/features/task"

	"gorm.io/gorm"
)

// struct task gorm model
type Task struct {
	gorm.Model
	Name        string
	ProjectID   uint
	Description string
	StatusTask  string
}

func CoreToModel(input task.Core) Task {
	return Task{
		Name:        input.Name,
		ProjectID:   input.ProjectID,
		Description: input.Description,
		StatusTask:  input.StatusTask,
	}
}
