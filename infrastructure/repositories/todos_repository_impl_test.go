package repositories

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"todoapp.com/domain/models"
	"todoapp.com/infrastructure/connectors/postgre"
)

func TestCreateShouldAddIdToThePassedModel(t *testing.T) {
	// Arrange
	todo := &models.Todo{
		ID:          nil,
		Name:        "test name 1",
		Description: "test description 1",
		IsCompleted: false,
	}
	testContext := context.Background()

	// Act
	// TODO: Check how to share DB context (or remove test)
	assert.FailNow(t, fmt.Sprintf("DB is: %v", postgre.DB))
	testTodosRepository := NewTodosRepository(postgre.DB)
	error := testTodosRepository.Create(testContext, todo)
	if error != nil {
		assert.FailNow(t, "Test todo Creation failed.")
	}

	// Assert
	assert.NotNil(t, todo.ID)
}
