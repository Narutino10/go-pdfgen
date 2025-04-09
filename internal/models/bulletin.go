package models

type Note struct {
	Matiere string
	Note    float64
	Coef    float64
}

type Bulletin struct {
	Eleve         string
	Classe        string
	Notes         []Note
	Noyenne       float64
	Appreciation  string
	Trimestre     string
	AnneeScolaire string
}

func FakeBulletin() Bulletin {
	notes := []Note{
		{"Mathématiques", 15.5, 4},
		{"Français", 14.0, 3},
		{"Histoire", 13.2, 2},
		{"SVT", 16.0, 2},
		{"Physique", 12.8, 3},
	}
	return Bulletin{
		Eleve:         "Alice Dupont",
		Classe:        "Terminale S",
		Trimestre:     "2",
		AnneeScolaire: "2024-2025",
		Notes:         notes,
		Noyenne:       14.3,
		Appreciation:  "Très bon trimestre, continue ainsi !",
	}
}
