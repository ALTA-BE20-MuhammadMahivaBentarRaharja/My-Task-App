package task

import (
	"time"
)

type Core struct {
	ID          uint
	Name        string
	ProjectID   uint
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	StatusTask  string
}

// interface untuk Data Layer
type TaskDataInterface interface {
	Insert(input Core) error
	SelectAllTasksByProjectId(projectId, userIdLogin int) ([]Core, error)
	SelectById(id int) (*Core, error)
	Update(id int, input Core) error
	Delete(id int) error
}

// interface untuk Service Layer
type TaskServiceInterface interface {
	Create(input Core, userIdLogin int) error
	GetAllTasksByProjectId(projectId, userIdLogin int) ([]Core, error)
	Update(id int, input Core, userIdLogin int) error
	Delete(id int, userIdLogin int) error
}
