package main

import (
	"github.com/Nautino10/go-pdfgen/internal/models"
	"github.com/Nautino10/go-pdfgen/internal/pdf"
)

func main() {
	writer := pdf.NewPDFWriter()
	writer.AddPage()

	bulletin := models.FakeBulletin()
	pdf.DrawBulletin(bulletin, writer)

	err := writer.Save("bulletin.pdf")
	if err != nil {
		panic(err)
	}
}
