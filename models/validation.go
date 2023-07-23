package models

import (
	"github.com/gobuffalo/validate/v3"
)

type ValidationErrors struct {
	Wrapped    error
	Validation *validate.Errors
}

// Error implements the error interface for ValidationErrors.
func (e *ValidationErrors) Error() string {
	if e.Wrapped == nil {
		return e.Validation.Error()
	}
	return e.Wrapped.Error()
}

func WrapValidation(errors *validate.Errors, err error) *ValidationErrors {
	if err == nil && (errors == nil || !errors.HasAny()) {
		return nil
	}
	return &ValidationErrors{err, errors}
}
