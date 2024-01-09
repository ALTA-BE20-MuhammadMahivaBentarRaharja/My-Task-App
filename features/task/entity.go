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
	Update(id int, input Core) error
	Delete(id int) error
}

// interface untuk Service Layer
type TaskServiceInterface interface {
	Create(input Core) error
	Update(id int, input Core) error
	Delete(id int) error
}
