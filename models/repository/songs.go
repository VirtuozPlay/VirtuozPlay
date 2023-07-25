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
		Title:  "Can Can",
		Imgurl: "../../images/cancan.jpg",
		Music:  "/assets/music/cancan/cancan.mp3",
		Notes: []graph.SongNote{
			{Measure: 1, Note: "A", Fret: 1, String: 1, Start: 0, End: 1000},
			{Measure: 2, Note: "B", Fret: 2, String: 2, Start: 1000, End: 2000},
			{Measure: 3, Note: "C", Fret: 3, String: 3, Start: 2000, End: 3000},
			{Measure: 4, Note: "D", Fret: 4, String: 1, Start: 3000, End: 4000},
			{Measure: 5, Note: "E", Fret: 5, String: 2, Start: 4000, End: 5000},
			{Measure: 6, Note: "F", Fret: 4, String: 3, Start: 5000, End: 6000},
			{Measure: 7, Note: "G#", Fret: 3, String: 5, Start: 6000, End: 7000},
			{Measure: 7, Note: "A", Fret: 2, String: 6, Start: 7000, End: 8000},
		},
	},
	"QyPqpmFqWC4uGInmkodgP": {
		ID:     1002,
		NanoID: "QyPqpmFqWC4uGInmkodgP",
		Title:  "Corinna Corinna",
		Imgurl: "../../images/corinna.jpg",
		Music:  "/assets/music/corinna/corinna.mp3",
		Notes: []graph.SongNote{
			{Measure: 1, Note: "A", Fret: 10, String: 1, Octave: 4, Start: 0, End: 1000},
			{Measure: 2, Note: "A", Fret: 11, String: 2, Octave: 4, Start: 1000, End: 2000},
			{Measure: 3, Note: "C", Fret: 10, String: 1, Octave: 4, Start: 2000, End: 3000},
			{Measure: 4, Note: "A", Fret: 11, String: 2, Octave: 4, Start: 3000, End: 4000},
			{Measure: 5, Note: "C", Fret: 9, String: 3, Octave: 5, Start: 4000, End: 5000},
			{Measure: 6, Note: "A", Fret: 8, String: 4, Octave: 5, Start: 5000, End: 6000},
			{Measure: 7, Note: "C", Fret: 1, String: 3, Octave: 5, Start: 6000, End: 7000},
			{Measure: 7, Note: "C", Fret: 2, String: 4, Octave: 5, Start: 7000, End: 8000},
		},
	},
	"SciuYTreZAmlKJhg4Tk0x": {
		ID:     1003,
		NanoID: "SciuYTreZAmlKJhg4Tk0x",
		Title:  "Sweet Home Alabama",
		Imgurl: "../../images/alabama.jpg",
		Music:  "/assets/music/alabama/alabama.mp3",
		Notes: []graph.SongNote{
			{Measure: 1, Note: "G", Fret: 10, String: 3, Octave: 4, Start: 0, End: 1000},
			{Measure: 2, Note: "A", Fret: 11, String: 4, Octave: 4, Start: 1000, End: 2000},
			{Measure: 3, Note: "F", Fret: 10, String: 5, Octave: 4, Start: 2000, End: 3000},
			{Measure: 4, Note: "B", Fret: 11, String: 3, Octave: 4, Start: 3000, End: 4000},
			{Measure: 5, Note: "C", Fret: 11, String: 4, Octave: 5, Start: 4000, End: 5000},
			{Measure: 6, Note: "D", Fret: 11, String: 5, Octave: 5, Start: 5000, End: 6000},
			{Measure: 7, Note: "E", Fret: 1, String: 4, Octave: 5, Start: 6000, End: 7000},
			{Measure: 7, Note: "A", Fret: 2, String: 3, Octave: 5, Start: 7000, End: 8000},
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
