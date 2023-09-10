package todos

import (
	"github.com/gofiber/fiber/v2"
	dto "todoapp.com/application/dtos/todo"
	model "todoapp.com/domain/models/todo"
	"todoapp.com/infrastructure/repositories/todos"
)

func GetAll() []dto.TodoDTO {
	entities := todos.GetAll()
	var dtos []dto.TodoDTO

	for _, entity := range *entities {
		var dto = new(dto.TodoDTO)
		dto.From(&entity)
		dtos = append(dtos, *dto)
	}

	return dtos
}

func Create(t *dto.TodoDTO) (*dto.TodoDTO, error) {
	if !t.ValidateCreate() {
		return nil, fiber.NewError(400, "Invalid Todo.")
	}

	var model = new(model.Todo)
	t.To(model)

	createdEntity, error := todos.Create(model)
	if error != nil {
		return nil, error
	}
	var createdDTO = new(dto.TodoDTO)
	createdDTO.From(createdEntity)

	return createdDTO, nil
}

func Update(t *dto.TodoDTO) (bool, error) {
	if !t.ValidateUpdate() {
		return false, fiber.NewError(400, "Invalid Todo.")
	}

	var model = new(model.Todo)
	t.To(model)

	isUpdated, error := todos.Update(model)
	if error != nil {
		return isUpdated, error
	}

	return isUpdated, nil
}

func Delete(id uint) (bool, error) {
	isDeleted, error := todos.Delete(id)
	if error != nil {
		return isDeleted, error
	}

	return isDeleted, nil
}
