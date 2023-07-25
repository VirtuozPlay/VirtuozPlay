package graph

import (
	"virtuozplay/graph/model"
	db "virtuozplay/models"

	"github.com/vektah/gqlparser/v2/gqlerror"
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
	if err = wrapError(err); performance == nil || err != nil {
		return nil, err
	}

	notes := make([]*model.PerformanceNote, len(performance.Notes))

	for i, note := range performance.Notes {
		notes[i] = &model.PerformanceNote{
			At:       int(note.At),
			Duration: int(note.Duration),
			Value:    note.Step,
		}
	}
	createdAt := performance.CreatedAt.String()
	song, err := ToGraphQLSong(performance.Song, nil)
	if err != nil {
		return nil, err
	}

	userName := "DummyUser"
	return &model.Performance{
		ID:        string(performance.NanoID),
		CreatedAt: &createdAt,
		Notes:     notes,
		// FIXME VERY TEMPORARY, please replace by proper DB relation
		Duration: 42069,
		// TODO VERY TEMPORARY, please replace by proper DB relation
		Author: &model.User{
			ID:   string(performance.NanoID),
			Name: &userName,
		},
		Song: song,
	}, nil
}

// ToGraphQLSong converts an instance of model.Song to its GraphQL representation.
func ToGraphQLSong(song *db.Song, err error) (*model.Song, error) {
	if err = wrapError(err); song == nil || err != nil {
		return nil, err
	}

	notes := make([]*model.SongNote, len(song.Notes))

	for i, note := range song.Notes {
		n := note
		notes[i] = &n
	}

	return &model.Song{
		ID:     string(song.NanoID),
		Title:  song.Title,
		Imgurl: song.Imgurl,
		Music:  song.Music,
		Notes:  notes,
	}, nil
}
