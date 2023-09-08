package todoscontroller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	todosservice "todoapp.com/application/services"
)

func SetupTodosController(router fiber.Router) {
	todosRouter := router.Group("/todos")

	todosRouter.Get("/", func(context *fiber.Ctx) error {
		todoDTOs := todosservice.GetTodos()

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
}
