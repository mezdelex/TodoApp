package services

import (
	"context"

	"todoapp.com/application/dtos"
	"todoapp.com/application/errors"
	"todoapp.com/domain/interfaces"
	"todoapp.com/domain/models"
)

type TodosServiceImpl struct {
	todosRepository interfaces.TodosRepository
}

func NewTodosService(todosRepository interfaces.TodosRepository) interfaces.TodosService {
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
		return errors.Errors{}.FiberValidationError("Todo")
	}

	model := &models.Todo{}
	dto.To(model)

	error := ts.todosRepository.Create(context, model)
	(*dto).ID = (*model).ID

	return error
}

func (ts *TodosServiceImpl) Update(context context.Context, dto *dtos.TodoDTO) error {
	if !dto.ValidateUpdateAndDelete() {
		return errors.Errors{}.FiberValidationError("Todo")
	}

	model := &models.Todo{}
	dto.To(model)

	return ts.todosRepository.Update(context, model)
}

func (ts *TodosServiceImpl) Delete(context context.Context, dto *dtos.TodoDTO) error {
	if !dto.ValidateUpdateAndDelete() {
		return errors.Errors{}.FiberValidationError("Todo")
	}

	model := &models.Todo{}
	dto.To(model)

	return ts.todosRepository.Delete(context, model)
}
