// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Note struct {
	// The offset of the note's start from the beginning of the performance, in milliseconds.
	At int `json:"at"`
	// The duration of the note, in milliseconds.
	Duration int `json:"duration"`
	// Human-readable representation of the note (e.g. 'C#', 'D', 'Fb', etc.)
	Value string `json:"value"`
}

type NoteInput struct {
	// The offset of the note's start from the beginning of the performance, in milliseconds.
	At int `json:"at"`
	// The duration of the note, in milliseconds.
	Duration int `json:"duration"`
	// Human-readable representation of the note (e.g. 'C#', 'D', 'Fb', etc.)
	Value string `json:"value"`
}

type Performance struct {
	ID        string  `json:"id"`
	Author    *string `json:"author,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	// The total duration of the performance, in milliseconds.
	Duration int `json:"duration"`
	// An array of *all* notes in the performance, sorted by their start time.
	Notes []*Note `json:"notes"`
}

type VirtuozPlay struct {
	Version string `json:"version"`
}
