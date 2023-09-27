package interfaces

import (
	"context"

	"todoapp.com/application/dtos"
)

type UsersService interface {
	GetAll(context context.Context) []dtos.UserDTO
	GetById(context context.Context, id *uint) dtos.UserDTO
	Create(context context.Context, dto *dtos.UserDTO) error
	Update(context context.Context, dto *dtos.UserDTO) error
	Delete(context context.Context, dto *dtos.UserDTO) error
}

type ExtraUsersService interface {
	GetByEmail(context context.Context, email *string) dtos.UserDTO
}
