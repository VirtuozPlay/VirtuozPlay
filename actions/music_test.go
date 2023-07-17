package actions

import (
	"fmt"
	"os"
	"testing"
)

// HomeHandler is a default handler to serve up
// a home page.
func Test_ParseMusicXML(t *testing.T) {
	file, err := os.ReadFile("../front/assets/music/charles-fox/score.xml")
	if err != nil {
		t.Fatal(err)
	}

	result, err := ParseMusicXML(file)

	if err != nil {
		t.Fatal(err)
	}

	// Browse measurements
	for _, part := range result.Parts {
		for _, measure := range part.Measures {
			// Browse notes in each bar
			for _, note := range measure.Notes {
				// Access the duration of each note
				duration := note.Duration
				// Store information or perform desired processing
				fmt.Println("Mesure:", measure.Number, "Notes: {", "Pitch: {", "Accidental:", note.Pitch.Accidental, "Step:", note.Pitch.Step, "Octave:", note.Pitch.Octave, "}", "Duration:", duration, "Voice:", note.Voice, "Type:", note.Type, "Rest:", note.Rest, "Chord:", note.Chord, " }")
				// fmt.Print( note.Duration)
			}
		}
	}

	// fmt.Printf("%v", result.Parts[0].Measures[0].Notes[0].Duration)
	// fmt.Printf("%v", result.Parts[0].Measures)

}

func Test_CompareNotes(t *testing.T) {
	err := CompareNotes()

	if err != nil {
		t.Fatal(err)
	}
}
