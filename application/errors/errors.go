package errors

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Errors struct{}

func (_ Errors) CannotReadFileError(fileName string) error {
	return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Could not read %s file.", fileName))
}

func (_ Errors) FiberValidationError(itemName string) error {
	return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Invalid %s.", itemName))
}

func (_ Errors) IncorrectPasswordError() error {
	return fiber.NewError(fiber.StatusUnauthorized, fmt.Sprintln("Incorrect password."))
}

func (_ Errors) ItemNotFoundError(itemName string) error {
	return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("%s not found.", itemName))
}
