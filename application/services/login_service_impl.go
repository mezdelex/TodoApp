package services

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"todoapp.com/application/dtos"
	"todoapp.com/application/errors"
	"todoapp.com/domain/interfaces"
	"todoapp.com/domain/models"
)

type LoginServiceImpl struct {
	userRepository interfaces.UsersRepositoryEmail
	config         *models.Config
}

func NewLoginService(userRepository interfaces.UsersRepositoryEmail, config *models.Config) *LoginServiceImpl {
	return &LoginServiceImpl{
		userRepository: userRepository,
		config:         config,
	}
}

func (ls *LoginServiceImpl) Login(context context.Context, login *dtos.LoginDTO) error {
	dbUser := ls.userRepository.GetByEmail(context, &login.Email)

	if dbUser.Email == "" {
		return errors.Errors{}.ItemNotFoundError("User")
	}
	if dbUser.Password != (*login).Password {
		return errors.Errors{}.IncorrectPasswordError()
	}

	return ls.GenerateToken(login)
}

func (ls *LoginServiceImpl) GenerateToken(login *dtos.LoginDTO) error {
	encodedKey, error := os.ReadFile(ls.config.PrivateKeyPath)
	if error != nil {
		return errors.Errors{}.CannotReadFileError("OPENSSH private key")
	}

	privateKey, error := jwt.ParseRSAPrivateKeyFromPEM(encodedKey)
	if error != nil {
		return errors.Errors{}.ItemNotParsedError("OPENSSH private key")
	}

	claims := jwt.MapClaims{
		"email":      login.Email,
		"expiration": time.Now().UTC().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedToken, error := token.SignedString(privateKey)
	login.Token = &signedToken

	return error
}
