package models

import (
	"time"
	graph "virtuozplay/graph/model"
)

type Song struct {
	ID           int64         `db:"id"`                                // The database ID of the song (do not expose to users!).
	NanoID       NanoID        `db:"nano_id"`                           // NanoID is the user-facing ID of the performance, generated using Go Nanoid.
	CreatedAt    time.Time     `db:"created_at"`                        //
	UpdatedAt    time.Time     `db:"updated_at"`                        //
	Title        string        `db:"title"`                             // Title is the title of the song.
	Performances []Performance `has_many:"performance" fk_id:"song_id"` //
	// FIXME:     vvvvvvvv Replace with proper notes encoding
	Notes []graph.SongNote `db:"-"`
	// FIXME:     ^^^^^^^^ Replace with proper notes encoding
}

func (p Song) TableName() string {
	return "song"
}
