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
- ✅ Rendu final 100% compatible Acrobat/Chrome

---

## 📁 Structure du projet

```
go-pdfgen/
├── main.go                    # Interface utilisateur CLI
├── internal/
│   ├── pdf/                   # Moteur PDF fait maison
│   │   └── writer.go
│   ├── models/                # Données métiers structurées
│   │   ├── bulletin.go
│   │   ├── invoice.go
│   │   └── report.go
│   ├── utils/                 # Fonctions de saisie terminale
│   │   └── formulaires.go
└── README.md
```

---

## 🚀 Exemple d'utilisation

```go
writer := pdf.NewPDFWriter()
writer.AddPage()
writer.AddText(100, 700, 14, "Élève : Ibrahim OUAHABI")
writer.AddText(100, 680, 12, "Classe : 5IW-3")
writer.Save("test.pdf")
```

---

## 🧪 Lancer le projet

```bash
git clone https://github.com/Narutino10/go-pdfgen.git
cd go-pdfgen
go run main.go
```

> 💡 Un menu interactif vous permet de générer :
> - un bulletin
> - une facture
> - un rapport
> le tout depuis le terminal.

---

## 🔧 Améliorations futures

- [ ] Ajout de styles : bordures, lignes, tableaux
- [ ] Multipages / génération longue
- [ ] Choix de polices dynamiques
- [ ] Export en tant que vraie librairie Go (module)

---

## 🧑‍💻 Auteurs

**Ibrahim OUAHABI**  
**Cheick LANIKPEKOUN**  
Projet réalisé dans le cadre du Master 2 IW3 – ESGI (2025)

