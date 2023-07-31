package models

import "github.com/gobuffalo/pop/v6"

// Tests that Performance implements the required interfaces
func (ms *ModelSuite) Test_performanceInterfaceImplementations() {
	ms.Implements((*pop.TableNameAble)(nil), &Performance{})
	ms.Implements((*pop.BeforeSaveable)(nil), &Performance{})
	ms.Implements((*pop.AfterCreateable)(nil), &Performance{})
	ms.Implements((*pop.AfterUpdateable)(nil), &Performance{})
	ms.Implements((*pop.AfterFindable)(nil), &Performance{})
	ms.Implements((*Value)(nil), &Performance{})
	ms.Implements((*PreLoadable)(nil), &Performance{})
}

func (ms *ModelSuite) Test_decodeEncodeAllNotes() {
	ms.LoadFixture("performance_all_possible_notes")

	var perf Performance
	ms.NoError(ms.Model.DB.First(&perf))

	ms.Equal(int64(1), perf.ID)
	ms.Equal(NanoID("perf-1"), perf.NanoID)
	ms.Equal(Base64Notes, perf.NotesEncoding)
	ms.Equal(21, len(perf.Notes))

	expectedNotes := []Note{
		{At: 5, Duration: 2143, Step: "A", Octave: 6},
		{At: 10, Duration: 3186, Step: "A#", Octave: 5},
		{At: 20, Duration: 6669, Step: "Ab", Octave: 7},
		{At: 105, Duration: 5547, Step: "B", Octave: 7},
		{At: 110, Duration: 4548, Step: "B#", Octave: 4},
		{At: 120, Duration: 3213, Step: "Bb", Octave: 4},
		{At: 205, Duration: 5456, Step: "C", Octave: 5},
		{At: 210, Duration: 3512, Step: "C#", Octave: 3},
		{At: 220, Duration: 4163, Step: "Cb", Octave: 6},
		{At: 305, Duration: 9481, Step: "D", Octave: 0},
		{At: 310, Duration: 6504, Step: "D#", Octave: 0},
		{At: 320, Duration: 9735, Step: "Db", Octave: 0},
		{At: 405, Duration: 7964, Step: "E", Octave: 6},
		{At: 410, Duration: 7352, Step: "E#", Octave: 6},
		{At: 420, Duration: 3174, Step: "Eb", Octave: 8},
		{At: 505, Duration: 2229, Step: "F", Octave: 1},
		{At: 510, Duration: 5507, Step: "F#", Octave: 7},
		{At: 520, Duration: 8258, Step: "Fb", Octave: 4},
		{At: 605, Duration: 7375, Step: "G", Octave: 7},
		{At: 610, Duration: 5614, Step: "G#", Octave: 1},
		{At: 620, Duration: 7122, Step: "Gb", Octave: 4},
	}

	ms.Equal(expectedNotes, perf.Notes)
	validationErrs, err := perf.Validate(ms.Model.DB)
	ms.NoErrorf(err, "Performance 'test-1' should be valid")
	ms.Falsef(validationErrs.HasAny(), "Performance 'test-1' should be valid, got errors: %v", validationErrs.Error())

	base64EncodedNotes := perf.EncodedNotes
	perf.encode()

	ms.NotEqualf(perf.EncodedNotes, base64EncodedNotes, "EncodedNotes should have been re-encoded")
	ms.Equal(BinaryNotes, perf.NotesEncoding)
	ms.Equal(base64EncodedNotes, Base64Notes.encodeBytes(perf.EncodedNotes))
}

func (ms *ModelSuite) Test_badPerfEncoding() {
	ms.LoadFixture("performance_bad_encoding")
	var perf Performance
	ms.Error(ms.Model.DB.First(&perf))
}

func (ms *ModelSuite) Test_createPerformance() {
	// Load reference song
	ms.LoadFixture("song_empty")
	var song Song
	ms.NoError(ms.Model.DB.First(&song))

	perf := Performance{
		ID:     1200,
		NanoID: "test-12",
		Notes:  []Note{{At: 0, Duration: 128, Step: "A"}},
		Song:   &song,
		SongID: song.ID,
	}
	validationErrs, err := ms.DB.ValidateAndCreate(&perf)
	ms.NoError(err)
	ms.False(validationErrs.HasAny())

	perf.NanoID = "test-12-updated"

	validationErrs, err = ms.DB.ValidateAndUpdate(&perf)
	ms.NoError(err)
	ms.False(validationErrs.HasAny())
}

