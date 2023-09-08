package todosrepository

import (
	todo "todoapp.com/domain/models"
	postgre_connector "todoapp.com/infrastructure"
)

func GetTodos() []todo.Todo {
	db := postgre_connector.DB
	var todos []todo.Todo

	db.Find(&todos)

	return todos
}
