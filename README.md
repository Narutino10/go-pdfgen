# 📄 pdfgen-lib

> Librairie Go de génération de fichiers PDF dynamiques, 100% codée maison, sans dépendance externe.

---

## ✨ Objectif

Créer une **librairie modulaire, simple à utiliser et performante** en Go permettant la génération de fichiers PDF depuis des données structurées (structs, JSON), sans utiliser de bibliothèque externe comme `gofpdf` ou `unipdf`.

---

## 🧠 Pourquoi ce projet ?

- Comprendre le format PDF et sa structure bas niveau
- Apprendre à manipuler des fichiers binaires en Go
- Créer un outil **professionnel**, **personnalisable** et **open-source**
- Automatiser des tâches concrètes de génération documentaire

---

## 📚 Fonctionnalités prévues

- [x] Génération brute de fichiers PDF avec du texte
- [ ] Positionnement des blocs texte
- [ ] Prise en charge de plusieurs pages
- [ ] Gestion de tableaux et colonnes
- [ ] Génération de documents types :
  - [ ] Bulletins scolaires 📘
  - [ ] Rapports d’activités 📊
  - [ ] Factures automatisées 🧾

---

## 🏗 Structure du projet

```
pdfgen-lib/
├── main.go                # Entrée principale pour tester la librairie
├── internal/
│   ├── pdf/               # Moteur PDF fait maison
│   │   └── writer.go
│   └── models/            # Données structurées à formater
│       ├── bulletin.go
│       ├── invoice.go
│       └── report.go
└── README.md
```

---

## 🚀 Exemple d'utilisation

```go
pdf := pdf.NewPDFWriter()
pdf.AddText(100, 700, "Bonjour le monde")
pdf.Save("test.pdf")
```

---

## ⚒️ Installation

```bash
git clone https://github.com/Narutino10/go-pdfgen.git
cd go-pdfgen
go run main.go
```

---

## 🔧 À venir

- 🎨 Rendu plus avancé (fonctions de style, lignes, tableaux)
- 📄 Export multi-pages
- 📦 Packaging en module Go

---

## 🧑‍💻 Auteure

**Ibrahim OUAHABI**  **Cheick LANIKPEKOUN**
Projet réalisé dans le cadre de l’ESGI (2025), Master 2 IW3

