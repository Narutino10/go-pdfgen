# 📄 pdfgen-lib

> Librairie Go de génération de fichiers PDF dynamiques, 100% codée maison, sans dépendance externe.

---

## ✨ Objectif

Créer une **librairie modulaire, simple à utiliser et performante** en Go permettant la génération de fichiers PDF depuis des données structurées (structs, saisie terminale), sans utiliser de bibliothèques externes comme `gofpdf` ou `unidoc`.

---

## 🧠 Pourquoi ce projet ?

- Comprendre la structure interne des fichiers PDF (catalog, pages, contenu)
- Manipuler des flux binaires en Go (fichiers `.pdf` générés manuellement)
- Créer un outil **ouvert**, **réutilisable** et **compréhensible**
- Proposer une alternative légère pour générer automatiquement :
  - Bulletins scolaires 📘
  - Factures 🧾
  - Rapports d’activités 📊

---

## 📚 Fonctionnalités actuelles

- ✅ Génération manuelle de PDF (Header, Catalog, Pages, Text)
- ✅ Positionnement des blocs texte avec `AddText(x, y, fontSize, string)`
- ✅ Nettoyage automatique des caractères spéciaux (`€, é, 🚀`…)
- ✅ Interface de **saisie terminale interactive** :
  - `SaisieBulletin()`
  - `SaisieFacture()`
  - `SaisieRapport()`
- ✅ **Lecture de PDF existants** (MVP) :
  - Extraction du texte des PDFs générés par cette librairie
  - Analyse des informations du document (nombre de pages, objets, etc.)
  - Support des caractères spéciaux et accents français
- ✅ Rendu final 100% compatible Acrobat/Chrome

---

## 📁 Structure du projet

```
go-pdfgen/
├── main.go                    # Interface utilisateur CLI
├── internal/
│   ├── pdf/                   # Moteur PDF fait maison
│   │   ├── writer.go          # Écriture/génération PDF
│   │   ├── reader.go          # Lecture/extraction PDF
│   │   └── ...                # Autres composants PDF
│   ├── models/                # Données métiers structurées
│   │   ├── bulletin.go
│   │   ├── invoice.go
│   │   └── report.go
│   └── utils/                 # Fonctions utilitaires
│       ├── formulaires.go     # Saisie interactive
│       └── lecteur.go         # Lecture de PDF
└── README.md
```

---

## 🚀 Exemples d'utilisation

### Génération de PDF
```go
writer := pdf.NewPDFWriter()
writer.AddPage()
writer.AddText(100, 700, 14, "Élève : Ibrahim OUAHABI")
writer.AddText(100, 680, 12, "Classe : 5IW-3")
writer.Save("test.pdf")
```

### Lecture de PDF
```go
// Lecture du contenu textuel
texte, err := utils.LirePDF("document.pdf")
if err != nil {
    log.Fatal(err)
}
fmt.Println(texte)

// Analyse du document
info, err := utils.AnalyserPDF("document.pdf")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Nombre de pages : %v\n", info["nombre_pages"])
```

---

## 🧪 Lancer le projet

```bash
git clone https://github.com/Narutino10/go-pdfgen.git
cd go-pdfgen
go run main.go
```

> 💡 Un menu interactif vous permet de :
> - **Générer** un bulletin, une facture ou un rapport
> - **Lire** un PDF existant et extraire son contenu textuel
> le tout depuis le terminal.

> ⚠️ **Note importante** : La fonctionnalité de lecture PDF est actuellement en MVP (Minimum Viable Product). Elle fonctionne principalement avec les PDFs générés par cette librairie. Le support pour tous types de PDFs sera ajouté prochainement.

---

## 🔧 Améliorations futures

### Génération PDF
- [ ] Ajout de styles : bordures, lignes, tableaux
- [ ] Multipages / génération longue
- [ ] Choix de polices dynamiques
- [ ] Export en tant que vraie librairie Go (module)

### Lecture PDF
- [ ] Support complet de tous les formats PDF standards
- [ ] Lecture des images et éléments graphiques
- [ ] Extraction de métadonnées avancées
- [ ] Support des PDF avec compression/chiffrement

---

## 🧑‍💻 Auteurs

**Ibrahim OUAHABI**  
**Cheick LANIKPEKOUN**  
Projet réalisé dans le cadre du Master 2 IW3 – ESGI (2025)

