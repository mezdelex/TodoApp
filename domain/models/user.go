package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       *uint
	Name     string
	Email    string `gorm:"constraint:unique"`
	Password string

	Todos []Todo
}
