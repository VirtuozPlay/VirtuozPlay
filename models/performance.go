package models

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
	"time"

	"github.com/go-interpreter/wagon/wasm/leb128"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	log "github.com/sirupsen/logrus"
)

type NotesEncoding int16

const (
	BinaryNotes NotesEncoding = iota
	Base64Notes NotesEncoding = iota
)

// Performance is the compact representation of a Performance.
// It is used to store and retrieve Performances from the database.
// Use the Encode and Decode methods to convert between Performance and Performance.
type Performance struct {
	ID            int64         `db:"id"`             // The database ID of the performance (do not expose to users!).
	NanoID        NanoID        `db:"nano_id"`        // NanoID is the user-facing ID of the performance, generated using Go Nanoid.
	CreatedAt     time.Time     `db:"created_at"`     //
	UpdatedAt     time.Time     `db:"updated_at"`     //
	SongID        int64         `db:"song_id"`        //
	Song          *Song         `belongs_to:"song"`   // Always linked to a Song
	Notes         []Note        `db:"-"`              // Notes is the list of notes played during the performance.
	NotesCount    int           `db:"notes_count"`    // NotesCount is the number of notes in the performance.
	NotesEncoding NotesEncoding `db:"notes_encoding"` // NotesEncoding is the encoding used for the EncodedNotes field.
	EncodedNotes  []byte        `db:"notes"`          // EncodedNotes is the encoded notes, see the documentation of Decode for more details.
}

type Note struct {
	At       int32  // At is the offset of th:wre note's start from the beginning of the performance, in milliseconds.
	Duration int32  // Duration is the duration of the note, in milliseconds.
	Step     string // Human-readable representation of the note (e.g. "C#", "D", "Fb", etc.)
	Octave   int32
}

func (p Performance) TableName() string {
	return "performance"
}

const performanceNotesValidatorName = "performance_notes"

func (p *Performance) Validate(*pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Name: "performance_id", Field: string(p.NanoID), Message: "performance ID is required"},
		&NotesAreValid{Name: performanceNotesValidatorName, Field: p.Notes},
	), nil
}

func (p *Performance) BeforeSave(*pop.Connection) error {
	p.encode()
	return nil
}

func (p *Performance) AfterCreate(*pop.Connection) error {
	return p.decode()
}

func (p *Performance) AfterUpdate(*pop.Connection) error {
	return p.decode()
}

func (p *Performance) AfterFind(*pop.Connection) error {
	return p.decode()
}

// AppendNote adds a new note to the performance.
// The function returns an error if the note is invalid or its timestamp occurs before the latest note in the performance.
func (p *Performance) AppendNote(index int, at int, duration int, step string, octave int) error {
	note := Note{
		At:       int32(at),
		Duration: int32(duration),
		Step:     step,
		Octave:   int32(octave),
	}

	errMsg := note.validate(index)
	if errMsg != "" {
		return errors.New(errMsg)
	}
	if len(p.Notes) > 0 {
		lastNote := p.Notes[len(p.Notes)-1]

		// if the note is identical to the last note, we can ignore it
		if lastNote.At == note.At && lastNote.Duration == note.Duration &&
			lastNote.Octave == note.Octave && lastNote.Step == note.Step {
			return nil
		}

		if lastNote.At > note.At {
			return fmt.Errorf("invalid note at index %v: occurs before the latest note in the performance", index)
		}
	}
	p.Notes = append(p.Notes, note)

	return nil
}

func (p *Performance) ResolvePreloads(preloads ...string) []string {
	for _, preload := range preloads {
		if strings.EqualFold(preload, "song") {
			return []string{"Song"}
		}
	}
	return nil
}

func (p *Performance) encode() {
	p.NotesCount = len(p.Notes)
	p.NotesEncoding = BinaryNotes
	sort.Slice(p.Notes, func(i, j int) bool {
		return compareNotesForSorting(p.Notes[i], p.Notes[j]) < 0
	})
	p.EncodedNotes = encodeNotes(p.Notes, p.NotesEncoding)
}

func compareNotesForSorting(a, b Note) int {
	// 1. compare start time
	if a.At != b.At {
		return int(a.At - b.At)
	}
	// 2. compare duration
	if a.Duration != b.Duration {
		return int(a.Duration - b.Duration)
	}
	// 3. compare octave
	if a.Octave != b.Octave {
		return int(a.Octave - b.Octave)
	}
	// 4. compare step
	return strings.Compare(a.Step, b.Step)
}

