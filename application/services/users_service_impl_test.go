package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"todoapp.com/application/dtos"
	"todoapp.com/domain/models"
)

type MockedUsersRepository struct {
	mock mock.Mock
}

func (m *MockedUsersRepository) GetAll(context context.Context) []models.User {
	args := m.mock.Called(context)

	return args.Get(0).([]models.User)
}

func (m *MockedUsersRepository) GetById(context context.Context, id *uint) models.User {
	args := m.mock.Called(context, id)

	return args.Get(0).(models.User)
}

func (m *MockedUsersRepository) GetByEmail(context context.Context, email string) models.User {
	args := m.mock.Called(context, email)

	return args.Get(0).(models.User)
}

func (m *MockedUsersRepository) Create(context context.Context, model *models.User) error {
	id1 := uint(1)
	args := m.mock.Called(context, model)
	model.ID = &id1

	return args.Error(0)
}

func (m *MockedUsersRepository) Update(context context.Context, model *models.User) error {
	args := m.mock.Called(context, model)

	return args.Error(0)
}

func (m *MockedUsersRepository) Delete(context context.Context, model *models.User) error {
	args := m.mock.Called(context, model)

	return args.Error(0)
}

// Just to implement UsersRepository interface
func (m *MockedUsersRepository) CleanUp(context context.Context) int64 {
	return 0
}

func TestUsersGetAllShouldReturnTestUserDTOs(t *testing.T) {
	// Arrange
	id1, id2 := uint(1), uint(2)
	testUsers := []models.User{
		{
			ID:       &id1,
			Name:     "test name 1",
			Email:    "a@a.com",
			Password: "aaaaaaaa",
		},
		{
			ID:       &id2,
			Name:     "test name 2",
			Email:    "b@b.com",
			Password: "bbbbbbbb",
		},
	}
	testUserDTOs := []dtos.UserDTO{
		{
			ID:       &id1,
			Name:     "test name 1",
			Email:    "a@a.com",
			Password: "",
		},
		{
			ID:       &id2,
			Name:     "test name 2",
			Email:    "b@b.com",
			Password: "",
		},
	}
	testContext := context.Background()
	mockedUsersRepository := new(MockedUsersRepository)
	mockedUsersRepository.mock.On("GetAll", testContext).Return(testUsers)

	// Act
	testUsersService := NewUsersService(mockedUsersRepository)
	result := testUsersService.GetAll(testContext)

	// Assert
	assert.Equal(t, testUserDTOs, result)
}

func TestUsersGetByIdShouldReturnTestUserDTO(t *testing.T) {
	// Arrange
	id1 := uint(1)
	testUser := models.User{
		ID:       &id1,
		Name:     "test name 1",
		Email:    "a@a.com",
		Password: "aaaaaaaa",
	}
	testUserDTO := dtos.UserDTO{
		ID:       &id1,
		Name:     "test name 1",
		Email:    "a@a.com",
		Password: "",
	}
	testContext := context.Background()
	mockedUsersRepository := new(MockedUsersRepository)
	mockedUsersRepository.mock.On("GetById", testContext, &id1).Return(testUser)

	// Act
	testUsersService := NewUsersService(mockedUsersRepository)
	result := testUsersService.GetById(testContext, &id1)

	// Assert
	assert.Equal(t, testUserDTO, result)
}

func TestUsersGetByEmailShouldReturnTestUserDTO(t *testing.T) {
	// Arrange
	id1 := uint(1)
	testEmail := "a@a.com"
	testUser := models.User{
		ID:       &id1,
		Name:     "test name 1",
		Email:    "a@a.com",
		Password: "aaaaaaaa",
	}
	testUserDTO := dtos.UserDTO{
		ID:       &id1,
		Name:     "test name 1",
		Email:    "a@a.com",
		Password: "",
	}
	testContext := context.Background()
	mockedUsersRepository := new(MockedUsersRepository)
	mockedUsersRepository.mock.On("GetByEmail", testContext, &testEmail).Return(testUser)

	// Act
	testUsersService := NewUsersService(mockedUsersRepository)
	result := testUsersService.GetByEmail(testContext, &testEmail)

	// Assert
	assert.Equal(t, testUserDTO, result)
}

func TestUsersCreateShouldReturnNoErrorOnCreateAndGeneratedId(t *testing.T) {
	// Arrange
	id1 := uint(1)
	testUser := &models.User{
		Name:     "test name 1",
		Email:    "a@a.com",
		Password: "aaaaaaaa",
	}
	testUserDTO := &dtos.UserDTO{
		Name:     "test name 1",
		Email:    "a@a.com",
		Password: "aaaaaaaa",
	}
	testContext := context.Background()
	mockedUsersRepository := new(MockedUsersRepository)
	mockedUsersRepository.mock.On("Create", testContext, testUser).Return(nil)

	// Act
	testUsersService := NewUsersService(mockedUsersRepository)
	error := testUsersService.Create(testContext, testUserDTO)

	// Assert
	assert.Equal(t, error, nil)
	assert.NotNil(t, (*testUserDTO).ID)
	assert.Equal(t, (*testUserDTO).ID, &id1)
}

func TestUsersUpdateShouldReturnNoErrorOnUpdate(t *testing.T) {
	// Arrange
	id1 := uint(1)
	testUser := &models.User{
		ID:       &id1,
		Name:     "test updated 1",
		Email:    "a@a.com",
		Password: "aaaaaaaa",
	}
	testUserDTO := &dtos.UserDTO{
		ID:       &id1,
		Name:     "test updated 1",
		Email:    "a@a.com",
		Password: "aaaaaaaa",
	}
	testContext := context.Background()
	mockedUsersRepository := new(MockedUsersRepository)
	mockedUsersRepository.mock.On("Update", testContext, testUser).Return(nil)

	// Act
	testUsersService := NewUsersService(mockedUsersRepository)
	error := testUsersService.Update(testContext, testUserDTO)

	// Assert
	assert.Nil(t, error)
}

func TestUsersDeleteShouldReturnNoErrorOnDelete(t *testing.T) {
	// Arrange
	id1 := uint(1)
	testUser := &models.User{
		ID:       &id1,
		Name:     "test updated 1",
		Email:    "a@a.com",
		Password: "aaaaaaaa",
	}
	testUserDTO := &dtos.UserDTO{
		ID:       &id1,
		Name:     "test updated 1",
		Email:    "a@a.com",
		Password: "aaaaaaaa",
	}
	testContext := context.Background()
	mockedUsersRepository := new(MockedUsersRepository)
	mockedUsersRepository.mock.On("Delete", testContext, testUser).Return(nil)

	// Act
	testUsersService := NewUsersService(mockedUsersRepository)
	error := testUsersService.Delete(testContext, testUserDTO)

	// Assert
	assert.Nil(t, error)
}
