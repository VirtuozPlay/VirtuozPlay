package repository

import (
	"github.com/gobuffalo/pop/v6"
	graph "virtuozplay/graph/model"
	"virtuozplay/models"
)

// Songs is the repository for the models.Song instances.
// Provides methods for find and CRUD operations of performances.
type Songs struct {
	DatabaseRepository[models.Song]
}

// NewSongsRepository creates a new repository for songs
func NewSongsRepository(db *pop.Connection) Songs {
	return Songs{NewDatabaseRepository[models.Song](db)}
}

// FIXME: Remove this when we have a proper loading for songs
var HardcodedSongs = map[models.NanoID]*models.Song{
	"h8rHA-Q0dD5dBbY1L2Fzf": {
		ID:     1001,
		NanoID: "h8rHA-Q0dD5dBbY1L2Fzf",
		Title:  "Cancan",
		Notes: []graph.SongNote{
			{Measure: 1, Note: "C", Fret: 10, String: 2, Start: 0, End: 1000},
			{Measure: 2, Note: "C", Fret: 11, String: 2, Start: 1000, End: 2000},
			{Measure: 3, Note: "C", Fret: 10, String: 2, Start: 2000, End: 3000},
			{Measure: 4, Note: "C", Fret: 11, String: 2, Start: 3000, End: 4000},
			{Measure: 5, Note: "C", Fret: 17, String: 2, Start: 4000, End: 5000},
			{Measure: 6, Note: "C", Fret: 18, String: 2, Start: 5000, End: 6000},
			{Measure: 7, Note: "C", Fret: 17, String: 2, Start: 6000, End: 7000},
			{Measure: 7, Note: "C", Fret: 18, String: 2, Start: 7000, End: 8000},
		},
	},
	"QyPqpmFqWC4uGInmkodgP": {
		ID:     1002,
		NanoID: "QyPqpmFqWC4uGInmkodgP",
		Title:  "Corinna",
		Notes: []graph.SongNote{
			{Measure: 1, Note: "A", Fret: 10, String: 2, Octave: 4, Start: 0, End: 1000},
			{Measure: 2, Note: "A", Fret: 11, String: 2, Octave: 4, Start: 1000, End: 2000},
			{Measure: 3, Note: "C", Fret: 10, String: 2, Octave: 4, Start: 2000, End: 3000},
			{Measure: 4, Note: "A", Fret: 11, String: 2, Octave: 4, Start: 3000, End: 4000},
			{Measure: 5, Note: "C", Fret: 11, String: 2, Octave: 5, Start: 4000, End: 5000},
			{Measure: 6, Note: "A", Fret: 11, String: 2, Octave: 5, Start: 5000, End: 6000},
			{Measure: 7, Note: "C", Fret: 1, String: 2, Octave: 5, Start: 6000, End: 7000},
			{Measure: 7, Note: "C", Fret: 2, String: 2, Octave: 5, Start: 7000, End: 8000},
		},
	},
}

func (r *Songs) FindByNanoID(nanoID models.NanoID, preloadFields ...string) (*models.Song, error) {
	// FIXME: Remove this when we have a proper loading for songs
	if song, ok := HardcodedSongs[nanoID]; ok {
		return song, nil
	}
	return r.DatabaseRepository.FindByNanoID(nanoID, preloadFields...)
}
