package data

import (
	"errors"
	"my-task-app/features/project"
	"my-task-app/features/task"
	"my-task-app/features/task/data"

	"gorm.io/gorm"
)

type projectQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) project.ProjectDataInterface {
	return &projectQuery{
		db: db,
	}
}

// Insert implements project.ProjectDataInterface.
func (repo *projectQuery) Insert(input project.Core) error {
	// proses mapping dari struct entities core ke model gorm
	projectInputGorm := Project{
		Name:        input.Name,
		UserID:      input.UserID,
		Description: input.Description,
	}
	// simpan ke DB
	tx := repo.db.Create(&projectInputGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}

// SelectAll implements project.ProjectDataInterface.
func (repo *projectQuery) SelectAll(userIdLogin int) ([]project.Core, error) {
	var projectsDataGorm []Project
	tx := repo.db.Where("user_id = ?", userIdLogin).Find(&projectsDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// proses mapping dari struct gorm model ke struct core
	var projectsDataCore []project.Core
	for _, value := range projectsDataGorm {
		var projectCore = project.Core{
			ID:          value.ID,
			Name:        value.Name,
			UserID:      value.UserID,
			Description: value.Description,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
		}
		projectsDataCore = append(projectsDataCore, projectCore)
	}

	return projectsDataCore, nil
}

// SelectById implements project.ProjectDataInterface.
func (repo *projectQuery) SelectById(id, userIdLogin int) (*project.Core, error) {
	// Dapatkan data project berdasarkan id dari database
	var projectDataGorm Project
	tx := repo.db.Preload("Tasks").Where("user_id = ?", userIdLogin).First(&projectDataGorm, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// proses mapping dari struct gorm model ke struct core
	var sliceOfTaskCore = []task.Core{}
	for _, v := range projectDataGorm.Tasks {
		sliceOfTaskCore = append(sliceOfTaskCore, task.Core{
			ID:          v.ID,
			Name:        v.Name,
			ProjectID:   v.ProjectID,
			Description: v.Description,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
			StatusTask:  v.StatusTask,
		})
	}
	var projectDataCore = project.Core{
		ID:          projectDataGorm.ID,
		Name:        projectDataGorm.Name,
		UserID:      projectDataGorm.UserID,
		Description: projectDataGorm.Description,
		CreatedAt:   projectDataGorm.CreatedAt,
		UpdatedAt:   projectDataGorm.UpdatedAt,
		Tasks:       sliceOfTaskCore,
	}

	return &projectDataCore, nil
}

func (repo *projectQuery) SelectAllTasksByProjectId(projectId int) ([]task.Core, error) {
	var tasksDataGorm []data.Task
	tx := repo.db.Where("project_id = ?", projectId).Find(&tasksDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// proses mapping dari struct gorm model ke struct core
	var tasksDataCore []task.Core
	for _, value := range tasksDataGorm {
		var taskCore = task.Core{
			ID:          value.ID,
			Name:        value.Name,
			ProjectID:   value.ProjectID,
			Description: value.Description,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
			StatusTask:  value.StatusTask,
		}
		tasksDataCore = append(tasksDataCore, taskCore)
	}

	return tasksDataCore, nil
}

// Update implements project.ProjectDataInterface.
func (repo *projectQuery) Update(userIdLogin int, id int, input project.Core) error {
	dataGorm := CoreToModel(input)
	tx := repo.db.Model(&Project{}).Where("id = ? AND user_id = ?", id, userIdLogin).Updates(dataGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}

// Delete implements project.ProjectDataInterface.
func (repo *projectQuery) Delete(id, userIdLogin int) error {
	// Hapus project dari database
	tx := repo.db.Where("id = ? AND user_id = ?", id, userIdLogin).Delete(&Project{})
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found")
	}
	return nil
}

// Delete implements task.TaskDataInterface.
func (repo *projectQuery) DeleteTask(id int) error {
	// Hapus task dari database
	tx := repo.db.Delete(&data.Task{}, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found")
	}
	return nil
}
