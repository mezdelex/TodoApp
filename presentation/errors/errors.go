package errors

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"todoapp.com/presentation/messages"
)

type Errors struct{}

func (_ Errors) HandleFiberError(item interface{}, context *fiber.Ctx, error error) error {
	var e *fiber.Error
	errors.As(error, &e)

	return context.Status(e.Code).JSON(fiber.Map{
		"data":    item,
		"message": e.Message,
		"status":  messages.Status{}.Error(),
	})
}

func (_ Errors) IdConflictError(context *fiber.Ctx, itemName string) error {
	return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": messages.Messages{}.IdConflictErrorMessage(itemName), "status": messages.Status{}.Error()})
}

func (_ Errors) ParsingError(context *fiber.Ctx, itemName string) error {
	return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": messages.Messages{}.ParsingErrorMessage(itemName), "status": messages.Status{}.Error()})
}

func (_ Errors) RouteConversionError(context *fiber.Ctx, routeElement string) error {
	return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": messages.Messages{}.RouteFormatErrorMessage(routeElement), "status": messages.Status{}.Error()})
}
