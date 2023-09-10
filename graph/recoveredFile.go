package graph

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
)

type MusicXML struct {
	Title string  `xml:"title"`
	Parts []Parts `xml:"part"`
}

// Part represents a part in a piece of music
type Parts struct {
	Measures []Measure `xml:"measure"`
}

// Measure represents a measure in a piece of music
type Measure struct {
	Number int    `xml:"number,attr"`
	Beat   int    `xml:"beat"`
	Notes  []Note `xml:"note"`
	// Key    []Key  `xml:"key"`
	// Time   []Time `xml:"time"`
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
	Abscissa  float64   `xml:"default-x,attr"`
}

type Pitch struct {
	Step   string `xml:"step"`
	Octave int    `xml:"octave"`
	Alter  int    `xml:"alter"`
}

type Notations struct {
	Technical Technical `xml:"technical"`
}

type Technical struct {
	String int `xml:"string"`
	Fret   int `xml:"fret"`
}

func ListFilesInDirectory(directoryPath string) ([]string, error) {
	var files []string

	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func ParseXMLFiles(files []string) ([]MusicXML, error) {
	var musicXMLList []MusicXML

	for i, file := range files {
		fmt.Println(i)
		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		if err := xml.NewDecoder(f).Decode(&musicXMLList); err != nil {
			return nil, err
		}

	}

	return musicXMLList, nil
}
