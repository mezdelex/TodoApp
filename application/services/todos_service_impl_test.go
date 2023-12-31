package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"todoapp.com/application/dtos"
	"todoapp.com/domain/models"
)

type MockedTodosRepository struct {
	mock mock.Mock
}

func (m *MockedTodosRepository) GetAll(context context.Context) []models.Todo {
	args := m.mock.Called(context)

	return args.Get(0).([]models.Todo)
}

func (m *MockedTodosRepository) Create(context context.Context, model *models.Todo) error {
	id1 := uint(1)
	args := m.mock.Called(context, model)
	model.ID = &id1

	return args.Error(0)
}

func (m *MockedTodosRepository) Update(context context.Context, model *models.Todo) error {
	args := m.mock.Called(context, model)

	return args.Error(0)
}

func (m *MockedTodosRepository) Delete(context context.Context, model *models.Todo) error {
	args := m.mock.Called(context, model)

	return args.Error(0)
}

func TestTodosGetAllShouldReturnTestTodoDTOs(t *testing.T) {
	// Arrange
	id1, id2 := uint(1), uint(2)
	testTodos := []models.Todo{
		{
			ID:          &id1,
			Name:        "test name 1",
			Description: "test description 1",
			IsCompleted: false,
		},
		{
			ID:          &id2,
			Name:        "test name 2",
			Description: "test description 2",
			IsCompleted: false,
		},
	}
	testTodoDTOs := []dtos.TodoDTO{
		{
			ID:          &id1,
			Name:        "test name 1",
			Description: "test description 1",
			IsCompleted: false,
		},
		{
			ID:          &id2,
			Name:        "test name 2",
			Description: "test description 2",
			IsCompleted: false,
		},
	}
	testContext := context.Background()
	mockedTodosRepository := new(MockedTodosRepository)
	mockedTodosRepository.mock.On("GetAll", testContext).Return(testTodos)

	// Act
	testTodosService := NewTodosService(mockedTodosRepository)
	result := testTodosService.GetAll(testContext)

	// Assert
	assert.Equal(t, testTodoDTOs, result)
}

func TestTodosCreateShouldReturnNoErrorOnCreateAndGeneratedId(t *testing.T) {
	// Arrange
	id1 := uint(1)
	testTodo := &models.Todo{
		Name:        "[test] name",
		Description: "[test] description",
		IsCompleted: false,
	}
	testTodoDTO := &dtos.TodoDTO{
		Name:        "[test] name",
		Description: "[test] description",
		IsCompleted: false,
	}
	testContext := context.Background()
	mockedTodosRepository := new(MockedTodosRepository)
	mockedTodosRepository.mock.On("Create", testContext, testTodo).Return(nil)

	// Act
	testTodosService := NewTodosService(mockedTodosRepository)
	error := testTodosService.Create(testContext, testTodoDTO)

	// Assert
	assert.Equal(t, error, nil)
	assert.NotNil(t, (*testTodoDTO).ID)
	assert.Equal(t, (*testTodoDTO).ID, &id1)
}

func TestTodosUpdateShouldReturnNoErrorOnUpdate(t *testing.T) {
	// Arrange
	id1 := uint(1)
	testTodo := &models.Todo{
		ID:          &id1,
		Name:        "[test] name",
		Description: "[test] description",
		IsCompleted: false,
	}
	testTodoDTO := &dtos.TodoDTO{
		ID:          &id1,
		Name:        "[test] name",
		Description: "[test] description",
		IsCompleted: false,
	}
	testContext := context.Background()
	mockedTodosRepository := new(MockedTodosRepository)
	mockedTodosRepository.mock.On("Update", testContext, testTodo).Return(nil)

	// Act
	testTodosService := NewTodosService(mockedTodosRepository)
	error := testTodosService.Update(testContext, testTodoDTO)

	// Assert
	assert.Nil(t, error)
}

func TestTodosDeleteShouldReturnNoErrorOnDelete(t *testing.T) {
	// Arrange
	id1 := uint(1)
	testTodo := &models.Todo{
		ID:          &id1,
		Name:        "[test] name",
		Description: "[test] description",
		IsCompleted: false,
	}
	testTodoDTO := &dtos.TodoDTO{
		ID:          &id1,
		Name:        "[test] name",
		Description: "[test] description",
		IsCompleted: false,
	}
	testContext := context.Background()
	mockedTodosRepository := new(MockedTodosRepository)
	mockedTodosRepository.mock.On("Delete", testContext, testTodo).Return(nil)

	// Act
	testTodosService := NewTodosService(mockedTodosRepository)
	error := testTodosService.Delete(testContext, testTodoDTO)

	// Assert
	assert.Nil(t, error)
}
