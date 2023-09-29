package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"todoapp.com/application/dtos"
	"todoapp.com/domain/models"
)

type MockedUsersService struct {
	mock mock.Mock
}

func (m *MockedUsersService) GetByEmail(context context.Context, email *string) dtos.UserDTO {
	args := m.mock.Called(context, email)

	return args.Get(0).(dtos.UserDTO)
}

// TODO: test GenerateToken first
func TestLoginGenerateTokenShouldReturnMockedTokenAndNoError(t *testing.T) {

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
	testConfig := &models.Config{
		PrivateKeyPath: "./mocks/id_ed25519",
		PublicKeyPath:  "./mocks/id_ed25519.pub",
	}

	// Act
	testLoginService := NewLoginService(MockedUsersService, testConfig)
	error := testLoginService.Login(testContext, testLoginDTO)

	// Assert
	assert.NotNil(t, error)
}
