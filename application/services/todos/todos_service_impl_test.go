package todos

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	dtos "todoapp.com/application/dtos/todo"
	models "todoapp.com/domain/models/todo"
)

// Shared Arrange
var id1, id2 uint = uint(1), uint(2)
var testTodos []models.Todo = []models.Todo{
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

type MockedTodosRepository struct {
	mock mock.Mock
}

func (m *MockedTodosRepository) GetAll(context context.Context) []models.Todo {
	m.mock.Called(context)
	return testTodos
}

func (m *MockedTodosRepository) Create(context context.Context, model *models.Todo) error {
	m.mock.Called(context)
	m.mock.Called(model)
	id := uint(1)
	model.ID = &id
	return nil
}

func (m *MockedTodosRepository) Update(context context.Context, model *models.Todo) error {
	m.mock.Called(context)
	m.mock.Called(model)
	return nil
}

func (m *MockedTodosRepository) Delete(context context.Context, model *models.Todo) error {
	m.mock.Called(context)
	m.mock.Called(model)
	return nil
}

// TODO: Useless here; maybe segregate
func (m *MockedTodosRepository) CleanUp(context context.Context) int64 {
	m.mock.Called(context)
	return 0
}

func TestGetAllShouldReturnTestTodoDTOs(t *testing.T) {
	// Arrange
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
