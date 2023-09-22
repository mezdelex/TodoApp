package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"todoapp.com/application/dtos"
	"todoapp.com/domain/interfaces"
	customErrors "todoapp.com/presentation/errors"
	"todoapp.com/presentation/messages"
)

type UsersController struct {
	usersService interfaces.UsersService
}

func NewUsersController(usersService interfaces.UsersService) *UsersController {
	return &UsersController{usersService: usersService}
}

func (uc *UsersController) Route(router fiber.Router) {
	usersRouter := router.Group("/users")

	usersRouter.Get("/", uc.GetAll)
	usersRouter.Get("/:id", uc.Get)
	usersRouter.Post("/", uc.Create)
	usersRouter.Put("/:id", uc.Update)
	usersRouter.Delete("/:id", uc.Delete)
}

func (uc *UsersController) GetAll(context *fiber.Ctx) error {
	userDTOs := uc.usersService.GetAll(context.Context())

	if len(userDTOs) == 0 {
		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"data":    userDTOs,
			"message": messages.Messages{}.CollectionEmptyMessage("User"),
			"status":  messages.Status{}.Success(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    userDTOs,
		"message": messages.Messages{}.ReturningItemsMessage(len(userDTOs), "user"),
		"status":  messages.Status{}.Success(),
	})
}

func (uc *UsersController) Get(context *fiber.Ctx) error {
	i, error := strconv.Atoi(context.Params("id"))
	if error != nil {
		return customErrors.Errors{}.RouteConversionError(context, "id")
	}

	id := uint(i)
	userDTO := uc.usersService.Get(context.Context(), &id)

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    userDTO,
		"message": messages.Messages{}.ReturningItemsMessage(1, "user"),
		"status":  messages.Status{}.Success(),
	})
}

func (uc *UsersController) Create(context *fiber.Ctx) error {
	newUser := &dtos.UserDTO{}
	error := context.BodyParser(newUser)
	if error != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": messages.Messages{}.ParsingErrorMessage("User"),
			"status":  messages.Status{}.Error(),
		})
	}

	error = uc.usersService.Create(context.Context(), newUser)
	if error != nil {
		return customErrors.Errors{}.HandleFiberError(newUser, context, error)
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    newUser,
		"message": messages.Messages{}.ItemCreatedMessage(newUser),
		"status":  messages.Status{}.Success(),
	})

}

func (uc *UsersController) Update(context *fiber.Ctx) error {
	i, error := strconv.Atoi(context.Params("id"))
	if error != nil {
		return customErrors.Errors{}.RouteConversionError(context, "id")
	}

	userToUpdate := &dtos.UserDTO{}
	error = context.BodyParser(userToUpdate)
	if error != nil {
		return customErrors.Errors{}.ParsingError(context, "User")
	}

	if userToUpdate.ID == nil || uint(i) != *(*userToUpdate).ID {
		return customErrors.Errors{}.IdConflictError(context, "User")
	}

	error = uc.usersService.Update(context.Context(), userToUpdate)
	if error != nil {
		return customErrors.Errors{}.HandleFiberError(userToUpdate, context, error)
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    userToUpdate,
		"message": messages.Messages{}.ItemCreatedMessage(userToUpdate),
		"status":  messages.Status{}.Success(),
	})

}

func (uc *UsersController) Delete(context *fiber.Ctx) error {
	i, error := strconv.Atoi(context.Params("id"))
	if error != nil {
		return customErrors.Errors{}.RouteConversionError(context, "id")
	}

	userToDelete := &dtos.UserDTO{}
	error = context.BodyParser(userToDelete)
	if error != nil {
		return customErrors.Errors{}.ParsingError(context, "User")
	}

	id := uint(i)
	if userToDelete.ID == nil || id != *(*userToDelete).ID {
		return customErrors.Errors{}.IdConflictError(context, "User")
	}

	error = uc.usersService.Delete(context.Context(), userToDelete)
	if error != nil {
		return customErrors.Errors{}.HandleFiberError(userToDelete, context, error)
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    userToDelete,
		"message": messages.Messages{}.ItemDeletedSuccessfullyMessage("User", id),
		"status":  messages.Status{}.Success(),
	})
}
