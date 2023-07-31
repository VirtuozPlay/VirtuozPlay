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
		ID:        1001,
		NanoID:    "h8rHA-Q0dD5dBbY1L2Fzf",
		Title:     "Can Can",
		ImgURL:    "../../images/cancan.jpg",
		URL:       "cancan",
		MusicPath: "/assets/music/cancan/cancan.mp3",
		Notes: []graph.SongNote{
			{Measure: 1, Note: "A", Fret: 1, String: 1, Beat: 4, Alter: 0, Octave: 3, Duration: 2, Type: "croche", Default: 10.0},
			{Measure: 2, Note: "B", Fret: 2, String: 2, Beat: 4, Alter: 0, Octave: 3, Duration: 2, Type: "croche", Default: 10.0},
			{Measure: 3, Note: "C", Fret: 3, String: 3, Beat: 4, Alter: 0, Octave: 3, Duration: 2, Type: "croche", Default: 10.0},
			{Measure: 4, Note: "D", Fret: 4, String: 1, Beat: 4, Alter: 0, Octave: 3, Duration: 2, Type: "croche", Default: 10.0},
			{Measure: 5, Note: "E", Fret: 5, String: 2, Beat: 4, Alter: 0, Octave: 3, Duration: 2, Type: "croche", Default: 10.0},
			{Measure: 6, Note: "F", Fret: 4, String: 3, Beat: 4, Alter: 0, Octave: 3, Duration: 2, Type: "croche", Default: 10.0},
			{Measure: 7, Note: "G#", Fret: 3, String: 5, Beat: 4, Alter: 0, Octave: 3, Duration: 2, Type: "croche", Default: 10.0},
			{Measure: 7, Note: "A", Fret: 2, String: 6, Beat: 4, Alter: 0, Octave: 3, Duration: 2, Type: "croche", Default: 10.0},
		},
	},
	"QyPqpmFqWC4uGInmkodgP": {
		ID:        1002,
		NanoID:    "QyPqpmFqWC4uGInmkodgP",
		Title:     "Corinna Corinna",
		ImgURL:    "../../images/corinna.jpg",
		URL:       "corinna",
		MusicPath: "/assets/music/corinna/corinna.mp3",
		Notes: []graph.SongNote{
			{Measure: 1, Note: "A", Fret: 10, String: 1, Octave: 4, Beat: 4, Alter: 0, Duration: 2, Type: "croche", Default: 10.0},
			{Measure: 2, Note: "A", Fret: 11, String: 2, Octave: 4, Beat: 4, Alter: 0, Duration: 2, Type: "croche", Default: 10.0},
			{Measure: 3, Note: "C", Fret: 10, String: 1, Octave: 4, Beat: 4, Alter: 0, Duration: 2, Type: "croche", Default: 10.0},
			{Measure: 4, Note: "A", Fret: 11, String: 2, Octave: 4, Beat: 4, Alter: 0, Duration: 2, Type: "croche", Default: 10.0},
			{Measure: 5, Note: "C", Fret: 9, String: 3, Octave: 5, Beat: 4, Alter: 0, Duration: 2, Type: "croche", Default: 10.0},
			{Measure: 6, Note: "A", Fret: 8, String: 4, Octave: 5, Beat: 4, Alter: 0, Duration: 2, Type: "croche", Default: 10.0},
			{Measure: 7, Note: "C", Fret: 1, String: 3, Octave: 5, Beat: 4, Alter: 0, Duration: 2, Type: "croche", Default: 10.0},
			{Measure: 7, Note: "C", Fret: 2, String: 4, Octave: 5, Beat: 4, Alter: 0, Duration: 2, Type: "croche", Default: 10.0},
		},
	},
	"SciuYTreZAmlKJhg4Tk0x": {
		ID:        1003,
		NanoID:    "SciuYTreZAmlKJhg4Tk0x",
		Title:     "Sweet Home Alabama",
		ImgURL:    "../../images/alabama.jpg",
		URL:       "alabama",
		MusicPath: "/assets/music/alabama/alabama.mp3",
		Notes: []graph.SongNote{
			{Measure: 1, Note: "G", Fret: 10, String: 3, Octave: 4, Beat: 4, Alter: 0, Duration: 2, Type: "croche", Default: 10.0},
			{Measure: 2, Note: "A", Fret: 11, String: 4, Octave: 4, Beat: 4, Alter: 0, Duration: 2, Type: "croche", Default: 10.0},
			{Measure: 3, Note: "F", Fret: 10, String: 5, Octave: 4, Beat: 4, Alter: 0, Duration: 2, Type: "croche", Default: 10.0},
			{Measure: 4, Note: "B", Fret: 11, String: 3, Octave: 4, Beat: 4, Alter: 0, Duration: 2, Type: "croche", Default: 10.0},
			{Measure: 5, Note: "C", Fret: 11, String: 4, Octave: 5, Beat: 4, Alter: 0, Duration: 2, Type: "croche", Default: 10.0},
			{Measure: 6, Note: "D", Fret: 11, String: 5, Octave: 5, Beat: 4, Alter: 0, Duration: 2, Type: "croche", Default: 10.0},
			{Measure: 7, Note: "E", Fret: 1, String: 4, Octave: 5, Beat: 4, Alter: 0, Duration: 2, Type: "croche", Default: 10.0},
			{Measure: 7, Note: "A", Fret: 2, String: 3, Octave: 5, Beat: 4, Alter: 0, Duration: 2, Type: "croche", Default: 10.0},
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
