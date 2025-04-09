package main

import (
	"github.com/Nautino10/go-pdfgen/internal/pdf"
)

func main() {
	// Créer un nouveau PDF
	writer := pdf.NewPDFWriter()

	// Ajouter une page
	writer.AddPage()

	// Ajouter du texte avec position et taille
	writer.AddText(100, 700, 18, "Hello depuis une vraie page PDF !")
	writer.AddText(100, 680, 12, "Test de texte dynamique généré sans bibliothèque.")
	writer.AddText(100, 660, 10, "🚀 Bravo ! Ton moteur PDF fonctionne.")

	// Sauvegarder le fichier
	err := writer.Save("page-test.pdf", "output-directory")
	if err != nil {
		panic(err)
	}
}
