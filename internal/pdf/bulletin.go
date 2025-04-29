// internal/pdf/bulletin.go
package pdf

import (
	"fmt"

	"github.com/Nautino10/go-pdfgen/internal/models"
)

func DrawBulletin(b models.Bulletin, pdf *PDFWriter) {
	y := 750
	pdf.AddText(100, y, 18, "Bulletin scolaire")
	y -= 30
	pdf.AddText(100, y, 12, fmt.Sprintf("Élève : %s", b.Eleve))
	y -= 20
	pdf.AddText(100, y, 12, fmt.Sprintf("Classe : %s", b.Classe))
	y -= 20
	pdf.AddText(100, y, 12, fmt.Sprintf("Trimestre : %s - Année : %s", b.Trimestre, b.AnneeScolaire))
	y -= 30

	for _, note := range b.Notes {
		line := fmt.Sprintf("%s : %.2f (coef %.1f)", note.Matiere, note.Note, note.Coef)
		pdf.AddText(100, y, 10, line)
		y -= 15
	}

	pdf.AddText(100, y, 12, fmt.Sprintf("Moyenne générale : %.2f", b.Noyenne))
	y -= 20
	pdf.AddText(100, y, 12, fmt.Sprintf("Appréciation : %s", b.Appreciation))
}
