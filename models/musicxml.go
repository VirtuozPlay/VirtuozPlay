package models

import (
	"encoding/xml"
	"fmt"
	"os"
	"virtuozplay/graph/model"
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
	Number  int       `xml:"number,attr"`
	Beat    int       `xml:"beat"`
	NoteXml []NoteXml `xml:"note"`
	Key     []Key     `xml:"key"`
	Time    []Time    `xml:"time"`
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

// NoteXml Note represents a note in a measure
type NoteXml struct {
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

// CompareNotes comparison with tuners
func CompareNotes(userNotes []*model.InputNote) error { // TODO: Rajouter la variable à la fonction song *Song,

	//_ = song                                                         // TODO: Remplacer par le path
	xmlData, err := os.ReadFile("./assets/musicXml/cleanCancan.xml") // To update

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

	var xmlFileNotes []struct {
		Name     string
		Duration int
		Octave   int
	}

	for _, part := range musicXML.Parts {
		for _, measure := range part.Measures {
			// Browse notes in each bar
			for _, note := range measure.NoteXml {
				// Access the duration of each note
				duration := note.Duration
				// Store information or perform desired processing
				// serverNotes := append("name", note.Pitch.Step, "duration", duration)
				xmlFileNotes = append(xmlFileNotes, struct {
					Name     string
					Duration int
					Octave   int
				}{Name: note.Pitch.Step, Duration: duration, Octave: note.Pitch.Octave})

			}
		}
	}

	fmt.Println(xmlFileNotes)

	matchedNotes := 0
	totalNotes := len(userNotes)

	var newUserNotes []struct {
		Name     string
		Duration int
		Octave   int
	}
	fmt.Println(userNotes)
	for _, note := range userNotes {
		newUserNotes = append(newUserNotes, struct {
			Name     string
			Duration int
			Octave   int
		}{
			Name:     note.Value,
			Duration: note.Duration,
			Octave:   note.Octave, // Set the appropriate value for the octave here
		})
	}

	fmt.Println(newUserNotes)

	// Create a boolean array to store comparison results
	results := make([]bool, totalNotes)

	// Browse user notes and extracted notes for comparison

	for i := 0; i < len(userNotes); i++ {
		userNote := userNotes[i]
		for j := 0; j < len(xmlFileNotes); j++ {
			xmlNote := xmlFileNotes[j]

			// Compare userNote with xmlNote
			if xmlNote.Name == userNote.Value &&
				xmlNote.Octave == userNotes[i].Octave &&
				(xmlNote.Duration >= int(0.95*float64(userNote.Duration))) &&
				(xmlNote.Duration <= int(1.05*float64(userNote.Duration))) {
				matchedNotes++
				results[i] = true
				break
			} else {
				results[i] = false
			}
		}
	}

	// for i := 0; i < len(xmlFileNotes) && i < len(userNotes); i++ {
	// 	if xmlFileNotes[i].Name == userNotes[i].Value &&
	// 		xmlFileNotes[i].Octave == userNotes[i].Octave &&
	// 		(xmlFileNotes[i].Duration >= int(0.95*float64(userNotes[i].Duration))) &&
	// 		(xmlFileNotes[i].Duration <= int(1.05*float64(userNotes[i].Duration))) {
	// 		matchedNotes++
	// 		results[i] = true
	// 	} else {
	// 		results[i] = false
	// 	}
	// }
	fmt.Println(results)

	//fmt.Printf("Corresponding notes: %d / %d\n", matchedNotes, totalNotes)
	// fmt.Printf("Precision: %.2f%%\n", precision)
	return nil
}
