package data

import (
	"my-task-app/features/task"

	"gorm.io/gorm"
)

// struct task gorm model
type Task struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	ProjectID   uint   `json:"project_id" form:"project_id"`
	Description string `json:"description" form:"description"`
	StatusTask  string `json:"status_task" form:"status_task"`
}

func CoreToModel(input task.Core) Task {
	return Task{
		Name:        input.Name,
		ProjectID:   input.ProjectID,
		Description: input.Description,
		StatusTask:  input.StatusTask,
	}
}
