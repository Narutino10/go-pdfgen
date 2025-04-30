package pdf

// Représente un modèle configurable de PDF
type TemplateConfig struct {
	Title      string
	FontSize   int
	MarginTop  int
	Align      string // "left", "center", "right"
	PageWidth  int    // en points (ex: 595 pour A4)
	MarginLeft int
}

// Applique un template avec un titre aligné
func (p *PDFWriter) ApplyTemplate(t TemplateConfig) {
	switch t.Align {
	case "center":
		p.AddTextCenter(t.MarginTop, t.FontSize, t.Title)
	case "right":
		p.AddTextRight(t.MarginTop, t.FontSize, t.Title)
	default:
		p.AddTextLeft(t.MarginTop, t.FontSize, t.Title)
	}
}

// Affiche du texte aligné à gauche
func (p *PDFWriter) AddTextLeft(y, fontSize int, text string) {
	x := 40 // marge gauche fixe
	p.AddText(x, y, fontSize, text)
}

// Affiche du texte centré
func (p *PDFWriter) AddTextCenter(y, fontSize int, text string) {
	pageWidth := 595
	textWidth := len(text) * fontSize / 2
	x := (pageWidth - textWidth) / 2
	p.AddText(x, y, fontSize, text)
}

// Affiche du texte aligné à droite
func (p *PDFWriter) AddTextRight(y, fontSize int, text string) {
	pageWidth := 595
	textWidth := len(text) * fontSize / 2
	x := pageWidth - textWidth - 40
	p.AddText(x, y, fontSize, text)
}
