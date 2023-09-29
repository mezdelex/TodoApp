package dtos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginValidateShouldReturnFalseIfEmailIsNil(t *testing.T) {
	// Arrange
	loginDTO := LoginDTO{
		Password: "aaaaaaaa",
	}

	// Act
	result := loginDTO.Validate()

	// Assert
	assert.False(t, result)
}

func TestLoginValidateShouldReturnFalseIfEmailIsEmpty(t *testing.T) {
	// Arrange
	loginDTO := LoginDTO{
		Email:    "",
		Password: "aaaaaaaa",
	}

	// Act
	result := loginDTO.Validate()

	// Assert
	assert.False(t, result)
}

func TestLoginValidateShouldReturnFalseIfEmailIsMalformed(t *testing.T) {
	// Arrange
	loginDTO := LoginDTO{
		Email:    "a@.com",
		Password: "aaaaaaaa",
	}

	// Act
	result := loginDTO.Validate()

	// Assert
	assert.False(t, result)
}

func TestLoginValidateShouldReturnFalseIfPasswordIsNil(t *testing.T) {
	// Arrange
	loginDTO := LoginDTO{
		Email: "a@a.com",
	}

	// Act
	result := loginDTO.Validate()

	// Assert
	assert.False(t, result)
}

func TestLoginValidateShouldReturnFalseIfPasswordIsEmpty(t *testing.T) {
	// Arrange
	loginDTO := LoginDTO{
		Email:    "a@a.com",
		Password: "",
	}

	// Act
	result := loginDTO.Validate()

	// Assert
	assert.False(t, result)
}

func TestLoginValidateShouldReturnFalseIfPasswordIsShorterThanEightCharacters(t *testing.T) {
	// Arrange
	loginDTO := LoginDTO{
		Email:    "a@a.com",
		Password: "aaaa",
	}

	// Act
	result := loginDTO.Validate()

	// Assert
	assert.False(t, result)
}

func TestLoginValidateShouldReturnFalseIfPasswordIsLongerThanSixteenCharacters(t *testing.T) {
	// Arrange
	loginDTO := LoginDTO{
		Email:    "a@a.com",
		Password: "aaaaaaaaaaaaaaaaa",
	}

	// Act
	result := loginDTO.Validate()

	// Assert
	assert.False(t, result)
}

func TestLoginValidateShouldReturnTrueIfAllRequiredValuesArePresent(t *testing.T) {
	// Arrange
	loginDTO := LoginDTO{
		Email:    "a@a.com",
		Password: "aaaaaaaaaaa",
	}

	// Act
	result := loginDTO.Validate()

	// Assert
	assert.True(t, result)
}
