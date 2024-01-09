package data

import (
	"my-task-app/features/project"
	_taskData "my-task-app/features/task/data"
	_userData "my-task-app/features/user/data"

	"gorm.io/gorm"
)

// struct user gorm model
type Project struct {
	gorm.Model
	Name        string
	UserID      uint
	Description string
	User        _userData.User
	Tasks       []_taskData.Task
}

func CoreToModel(input project.Core) Project {
	return Project{
		Name:        input.Name,
		UserID:      input.UserID,
		Description: input.Description,
	}
}
