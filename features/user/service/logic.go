package service

import (
	"errors"
	"log"
	"my-task-app/app/middlewares"
	"my-task-app/features/user"
)

type userService struct {
	userData user.UserDataInterface
}

// dependency injection
func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
	}
}

// Create implements user.UserServiceInterface.
func (service *userService) Create(input user.Core) error {
	// logic validation
	if input.Email == "" {
		return errors.New("[validation] email harus diisi")
	}
	err := service.userData.Insert(input)
	return err
}

// GetById implements user.UserServiceInterface.
func (service *userService) GetById(userIdLogin int) (*user.Core, error) {
	result, err := service.userData.SelectById(userIdLogin)
	return result, err
}

// Update implements user.UserServiceInterface.
func (service *userService) Update(userIdLogin int, input user.Core) error {
	//validasi
	if userIdLogin <= 0 {
		return errors.New("invalid id")
	}
	//validasi inputan
	// ...
	err := service.userData.Update(userIdLogin, input)
	return err
}

// Delete implements user.UserServiceInterface.
func (service *userService) Delete(userIdLogin int) error {
	//validasi
	if userIdLogin <= 0 {
		return errors.New("invalid id")
	}
	err := service.userData.Delete(userIdLogin)
	return err
}

// Login implements user.UserServiceInterface.
func (service *userService) Login(email string, password string) (data *user.Core, token string, err error) {
	if email == "" || password == "" {
		return nil, "", errors.New("email dan password wajib diisi")
	}
	// check apakah passwrd lebih dari 8 karakter atau terdiri Uppercase, lowercase,number, symbol
	data, err = service.userData.Login(email, password)
	if err != nil {
		return nil, "", err
	}
	log.Println("id user:", data.ID)
	token, errJwt := middlewares.CreateToken(int(data.ID))
	if errJwt != nil {
		return nil, "", errJwt
	}
	return data, token, err
}
