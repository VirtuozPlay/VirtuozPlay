package models

import (
	"bytes"
	"encoding/base64"
	"github.com/go-interpreter/wagon/wasm/leb128"
	log "github.com/sirupsen/logrus"
	"io"
	"time"
)

type NotesEncoding int16

const (
	BinaryNotes NotesEncoding = iota
	Base64Notes NotesEncoding = iota
)

// RawPerformance is the compact representation of a Performance.
// It is used to store and retrieve Performances from the database.
// Use the Encode and Decode methods to convert between RawPerformance and Performance.
type RawPerformance struct {
	ID            int64         `db:"id"`             // The database ID of the performance (do not expose to users!).
	NanoID        string        `db:"nano_id"`        // NanoID is the user-facing ID of the performance, generated using Go Nanoid.
	CreatedAt     time.Time     `db:"created_at"`     //
	UpdatedAt     time.Time     `db:"updated_at"`     //
	NotesCount    int           `db:"notes_count"`    // NotesCount is the number of notes in the performance.
	NotesEncoding NotesEncoding `db:"notes_encoding"` // NotesEncoding is the encoding used for the Notes field.
	Notes         []byte        `db:"notes"`          // Notes is the encoded notes, see the documentation of Decode for more details.
}

// Performance is the full representation of a player's stats and notes played during a play session.
type Performance struct {
	ID        int64     `json:"-"`          // The database ID of the performance (do not expose to users!).
	NanoID    string    `json:"id"`         // NanoID is the user-facing ID of the performance, generated using Go Nanoid.
	CreatedAt time.Time `json:"created_at"` //
	UpdatedAt time.Time `json:"updated_at"` //
	Notes     []Note    `json:"notes"`      // Notes is the list of notes played during the performance.
}

type Note struct {
	At       int32  `json:"at"`       // At is the offset of the note's start from the beginning of the performance, in milliseconds.
	Duration int32  `json:"duration"` // Duration is the duration of the note, in milliseconds.
	Value    string `json:"value"`    // Human-readable representation of the note (e.g. "C#", "D", "Fb", etc.)
}

func (p RawPerformance) TableName() string {
	return "performance"
}

// Decode converts a RawPerformance into a Performance.
//
// Notes are encoded as a slice of bytes, where each note is encoded as follows:
// - The `at` field is encoded as a LEB128 unsigned integer.
// - The `duration` field is encoded as a LEB128 unsigned integer.
// - The `value` is a single byte, the low 3 bits of which are the letter, and the high 2 bits are the modifier (see decodeNote()).
func (p *RawPerformance) Decode() (*Performance, error) {
	notes, err := decodeNotes(p.Notes, p.NotesCount, p.NotesEncoding)
	if err != nil {
		return nil, err
	}
	return &Performance{
		ID:        p.ID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Notes:     notes,
	}, nil
}

// Encode converts a Performance into a RawPerformance.
func (p *Performance) Encode() RawPerformance {
	encoding := BinaryNotes
	return RawPerformance{
		ID:            p.ID,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
		NotesCount:    len(p.Notes),
		NotesEncoding: encoding,
		Notes:         encodeNotes(p.Notes, encoding),
	}
}

// decodeNotes decodes a slice of bytes into a slice of `count` notes.
// The encoding parameter specifies the representation of the notes in the bytes slice.
func decodeNotes(notes []byte, count int, encoding NotesEncoding) ([]Note, error) {
	n, err := encoding.decodeBytes(notes)

	if err != nil {
		return nil, err
	}

	src := bytes.NewReader(n)
	decoded := make([]Note, count)

	for i := 0; i < count; i++ {
		note, err := decodeNote(src)
		if err != nil {
			return nil, err
		}
		decoded[i] = note
	}

	if src.Len() > 0 {
		log.Warnf("%v notes were decoded, but %v bytes remain in the buffer", count, src.Len())
	}

	return decoded, nil
}

