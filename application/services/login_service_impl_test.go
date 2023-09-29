package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"todoapp.com/application/dtos"
	"todoapp.com/domain/models"
)

func TestLoginGenerateTokenShouldReturnTokenAndNoError(t *testing.T) {
	// Arrange
	testLoginDTO := &dtos.LoginDTO{
		Email:    "a@a.com",
		Password: "aaaaaaaa",
	}
	testConfig := &models.Config{
		PrivateKeyPath: "../../testing/id_rsa",
		PublicKeyPath:  "../../testing/id_rsa.pub",
	}
	MockedUsersRepository := new(MockedUsersRepository)

	// Act
	testLoginService := NewLoginService(MockedUsersRepository, testConfig)
	error := testLoginService.GenerateToken(testLoginDTO)

	// Assert
	assert.Nil(t, error)
	assert.NotNil(t, testLoginDTO.Token)
	assert.NotEmpty(t, testLoginDTO.Token)
}

func TestLoginLoginShouldReturnErrorIfGivenPasswordDoesNotMatch(t *testing.T) {
	// Arrange
	id1 := uint(1)
	testUser := models.User{
		ID:       &id1,
		Name:     "test 1",
		Email:    "a@a.com",
		Password: "aaaaaaaa",
	}
	testLoginDTO := &dtos.LoginDTO{
		Email:    "a@a.com",
		Password: "bbbbb",
	}
	testContext := context.Background()
	MockedUserRepository := new(MockedUsersRepository)
	MockedUserRepository.mock.On("GetByEmail", testContext, &testLoginDTO.Email).Return(testUser)
	testConfig := &models.Config{
		PrivateKeyPath: "../../testing/id_rsa",
		PublicKeyPath:  "../../testing/id_rsa.pub",
	}

	// Act
	testLoginService := NewLoginService(MockedUserRepository, testConfig)
	error := testLoginService.Login(testContext, testLoginDTO)

	// Assert
	assert.NotNil(t, error)
}

func TestLoginLoginShouldReturnErrorIfUserIsNotFound(t *testing.T) {
	// Arrange
	id1 := uint(1)
	testUser := models.User{
		ID:       &id1,
		Name:     "test 1",
		Email:    "a@a.com",
		Password: "aaaaaaaa",
	}
	testLoginDTO := &dtos.LoginDTO{
		Email:    "b@b.com",
		Password: "aaaaaaaa",
	}
	testContext := context.Background()
	MockedUserRepository := new(MockedUsersRepository)
	MockedUserRepository.mock.On("GetByEmail", testContext, &testLoginDTO.Email).Return(testUser)
	testConfig := &models.Config{
		PrivateKeyPath: "../../testing/id_rsa",
		PublicKeyPath:  "../../testing/id_rsa.pub",
	}

	// Act
	testLoginService := NewLoginService(MockedUserRepository, testConfig)
	error := testLoginService.Login(testContext, testLoginDTO)

	// Assert
	assert.NotNil(t, error)
}

func TestLoginLoginShouldNotReturnAnyErrorsIfEverythingIsOk(t *testing.T) {
	// Arrange
	id1 := uint(1)
	testUser := models.User{
		ID:       &id1,
		Name:     "test 1",
		Email:    "a@a.com",
		Password: "aaaaaaaa",
	}
	testLoginDTO := &dtos.LoginDTO{
		Email:    "a@a.com",
		Password: "aaaaaaaa",
	}
	testContext := context.Background()
	MockedUserRepository := new(MockedUsersRepository)
	MockedUserRepository.mock.On("GetByEmail", testContext, &testLoginDTO.Email).Return(testUser)
	testConfig := &models.Config{
		PrivateKeyPath: "../../testing/id_rsa",
		PublicKeyPath:  "../../testing/id_rsa.pub",
	}

	// Act
	testLoginService := NewLoginService(MockedUserRepository, testConfig)
	error := testLoginService.Login(testContext, testLoginDTO)

	// Assert
	assert.Nil(t, error)
}
