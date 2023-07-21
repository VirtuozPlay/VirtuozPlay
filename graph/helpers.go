package graph

import (
	"github.com/gobuffalo/validate/v3"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"virtuozplay/graph/model"
	db "virtuozplay/models"
)

// This file will not be regenerated automatically.

// WrapValidationErrors the errors of pop.ValidateAndCreate, pop.ValidateAndUpdate,
// or pop.ValidateAndSave to the error format expected by gqlgen.
func WrapValidationErrors(errors *validate.Errors, err error) error {
	if err != nil {
		return err
	}
	if errors == nil || !errors.HasAny() {
		return nil
	}
	keys := errors.Keys()
	gqlErrors := make(gqlerror.List, 0) // they are at least one error per key
	for _, key := range keys {
		for _, subError := range errors.Get(key) {
			gqlErrors = append(gqlErrors, &gqlerror.Error{
				Message: subError,
			})
		}
	}
	return gqlErrors
}

func ToGraphQLPerformance(performance *db.Performance) (*model.Performance, error) {
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
		ID:        performance.NanoID,
		CreatedAt: &createdAt,
		Notes:     notes,
	}, nil
}
