package dtos

import (
	"regexp"
	"strings"
)

type LoginDTO struct {
	Email    string
	Password string
	Token    *string
}

// Validator
func (dto *LoginDTO) Validate() bool {
	rfc2822EmailPattern := `[a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?`
	trimmedEmail := strings.TrimSpace((*dto).Email)
	trimmedPassword := strings.TrimSpace((*dto).Password)

	isOk, _ := regexp.MatchString(rfc2822EmailPattern, trimmedEmail)

	// TODO: check password for certain special characters
	if !isOk || len(trimmedPassword) < 8 || len(trimmedPassword) > 16 {
		return false
	}

	return true
}
