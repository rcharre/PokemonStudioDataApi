package validation

import (
	"fmt"
	"strings"
)

type Validation struct {
	Key   string
	Value any
}

func NewValidation(key string, value any) *Validation {
	return &Validation{
		key,
		value,
	}
}

func (v *Validation) String() string {
	return fmt.Sprint(v.Key, "=", v.Value)
}

type ValidationError struct {
	Validations []*Validation
}

func NewValidationError(validations []*Validation) *ValidationError {
	return &ValidationError{
		validations,
	}
}

func (e *ValidationError) Error() string {
	var sb = strings.Builder{}

	sb.WriteString("validation error:\n")
	for _, validation := range e.Validations {
		sb.WriteString(validation.String())
		sb.WriteString("\n")
	}
	return sb.String()
}
