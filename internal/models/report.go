package models

type Report struct {
	Titre        string
	Auteur       string
	Date         string
	Introduction string
	Sections     []string
	Conclusion   string
}

func FakeReport() Report {
	return Report{
		Titre:        "Rapport d'Activité Annuel",
		Auteur:       "Jean Dupuis",
		Date:         "09/04/2025",
		Introduction: "Ce rapport présente l'activité de l'entreprise sur l'année écoulée, ses réussites, ses défis et les perspectives pour l'année suivante.",
		Sections: []string{
			"1. Résultats financiers",
			"2. Développement commercial",
			"3. Projets en cours",
			"4. Ressources humaines",
		},
		Conclusion: "L'année a été globalement positive. De nombreux projets structurants ont été menés à bien et la croissance reste stable.",
	}
}
