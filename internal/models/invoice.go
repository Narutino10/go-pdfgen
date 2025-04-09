package models

type LigneFacture struct {
	Description  string
	Quantite     int
	PrixUnitaire float64
}

type Facture struct {
	Numero   string
	Date     string
	Client   string
	Lignes   []LigneFacture
	TotalHT  float64
	TVA      float64
	TotalTTC float64
}

func FakeInvoice() Facture {
	lignes := []LigneFacture{
		{"Création de site web", 1, 1200.00},
		{"Maintenance mensuelle", 3, 150.00},
		{"Hébergement", 6, 10.00},
	}

	var totalHT float64
	for _, l := range lignes {
		totalHT += float64(l.Quantite) * l.PrixUnitaire
	}
	tva := totalHT * 0.2
	return Facture{
		Numero:   "FAC-2025-042",
		Date:     "09/04/2025",
		Client:   "Entreprise ABC",
		Lignes:   lignes,
		TotalHT:  totalHT,
		TVA:      tva,
		TotalTTC: totalHT + tva,
	}
}
