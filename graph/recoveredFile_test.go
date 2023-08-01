package graph

import (
	"fmt"
	"os"
	"testing"
)

func Test_ListFilesInDirectory(t *testing.T) {
	files, err := ListFilesInDirectory("../assets/musicXml")
	if err != nil {
		t.Fatalf("Erreur lors de l'appel de la fonction listFilesInDirectory : %v", err)
	}
	if len(files) == 0 {
		t.Errorf("Aucun fichier n'a été trouvé dans le répertoire.")
	}

}

func Test_ListFilesInDirectory_InvalidPath(t *testing.T) {
	// Créer un répertoire temporaire pour les tests
	tempDir, err := os.MkdirTemp("", "testdir")
	// tempDir, err := ioutil.TempDir("", "testdir")
	if err != nil {
		t.Fatalf("Erreur lors de la création du répertoire temporaire : %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Appelez la fonction avec un répertoire invalide qui n'existe pas
	invalidPath := tempDir + "/non_existent_directory"
	files, err := ListFilesInDirectory(invalidPath)

	fmt.Println(files)

	if err == nil {
		t.Fatal("La fonction ListFilesInDirectory n'a pas renvoyé d'erreur pour un répertoire invalide.")
	}

	if _, ok := err.(*os.PathError); !ok {
		t.Fatalf("Erreur inattendue. Attendu : *os.PathError, Obtenu : %T", err)
	}
}

func Test_parseXMLFiles(t *testing.T) {
	files, err := ListFilesInDirectory("../assets/musicXml")
	if err != nil {
		t.Fatalf("Erreur lors de l'appel de la fonction listFilesInDirectory : %v", err)
	}

	if len(files) == 0 {
		t.Errorf("Aucun fichier n'a été trouvé dans le répertoire.")
	}

	// files := []string{tmpfile.Name()}
	musicXMLList, err := ParseXMLFiles(files)
	if err != nil {
		t.Fatalf("Erreur lors de l'appel de la fonction parseXMLFiles : %v", err)
	}

	if len(musicXMLList) == 0 {
		t.Fatalf("Aucun fichier XML n'a été correctement analysé.")
	}
	for _, musicXML := range musicXMLList {
		for _, part := range musicXML.Parts {
			for _, measure := range part.Measures {
				for _, note := range measure.Notes {
					// Access the duration of each note
					duration := note.Duration
					// Store information or perform desired processing
					fmt.Println("Mesure:", measure.Number, "Beat: ", measure.Beat, "Default-x: ", note.Abscissa, " Notes: {", "Pitch: {", "Step:", note.Pitch.Step, " Alter: ", note.Pitch.Alter, " Octave:", note.Pitch.Octave, "}", "Duration:", duration, "String:", note.Notations.Technical.String, "Fret:", note.Notations.Technical.Fret, " }")
				}
			}
		}

	}

}

func Test_ParseXMLFiles_InvalidXML(t *testing.T) {
	// Créer un fichier temporaire avec un contenu non-XML

	xmlContent := "This file does not exist"
	tmpfile, err := os.CreateTemp("", "test.xml")
	// tmpfile, err := ioutil.TempFile("", "test.xml")
	if err != nil {
		t.Fatalf("Erreur lors de la création du fichier temporaire : %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(xmlContent)); err != nil {
		t.Fatalf("Erreur lors de l'écriture du contenu dans le fichier temporaire : %v", err)
	}

	files := []string{tmpfile.Name()}

	_, err = ParseXMLFiles(files)
	if err == nil {
		t.Fatal("La fonction ParseXMLFiles n'a pas renvoyé d'erreur pour un XML non valide.")
	}

}
