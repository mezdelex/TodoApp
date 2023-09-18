package todo

import "strings"

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
