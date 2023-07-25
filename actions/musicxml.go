package actions

import (
	"encoding/xml"
	"fmt"
	"os"
)

type MusicXML struct {
	Parts []Part `xml:"part"`
}

// Part represents a part in a piece of music
type Part struct {
	Measures []Measure `xml:"measure"`
}

// Measure represents a measure in a piece of music
type Measure struct {
	Number int    `xml:"number,attr"`
	Beat   int    `xml:"beat"`
	Notes  []Note `xml:"note"`
	Key    []Key  `xml:"key"`
	Time   []Time `xml:"time"`
}

// Key represents a key signature change
type Key struct {
	Fifths int    `xml:"fifths"`
	Mode   string `xml:"mode"`
}

// Time represents a time signature change
type Time struct {
	Beats    int `xml:"beats"`
	BeatType int `xml:"beat-type"`
}

// Note represents a note in a measure
type Note struct {
	Pitch     Pitch     `xml:"pitch"`
	Duration  int       `xml:"duration"`
	Voice     int       `xml:"voice"`
	Type      string    `xml:"type"`
	Rest      xml.Name  `xml:"rest"`
	Chord     xml.Name  `xml:"chord"`
	Notations Notations `xml:"notations"`
	Default   float64   `xml:"default-x,attr"`
}

// Pitch represents the pitch of a note
type Pitch struct {
	Step   string `xml:"step"`
	Octave int    `xml:"octave"`
	Alter  int    `xml:"alter"`
}

type Notations struct {
	Technical Technical `xml:"technical"`
}

type Technical struct {
	String int    `xml:"string"`
	Fret   string `xml:"fret"`
}

type JsonNote struct {
	Octave    int    `json:"octave"`
	Step      string `json:"step"`
	Timestamp int64  `json:"timestamp"`
	Duration  int64  `json:"duration"`
}

func ParseMusicXML(xmlData []byte) (*MusicXML, error) {
	musicXML := &MusicXML{}

	err := xml.Unmarshal(xmlData, musicXML)
	if err != nil {
		// Handle the error
		return nil, err
	}

	return musicXML, nil
}

// comparison with tuners

func CompareNotes() error {
	userNotes := []struct {
		Name     string
		Duration int
	}{
		{Name: "", Duration: 2},
		{Name: "G", Duration: 1},
		{Name: "E", Duration: 1},
		{Name: "A", Duration: 1},
		{Name: "H", Duration: 2},
		{Name: "A", Duration: 1},
		{Name: "F", Duration: 1},
		{Name: "G", Duration: 1},
		{Name: "A", Duration: 1},
	}

	xmlData, err := os.ReadFile("../assets/musicXml/cleanCancan.xml") // To update
	if err != nil {
		// Managing errors
		fmt.Printf("%s\n", err.Error())
		return err
	}
	musicXML, err := ParseMusicXML(xmlData) // Call the ParseMusicXML function
	if err != nil {
		// Managing errors
		fmt.Printf("%s\n", err.Error())
		return err
	}

	serverNotes := []struct {
		Name     string
		Duration int
	}{}

	for _, part := range musicXML.Parts {
		for _, measure := range part.Measures {
			// Browse notes in each bar
			for _, note := range measure.Notes {
				// Access the duration of each note
				duration := note.Duration
				// Store information or perform desired processing
				// serverNotes := append("name", note.Pitch.Step, "duration", duration)
				serverNotes = append(serverNotes, struct {
					Name     string
					Duration int
				}{Name: note.Pitch.Step, Duration: duration})

			}
		}
	}

	fmt.Println(serverNotes)

	matchedNotes := 0
	totalNotes := len(userNotes)

	// Browse user notes and extracted notes for comparison

	for i := 0; i < len(serverNotes) && i < len(userNotes); i++ {
		if serverNotes[i].Name == userNotes[i].Name &&
			(serverNotes[i].Duration >= int(0.95*float64(userNotes[i].Duration))) &&
			(serverNotes[i].Duration <= int(1.05*float64(userNotes[i].Duration))) {
			matchedNotes++
		}
	}

	// Calculate precision
	// precision := float64(matchedNotes) / float64(totalNotes) * 100

	fmt.Printf("Corresponding notes: %d / %d\n", matchedNotes, totalNotes)
	// fmt.Printf("Precision: %.2f%%\n", precision)
	return nil
}
