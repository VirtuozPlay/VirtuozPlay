package models

import (
	"errors"
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

// WrapValidation wraps an error and/or validation errors into a single error.
func WrapValidation(errors *validate.Errors, err error) error {
	if err == nil && (errors == nil || !errors.HasAny()) {
		return nil
	}
	return &ValidationErrors{err, errors}
}

func UnwrapErrors(err error) error {
	if err == nil {
		return nil
	}
	var verr *ValidationErrors
	if errors.As(err, &verr) {
		if verr.Validation != nil {
			return verr.Validation
		}
		return verr.Wrapped
	}
	return err
}