func encodeNotes(notes []Note, encoding NotesEncoding) []byte {
	dst := new(bytes.Buffer)

	for _, note := range notes {
		encodeNote(dst, note)
	}

	return encoding.encodeBytes(dst.Bytes())
}

// decodeNote decodes a single note from the given reader.
// The encoding is assumed to by BinaryNotes.
func decodeNote(src io.Reader) (Note, error) {
	var at, duration uint32
	var err error

	// Read the note's components
	// 1. read the note's start time
	if at, err = leb128.ReadVarUint32(src); err != nil {
		return Note{}, err
	}
	// 2. read the note's duration
	if duration, err = leb128.ReadVarUint32(src); err != nil {
		return Note{}, err
	}
	// 3. read the note's (encoded) value
	valueBuf := []byte{0}
	if _, err = io.ReadFull(src, valueBuf); err != nil {
		return Note{}, err
	}
	value := decodeNoteValue(valueBuf[0])

	return Note{
		At:       int32(at),
		Duration: int32(duration),
		Value:    value,
	}, nil
}

// encodeNote encodes a single note into the given writer.
func encodeNote(dst io.Writer, note Note) {
	// Write the note's components
	// 1. write the note's start time
	_, _ = leb128.WriteVarUint32(dst, uint32(note.At))
	// 2. write the note's duration
	_, _ = leb128.WriteVarUint32(dst, uint32(note.Duration))
	// 3. write the note's (encoded) value
	_, _ = dst.Write([]byte{encodeNoteValue(note.Value)})
}

// decodeNoteValue transform a byte-encode note into its string representation.
// The `rawValue` byte is composed of two parts:
// - the low 3 bits are the note's letter (A-G) encoded as an offset from 'A'
// - the high 2 bits are the note's modifier (0: natural, 1: sharp (#), 2: flat (b))
// (higher bits are reserved for future representations of the modifier)
func decodeNoteValue(rawValue byte) string {
	letter := 'A' + rune(rawValue&0b111)
	modifier := rawValue >> 3

	switch modifier {
	case byte(1):
		return string(letter) + "#"
	case byte(2):
		return string(letter) + "b"
	case byte(0):
		fallthrough
	default:
		return string(letter)
	}
}

// encodeNoteValue transforms a note's string representation into its byte-encoded representation.
func encodeNoteValue(note string) byte {
	result := note[0] - 'A'

	if len(note) > 1 {
		switch note[1] {
		case '#':
			result |= 1 << 3
		case 'b':
			result |= 2 << 3
		}
	}
	return result
}

// decodeBytes converts the given bytes according to the encoding.
// Currently, only Base64Notes and BinaryNotes are supported.
// If the encoding is BinaryNotes, the bytes are returned as-is.
//
// When Base64Notes is passed, the `from` slice is assumed to be a base64-encoded string.
func (encoding NotesEncoding) decodeBytes(from []byte) ([]byte, error) {
	switch encoding {
	case Base64Notes:
		dst := make([]byte, base64.StdEncoding.DecodedLen(len(from)))
		count, err := base64.StdEncoding.Decode(dst, from)

		if err != nil {
			return nil, err
		}
		return dst[:count], nil
	case BinaryNotes:
		fallthrough
	default:
		return from, nil
	}
}

// decodeBytes converts the given bytes according to the encoding.
// Currently, only Base64Notes and BinaryNotes are supported.
// If the encoding is BinaryNotes, the bytes are returned as-is.
//
// When Base64Notes is passed, the function returns a base64-encoded string.
func (encoding NotesEncoding) encodeBytes(from []byte) []byte {
	switch encoding {
	case Base64Notes:
		dst := make([]byte, base64.StdEncoding.EncodedLen(len(from)))
		base64.StdEncoding.Encode(dst, from)
		return dst
	case BinaryNotes:
		fallthrough
	default:
		return from
	}
}
