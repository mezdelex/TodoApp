package services

import (
	"context"

	"todoapp.com/application/dtos"
	"todoapp.com/application/errors"
	"todoapp.com/domain/interfaces"
	"todoapp.com/domain/models"
)

type UsersServiceImpl struct {
	usersRepository interfaces.BaseUsersRepositoryWithEmail
}

func NewUsersService(usersRepository interfaces.BaseUsersRepositoryWithEmail) *UsersServiceImpl {
	return &UsersServiceImpl{usersRepository: usersRepository}
}

func (us *UsersServiceImpl) GetAll(context context.Context) []dtos.UserDTO {
	dtosSlice := []dtos.UserDTO{}
	entities := us.usersRepository.GetAll(context)

	for _, entity := range entities {
		dto := &dtos.UserDTO{}
		dto.From(&entity)
		dtosSlice = append(dtosSlice, *dto)
	}

	return dtosSlice
}

func (us *UsersServiceImpl) GetById(context context.Context, id *uint) dtos.UserDTO {
	dto := &dtos.UserDTO{}
	entity := us.usersRepository.GetById(context, id)

	dto.From(&entity)

	return (*dto)
}

func (us *UsersServiceImpl) GetByEmail(context context.Context, email *string) dtos.UserDTO {
	dto := &dtos.UserDTO{}
	entity := us.usersRepository.GetByEmail(context, email)

	dto.From(&entity)

	return (*dto)
}

func (us *UsersServiceImpl) Create(context context.Context, dto *dtos.UserDTO) error {
	if !dto.ValidateCreate() {
		return errors.Errors{}.FiberValidationError("User")
	}

	model := &models.User{}
	dto.To(model)

	error := us.usersRepository.Create(context, model)
	(*dto).ID = (*model).ID

	return error
}

func (us *UsersServiceImpl) Update(context context.Context, dto *dtos.UserDTO) error {
	if !dto.ValidateUpdateAndDelete() {
		return errors.Errors{}.FiberValidationError("User")
	}

	model := &models.User{}
	dto.To(model)

	return us.usersRepository.Update(context, model)
}

func (us *UsersServiceImpl) Delete(context context.Context, dto *dtos.UserDTO) error {
	if !dto.ValidateUpdateAndDelete() {
		return errors.Errors{}.FiberValidationError("User")
	}

	model := &models.User{}
	dto.To(model)

	return us.usersRepository.Delete(context, model)
}
