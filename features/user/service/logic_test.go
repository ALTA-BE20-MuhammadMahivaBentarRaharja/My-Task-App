package service

import (
	"errors"
	"my-task-app/features/user"
	"my-task-app/mocks"
	hashMock "my-task-app/utils/encrypts/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetById(t *testing.T) {
	//pembuatan object mock yg dibutuhkan
	repo := new(mocks.UserData)
	hash := new(hashMock.HashService)

	returnData := user.Core{
		ID:          1,
		Name:        "reza",
		Email:       "reza@mail.id",
		Password:    "123",
		Address:     "jakarta",
		PhoneNumber: "0851",
		Role:        "member",
	}

	t.Run("Success Get By Id", func(t *testing.T) {
		// mock return suatu func dari data layer
		repo.On("SelectById", mock.Anything).Return(&returnData, nil).Once()
		//create object service
		srv := New(repo, hash)
		//return dari service layer
		result, err := srv.GetById(1)

		assert.NoError(t, err)
		assert.Equal(t, returnData.Name, result.Name)
		assert.Equal(t, returnData.Email, result.Email)
		repo.AssertExpectations(t)
	})

	t.Run("User Not Found", func(t *testing.T) {
		// mock return suatu func dari data layer untuk simulasi user tidak ditemukan
		repo.On("SelectById", 1).Return(nil, errors.New("user not found")).Once()
		// create object service
		srv := New(repo, hash)
		// return dari service layer
		result, err := srv.GetById(1)

		assert.Error(t, err)
		assert.Nil(t, result)
		repo.AssertExpectations(t)
	})
}

/*
func TestGetAll(t *testing.T) {
	//pembuatan object mock yg dibutuhkan
	repo := new(mocks.UserData)
	hash := new(hashMock.HashService)

	returnData := []user.Core{
		{
			ID:          1,
			Name:        "alta",
			Email:       "alta@mail.id",
			Password:    "qwerty",
			Address:     "Surabaya",
			PhoneNumber: "081234",
			Role:        "user",
		},
		{
			ID:          2,
			Name:        "alta2",
			Email:       "alta2@mail.id",
			Password:    "qwerty",
			Address:     "Surabaya",
			PhoneNumber: "081234",
			Role:        "user",
		},
	}

	t.Run("Success Get All", func(t *testing.T) {
		// mock return suatu func dari data layer
		repo.On("SelectAll").Return(returnData, nil).Once()
		//create object service
		srv := New(repo, hash)
		//return dari service layer
		result, err := srv.GetAll()

		assert.NoError(t, err)
		assert.Equal(t, returnData[0].Name, result[0].Name)
		assert.Equal(t, returnData[1].Name, result[1].Name)
		assert.Equal(t, returnData[0].Email, result[0].Email)
		repo.AssertExpectations(t)
	})
}
*/
