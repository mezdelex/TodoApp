package repositories

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"todoapp.com/domain/models"
)

// TODO: Clean installed sql-mock and non used packages

// values of the DB
var mockedValues []models.Todo

type DBInstance struct {
	mockedDB mock.Mock
	err      error
}

func TestGetAllShouldReturnMockedDBValues(t *testing.T) {
	// Arrange
	mockedContext := context.Context{}
	repositoryWithMockedDB := NewTodosRepository(mockedDB)

	// Act
	result := repositoryWithMockedDB.GetAll(mockedContext)

	// Assert
	assert.Equal(t, mockedValues, result)
}
