package models

func (ms *ModelSuite) Test_decodeEncodeAllNotes() {
	ms.LoadFixture("performance_all_possible_notes")

	var rawPerf RawPerformance
	ms.NoError(ms.Model.DB.First(&rawPerf))

	perf, err := rawPerf.Decode()
	ms.NoError(err)

	ms.Assert().Equal(int64(1), perf.ID)
	ms.Assert().Equal(rawPerf.CreatedAt, perf.CreatedAt)
	ms.Assert().Equal(rawPerf.UpdatedAt, perf.UpdatedAt)
	ms.Assert().Equal(Base64Notes, rawPerf.NotesEncoding)
	ms.Assert().Equal(21, len(perf.Notes))
	ms.Assert().Equal([]Note{
		{At: 0, Duration: 3055, Value: "A"},
		{At: 0, Duration: 3447, Value: "A#"},
		{At: 0, Duration: 4768, Value: "Ab"},
		{At: 0, Duration: 863, Value: "B"},
		{At: 0, Duration: 2157, Value: "B#"},
		{At: 0, Duration: 936, Value: "Bb"},
		{At: 0, Duration: 4674, Value: "C"},
		{At: 0, Duration: 2880, Value: "C#"},
		{At: 0, Duration: 8213, Value: "Cb"},
		{At: 0, Duration: 6929, Value: "D"},
		{At: 0, Duration: 2925, Value: "D#"},
		{At: 0, Duration: 1171, Value: "Db"},
		{At: 0, Duration: 4627, Value: "E"},
		{At: 0, Duration: 3154, Value: "E#"},
		{At: 0, Duration: 8494, Value: "Eb"},
		{At: 0, Duration: 4596, Value: "F"},
		{At: 0, Duration: 7759, Value: "F#"},
		{At: 0, Duration: 5823, Value: "Fb"},
		{At: 0, Duration: 2918, Value: "G"},
		{At: 0, Duration: 4153, Value: "G#"},
		{At: 0, Duration: 8239, Value: "Gb"},
	}, perf.Notes)

	reEncoded := perf.Encode()

	ms.Assert().NotEqualf(&rawPerf.Notes, &reEncoded.Notes, "Notes should have been re-encoded")
	ms.Assert().Equal(rawPerf.ID, reEncoded.ID)
	ms.Assert().Equal(rawPerf.CreatedAt, reEncoded.CreatedAt)
	ms.Assert().Equal(rawPerf.UpdatedAt, reEncoded.UpdatedAt)
	ms.Assert().Equal(BinaryNotes, reEncoded.NotesEncoding)
	ms.Assert().Equal(rawPerf.Notes, Base64Notes.encodeBytes(reEncoded.Notes))
}

func (ms *ModelSuite) Test_badPerfEncoding() {
	ms.LoadFixture("performance_bad_encoding")

	var rawPerf RawPerformance
	ms.NoError(ms.Model.DB.First(&rawPerf))

	_, err := rawPerf.Decode()
	ms.Error(err)
}

func (ms *ModelSuite) Test_decodeBinaryNotes() {
	noteA, err1 := decodeNotes([]byte{0x00, 0x80, 0x01, 0x00}, 1, BinaryNotes)
	noteDb, err2 := decodeNotes([]byte{0xad, 0x01, 0xd3, 0xa8, 0x05, 0x13}, 1, BinaryNotes)
	noteCs, err3 := decodeNotes([]byte{0xc8, 0x01, 0x0f, 0x0a}, 1, BinaryNotes)

	ms.NoError(err1, "'A' note decoding failed")
	ms.NoError(err2, "'Db' note decoding failed")
	ms.NoError(err3, "'C#' note decoding failed")

	ms.Assert().Equal([]Note{{At: 0, Duration: 128, Value: "A"}}, noteA)
	ms.Assert().Equal([]Note{{At: 173, Duration: 87123, Value: "Db"}}, noteDb)
	ms.Assert().Equal([]Note{{At: 200, Duration: 15, Value: "C#"}}, noteCs)
}

// Tests the base64 decoding of notes
func (ms *ModelSuite) Test_decodeBase64Notes() {
	noteA, err1 := decodeNotes([]byte("AIABAA=="), 1, Base64Notes)
	noteDb, err2 := decodeNotes([]byte("rQHTqAUT"), 1, Base64Notes)
	noteCs, err3 := decodeNotes([]byte("yAEPCg=="), 1, Base64Notes)
	_, notBase64 := decodeNotes([]byte("Is this even valid base64?"), 1, Base64Notes)

	ms.NoError(err1, "'A' note decoding failed")
	ms.NoError(err2, "'Db' note decoding failed")
	ms.NoError(err3, "'C#' note decoding failed")

	ms.Assert().Equal([]Note{{At: 0, Duration: 128, Value: "A"}}, noteA)
	ms.Assert().Equal([]Note{{At: 173, Duration: 87123, Value: "Db"}}, noteDb)
	ms.Assert().Equal([]Note{{At: 200, Duration: 15, Value: "C#"}}, noteCs)

	ms.Error(notBase64)
}

func (ms *ModelSuite) Test_encodeBase64Notes() {
	noteA := string(encodeNotes([]Note{{At: 0, Duration: 128, Value: "A"}}, Base64Notes))
	noteDb := string(encodeNotes([]Note{{At: 173, Duration: 87123, Value: "Db"}}, Base64Notes))
	noteCs := string(encodeNotes([]Note{{At: 200, Duration: 15, Value: "C#"}}, Base64Notes))

	ms.Assert().Equal("AIABAA==", noteA)
	ms.Assert().Equal("rQHTqAUT", noteDb)
	ms.Assert().Equal("yAEPCg==", noteCs)
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
