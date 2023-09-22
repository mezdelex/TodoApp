package dtos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodoValidateUpdateAndDeleteShouldReturnFalseOnNilId(t *testing.T) {
	// Arrange
	todoDTO := TodoDTO{
		Name:        "Test name",
		Description: "Test description",
		IsCompleted: false,
	}

	// Act
	result := todoDTO.ValidateUpdateAndDelete()

	// Assert
	assert.False(t, result)
}

func TestTodoValidateCreateShouldReturnFalseOnNonNilId(t *testing.T) {
	// Arrange
	id := uint(1)
	todoDTO := TodoDTO{
		ID:          &id,
		Name:        "Test name",
		Description: "Test description",
		IsCompleted: false,
	}

	// Act
	result := todoDTO.ValidateCreate()

	// Assert
	assert.False(t, result)
}

func TestTodoValidateShouldReturnFalseIfNameIsNil(t *testing.T) {
	// Arrange
	todoDTO := TodoDTO{
		Description: "Test description",
		IsCompleted: false,
	}

	// Act
	result := todoDTO.Validate()

	// Assert
	assert.False(t, result)
}

func TestTodoValidateShouldReturnFalseIfNameIsEmpty(t *testing.T) {
	// Arrange
	todoDTO := TodoDTO{
		Name:        "",
		Description: "Test description",
		IsCompleted: false,
	}

	// Act
	result := todoDTO.Validate()

	// Assert
	assert.False(t, result)
}

func TestTodoValidateShouldReturnFalseIfDescriptionIsNil(t *testing.T) {
	// Arrange
	todoDTO := TodoDTO{
		Name:        "Test name",
		IsCompleted: false,
	}

	// Act
	result := todoDTO.Validate()

	// Assert
	assert.False(t, result)
}

func TestTodoValidateShouldReturnFalseIfDescriptionIsEmpty(t *testing.T) {
	// Arrange
	todoDTO := TodoDTO{
		Name:        "Test name",
		Description: "",
		IsCompleted: false,
	}

	// Act
	result := todoDTO.Validate()

	// Assert
	assert.False(t, result)
}

func TestTodoValidateShouldReturnFalseIfIsCompletedIsNotFalse(t *testing.T) {
	// Arrange
	todoDTO := TodoDTO{
		Name:        "Test name",
		Description: "Test description",
		IsCompleted: true,
	}

	// Act
	result := todoDTO.Validate()

	// Assert
	assert.False(t, result)
}

func TestTodoValidateShouldReturnTrueIfAllRequiredValuesArePresent(t *testing.T) {
	// Arrange
	todoDTO := TodoDTO{
		Name:        "Test name",
		Description: "Test description",
		IsCompleted: false,
	}

	// Act
	result := todoDTO.Validate()

	// Assert
	assert.True(t, result)
}
