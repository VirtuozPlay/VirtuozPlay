package graph

import (
	"fmt"
	"os"
	"strings"
)


func findFileByTitle(directoryPath, fileTitle string) (string, error) {
	// Lire le contenu du répertoire
	files, err := os.ReadDir( "../assets/musicXml")
	if err != nil {
		return "", err
	}
	
	var foundFiles []string // Tableau pour stocker les noms de fichiers trouvés
	
	// Browse directory files
	for _, file := range files {
		foundFiles = append(foundFiles, file.Name())
		fmt.Printf("Liste des fichiers : %s\n", foundFiles)
		// Vérifier si le fichier est un fichier régulier
		if file.Type().IsRegular() {
			// Vérifier si le nom du fichier contient le titre de la chanson (en supposant que le titre est dans le nom du fichier)
			if strings.Contains(strings.ToLower(file.Name()), strings.ToLower(fileTitle)) {
				return file.Name(), nil
			}
		}
	}

	// Si aucun fichier correspondant n'est trouvé, renvoyer une erreur
	return "", fmt.Errorf("Aucun fichier trouvé pour la chanson : %s", fileTitle)
}