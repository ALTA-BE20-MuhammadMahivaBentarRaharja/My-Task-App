package handler

type TaskResponse struct {
	ID          uint   `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	ProjectID   uint   `json:"project_id" form:"project_id"`
	Description string `json:"description" form:"description"`
	StatusTask  string `json:"status_task" form:"status_task"`
}

// func CoreToResponse(data *task.Core) TaskResponse {
// 	var result = TaskResponse{
// 		ID:          data.ID,
// 		Name:        data.Name,
// 		ProjectID:   data.ProjectID,
// 		Description: data.Description,
// 	}
// 	return result
// }

// func CoreToResponseList(data []task.Core) []TaskResponse {
// 	var results []TaskResponse
// 	for _, v := range data {
// 		var result = TaskResponse{
// 			ID:          v.ID,
// 			Name:        v.Name,
// 			ProjectID:   v.ProjectID,
// 			Description: v.Description,
// 		}
// 		results = append(results, result)
// 	}
// 	return results
// }
