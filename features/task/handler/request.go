package handler

import (
	"my-task-app/features/task"
)

type TaskRequest struct {
	Name        string `json:"name" form:"name"`
	ProjectID   uint   `json:"project_id" form:"project_id"`
	Description string `json:"description" form:"description"`
	StatusTask  string `json:"status_task" form:"status_task"`
}

func RequestToCore(input TaskRequest) task.Core {
	return task.Core{
		Name:        input.Name,
		ProjectID:   input.ProjectID,
		Description: input.Description,
		StatusTask:  input.StatusTask,
	}
}
