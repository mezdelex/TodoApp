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

func TestUsersCreateUpdateAndDeleteIntegration(t *testing.T) {
	// Arrange
	user := &models.User{
		ID:       nil,
		Name:     "[test] name 1",
		Email:    "a@a.com",
		Password: "aassddee",
	}
	testContext := context.Background()

	// Act
	error := environments.LoadEnv()
	if error != nil {
		t.Skip()
	}
	db := connectors.Postgre{}.Connect()
	testUsersRepository := NewUsersRepository(db)
	error = testUsersRepository.Create(testContext, user)
	if error != nil {
		assert.FailNow(t, "Test User creation failed.")
	}

	// Assert
	assert.NotNil(t, user.ID)
	testUsersUpdate(t, user, testContext, testUsersRepository)
}

func testUsersUpdate(t *testing.T, user *models.User, testContext context.Context, testUsersRepository interfaces.UsersRepository) {
	// Arrange
	user.Name = "[test] name modified"

	// Act
	error := testUsersRepository.Update(testContext, user)
	if error != nil {
		assert.FailNow(t, "Test User update failed.")
	}

	// Assert
	assert.Nil(t, error)
	testUsersDelete(t, user, testContext, testUsersRepository)
}

func testUsersDelete(t *testing.T, user *models.User, testContext context.Context, testUsersRepository interfaces.UsersRepository) {
	// Arrange
	// Act
	error := testUsersRepository.Delete(testContext, user)
	if error != nil {
		assert.FailNow(t, "Test User delete failed.")
	}

	// Assert
	assert.Nil(t, error)
	testUsersCleanUp(t, testContext, testUsersRepository)
}

func testUsersCleanUp(t *testing.T, testContext context.Context, testUsersRepository interfaces.UsersRepository) {
	// Arrange
	// Act
	rowsAffected := testUsersRepository.CleanUp(testContext)

	// Assert
	assert.NotZero(t, rowsAffected)
}
