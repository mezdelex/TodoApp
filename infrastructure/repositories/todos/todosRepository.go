package todos

import (
	"github.com/gofiber/fiber/v2"
	"todoapp.com/domain/models/todo"
	"todoapp.com/infrastructure/connectors/postgre"
)

func GetAll() *[]todo.Todo {
	db := postgre.DB
	var todos = new([]todo.Todo)

	db.Where("deleted_at is null").Find(todos)

	return todos
}

func Get(id uint) *todo.Todo {
	db := postgre.DB
	var todo = new(todo.Todo)

	db.Where("id = ? and deleted_at is null", id).Find(todo)

	return todo
}

func Create(t *todo.Todo) (*todo.Todo, error) {
	db := postgre.DB

	error := db.Create(t).Error
	if error != nil {
		return nil, error
	}

	return t, nil
}

func Update(t *todo.Todo) (bool, error) {
	db := postgre.DB

	entity := Get(*t.ID)
	if entity == nil {
		return false, fiber.NewError(404, "Couldn't find the requested Todo.")
	}

	(*entity).Name = (*t).Name
	(*entity).Description = (*t).Description
	(*entity).IsCompleted = (*t).IsCompleted

	error := db.Save(entity).Error
	if error != nil {
		return false, error
	}

	return true, nil
}

func Delete(id uint) (bool, error) {
	db := postgre.DB

	entity := Get(id)
	if (*entity).ID == nil {
		return false, fiber.NewError(404, "Couldn't find the requested Todo.")
	}

	error := db.Delete(entity, id).Error
	if error != nil {
		return false, error
	}

	return true, nil
}
