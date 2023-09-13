package errors

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Errors struct{}

func (_ Errors) FiberValidationError(itemName string) error {
	return fiber.NewError(400, fmt.Sprintf("Invalid %s.", itemName))
}