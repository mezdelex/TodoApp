package todos

import (
	"github.com/stroiman/go-automapper"
	"todoapp.com/application/dtos/todo"
	"todoapp.com/infrastructure/repositories/todos"
)

func GetAll() []todo.TodoDTO {
	entities := todos.GetAll()
	var dtos []todo.TodoDTO

	for entity := range entities {
		var dto todo.TodoDTO
		automapper.Map(&entity, &dto)
		dtos = append(dtos, dto)
	}

	return dtos
}
