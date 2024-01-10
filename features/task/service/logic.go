package service

import (
	"errors"
	"my-task-app/features/project"
	"my-task-app/features/task"
)

type taskService struct {
	taskData       task.TaskDataInterface
	projectService project.ProjectServiceInterface
}

// dependency injection
func New(repo task.TaskDataInterface, project project.ProjectServiceInterface) task.TaskServiceInterface {
	return &taskService{
		taskData:       repo,
		projectService: project,
	}
}

// Create implements user.UserServiceInterface.
func (service *taskService) Create(input task.Core, userIdLogin int) error {
	// logic validation
	if input.Name == "" {
		return errors.New("[validation] name harus diisi")
	}
	if input.Description == "" {
		return errors.New("[validation] description harus diisi")
	}
	if input.ProjectID <= 0 {
		return errors.New("Project ID harus diisi")
	}

	result, errGet := service.projectService.GetById(int(input.ProjectID), userIdLogin)
	if errGet != nil {
		return errors.New("Error, Product ID tidak ditemukan.")
	}
	if result == nil {
		return errors.New("Error, Project ID tidak ditemukan.")
	}

	err := service.taskData.Insert(input)
	return err
}

// Update implements task.TaskServiceInterface.
func (service *taskService) Update(id int, input task.Core, userIdLogin int) error {
	//validasi
	if id <= 0 {
		return errors.New("invalid id")
	}

	resultPrj, errPrj := service.taskData.SelectById(id)
	if errPrj != nil {
		return errors.New("Error, data tidak ditemukan.")
	}

	result, errGet := service.projectService.GetById(int(resultPrj.ProjectID), userIdLogin)
	if errGet != nil {
		return errors.New("Error, Product ID tidak ditemukan.")
	}
	if result == nil {
		return errors.New("Error, Project ID tidak ditemukan.")
	}

	err := service.taskData.Update(id, input)
	return err
}

// Delete implements task.TaskServiceInterface.
func (service *taskService) Delete(id, userIdLogin int) error {
	//validasi
	if id <= 0 {
		return errors.New("invalid id")
	}

	resultPrj, errPrj := service.taskData.SelectById(id)
	if errPrj != nil {
		return errors.New("Error, data tidak ditemukan.")
	}

	result, errGet := service.projectService.GetById(int(resultPrj.ProjectID), userIdLogin)
	if errGet != nil {
		return errors.New("Error, Product ID tidak ditemukan.")
	}
	if result == nil {
		return errors.New("Error, Project ID tidak ditemukan.")
	}

	err := service.taskData.Delete(id)
	return err
}
