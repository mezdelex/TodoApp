package interfaces

import (
	"context"

	"todoapp.com/domain/models"
)

type (
	UsersRepository interface {
		BaseUsersRepository
		UsersRepositoryEmail
		UsersRepositoryCleanUp
	}

	BaseUsersRepositoryWithEmail interface {
		BaseUsersRepository
		UsersRepositoryEmail
	}

	BaseUsersRepository interface {
		GetAll(context context.Context) []models.User
		GetById(context context.Context, id *uint) models.User
		Create(context context.Context, model *models.User) error
		Update(context context.Context, model *models.User) error
		Delete(context context.Context, model *models.User) error
	}

	UsersRepositoryEmail interface {
		GetByEmail(context context.Context, email *string) models.User
	}

	UsersRepositoryCleanUp interface {
		CleanUp(context context.Context) int64
	}
)
