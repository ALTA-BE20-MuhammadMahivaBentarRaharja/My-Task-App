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
