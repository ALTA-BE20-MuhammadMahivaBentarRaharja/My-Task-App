package user

import "time"

type Core struct {
	ID          uint
	Name        string
	Email       string
	Password    string
	Address     string
	PhoneNumber string
	Role        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// interface untuk Data Layer
type UserDataInterface interface {
	Insert(input Core) error
	SelectById(id int) (*Core, error)
	Update(id int, input Core) error
	Delete(id int) error
	Login(email, password string) (data *Core, err error)
}

// interface untuk Service Layer
type UserServiceInterface interface {
	Create(input Core) error
	GetById(id int) (*Core, error)
	Update(id int, input Core) error
	Delete(id int) error
	Login(email, password string) (data *Core, token string, err error)
}
