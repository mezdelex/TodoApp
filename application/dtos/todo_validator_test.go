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

func TestValidateCreateShouldReturnFalseOnNonNilId(t *testing.T) {
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

func TestValidateShouldReturnFalseIfNameIsNil(t *testing.T) {
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

func TestValidateShouldReturnFalseIfNameIsEmpty(t *testing.T) {
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

func TestValidateShouldReturnFalseIfDescriptionIsNil(t *testing.T) {
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

func TestValidateShouldReturnFalseIfDescriptionIsEmpty(t *testing.T) {
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

func TestValidateShouldReturnFalseIfIsCompletedIsNotFalse(t *testing.T) {
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

func TestValidateShouldReturnTrueIfAllRequiredValuesArePresent(t *testing.T) {
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
