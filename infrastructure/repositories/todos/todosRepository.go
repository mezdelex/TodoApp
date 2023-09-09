package todos

import (
	"todoapp.com/domain/models/todo"
	"todoapp.com/infrastructure/connectors/postgre"
)

func GetAll() []todo.Todo {
	db := postgre.DB
	var todos []todo.Todo

	db.Find(&todos)

	return todos
}
