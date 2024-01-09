package data

import (
	"my-task-app/features/user"

	"gorm.io/gorm"
)

// struct user gorm model
type User struct {
	gorm.Model
	// ID          uint `gorm:"primaryKey"`
	// CreatedAt   time.Time
	// UpdatedAt   time.Time
	// DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string
	Email       string `gorm:"unique"`
	Password    string
	Address     string
	PhoneNumber string
	Role        string
}

func CoreToModel(input user.Core) User {
	return User{
		Name:        input.Name,
		Email:       input.Email,
		Address:     input.Address,
		PhoneNumber: input.PhoneNumber,
		Role:        input.Role,
	}
}

func (u User) ModelToCore() user.Core {
	return user.Core{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		Password:    u.Password,
		Address:     u.Address,
		PhoneNumber: u.PhoneNumber,
		Role:        u.Role,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}
