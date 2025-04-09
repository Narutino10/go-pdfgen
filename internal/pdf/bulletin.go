// internal/pdf/bulletin.go
package pdf

import (
	"fmt"

	"github.com/Nautino10/go-pdfgen/internal/models"
)

func DrawBulletin(b models.Bulletin, pdf *PDFWriter) {
	// Titre
	pdf.AddObject(`<< /Length 50 >>
stream
BT /F1 18 Tf 100 750 Td (Bulletin scolaire) Tj ET
endstream`)

	// Infos générales
	pdf.AddObject(fmt.Sprintf(`<< /Length 100 >>
stream
BT /F1 12 Tf 100 730 Td (Élève : %s) Tj
110 710 Td (Classe : %s) Tj
120 690 Td (Trimestre : %s - Année : %s) Tj ET
endstream`, b.Eleve, b.Classe, b.Trimestre, b.AnneeScolaire))

	y := 660
	for _, note := range b.Notes {
		line := fmt.Sprintf("%s : %.2f (coef %.1f)", note.Matiere, note.Note, note.Coef)
		pdf.AddObject(fmt.Sprintf(`<< /Length %d >>
stream
BT /F1 10 Tf 100 %d Td (%s) Tj ET
endstream`, len(line)+30, y, line))
		y -= 20
	}

	// Moyenne et appréciation
	pdf.AddObject(fmt.Sprintf(`<< /Length 100 >>
stream
BT /F1 12 Tf 100 %d Td (Moyenne : %.2f) Tj
110 %d Td (Appréciation : %s) Tj ET
endstream`, y, b.Noyenne, y-20, b.Appreciation))
}
