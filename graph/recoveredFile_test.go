package graph

import (
	"fmt"
	"testing"
)


func Test_findFileByTitle(t *testing.T) {
	fileName, err := findFileByTitle("../assets/musicXml", "cleanCancan")

	if err != nil {
		t.Fatal(err)
	}

	// Utilisez fileName pour les assertions ou les vérifications supplémentaires que vous souhaitez effectuer
	fmt.Printf("Fichier trouvé : %s\n", fileName)
}