package dtos

import (
	"regexp"
	"strings"

	"todoapp.com/domain/models"
)

type UserDTO struct {
	ID       *uint
	Name     string
	Email    string
	Password string
}

// Mapper
func (u *UserDTO) To(model *models.User) {
	(*model).ID = (*u).ID
	(*model).Name = (*u).Name
	(*model).Email = (*u).Email
}

func (u *UserDTO) From(model *models.User) {
	(*u).ID = (*model).ID
	(*u).Name = (*model).Name
	(*u).Email = (*model).Email
}

// Validator
func (dto *UserDTO) Validate() bool {
	rfc2822EmailPattern := `[a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?`
	trimmedEmail := strings.TrimSpace((*dto).Email)
	trimmedPassword := strings.TrimSpace((*dto).Password)

	isOk, _ := regexp.MatchString(rfc2822EmailPattern, trimmedEmail)

	// TODO: check password for certain special characters
	if strings.TrimSpace((*dto).Name) == "" || trimmedEmail == "" || !isOk || len(trimmedPassword) < 8 || len(trimmedPassword) > 16 {
		return false
	}

	return true
}

func (dto *UserDTO) ValidateCreate() bool {
	if (*dto).ID != nil || !dto.Validate() {
		return false
	}

	return true
}

func (dto *UserDTO) ValidateUpdateAndDelete() bool {
	if (*dto).ID == nil || !dto.Validate() {
		return false
	}

	return true
}
