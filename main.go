package main

import (
	"go-pdfgen/internal/models"
	"go-pdfgen/internal/pdf"
)

func main() {
	pdfWriter := pdf.NewPDFWriter()
	bulletin := models.FakeBulletin()

	pdf.DrawBulletin(bulletin, pdfWriter)

	err := pdfWriter.Save("bulletin.pdf")
	if err != nil {
		panic(err)
	}
}
