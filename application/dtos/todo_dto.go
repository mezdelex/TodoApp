package dtos

import (
	"strings"

	"todoapp.com/domain/models"
)

type TodoDTO struct {
	ID          *uint
	Name        string
	Description string
	IsCompleted bool
}

// Mapper
func (t *TodoDTO) To(model *models.Todo) {
	(*model).ID = (*t).ID
	(*model).Name = (*t).Name
	(*model).Description = (*t).Description
	(*model).IsCompleted = (*t).IsCompleted
}

func (t *TodoDTO) From(model *models.Todo) {
	(*t).ID = (*model).ID
	(*t).Name = (*model).Name
	(*t).Description = (*model).Description
	(*t).IsCompleted = (*model).IsCompleted
}

// Validator
func (dto *TodoDTO) Validate() bool {
	if strings.TrimSpace((*dto).Name) == "" || strings.TrimSpace((*dto).Description) == "" || (*dto).IsCompleted {
		return false
	}

	return true
}

func (dto *TodoDTO) ValidateCreate() bool {
	if (*dto).ID != nil || !dto.Validate() {
		return false
	}

	return true
}

func (dto *TodoDTO) ValidateUpdateAndDelete() bool {
	if (*dto).ID == nil || !dto.Validate() {
		return false
	}

	return true
}
