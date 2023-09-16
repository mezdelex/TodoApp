package todos

import (
	"context"

	"todoapp.com/application/dtos/todo"
)

type TodosService interface {
	GetAll(context context.Context) []todo.TodoDTO
	Create(context context.Context, dto *todo.TodoDTO) error
	Update(context context.Context, dto *todo.TodoDTO) error
	Delete(context context.Context, dto *todo.TodoDTO) error
}
