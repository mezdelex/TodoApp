package interfaces

import (
	"context"

	"todoapp.com/domain/models"
)

type UsersRepository interface {
	GetAll(context context.Context) []models.User
	// GetAllFiltered(context context.Context, params *models.User) []models.User
	Get(context context.Context, id *uint) models.User
	Create(context context.Context, model *models.User) error
	Update(context context.Context, model *models.User) error
	Delete(context context.Context, model *models.User) error
	CleanUp(context context.Context) int64
}
