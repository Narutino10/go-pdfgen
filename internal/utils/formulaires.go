package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Nautino10/go-pdfgen/internal/models"
)

func SaisieBulletin() models.Bulletin {
	reader := bufio.NewReader(os.Stdin)

	// Saisie de l'année scolaire
	fmt.Print("Année scolaire (ex : 2024-2025) : ")
	annee, _ := reader.ReadString('\n')
	annee = strings.TrimSpace(annee)

	// Saisie de l'année d'étude
	fmt.Print("Année d'étude (ex : 5 pour 5e année) : ")
	niveauStr, _ := reader.ReadString('\n')
	niveauStr = strings.TrimSpace(niveauStr)
	niveau, _ := strconv.Atoi(niveauStr)
	classeLabel := fmt.Sprintf("%dIW", niveau)

	// Saisie du numéro de la classe
	fmt.Print("Numéro de classe (ex : 3) : ")
	numStr, _ := reader.ReadString('\n')
	numStr = strings.TrimSpace(numStr)

	// Construction nom de classe
	nomClasse := fmt.Sprintf("%s-%s", classeLabel, numStr)

	// Saisie du semestre
	fmt.Print("Semestre (1, 2 ou 3) : ")
	semestre, _ := reader.ReadString('\n')
	semestre = strings.TrimSpace(semestre)

	// Génération automatique des matières selon le niveau
	matières := []string{}
	switch niveau {
	case 5:
		matières = []string{"Go", "Airtable", "Docker", "Git", "Architecture Logicielle"}
	case 4:
		matières = []string{"Java", "DevOps", "Base de données", "API REST"}
	default:
		matières = []string{"Culture numérique", "Anglais", "Projet tutoré"}
	}

	// Saisie des notes
	var notes []models.Note
	var total, totalCoef float64
	fmt.Println("\n🎓 Veuillez saisir les notes pour les matières suivantes :")
	for _, matiere := range matières {
		fmt.Printf("%s - Note : ", matiere)
		noteStr, _ := reader.ReadString('\n')
		note, _ := strconv.ParseFloat(strings.TrimSpace(noteStr), 64)

		fmt.Print("Coefficient : ")
		coefStr, _ := reader.ReadString('\n')
		coef, _ := strconv.ParseFloat(strings.TrimSpace(coefStr), 64)

		notes = append(notes, models.Note{
			Matiere: matiere,
			Note:    note,
			Coef:    coef,
		})

		total += note * coef
		totalCoef += coef
	}

	// Calcul moyenne
	moyenne := 0.0
	if totalCoef > 0 {
		moyenne = total / totalCoef
	}

	// Appréciation automatique
	appreciation := ""
	switch {
	case moyenne >= 16:
		appreciation = "Excellent travail, continue ainsi."
	case moyenne >= 14:
		appreciation = "Très bon trimestre, bravo !"
	case moyenne >= 12:
		appreciation = "Bon travail dans l'ensemble."
	case moyenne >= 10:
		appreciation = "Trimestre satisfaisant, peut mieux faire."
	default:
		appreciation = "Résultats insuffisants, doit s'impliquer davantage."
	}

	// Saisie du nom de l'élève
	fmt.Print("\nNom complet de l'élève : ")
	eleve, _ := reader.ReadString('\n')
	eleve = strings.TrimSpace(eleve)

	return models.Bulletin{
		Eleve:         eleve,
		Classe:        nomClasse,
		Trimestre:     semestre,
		AnneeScolaire: annee,
		Notes:         notes,
		Noyenne:       moyenne,
		Appreciation:  appreciation,
	}
}

func SaisieFacture() models.Facture {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Numéro de facture : ")
	numero, _ := reader.ReadString('\n')
	numero = strings.TrimSpace(numero)

	fmt.Print("Date : ")
	date, _ := reader.ReadString('\n')
	date = strings.TrimSpace(date)

	fmt.Print("Client : ")
	client, _ := reader.ReadString('\n')
	client = strings.TrimSpace(client)

	var lignes []models.LigneFacture
	var totalHT float64

	for {
		fmt.Print("Description (laisser vide pour terminer) : ")
		desc, _ := reader.ReadString('\n')
		desc = strings.TrimSpace(desc)
		if desc == "" {
			break
		}

		fmt.Print("Quantité : ")
		quantStr, _ := reader.ReadString('\n')
		quant, _ := strconv.Atoi(strings.TrimSpace(quantStr))

		fmt.Print("Prix unitaire : ")
		prixStr, _ := reader.ReadString('\n')
		prix, _ := strconv.ParseFloat(strings.TrimSpace(prixStr), 64)

		lignes = append(lignes, models.LigneFacture{
			Description:  desc,
			Quantite:     quant,
			PrixUnitaire: prix,
		})
		totalHT += float64(quant) * prix
	}

	tva := totalHT * 0.2

	return models.Facture{
		Numero:   numero,
		Date:     date,
		Client:   client,
		Lignes:   lignes,
		TotalHT:  totalHT,
		TVA:      tva,
		TotalTTC: totalHT + tva,
	}
}

func SaisieRapport() models.Report {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Titre : ")
	titre, _ := reader.ReadString('\n')
	titre = strings.TrimSpace(titre)

	fmt.Print("Auteur : ")
	auteur, _ := reader.ReadString('\n')
	auteur = strings.TrimSpace(auteur)

	fmt.Print("Date : ")
	date, _ := reader.ReadString('\n')
	date = strings.TrimSpace(date)

	fmt.Print("Introduction : ")
	intro, _ := reader.ReadString('\n')
	intro = strings.TrimSpace(intro)

	var sections []string
	for {
		fmt.Print("Ajouter une section (laisser vide pour terminer) : ")
		section, _ := reader.ReadString('\n')
		section = strings.TrimSpace(section)
		if section == "" {
			break
		}
		sections = append(sections, section)
	}

	fmt.Print("Conclusion : ")
	conclusion, _ := reader.ReadString('\n')
	conclusion = strings.TrimSpace(conclusion)

	return models.Report{
		Titre:        titre,
		Auteur:       auteur,
		Date:         date,
		Introduction: intro,
		Sections:     sections,
		Conclusion:   conclusion,
	}
}
