package interfaces

import (
	"context"

	"todoapp.com/application/dtos"
)

type UsersService interface {
	GetAll(context context.Context) []dtos.UserDTO
	Get(context context.Context, id *uint) dtos.UserDTO
	Create(context context.Context, dto *dtos.UserDTO) error
	Update(context context.Context, dto *dtos.UserDTO) error
	Delete(context context.Context, dto *dtos.UserDTO) error
}
