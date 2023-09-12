package interfaces

import (
	"context"

	"todoapp.com/domain/models"
)

type TodosRepository interface {
	GetAll(context *context.Context) []models.Todo
	Create(context *context.Context, model *models.Todo) error
	Update(context *context.Context, model *models.Todo) error
	Delete(context *context.Context, model *models.Todo) error
}
