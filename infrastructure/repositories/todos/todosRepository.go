package todos

import (
	"todoapp.com/domain/models/todo"
	"todoapp.com/infrastructure/connectors/postgre"
)

func GetAll() *[]todo.Todo {
	db := postgre.DB
	var todos = new([]todo.Todo)

	db.Find(todos)

	return todos
}

func Create(t *todo.Todo) (*todo.Todo, error) {
	db := postgre.DB

	error := db.Create(t).Error
	if error != nil {
		return nil, error
	}

	return t, nil
}
