package controllers

import (
	"github.com/gofiber/fiber/v2"
	"todoapp.com/application/dtos"
	"todoapp.com/domain/interfaces"
	customErrors "todoapp.com/presentation/errors"
	"todoapp.com/presentation/messages"
)

type LoginController struct {
	loginService interfaces.LoginService
}

func NewLoginController(loginService interfaces.LoginService) *LoginController {
	return &LoginController{loginService: loginService}
}

func (lc *LoginController) Route(router fiber.Router) {
	usersRouter := router.Group("/login")

	usersRouter.Post("/", lc.Login)
}

func (lc *LoginController) Login(context *fiber.Ctx) error {
	newCredentials := &dtos.LoginDTO{}
	error := context.BodyParser(newCredentials)
	if error != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": messages.Messages{}.ParsingErrorMessage("Credentials"),
			"status":  messages.Status{}.Error(),
		})
	}

	error = lc.loginService.Login(context.Context(), newCredentials)
	if error != nil {
		return customErrors.Errors{}.HandleFiberError(newCredentials, context, error)
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    newCredentials.Token,
		"message": messages.Messages{}.LoggedInSuccessfullyMessage(),
		"status":  messages.Status{}.Success(),
	})
}
