package handler

import "my-task-app/features/user"

type UserResponse struct {
	ID      uint   `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Address string `json:"address" form:"address"`
}

// func CoreToResponse(data user.Core) UserResponse {
// 	return UserResponse{
// 		ID:    data.ID,
// 		Name:  data.Name,
// 		Email: data.Email,
// 	}
// }

func CoreToResponse(data *user.Core) UserResponse {
	var result = UserResponse{
		ID:      data.ID,
		Name:    data.Name,
		Email:   data.Email,
		Address: data.Address,
	}
	return result
}

// func CoreToResponseList(data []user.Core) []UserResponse {
// 	var results []UserResponse
// 	for _, v := range data {
// 		results = append(results, CoreToResponse(v))
// 	}
// 	return results
// }

func CoreToResponseList(data []user.Core) []UserResponse {
	var results []UserResponse
	for _, v := range data {
		var result = UserResponse{
			ID:      v.ID,
			Name:    v.Name,
			Email:   v.Email,
			Address: v.Address,
		}
		results = append(results, result)
	}
	return results
}
