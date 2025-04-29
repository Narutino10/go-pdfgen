// internal/pdf/invoice.go
package pdf

import (
	"fmt"

	"github.com/Nautino10/go-pdfgen/internal/models"
)

func DrawInvoice(invoice models.Facture, pdf *PDFWriter) {
	y := 750
	pdf.AddText(100, y, 18, "Facture")
	y -= 30
	pdf.AddText(100, y, 12, fmt.Sprintf("Numéro : %s", invoice.Numero))
	y -= 20
	pdf.AddText(100, y, 12, fmt.Sprintf("Date : %s", invoice.Date))
	y -= 20
	pdf.AddText(100, y, 12, fmt.Sprintf("Client : %s", invoice.Client))
	y -= 30

	for _, ligne := range invoice.Lignes {
		line := fmt.Sprintf("%s - Quantité: %d - Prix Unitaire: %.2f €", ligne.Description, ligne.Quantite, ligne.PrixUnitaire)
		pdf.AddText(100, y, 10, line)
		y -= 15
	}

	pdf.AddText(100, y, 12, fmt.Sprintf("Total HT : %.2f €", invoice.TotalHT))
	y -= 20
	pdf.AddText(100, y, 12, fmt.Sprintf("TVA : %.2f €", invoice.TVA))
	y -= 20
	pdf.AddText(100, y, 12, fmt.Sprintf("Total TTC : %.2f €", invoice.TotalTTC))
}
