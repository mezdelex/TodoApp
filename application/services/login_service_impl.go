package services

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func (ls *LoginServiceImpl) Login(context context.Context, login *dtos.LoginDTO) error {
	user := ls.userService.(interfaces.ExtraUsersService).GetByEmail(context, &login.Email)

	if user.Password == "" {
		return errors.Errors{}.ItemNotFoundError("User")
	}
	if user.Password != (*login).Password {
		return errors.Errors{}.IncorrectPasswordError()
	}

	return ls.GenerateToken(login)
}

func (ls *LoginServiceImpl) GenerateToken(login *dtos.LoginDTO) error {
	byteSlice, error := os.ReadFile(ls.config.PrivateKeyPath)
	if error != nil {
		return errors.Errors{}.CannotReadFileError("OPENSSH private key")
	}
	// TODO: test if this works; might require some trial and error
	block, _ := pem.Decode(byteSlice)
	parsedResult, _ := x509.ParsePKCS8PrivateKey(block.Bytes)
	key := parsedResult.(*rsa.PrivateKey)

	claims := jwt.MapClaims{
		"email":      login.Email,
		"expiration": time.Now().UTC().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(&jwt.SigningMethodEd25519{}, claims)

	encodedToken, error := token.SignedString(key.N)
	login.Token = &encodedToken

	return error
}
