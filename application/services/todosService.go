package todosservice

import (
	"github.com/stroiman/go-automapper"
	tododto "todoapp.com/application/dtos"
	todosrepository "todoapp.com/infrastructure/repositories"
)

func GetTodos() []tododto.TodoDTO {
	entities := todosrepository.GetTodos()
	var dtos []tododto.TodoDTO

	for entity := range entities {
		var dto tododto.TodoDTO
		automapper.Map(&entity, &dto)
		dtos = append(dtos, dto)
	}

	return dtos
}
