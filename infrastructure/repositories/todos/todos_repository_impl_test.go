package todos

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	interfaces "todoapp.com/domain/interfaces/todos"
	"todoapp.com/domain/models/todo"
	"todoapp.com/infrastructure/connectors/postgre"
	"todoapp.com/infrastructure/environments"
)

func TestCreateUpdateAndDeleteIntegration(t *testing.T) {
	// Arrange
	todo := &todo.Todo{
		ID:          nil,
		Name:        "[test] name 1",
		Description: "[test] description 1",
		IsCompleted: false,
	}
	testContext := context.Background()

	// Act
	error := environments.LoadEnv()
	if error != nil {
		t.Skip()
	}
	db := postgre.Connect()
	testTodosRepository := NewTodosRepository(db)
	error = testTodosRepository.Create(testContext, todo)
	if error != nil {
		assert.FailNow(t, "Test todo Creation failed.")
	}

	// Assert
	assert.NotNil(t, todo.ID)
	testUpdate(t, todo, testContext, testTodosRepository)
}

func testUpdate(t *testing.T, todo *todo.Todo, testContext context.Context, testTodosRepository interfaces.TodosRepository) {
	// Arrange
	todo.Name = "[test] name modified"

	// Act
	error := testTodosRepository.Update(testContext, todo)
	if error != nil {
		assert.FailNow(t, "Test todo Update failed.")
	}

	// Assert
	assert.Nil(t, error)
	testDelete(t, todo, testContext, testTodosRepository)
}

func testDelete(t *testing.T, todo *todo.Todo, testContext context.Context, testTodosRepository interfaces.TodosRepository) {
	// Arrange
	// Act
	error := testTodosRepository.Delete(testContext, todo)
	if error != nil {
		assert.FailNow(t, "Test todo Delete failed.")
	}

	// Assert
	assert.Nil(t, error)
	testCleanUp(t, testContext, testTodosRepository)
}

func testCleanUp(t *testing.T, testContext context.Context, testTodosRepository interfaces.TodosRepository) {
	// Arrange
	// Act
	rowsAffected := testTodosRepository.CleanUp(testContext)

	// Assert
	assert.NotZero(t, rowsAffected)
}
