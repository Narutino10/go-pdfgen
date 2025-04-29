package main

import (
	"fmt"
	"log"

	"github.com/Nautino10/go-pdfgen/internal/pdf"
	"github.com/Nautino10/go-pdfgen/internal/utils"
)

func main() {
	for {
		var docType string

		fmt.Println("\n📄 Que voulez-vous générer ?")
		fmt.Println("1 - Bulletin scolaire")
		fmt.Println("2 - Facture")
		fmt.Println("3 - Rapport d'activité")
		fmt.Print("Votre choix (1, 2 ou 3) : ")
		fmt.Scanln(&docType)

		writer := pdf.NewPDFWriter()
		writer.AddPage()

		switch docType {
		case "1":
			bulletin := utils.SaisieBulletin()
			pdf.DrawBulletin(bulletin, writer)
			err := writer.Save("bulletin.pdf")
			if err != nil {
				log.Fatalf("Erreur sauvegarde Bulletin : %v", err)
			}
			fmt.Println("✅ Bulletin généré sous 'bulletin.pdf'")
		case "2":
			invoice := utils.SaisieFacture()
			pdf.DrawInvoice(invoice, writer)
			err := writer.Save("facture.pdf")
			if err != nil {
				log.Fatalf("Erreur sauvegarde Facture : %v", err)
			}
			fmt.Println("✅ Facture générée sous 'facture.pdf'")
		case "3":
			report := utils.SaisieRapport()
			pdf.DrawReport(report, writer)
			err := writer.Save("rapport.pdf")
			if err != nil {
				log.Fatalf("Erreur sauvegarde Rapport : %v", err)
			}
			fmt.Println("✅ Rapport généré sous 'rapport.pdf'")
		default:
			fmt.Println("⚠️ Choix invalide. Veuillez entrer 1, 2 ou 3.")
		}

		var again string
		fmt.Print("\n🔁 Voulez-vous générer un autre document ? (O/N) : ")
		fmt.Scanln(&again)
		if again != "O" && again != "o" {
			fmt.Println("\n👋 Merci d'avoir utilisé le générateur PDF !")
			break
		}
	}
}
