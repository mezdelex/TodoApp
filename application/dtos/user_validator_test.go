package dtos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: if password check considers special characters, these should change too
func TestUserValidateUpdateAndDeleteShouldReturnFalseOnNilId(t *testing.T) {
	// Arrange
	userDTO := UserDTO{
		Name:     "Test name",
		Email:    "a@a.com",
		Password: "aaaaaaaa",
	}

	// Act
	result := userDTO.ValidateUpdateAndDelete()

	// Assert
	assert.False(t, result)
}

func TestUserValidateCreateShouldReturnFalseOnNonNilId(t *testing.T) {
	// Arrange
	id := uint(1)
	userDTO := UserDTO{
		ID:       &id,
		Name:     "Test name",
		Email:    "a@a.com",
		Password: "aaaaaaaa",
	}

	// Act
	result := userDTO.ValidateCreate()

	// Assert
	assert.False(t, result)
}

// TODO: continue here
func TestUserValidateShouldReturnFalseIfNameIsNil(t *testing.T) {
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

func TestUserValidateShouldReturnFalseIfNameIsEmpty(t *testing.T) {
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

func TestUserValidateShouldReturnFalseIfDescriptionIsNil(t *testing.T) {
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

func TestUserValidateShouldReturnFalseIfDescriptionIsEmpty(t *testing.T) {
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

func TestUserValidateShouldReturnFalseIfIsCompletedIsNotFalse(t *testing.T) {
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

func TestUserValidateShouldReturnTrueIfAllRequiredValuesArePresent(t *testing.T) {
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
