package models

import (
	"errors"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
)

func (ms *ModelSuite) Test_ValidationErrors() {
	ms.Implements((*error)(nil), &ValidationErrors{})

	ms.Nil(WrapValidation(nil, nil))
	ms.Error(WrapValidation(nil, errors.New("yes")))
	ms.NotEmpty(WrapValidation(nil, errors.New("yes")).Error())

	valErrors := validate.Validate(&validators.StringIsPresent{
		Name:    "should_fail",
		Field:   "",
		Message: "should fail",
	})

	ms.Error(WrapValidation(valErrors, nil))
	ms.NotEmpty(WrapValidation(valErrors, nil).Error())
}
