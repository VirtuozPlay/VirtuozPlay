package graph

import (
	"github.com/vektah/gqlparser/v2/gqlerror"
	"virtuozplay/graph/model"
	db "virtuozplay/models"
)

// This file will not be regenerated automatically.

// wrapError the errors of pop.ValidateAndCreate, pop.ValidateAndUpdate,
// or pop.ValidateAndSave to the error format expected by gqlgen.
func wrapError(err error) error {
	if vErrs, ok := err.(*db.ValidationErrors); ok {
		if vErrs == nil {
			return nil
		}
		if vErrs.Wrapped != nil {
			return vErrs.Wrapped
		}
		if vErrs.Validation == nil || !vErrs.Validation.HasAny() {
			return nil
		}
		keys := vErrs.Validation.Keys()
		gqlErrors := make(gqlerror.List, 0) // they are at least one error per key
		for _, key := range keys {
			for _, subError := range vErrs.Validation.Get(key) {
				gqlErrors = append(gqlErrors, &gqlerror.Error{
					Message: subError,
				})
			}
		}
		return gqlErrors
	}
	return err
}

// ToGraphQLPerformance converts an instance of the Performance model to its GraphQL representation.
func ToGraphQLPerformance(performance *db.Performance, err error) (*model.Performance, error) {
	if err = wrapError(err); err != nil {
		return nil, err
	}
	if performance == nil {
		panic("performance is nil")
	}

	notes := make([]*model.Note, len(performance.Notes))

	for i, note := range performance.Notes {
		notes[i] = &model.Note{
			At:       int(note.At),
			Duration: int(note.Duration),
			Value:    note.Value,
		}
	}
	createdAt := performance.CreatedAt.String()

	return &model.Performance{
		ID:        string(performance.NanoID),
		CreatedAt: &createdAt,
		Notes:     notes,
	}, nil
}
