package todos

import (
	"context"

	"todoapp.com/domain/models/todo"
)

type TodosRepository interface {
	GetAll(context context.Context) []todo.Todo
	Create(context context.Context, model *todo.Todo) error
	Update(context context.Context, model *todo.Todo) error
	Delete(context context.Context, model *todo.Todo) error
	CleanUp(context context.Context) int64
}
