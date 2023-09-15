package dtos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateUpdateAndDeleteShouldReturnFalseOnNilId(t *testing.T) {
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

func TestValidateCreateReturnFalseOnNonNilId(t *testing.T) {
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

func TestValidateReturnFalseIfNameIsNil(t *testing.T) {
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

func TestValidateReturnFalseIfNameIsEmpty(t *testing.T) {
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

func TestValidateReturnFalseIfDescriptionIsNil(t *testing.T) {
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

func TestValidateReturnFalseIfDescriptionIsEmpty(t *testing.T) {
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

func TestValidateReturnFalseIfIsCompletedIsNotFalse(t *testing.T) {
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
