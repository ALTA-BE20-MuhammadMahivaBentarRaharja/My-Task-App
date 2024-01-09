package project

import (
	"my-task-app/features/task"
	"my-task-app/features/user"
	"time"
)

type Core struct {
	ID          uint
	Name        string
	UserID      uint
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        user.Core
	Tasks       []task.Core
}

// interface untuk Data Layer
type ProjectDataInterface interface {
	Insert(input Core) error
	SelectAll(userIdLogin int) ([]Core, error)
	SelectById(id, userIdLogin int) (*Core, error)
	Update(userIdLogin int, id int, input Core) error
	Delete(id int, userIdLogin int) error
}

// interface untuk Service Layer
type ProjectServiceInterface interface {
	Create(input Core) error
	GetAll(userIdLogin int) ([]Core, error)
	GetById(id, userIdLogin int) (*Core, error)
	Update(userIdLogin int, id int, input Core) error
	Delete(id int, userIdLogin int) error
}
