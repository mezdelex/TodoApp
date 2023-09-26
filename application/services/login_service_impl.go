package services

import (
	"context"
	"os"

	"todoapp.com/application/dtos"
	"todoapp.com/application/errors"
	"todoapp.com/domain/interfaces"
	"todoapp.com/domain/models"
)

type LoginServiceImpl struct {
	userService interfaces.UsersService
	config      *models.Config
}

func NewLoginService(userService interfaces.UsersService, config *models.Config) *LoginServiceImpl {
	return &LoginServiceImpl{
		userService: userService,
		config:      config,
	}
}

// TODO: Pass config struct
func (ls *LoginServiceImpl) Login(context context.Context, login *dtos.LoginDTO) error {
	// Get user from DB
	user := ls.userService.GetByEmail(context, &login.Email)
	if user.Password == "" {
		return errors.Errors{}.ItemNotFoundError("User")
	}

	// Check password
	if user.Password != (*login).Password {
		return errors.Errors{}.IncorrectPasswordError()
	}

	// Generate bearer token
	var error error
	login.Token, error = ls.generateToken(&login.Password)

	return error
}

func (ls *LoginServiceImpl) RefreshToken(context context.Context, login *dtos.LoginDTO) error {
	// TODO: refresh token logic
	return nil
}

func (ls *LoginServiceImpl) generateToken(passsword *string) (*string, error) {
	// jwt workflow
	os.Open("config.json")

	return nil, nil
}
