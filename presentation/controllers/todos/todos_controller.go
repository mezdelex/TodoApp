package todos

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"todoapp.com/application/dtos/todo"
	interfaces "todoapp.com/domain/interfaces/todos"
	customErrors "todoapp.com/presentation/errors"
	"todoapp.com/presentation/messages"
)

type TodosController struct {
	todosService interfaces.TodosService
}

func NewTodosController(todosService interfaces.TodosService) *TodosController {
	return &TodosController{todosService: todosService}
}

func (tc *TodosController) Route(router fiber.Router) {
	todosRouter := router.Group("/todos")

	todosRouter.Get("/", tc.GetAll)
	todosRouter.Post("/", tc.Create)
	todosRouter.Put("/:id", tc.Update)
	todosRouter.Delete("/:id", tc.Delete)
}

func (tc *TodosController) GetAll(context *fiber.Ctx) error {
	todoDTOs := tc.todosService.GetAll(context.Context())

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
}

func (tc *TodosController) Create(context *fiber.Ctx) error {
	newTodo := &todo.TodoDTO{}
	error := context.BodyParser(newTodo)
	if error != nil {
		return context.Status(400).JSON(fiber.Map{
			"message": messages.Messages{}.ParsingErrorMessage("Todo"),
			"status":  messages.Status{}.Error(),
		})
	}

	error = tc.todosService.Create(context.Context(), newTodo)
	if error != nil {
		return customErrors.Errors{}.HandleFiberError(newTodo, context, error)
	}

	return context.Status(201).JSON(fiber.Map{
		"data":    newTodo,
		"message": messages.Messages{}.ItemCreatedMessage(newTodo),
		"status":  messages.Status{}.Success(),
	})

}

func (tc *TodosController) Update(context *fiber.Ctx) error {
	i, error := strconv.Atoi(context.Params("id"))
	if error != nil {
		return customErrors.Errors{}.RouteConversionError(context, "id")
	}

	todoToUpdate := &todo.TodoDTO{}
	error = context.BodyParser(todoToUpdate)
	if error != nil {
		return customErrors.Errors{}.ParsingError(context, "Todo")
	}

	if todoToUpdate.ID == nil || uint(i) != *(*todoToUpdate).ID {
		return customErrors.Errors{}.IdConflictError(context, "Todo")
	}

	error = tc.todosService.Update(context.Context(), todoToUpdate)
	if error != nil {
		return customErrors.Errors{}.HandleFiberError(todoToUpdate, context, error)
	}

	return context.Status(200).JSON(fiber.Map{
		"data":    todoToUpdate,
		"message": messages.Messages{}.ItemCreatedMessage(todoToUpdate),
		"status":  messages.Status{}.Success(),
	})

}

func (tc *TodosController) Delete(context *fiber.Ctx) error {
	i, error := strconv.Atoi(context.Params("id"))
	if error != nil {
		return customErrors.Errors{}.RouteConversionError(context, "id")
	}

	todoToDelete := &todo.TodoDTO{}
	error = context.BodyParser(todoToDelete)
	if error != nil {
		return customErrors.Errors{}.ParsingError(context, "Todo")
	}

	id := uint(i)
	if todoToDelete.ID == nil || id != *(*todoToDelete).ID {
		return customErrors.Errors{}.IdConflictError(context, "Todo")
	}

	error = tc.todosService.Delete(context.Context(), todoToDelete)
	if error != nil {
		return customErrors.Errors{}.HandleFiberError(todoToDelete, context, error)
	}

	return context.Status(200).JSON(fiber.Map{
		"data":    todoToDelete,
		"message": messages.Messages{}.ItemDeletedSuccessfullyMessage("Todo", id),
		"status":  messages.Status{}.Success(),
	})
}
