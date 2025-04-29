// internal/pdf/report.go
package pdf

import (
	"fmt"

	"github.com/Nautino10/go-pdfgen/internal/models"
)

func DrawReport(report models.Report, pdf *PDFWriter) {
	y := 750
	pdf.AddText(100, y, 18, "Rapport d'Activité")
	y -= 30
	pdf.AddText(100, y, 12, fmt.Sprintf("Titre : %s", report.Titre))
	y -= 20
	pdf.AddText(100, y, 12, fmt.Sprintf("Auteur : %s", report.Auteur))
	y -= 20
	pdf.AddText(100, y, 12, fmt.Sprintf("Date : %s", report.Date))
	y -= 30

	pdf.AddText(100, y, 12, "Introduction :")
	y -= 20
	pdf.AddText(110, y, 10, report.Introduction)
	y -= 40

	pdf.AddText(100, y, 12, "Sections :")
	y -= 20
	for _, section := range report.Sections {
		pdf.AddText(110, y, 10, section)
		y -= 15
	}

	y -= 20
	pdf.AddText(100, y, 12, "Conclusion :")
	y -= 20
	pdf.AddText(110, y, 10, report.Conclusion)
}