func (p *Performance) decode() error {
	notes, err := decodeNotes(p.EncodedNotes, p.NotesCount, p.NotesEncoding)
	if err != nil {
		return err
	}
	sort.Slice(notes, func(i, j int) bool {
		return compareNotesForSorting(notes[i], notes[j]) < 0
	})
	p.Notes = notes
	return nil
}

// decodeNotes converts a Performance into a Performance.
//
// EncodedNotes are encoded as a slice of bytes, where each note is encoded as follows:
// - The `at` field is encoded as a LEB128 unsigned integer.
// - The `duration` field is encoded as a LEB128 unsigned integer.
// - The `octave` field is encoded as a LEB128 unsigned integer.
// - The `value` is a single byte, the low 3 bits of which are the letter, and the high 2 bits are the modifier (see decodeNote()).
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
	var at, duration, octave uint32
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
	// 3. read the note's octave
	if octave, err = leb128.ReadVarUint32(src); err != nil {
		return Note{}, err
	}
	// 4. read the note's (encoded) step
	stepBuf := []byte{0}
	if _, err = io.ReadFull(src, stepBuf); err != nil {
		return Note{}, err
	}
	step := decodeNoteValue(stepBuf[0])

	return Note{
		At:       int32(at),
		Duration: int32(duration),
		Step:     step,
		Octave:   int32(octave),
	}, nil
}

// encodeNote encodes a single note into the given writer.
func encodeNote(dst io.Writer, note Note) {
	// Write the note's components
	// 1. write the note's start time
	_, _ = leb128.WriteVarUint32(dst, uint32(note.At))
	// 2. write the note's duration
	_, _ = leb128.WriteVarUint32(dst, uint32(note.Duration))
	// 3. write the note's octave
	_, _ = leb128.WriteVarUint32(dst, uint32(note.Octave))
	// 4. write the note's (encoded) step
	_, _ = dst.Write([]byte{encodeNoteValue(note.Step)})
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

// NotesAreValid is a validator that can be passed to the `validate.Validate` function.
// It checks that all the notes are valid (i.e. that they have a positive duration, start time, etc.)
type NotesAreValid struct {
	Name  string
	Field []Note
	Limit uint
}

// NoteValidationLimit is the maximum number of errors returned by the Normalize function.
const NoteValidationLimit uint = 20

// IsValid checks for invalid notes (notes with negative duration or start time, invalid values, etc.)
// The function returns a slice of errors, where each error corresponds to a single invalid note.
// If the number of errors exceeds Limit (by default: NoteValidationLimit),
// the Validation stops and the function appends an error indicating the limit has been reached.
func (v *NotesAreValid) IsValid(errors *validate.Errors) {
	if v.Limit == 0 {
		v.Limit = NoteValidationLimit
	}

	count := uint(0)
	errorKey := validators.GenerateKey(v.Name)
	for i, note := range v.Field {
		if msg := note.validate(i); msg != "" {
			errors.Add(errorKey, msg)
			count++
		}
		if count >= v.Limit {
			errors.Add(errorKey, fmt.Sprintf("too many errors, aborting"))
			break
		}
	}
}

func (n *Note) validate(i int) string {
	if n.At < 0 {
		return fmt.Sprintf("invalid note at index %v: negative start time (%v)", i, n.At)
	}
	if n.Duration < 0 {
		return fmt.Sprintf("invalid note at index %v: negative duration (%v)", i, n.Duration)
	}
	if n.Octave < 0 {
		return fmt.Sprintf("invalid note at index %v: negative octave (%v)", i, n.Octave)
	}
	if n.Octave > 9 {
		return fmt.Sprintf("invalid note at index %v: octave above 9 (%v)", i, n.Octave)
	}
	if len(n.Step) == 0 {
		return fmt.Sprintf("invalid note at index %v: empty step", i)
	}
	if len(n.Step) > 2 {
		return fmt.Sprintf("invalid note at index %v: invalid step (%v)", i, n.Step)
	}
	if n.Step[0] < 'A' || n.Step[0] > 'G' {
		return fmt.Sprintf("invalid note at index %v: invalid note letter (%v)", i, n.Step)
	}
	if len(n.Step) == 2 {
		if n.Step[1] != '#' && n.Step[1] != 'b' {
			return fmt.Sprintf("invalid note at index %v: invalid note modifier (%v)", i, n.Step)
		}
	}
	return ""
}

