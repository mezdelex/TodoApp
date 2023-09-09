package todo

import model "todoapp.com/domain/models/todo"

type TodoDTO struct {
	ID          *uint
	Name        string
	Description string
	IsCompleted bool
}

func (t *TodoDTO) To(model *model.Todo) {
	(*model).ID = (*t).ID
	(*model).Name = (*t).Name
	(*model).Description = (*t).Description
	(*model).IsCompleted = (*t).IsCompleted
}

func (t *TodoDTO) From(model *model.Todo) {
	(*t).ID = (*model).ID
	(*t).Name = (*model).Name
	(*t).Description = (*model).Description
	(*t).IsCompleted = (*model).IsCompleted
}
