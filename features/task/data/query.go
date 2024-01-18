package data

import (
	"errors"
	"my-task-app/features/task"

	"gorm.io/gorm"
)

type taskQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) task.TaskDataInterface {
	return &taskQuery{
		db: db,
	}
}

// Insert implements task.TaskDataInterface.
func (repo *taskQuery) Insert(input task.Core) error {
	// proses mapping dari struct entities core ke model gorm
	taskInputGorm := Task{
		Name:        input.Name,
		ProjectID:   input.ProjectID,
		Description: input.Description,
		StatusTask:  input.StatusTask,
	}
	// simpan ke DB
	tx := repo.db.Create(&taskInputGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}

func (repo *taskQuery) SelectAllTasksByProjectId(projectId, userIdLogin int) ([]task.Core, error) {
	var tasksDataGorm []Task
	tx := repo.db.Where("project_id = ? AND user_id = ?", projectId, userIdLogin).Find(&tasksDataGorm)
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

// SelectById implements task.TaskDataInterface.
func (repo *taskQuery) SelectById(id int) (*task.Core, error) {
	// Dapatkan data task berdasarkan id dari database
	var taskDataGorm Task
	tx := repo.db.First(&taskDataGorm, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// proses mapping dari struct gorm model ke struct core
	var taskDataCore = task.Core{
		ID:          taskDataGorm.ID,
		Name:        taskDataGorm.Name,
		ProjectID:   taskDataGorm.ProjectID,
		Description: taskDataGorm.Description,
		CreatedAt:   taskDataGorm.CreatedAt,
		UpdatedAt:   taskDataGorm.UpdatedAt,
		StatusTask:  taskDataGorm.StatusTask,
	}
	return &taskDataCore, nil
}

// Update implements task.TaskDataInterface.
func (repo *taskQuery) Update(id int, input task.Core) error {
	dataGorm := CoreToModel(input)
	tx := repo.db.Model(&Task{}).Where("id = ?", id).Updates(dataGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found ")
	}
	return nil
}

// Delete implements task.TaskDataInterface.
func (repo *taskQuery) Delete(id int) error {
	// Hapus task dari database
	tx := repo.db.Delete(&Task{}, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("error record not found")
	}
	return nil
}
