package interfaces

import (
	"context"

	"todoapp.com/application/dtos"
)

type TodosService interface {
	GetAll(context *context.Context) []dtos.TodoDTO
	Create(context *context.Context, dto *dtos.TodoDTO) error
	Update(context *context.Context, dto *dtos.TodoDTO) error
	Delete(context *context.Context, dto *dtos.TodoDTO) error
}
