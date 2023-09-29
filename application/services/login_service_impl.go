package services

import (
	"context"
	"crypto/ed25519"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
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
	// byteSlice, error := os.ReadFile(ls.config.PrivateKeyPath)
	// TODO: change this
	raw, error := os.ReadFile("./mocks/id_ed25519")
	if error != nil {
		return errors.Errors{}.CannotReadFileError("OPENSSH private key")
	}

	block, _ := pem.Decode(raw)

	// TODO: Extract to dtos?
	type ed25519PrivKeyDTO struct {
		ObjectIdentifier struct{ ObjectIdentifier asn1.ObjectIdentifier }
		PrivateKey       []byte
		Version          int
	}

	fmt.Println(block)

	asn1PrivateKey := &ed25519PrivKeyDTO{}
	// TODO: continue here
	// TODO: continue here
	// TODO: continue here
	// TODO: continue here
	// TODO: continue here
	// This fails
	something, error := asn1.Unmarshal(block.Bytes, asn1PrivateKey)
	fmt.Println(something)
	fmt.Println(asn1PrivateKey)
	// 34 -> 32?
	privateKey := ed25519.NewKeyFromSeed(asn1PrivateKey.PrivateKey[2:])

	claims := jwt.MapClaims{
		"email":      login.Email,
		"expiration": time.Now().UTC().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(&jwt.SigningMethodEd25519{}, claims)
	fmt.Println(token)

	signedToken, error := token.SignedString(privateKey)
	login.Token = &signedToken

	return error
}
