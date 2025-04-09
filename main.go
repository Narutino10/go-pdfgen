// main.go
package main

import (
	"fmt"

	"github.com/Nautino10/go-pdfgen/internal/pdf"
)

func main() {
	writer := pdf.NewPDFWriter()

	// Objet 1 : /Catalog
	writer.AddObject(`<< /Type /Catalog /Pages 2 0 R >>`)

	// Objet 2 : /Pages
	writer.AddObject(`<< /Type /Pages /Kids [3 0 R] /Count 1 >>`)

	// Objet 3 : /Page
	writer.AddObject(`<< /Type /Page /Parent 2 0 R /MediaBox [0 0 595 842] /Contents 4 0 R /Resources << >> >>`)

	// Objet 4 : contenu texte minimal
	writer.AddObject(`<< /Length 44 >>
stream
BT /F1 24 Tf 100 700 Td (Hello, PDF World!) Tj ET
endstream`)

	// Sauvegarde du fichier PDF
	err := writer.Save("output.pdf")
	if err != nil {
		fmt.Println("Erreur lors de la sauvegarde :", err)
		return
	}

	fmt.Println("✅ PDF généré : output.pdf")
}
