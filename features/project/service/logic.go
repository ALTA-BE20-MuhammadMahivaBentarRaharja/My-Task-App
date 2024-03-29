package service

import (
	"errors"
	"my-task-app/features/project"
)

type projectService struct {
	projectData project.ProjectDataInterface
}

// dependency injection
func New(repo project.ProjectDataInterface) project.ProjectServiceInterface {
	return &projectService{
		projectData: repo,
	}
}

// Create implements user.UserServiceInterface.
func (service *projectService) Create(input project.Core) error {
	// logic validation
	if input.Name == "" {
		return errors.New("[validation] name harus diisi")
	}
	if input.Description == "" {
		return errors.New("[validation] description harus diisi")
	}
	err := service.projectData.Insert(input)
	return err
}

// GetAll implements project.ProjectServiceInterface.
func (service *projectService) GetAll(userIdLogin int) ([]project.Core, error) {
	// memanggil func yg ada di data layer
	results, err := service.projectData.SelectAll(userIdLogin)
	return results, err
}

// GetById implements project.ProjectServiceInterface.
func (service *projectService) GetById(id, userIdLogin int) (*project.Core, error) {
	result, err := service.projectData.SelectById(id, userIdLogin)
	return result, err
}

// Update implements project.ProjectServiceInterface.
func (service *projectService) Update(userIdLogin int, id int, input project.Core) error {
	//validasi
	if id <= 0 {
		return errors.New("invalid id")
	}
	err := service.projectData.Update(userIdLogin, id, input)
	return err
}

// Delete implements project.ProjectServiceInterface.
func (service *projectService) Delete(id, userIdLogin int) error {
	//validasi
	if id <= 0 {
		return errors.New("invalid id")
	}

	_, errSelect := service.projectData.SelectById(id, userIdLogin)
	if errSelect != nil {
		return errSelect
	}

	// Fetch tasks associated with the project
	tasks, errGet := service.projectData.SelectAllTasksByProjectId(id)
	if errGet != nil {
		return errGet
	}

	// Delete each task associated with the project
	for _, task := range tasks {
		errDel := service.projectData.DeleteTask(int(task.ID))
		if errDel != nil {
			return errDel
		}
	}

	err := service.projectData.Delete(id, userIdLogin)
	return err
}
