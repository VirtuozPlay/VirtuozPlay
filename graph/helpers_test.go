package graph

import (
	"fmt"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/stretchr/testify/suite"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"sort"
	"testing"
)

// This file will not be regenerated automatically.

type HelpersSuite struct {
	suite.Suite
}

func (suite *HelpersSuite) Test_WrapValidationErrorsNoErrors() {
	noErrors := validate.NewErrors()
	suite.NoError(WrapValidationErrors(nil, nil))
	suite.NoError(WrapValidationErrors(noErrors, nil))
}

func (suite *HelpersSuite) Test_WrapValidationErrorsPassThrough() {
	err := fmt.Errorf("some error")
	suite.Equal(err, WrapValidationErrors(nil, err))
}

func (suite *HelpersSuite) Test_WrapValidationErrorSingle() {
	err := WrapValidationErrors(validate.Validate(&validators.StringIsPresent{Name: "str", Field: "", Message: "str is required"}), nil)
	suite.Len(err, 1)

	if errors, ok := err.(gqlerror.List); ok {
		suite.EqualError(errors[0], "input: str is required")
		return
	}
	suite.Fail("expected gqlerror.List")
}

func (suite *HelpersSuite) Test_WrapValidationErrorMultiple() {
	err := WrapValidationErrors(validate.Validate(
		&validators.StringIsPresent{Name: "str", Field: "", Message: "str is required"},
		&validators.StringsMatch{Name: "str", Field: "", Field2: "toMatch", Message: "no match"},
		&validators.IntIsGreaterThan{Name: "int_test", Field: 21, Compared: 42, Message: "21 is indeed not greater than 42"}), nil)

	suite.Len(err, 3)
	if errors, ok := err.(gqlerror.List); ok {
		sort.Slice(errors, func(i, j int) bool {
			return errors[i].Error() < errors[j].Error()
		})
		suite.EqualError(errors[0], "input: 21 is indeed not greater than 42")
		suite.EqualError(errors[1], "input: no match")
		suite.EqualError(errors[2], "input: str is required")
		return
	}
	suite.Fail("expected gqlerror.List")
}

func Test_Helpers(t *testing.T) {
	suite.Run(t, new(HelpersSuite))
}
