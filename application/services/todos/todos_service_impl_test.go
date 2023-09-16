package todos

import (
	"context"
	// "testing"

	"github.com/stretchr/testify/mock"
	"todoapp.com/domain/models/todo"
)

// Shared Arrange
type MockedTodosRepository struct {
	mock mock.Mock
}

func (m *MockedTodosRepository) GetAll(context context.Context, testTodos []todo.Todo) []todo.Todo {
	return testTodos
}

// TODO: Continue here

// func TestGetAllShouldReturnTestTodos(t *testing.T) {
// 	// Arrange
// 	id1, id2 := uint(1), uint(2)
// 	testTodos := []models.Todo{
// 		models.Todo{
// 			ID:          &id1,
// 			Name:        "test name 1",
// 			Description: "test description 1",
// 			IsCompleted: false,
// 		},
// 		models.Todo{
// 			ID:          &id2,
// 			Name:        "test name 2",
// 			Description: "test description 2",
// 			IsCompleted: false,
// 		},
// 	}
// 	testContext := context.Background()
// 	mockedTodosRepository := new(MockedTodosRepository)
// 	mockedTodosRepository.mock.On("GetAll", testContext, testTodos).Return(testTodos)
//
// 	// Act
// 	testTodosService := NewTodosService(mockedTodosRepository)
// 	// Assert
//
// 	_ = testTodos
// }
