package service

import (
	"errors"
	"my-task-app/features/task"
)

type taskService struct {
	taskData task.TaskDataInterface
}

// dependency injection
func New(repo task.TaskDataInterface) task.TaskServiceInterface {
	return &taskService{
		taskData: repo,
	}
}

// Create implements user.UserServiceInterface.
func (service *taskService) Create(input task.Core) error {
	// logic validation
	if input.Name == "" {
		return errors.New("[validation] name harus diisi")
	}
	if input.Description == "" {
		return errors.New("[validation] description harus diisi")
	}
	err := service.taskData.Insert(input)
	return err
}

// Update implements task.TaskServiceInterface.
func (service *taskService) Update(id int, input task.Core) error {
	//validasi
	if id <= 0 {
		return errors.New("invalid id")
	}
	err := service.taskData.Update(id, input)
	return err
}

// Delete implements task.TaskServiceInterface.
func (service *taskService) Delete(id int) error {
	//validasi
	if id <= 0 {
		return errors.New("invalid id")
	}
	err := service.taskData.Delete(id)
	return err
}
