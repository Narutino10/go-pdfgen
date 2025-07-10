package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Nautino10/go-pdfgen/internal/pdf"
)

// LirePDF lit un fichier PDF et retourne son contenu textuel
func LirePDF(cheminFichier string) (string, error) {
	if _, err := os.Stat(cheminFichier); os.IsNotExist(err) {
		return "", fmt.Errorf("le fichier '%s' n'existe pas", cheminFichier)
	}

	if filepath.Ext(cheminFichier) != ".pdf" {
		return "", fmt.Errorf("le fichier doit avoir l'extension .pdf")
	}

	reader := pdf.NewPDFReader()
	
	err := reader.LoadPDF(cheminFichier)
	if err != nil {
		return "", fmt.Errorf("erreur lors du chargement du PDF: %v", err)
	}

	texte, err := reader.ExtractText()
	if err != nil {
		return "", fmt.Errorf("erreur lors de l'extraction du texte: %v", err)
	}

	return texte, nil
}

// AnalyserPDF analyse un fichier PDF et retourne des informations détaillées
func AnalyserPDF(cheminFichier string) (map[string]interface{}, error) {
	if _, err := os.Stat(cheminFichier); os.IsNotExist(err) {
		return nil, fmt.Errorf("le fichier '%s' n'existe pas", cheminFichier)
	}

	reader := pdf.NewPDFReader()
	
	err := reader.LoadPDF(cheminFichier)
	if err != nil {
		return nil, fmt.Errorf("erreur lors du chargement du PDF: %v", err)
	}

	info := reader.GetDocumentInfo()
	
	fileInfo, err := os.Stat(cheminFichier)
	if err == nil {
		info["taille_fichier"] = fileInfo.Size()
		info["nom_fichier"] = filepath.Base(cheminFichier)
	}

	return info, nil
}

// SaisieLecturePDF demande à l'utilisateur le chemin du PDF à lire
func SaisieLecturePDF() string {
	var cheminFichier string
	
	fmt.Println("\n📖 Lecture d'un fichier PDF")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	
	for {
		fmt.Print("📁 Entrez le chemin du fichier PDF à lire : ")
		fmt.Scanln(&cheminFichier)
		
		if _, err := os.Stat(cheminFichier); os.IsNotExist(err) {
			fmt.Printf("❌ Le fichier '%s' n'existe pas. Veuillez réessayer.\n", cheminFichier)
			continue
		}
		
		if filepath.Ext(cheminFichier) != ".pdf" {
			fmt.Println("❌ Le fichier doit avoir l'extension .pdf. Veuillez réessayer.")
			continue
		}
		
		break
	}
	
	return cheminFichier
}
