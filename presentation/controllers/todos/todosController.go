package todos

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"todoapp.com/application/dtos/todo"
	"todoapp.com/application/services/todos"
	"todoapp.com/presentation/messages"
)

func SetupTodosController(router fiber.Router) {
	todosRouter := router.Group("/todos")

	todosRouter.Get("/", func(context *fiber.Ctx) error {
		todoDTOs := todos.GetAll()

		if len(todoDTOs) == 0 {
			return context.Status(200).JSON(fiber.Map{
				"data":    todoDTOs,
				"message": messages.Messages{}.CollectionEmptyMessage("Todo"),
				"status":  messages.Status{}.Success(),
			})
		}

		return context.Status(200).JSON(fiber.Map{
			"data":    todoDTOs,
			"message": messages.Messages{}.ReturningItemsMessage(len(todoDTOs), "todo"),
			"status":  messages.Status{}.Success(),
		})
	})

	todosRouter.Post("/", func(context *fiber.Ctx) error {
		newTodo := new(todo.TodoDTO)
		error := context.BodyParser(newTodo)
		if error != nil {
			return context.Status(400).JSON(fiber.Map{
				"message": messages.Messages{}.ParsingErrorMessage("Todo"),
				"status":  messages.Status{}.Error(),
			})
		}

		createdTodo, error := todos.Create(newTodo)
		if error != nil {
			var e *fiber.Error
			errors.As(error, &e)

			return context.Status(e.Code).JSON(fiber.Map{
				"data":    createdTodo,
				"message": e.Message,
				"status":  messages.Status{}.Error(),
			})
		}

		return context.Status(201).JSON(fiber.Map{
			"data":    createdTodo,
			"message": messages.Messages{}.ItemCreatedMessage(createdTodo),
			"status":  "success",
		})
	})

	todosRouter.Put("/:id", func(context *fiber.Ctx) error {
		todoToUpdate := new(todo.TodoDTO)
		error := context.BodyParser(todoToUpdate)
		if error != nil {
			return context.Status(400).JSON(fiber.Map{
				"message": messages.Messages{}.ParsingErrorMessage("Todo"),
				"status":  messages.Status{}.Error(),
			})
		}

		i, error := strconv.Atoi(context.Params("id"))
		if error != nil {
			return context.Status(400).JSON(fiber.Map{
				"message": messages.Messages{}.RouteFormatErrorMessage("id"),
				"status":  messages.Status{}.Error(),
			})
		}
		if todoToUpdate.ID == nil || uint(i) != *todoToUpdate.ID {
			return context.Status(400).JSON(fiber.Map{
				"message": messages.Messages{}.UpdateIdsConflictMessage("Todo"),
				"status":  messages.Status{}.Error(),
			})
		}

		isUpdated, error := todos.Update(todoToUpdate)
		if error != nil {
			var e *fiber.Error
			errors.As(error, &e)

			return context.Status(e.Code).JSON(fiber.Map{
				"data":    todoToUpdate,
				"message": e.Message,
				"status":  messages.Status{}.Error(),
			})
		}

		return context.Status(200).JSON(fiber.Map{
			"data":    isUpdated,
			"message": messages.Messages{}.ValuesUpdatedSuccessfullyMessage("Todo"),
			"status":  messages.Status{}.Success(),
		})
	})

	todosRouter.Delete("/:id", func(context *fiber.Ctx) error {
		i, error := strconv.Atoi(context.Params("id"))
		if error != nil {
			return context.Status(400).JSON(fiber.Map{
				"message": messages.Messages{}.RouteFormatErrorMessage("id"),
				"status":  messages.Status{}.Error(),
			})
		}

		id := uint(i)
		isDeleted, error := todos.Delete(id)
		if error != nil {
			var e *fiber.Error
			errors.As(error, &e)

			return context.Status(e.Code).JSON(fiber.Map{
				"data":    isDeleted,
				"message": e.Message,
				"status":  messages.Status{}.Error(),
			})
		}

		return context.Status(200).JSON(fiber.Map{
			"data":    isDeleted,
			"message": messages.Messages{}.ItemDeletedSuccessfullyMessage("Todo", id),
			"status":  messages.Status{}.Success(),
		})
	})
}
