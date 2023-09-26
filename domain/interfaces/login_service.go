package interfaces

import (
	"context"

	"todoapp.com/application/dtos"
)

type LoginService interface {
	Login(context context.Context, login *dtos.LoginDTO)
	RefreshToken()
}
