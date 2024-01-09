package handler

import (
	"my-task-app/features/project"
	"my-task-app/features/task/handler"
)

type ProjectResponse struct {
	ID          uint   `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	UserID      uint   `json:"user_id" form:"user_id"`
	Description string `json:"description" form:"description"`
	Tasks       []handler.TaskResponse
}

type ProjectResponses struct {
	ID          uint   `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	UserID      uint   `json:"user_id" form:"user_id"`
	Description string `json:"description" form:"description"`
}

func CoreToResponsePreload(data *project.Core) []handler.TaskResponse {
	var sliceOfTaskCore = []handler.TaskResponse{}
	for _, v := range data.Tasks {
		sliceOfTaskCore = append(sliceOfTaskCore, handler.TaskResponse{
			ID:          v.ID,
			Name:        v.Name,
			ProjectID:   v.ProjectID,
			Description: v.Description,
			StatusTask:  v.StatusTask,
		})
	}
	return sliceOfTaskCore
}

func CoreToResponse(data *project.Core) ProjectResponse {
	var result = ProjectResponse{
		ID:          data.ID,
		Name:        data.Name,
		UserID:      data.UserID,
		Description: data.Description,
		Tasks:       CoreToResponsePreload(data),
	}
	return result
}

func CoreToResponseList(data []project.Core) []ProjectResponses {
	var results []ProjectResponses
	for _, v := range data {
		var result = ProjectResponses{
			ID:          v.ID,
			Name:        v.Name,
			UserID:      v.UserID,
			Description: v.Description,
		}
		results = append(results, result)
	}
	return results
}
