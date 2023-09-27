package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"todoapp.com/application/dtos"
)

type MockedUsersService struct {
	mock mock.Mock
}

func (m *MockedUsersService) GetByEmail(context context.Context, email *string) dtos.UserDTO {
	args := m.mock.Called(context, email)

	return args.Get(0).(dtos.UserDTO)
}

func TestLoginLoginShouldReturnErrorIfGivenPasswordDoesNotMatch(t *testing.T) {
	// Arrange
	id1 := uint(1)
	testUserDTO := &dtos.UserDTO{
		ID:       &id1,
		Name:     "test 1",
		Email:    "a@a.com",
		Password: "aaaaaaaa",
	}
	testLoginDTO := &dtos.LoginDTO{
		Email:    "a@a.com",
		Password: "bbbbb",
	}
	testEmail := "a@a.com"
	testContext := context.Background()
	MockedUsersService := new(MockedUsersService)
	MockedUsersService.mock.On("GetByEmail", testContext, &testEmail).Return(testUserDTO)

	// Act
	// TODO: mock config and GenerateToken
	// TODO: mock config and GenerateToken
	// TODO: mock config and GenerateToken
	// TODO: mock config and GenerateToken
	// TODO: mock config and GenerateToken
	testLoginService := NewLoginService(MockedUsersService)
	error := testLoginService.Login(testContext, testLoginDTO)

	// Assert
	assert.NotNil(t, error)
}

// User and Login service test can share mocked GenerateToken defined here since they share package.
