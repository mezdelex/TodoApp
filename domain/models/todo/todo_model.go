package todo

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ID          *uint
	Name        string
	Description string
	IsCompleted bool
}