func (ms *ModelSuite) Test_decodeBinaryNotes() {
	noteA, err1 := decodeNotes([]byte{0x00, 0x80, 0x01, 0x01, 0x00}, 1, BinaryNotes)
	noteDb, err2 := decodeNotes([]byte{0xad, 0x01, 0xd3, 0xa8, 0x05, 0x02, 0x13}, 1, BinaryNotes)
	noteCs, err3 := decodeNotes([]byte{0xc8, 0x01, 0x0f, 0x04, 0x0a}, 1, BinaryNotes)

	ms.NoError(err1, "'A' note decoding failed")
	ms.NoError(err2, "'Db' note decoding failed")
	ms.NoError(err3, "'C#' note decoding failed")

	ms.Equal([]Note{{At: 0, Duration: 128, Step: "A", Octave: 1}}, noteA)
	ms.Equal([]Note{{At: 173, Duration: 87123, Step: "Db", Octave: 2}}, noteDb)
	ms.Equal([]Note{{At: 200, Duration: 15, Step: "C#", Octave: 4}}, noteCs)
}

// Tests the base64 decoding of notes
func (ms *ModelSuite) Test_decodeBase64Notes() {
	noteA, err1 := decodeNotes([]byte("AIABAAA="), 1, Base64Notes)
	noteDb, err2 := decodeNotes([]byte("rQHTqAUAEw=="), 1, Base64Notes)
	noteCs, err3 := decodeNotes([]byte("yAEPAAo="), 1, Base64Notes)
	_, notBase64 := decodeNotes([]byte("Is this even valid base64?"), 1, Base64Notes)

	ms.NoError(err1, "'A' note decoding failed")
	ms.NoError(err2, "'Db' note decoding failed")
	ms.NoError(err3, "'C#' note decoding failed")

	ms.Equal([]Note{{At: 0, Duration: 128, Step: "A"}}, noteA)
	ms.Equal([]Note{{At: 173, Duration: 87123, Step: "Db"}}, noteDb)
	ms.Equal([]Note{{At: 200, Duration: 15, Step: "C#"}}, noteCs)

	ms.Error(notBase64)
}

func (ms *ModelSuite) Test_encodeBase64Notes() {
	noteA := string(encodeNotes([]Note{{At: 0, Duration: 128, Step: "A"}}, Base64Notes))
	noteDb := string(encodeNotes([]Note{{At: 173, Duration: 87123, Step: "Db"}}, Base64Notes))
	noteCs := string(encodeNotes([]Note{{At: 200, Duration: 15, Step: "C#"}}, Base64Notes))

	ms.Equal("AIABAAA=", noteA)
	ms.Equal("rQHTqAUAEw==", noteDb)
	ms.Equal("yAEPAAo=", noteCs)
}

// Tests that the decoder returns an error when the notes are truncated
func (ms *ModelSuite) Test_decodeTruncatedNotes() {
	_, err1 := decodeNotes([]byte{}, 1, BinaryNotes)
	_, err2 := decodeNotes([]byte{0x00}, 1, BinaryNotes)
	_, err3 := decodeNotes([]byte{0x00, 0x80}, 1, BinaryNotes)
	_, err4 := decodeNotes([]byte{0x00, 0x80, 0x01}, 1, BinaryNotes)

	ms.Error(err1)
	ms.Error(err2)
	ms.Error(err3)
	ms.Error(err4)
}

func (ms *ModelSuite) Test_decodeNotesWrongCount() {
	_, tooMuchData := decodeNotes([]byte{0x00, 0x80, 0x01, 0x00, 0xad, 0x01, 0xd3, 0xa8, 0x05, 0x13}, 3, BinaryNotes)
	_, notEnoughData := decodeNotes([]byte{0x00, 0x80, 0x01, 0x00, 0xad, 0x01, 0xd3, 0xa8, 0x05, 0x13}, 1, BinaryNotes)

	ms.Error(tooMuchData)
	ms.NoError(notEnoughData) // This is a warning and not an error
}

