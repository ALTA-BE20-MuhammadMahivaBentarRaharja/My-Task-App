package service

import (
	"errors"
	"log"
	"my-task-app/app/middlewares"
	"my-task-app/features/user"
	"my-task-app/utils/encrypts"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData    user.UserDataInterface
	hashService encrypts.HashInterface
	validate    *validator.Validate
}

// dependency injection
func New(repo user.UserDataInterface, hash encrypts.HashInterface) user.UserServiceInterface {
	return &userService{
		userData:    repo,
		hashService: hash,
		validate:    validator.New(),
	}
}

// Create implements user.UserServiceInterface.
func (service *userService) Create(input user.Core) error {
	// logic validation
	// validasi email manual
	// if input.Email == "" {
	// 	return errors.New("[validation] email harus diisi")
	// }
	// validasi email dengan menggunakan library validator
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	if input.Password != "" {
		hashedPass, errHash := service.hashService.HashPassword(input.Password)
		if errHash != nil {
			return errors.New("Error hash password.")
		}
		input.Password = hashedPass
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
	isValid := service.hashService.CheckPasswordHash(data.Password, password)
	if !isValid {
		return nil, "", errors.New("password tidak sesuai.")
	}
	log.Println("id user:", data.ID)
	token, errJwt := middlewares.CreateToken(int(data.ID))
	if errJwt != nil {
		return nil, "", errJwt
	}
	return data, token, err
}
