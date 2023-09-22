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
	var id uint = uint(1)
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

func TestUserValidateShouldReturnFalseIfNameIsNil(t *testing.T) {
	// Arrange
	userDTO := UserDTO{
		Email:    "a@a.com",
		Password: "aaaaaaaa",
	}

	// Act
	result := userDTO.Validate()

	// Assert
	assert.False(t, result)
}

func TestUserValidateShouldReturnFalseIfNameIsEmpty(t *testing.T) {
	// Arrange
	userDTO := UserDTO{
		Name:     "",
		Email:    "a@a.com",
		Password: "aaaaaaaa",
	}

	// Act
	result := userDTO.Validate()

	// Assert
	assert.False(t, result)
}

func TestUserValidateShouldReturnFalseIfEmailIsNil(t *testing.T) {
	// Arrange
	userDTO := UserDTO{
		Name:     "Test name",
		Password: "aaaaaaaa",
	}

	// Act
	result := userDTO.Validate()

	// Assert
	assert.False(t, result)
}

func TestUserValidateShouldReturnFalseIfEmailIsEmpty(t *testing.T) {
	// Arrange
	userDTO := UserDTO{
		Name:     "Test name",
		Email:    "",
		Password: "aaaaaaaa",
	}

	// Act
	result := userDTO.Validate()

	// Assert
	assert.False(t, result)
}

func TestUserValidateShouldReturnFalseIfEmailIsMalformed(t *testing.T) {
	// Arrange
	userDTO := UserDTO{
		Name:     "Test name",
		Email:    "a@.com",
		Password: "aaaaaaaa",
	}

	// Act
	result := userDTO.Validate()

	// Assert
	assert.False(t, result)
}

func TestUserValidateShouldReturnFalseIfIsPasswordIsNil(t *testing.T) {
	// Arrange
	userDTO := UserDTO{
		Name:  "Test name",
		Email: "a@a.com",
	}

	// Act
	result := userDTO.Validate()

	// Assert
	assert.False(t, result)
}

func TestUserValidateShouldReturnFalseIfIsPasswordIsEmpty(t *testing.T) {
	// Arrange
	userDTO := UserDTO{
		Name:     "Test name",
		Email:    "a@a.com",
		Password: "",
	}

	// Act
	result := userDTO.Validate()

	// Assert
	assert.False(t, result)
}

func TestUserValidateShouldReturnFalseIfIsPasswordIsShorterThanEightCharacters(t *testing.T) {
	// Arrange
	userDTO := UserDTO{
		Name:     "Test name",
		Email:    "a@a.com",
		Password: "aaaa",
	}

	// Act
	result := userDTO.Validate()

	// Assert
	assert.False(t, result)
}

func TestUserValidateShouldReturnFalseIfIsPasswordIsLongerThan16Characters(t *testing.T) {
	// Arrange
	userDTO := UserDTO{
		Name:     "Test name",
		Email:    "a@a.com",
		Password: "aaaaaaaaaaaaaaaaa",
	}

	// Act
	result := userDTO.Validate()

	// Assert
	assert.False(t, result)
}

func TestUserValidateShouldReturnTrueIfAllRequiredValuesArePresent(t *testing.T) {
	// Arrange
	userDTO := UserDTO{
		Name:     "Test name",
		Email:    "a@a.com",
		Password: "aaaaaaaaaaa",
	}

	// Act
	result := userDTO.Validate()

	// Assert
	assert.True(t, result)
}
