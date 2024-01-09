package handler

import (
	"my-task-app/features/project"
)

type ProjectRequest struct {
	Name        string `json:"name" form:"name"`
	UserID      uint   `json:"user_id" form:"user_id"`
	Description string `json:"description" form:"description"`
}

type ProjectRequestUpdate struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}

func RequestToCore(userIdLogin uint, input ProjectRequest) project.Core {
	return project.Core{
		Name:        input.Name,
		UserID:      userIdLogin,
		Description: input.Description,
	}
}

func RequestToCoreUpdate(input ProjectRequestUpdate) project.Core {
	return project.Core{
		Name:        input.Name,
		Description: input.Description,
	}
}
