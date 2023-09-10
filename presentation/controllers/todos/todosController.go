package todos

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"todoapp.com/application/dtos/todo"
	"todoapp.com/application/services/todos"
)

func SetupTodosController(router fiber.Router) {
	todosRouter := router.Group("/todos")

	todosRouter.Get("/", func(context *fiber.Ctx) error {
		todoDTOs := todos.GetAll()

		if len(todoDTOs) == 0 {
			return context.Status(404).JSON(fiber.Map{
				"data":    todoDTOs,
				"message": "No todos left.",
				"status":  "error",
			})
		}

		return context.Status(200).JSON(fiber.Map{
			"data":    todoDTOs,
			"message": fmt.Sprintf("Returning %d todo(s).", len(todoDTOs)),
			"status":  "success",
		})
	})

	todosRouter.Post("/", func(context *fiber.Ctx) error {
		newTodo := new(todo.TodoDTO)
		error := context.BodyParser(newTodo)
		if error != nil {
			return context.Status(400).JSON(fiber.Map{
				"data":    nil,
				"message": "The provided Todo could not be parsed.",
				"status":  "error",
			})
		}

		createdTodo, error := todos.Create(newTodo)
		if error != nil {
			var e *fiber.Error
			errors.As(error, &e)

			return context.Status(e.Code).JSON(fiber.Map{
				"data":    createdTodo,
				"message": e.Message,
				"status":  "error",
			})
		}

		return context.Status(201).JSON(fiber.Map{
			"data":    createdTodo,
			"message": fmt.Sprintf("%v was created successfully", createdTodo),
			"status":  "success",
		})
	})

	todosRouter.Put("/:id", func(context *fiber.Ctx) error {
		todoToUpdate := new(todo.TodoDTO)
		error := context.BodyParser(todoToUpdate)
		if error != nil {
			return context.Status(400).JSON(fiber.Map{
				"data":    nil,
				"message": "The provided Todo could not be parsed.",
				"status":  "error",
			})
		}

		i, _ := strconv.Atoi(context.Params("id"))
		if todoToUpdate.ID == nil || uint(i) != *todoToUpdate.ID {
			return context.Status(400).JSON(fiber.Map{
				"data":    nil,
				"message": "The route id and the Todo's id are not equal.",
				"status":  "error",
			})
		}

		isUpdated, error := todos.Update(todoToUpdate)
		if error != nil {
			var e *fiber.Error
			errors.As(error, &e)

			return context.Status(e.Code).JSON(fiber.Map{
				"data":    todoToUpdate,
				"message": e.Message,
				"status":  "error",
			})
		}

		return context.Status(200).JSON(fiber.Map{
			"data":    isUpdated,
			"message": "Todo values updated successfully.",
			"status":  "success",
		})
	})
}
