package todos

import (
	"context"

	dtos "todoapp.com/application/dtos/todo"
	"todoapp.com/application/errors"
	"todoapp.com/domain/interfaces/todos"
	models "todoapp.com/domain/models/todo"
)

type TodosServiceImpl struct {
	todosRepository todos.TodosRepository
}

func NewTodosService(todosRepository todos.TodosRepository) todos.TodosService {
	return &TodosServiceImpl{todosRepository: todosRepository}
}

func (ts *TodosServiceImpl) GetAll(context context.Context) []dtos.TodoDTO {
	dtosSlice := []dtos.TodoDTO{}
	entities := ts.todosRepository.GetAll(context)

	for _, entity := range entities {
		dto := &dtos.TodoDTO{}
		dto.From(&entity)
		dtosSlice = append(dtosSlice, *dto)
	}

	return dtosSlice
}

func (ts *TodosServiceImpl) Create(context context.Context, dto *dtos.TodoDTO) error {
	if !dto.ValidateCreate() {
		return errors.Errors.FiberValidationError(errors.Errors{}, "Todo")
	}

	model := &models.Todo{}
	dto.To(model)

	error := ts.todosRepository.Create(context, model)
	(*dto).ID = (*model).ID

	return error
}

func (ts *TodosServiceImpl) Update(context context.Context, dto *dtos.TodoDTO) error {
	if !dto.ValidateUpdateAndDelete() {
		return errors.Errors.FiberValidationError(errors.Errors{}, "Todo")
	}

	model := &models.Todo{}
	dto.To(model)

	return ts.todosRepository.Update(context, model)
}

func (ts *TodosServiceImpl) Delete(context context.Context, dto *dtos.TodoDTO) error {
	if !dto.ValidateUpdateAndDelete() {
		return errors.Errors.FiberValidationError(errors.Errors{}, "Todo")
	}

	model := &models.Todo{}
	dto.To(model)

	return ts.todosRepository.Delete(context, model)
}
