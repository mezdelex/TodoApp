package errors

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Errors struct{}

func (_ Errors) FiberValidationError(itemName string) error {
	return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Invalid %s.", itemName))
}
