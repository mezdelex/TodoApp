package repositories

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"todoapp.com/domain/interfaces"
	"todoapp.com/domain/models"
	"todoapp.com/infrastructure/connectors"
	"todoapp.com/infrastructure/environments"
)

func TestTodosCreateUpdateAndDeleteIntegration(t *testing.T) {
	// Arrange
	todo := &models.Todo{
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
	db := connectors.Postgre{}.Connect()
	testTodosRepository := NewTodosRepository(db)
	error = testTodosRepository.Create(testContext, todo)
	if error != nil {
		assert.FailNow(t, "Test Todo creation failed.")
	}

	// Assert
	assert.NotNil(t, todo.ID)
	testTodosUpdate(t, todo, testContext, testTodosRepository)
}

func testTodosUpdate(t *testing.T, todo *models.Todo, testContext context.Context, testTodosRepository interfaces.TodosRepository) {
	// Arrange
	todo.Name = "[test] name modified"

	// Act
	error := testTodosRepository.Update(testContext, todo)
	if error != nil {
		assert.FailNow(t, "Test Todo update failed.")
	}

	// Assert
	assert.Nil(t, error)
	testTodosDelete(t, todo, testContext, testTodosRepository)
}

func testTodosDelete(t *testing.T, todo *models.Todo, testContext context.Context, testTodosRepository interfaces.TodosRepository) {
	// Arrange
	// Act
	error := testTodosRepository.Delete(testContext, todo)
	if error != nil {
		assert.FailNow(t, "Test Todo delete failed.")
	}

	// Assert
	assert.Nil(t, error)
	testTodosCleanUp(t, testContext, testTodosRepository)
}

func testTodosCleanUp(t *testing.T, testContext context.Context, testTodosRepository interfaces.TodosRepository) {
	// Arrange
	// Act
	rowsAffected := testTodosRepository.(interfaces.ExtraTodosRepository).CleanUp(testContext)

	// Assert
	assert.NotZero(t, rowsAffected)
}
