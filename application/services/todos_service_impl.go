package services

import (
	"context"

	"todoapp.com/application/dtos"
	"todoapp.com/application/errors"
	servicesInterface "todoapp.com/application/services/interfaces"
	"todoapp.com/domain/models"
	repositoriesInterface "todoapp.com/infrastructure/repositories/interfaces"
)

type TodosServiceImpl struct {
	todosRepository repositoriesInterface.TodosRepository
}

func NewTodosService(todosRepository repositoriesInterface.TodosRepository) servicesInterface.TodosService {
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
