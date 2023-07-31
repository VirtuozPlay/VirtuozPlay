package graph

import (
	"errors"
	"fmt"
	"sort"
	"testing"
	"time"
	"virtuozplay/models"

	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/stretchr/testify/suite"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// This file will not be regenerated automatically.

type HelpersSuite struct {
	suite.Suite
}

func (hs *HelpersSuite) Test_WrapValidationErrorsNoErrors() {
	noErrors := models.WrapValidation(validate.NewErrors(), nil)
	hs.NoError(wrapError(nil))
	hs.NoError(wrapError(noErrors))
}

func (hs *HelpersSuite) Test_WrapValidationErrorsPassThrough() {
	err := fmt.Errorf("some error")
	hs.Equal(err, wrapError(err))
}

func (hs *HelpersSuite) Test_WrapValidationErrorSingle() {
	err := wrapError(models.WrapValidation(validate.Validate(&validators.StringIsPresent{Name: "str", Field: "", Message: "str is required"}), nil))
	hs.Len(err, 1)

	if errs, ok := err.(gqlerror.List); ok {
		hs.EqualError(errs[0], "input: str is required")
		return
	}
	hs.Fail("expected gqlerror.List")
}

func (hs *HelpersSuite) Test_WrapValidationErrorMultiple() {
	err := wrapError(models.WrapValidation(validate.Validate(
		&validators.StringIsPresent{Name: "str", Field: "", Message: "str is required"},
		&validators.StringsMatch{Name: "str", Field: "", Field2: "toMatch", Message: "no match"},
		&validators.IntIsGreaterThan{Name: "int_test", Field: 21, Compared: 42, Message: "21 is indeed not greater than 42"}), nil))

	hs.Len(err, 3)
	if errs, ok := err.(gqlerror.List); ok {
		sort.Slice(errs, func(i, j int) bool {
			return errs[i].Error() < errs[j].Error()
		})
		hs.EqualError(errs[0], "input: 21 is indeed not greater than 42")
		hs.EqualError(errs[1], "input: no match")
		hs.EqualError(errs[2], "input: str is required")
		return
	}
	hs.Fail("expected gqlerror.List")
}

func (hs *HelpersSuite) Test_ToGraphQLPerformance() {
	perf, err := ToGraphQLPerformance(nil, nil)
	hs.Nil(perf)
	hs.NoError(err)

	perf, err = ToGraphQLPerformance(nil, errors.New("some error"))
	hs.Nil(perf)
	hs.EqualError(err, "some error")

	now := time.Now()

	perf, err = ToGraphQLPerformance(&models.Performance{NanoID: "nano", CreatedAt: now, Notes: []models.Note{{
		At:       91,
		Duration: 7121,
		Step:     "Bb",
	}}}, nil)

	hs.NoError(err)
	hs.Equal("nano", perf.ID)
	hs.Equal(now.String(), *perf.CreatedAt)
	hs.Len(perf.Notes, 1)
	hs.Equal(91, perf.Notes[0].At)
	hs.Equal(7121, perf.Notes[0].Duration)
	hs.Equal("Bb", perf.Notes[0].Value)
}

func Test_Helpers(t *testing.T) {
	suite.Run(t, new(HelpersSuite))
}