func (ms *ModelSuite) Test_validatePerformance() {
	perf := Performance{
		ID:     42,
		NanoID: "test-42",
		Notes: []Note{
			{At: 0, Duration: 128, Step: "A", Octave: 0},      // valid
			{At: -42, Duration: 128, Step: "A", Octave: 0},    // invalid, negative `At`
			{At: 0, Duration: 128, Step: "A", Octave: -1},     // invalid, negative `Octave`
			{At: 0, Duration: 128, Step: "A", Octave: 10},     // invalid, `Octave` superior of 9
			{At: 0, Duration: 128, Step: "A", Octave: 0},      // valid
			{At: 0, Duration: -128, Step: "A", Octave: 0},     // invalid, negative `Duration`
			{At: 0, Duration: 128, Step: "A", Octave: 0},      // valid
			{At: 0, Duration: 128, Step: "", Octave: 0},       // invalid, empty `Step`
			{At: 0, Duration: 128, Step: "Abb", Octave: 0},    // invalid, step too long
			{At: 0, Duration: 128, Step: "Csharp", Octave: 0}, // invalid, step too long again
			{At: 0, Duration: 128, Step: "@", Octave: 0},      // invalid, letter before 'A'
			{At: 0, Duration: 128, Step: "H", Octave: 0},      // invalid, letter after 'G'
			{At: 0, Duration: 128, Step: "Ab", Octave: 0},     // valid
			{At: 0, Duration: 128, Step: "D#", Octave: 0},     // valid
			{At: 0, Duration: 128, Step: "D!", Octave: 0},     // invalid, bad modifier
		},
	}

	validationErrs, err := perf.Validate(ms.Model.DB)

	ms.NoError(err)
	ms.Equal(1, validationErrs.Count())
	ms.Len(validationErrs.Get(performanceNotesValidatorName), 10, "Performance should have exactly 10 Validation errors, got %v", validationErrs.Error())
}

func (ms *ModelSuite) Test_validatePerformanceTooManyErrors() {
	badNote := Note{At: -42, Duration: -128, Step: "..."}
	notes := make([]Note, 100)

	for i := range notes {
		notes[i] = badNote
	}
	perf := Performance{
		ID:     84,
		NanoID: "test-84",
		Notes:  notes,
	}

	validationErrs, err := perf.Validate(ms.Model.DB)

	ms.NoError(err)
	ms.Equal(validationErrs.Count(), 1)
	ms.Len(validationErrs.Get(performanceNotesValidatorName), int(NoteValidationLimit+1))
}

func (ms *ModelSuite) Test_normalizePerformance() {
	perf := Performance{
		ID:     21,
		NanoID: "test-21",
		Notes: []Note{
			{At: 999, Duration: 128, Step: "A"},
			{At: 10, Duration: 32, Step: "A"},
			{At: 10, Duration: 128, Step: "D"},
			{At: 10, Duration: 64, Step: "C"},
			{At: 10, Duration: 128, Step: "Db"},
			{At: 10, Duration: 128, Step: "B"},
		},
	}

	validationErrs, err := perf.Validate(ms.Model.DB)
	ms.NoError(err, "Performance 'test-21' should be valid before normalizing")
	ms.False(validationErrs.HasAny(), "Performance 'test-21' should be valid before normalizing")

	perf.encode()
	ms.NoError(perf.decode())

	validationErrs, err = perf.Validate(ms.Model.DB)
	ms.NoError(err, "Performance 'test-21' should be valid after normalizing")
	ms.False(validationErrs.HasAny(), "Performance 'test-21' should be valid after normalizing")

	ms.Equal([]Note{
		{At: 10, Duration: 32, Step: "A"},
		{At: 10, Duration: 64, Step: "C"},
		{At: 10, Duration: 128, Step: "B"},
		{At: 10, Duration: 128, Step: "D"},
		{At: 10, Duration: 128, Step: "Db"},
		{At: 999, Duration: 128, Step: "A"},
	}, perf.Notes)
}

func (ms *ModelSuite) Test_appendNote() {
	ms.LoadFixture("performance_append_notes")

	var perf Performance
	ms.NoError(ms.Model.DB.First(&perf))

	notesLen := len(perf.Notes)

	// Should not fail
	ms.NoError(perf.AppendNote(0, 10999, 67, "E", 0))
	ms.Equal(Note{At: 10999, Duration: 67, Step: "E"}, perf.Notes[notesLen])
	ms.Len(perf.Notes, notesLen+1)

	// Should not add same note twice
	ms.NoError(perf.AppendNote(0, 10999, 67, "E", 0))
	ms.Equal(Note{At: 10999, Duration: 67, Step: "E"}, perf.Notes[notesLen])
	ms.Len(perf.Notes, notesLen+1)

	// Same note, but with invalid start
	ms.Error(perf.AppendNote(0, 10, 67, "E", 0))

	// Same note, but with invalid value
	ms.Error(perf.AppendNote(0, 20999, 67, "Ez", 0))
}

func (ms *ModelSuite) Test_Performance_ResolvePreloads() {
	var perf Performance

	ms.Empty(perf.ResolvePreloads())
	ms.Contains(perf.ResolvePreloads("song"), "Song")
	ms.Contains(perf.ResolvePreloads("SONG"), "Song")
	ms.Len(perf.ResolvePreloads("Song", "sOnG", "other"), 1)
}
