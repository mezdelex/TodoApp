package todo

import "todoapp.com/domain/models/todo"

func (t *TodoDTO) To(model *todo.Todo) {
	(*model).ID = (*t).ID
	(*model).Name = (*t).Name
	(*model).Description = (*t).Description
	(*model).IsCompleted = (*t).IsCompleted
}

func (t *TodoDTO) From(model *todo.Todo) {
	(*t).ID = (*model).ID
	(*t).Name = (*model).Name
	(*t).Description = (*model).Description
	(*t).IsCompleted = (*model).IsCompleted
}
