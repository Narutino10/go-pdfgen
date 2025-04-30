package pdf

import (
	"fmt"
	"os"
	"strings"
)

type PDFWriter struct {
	objects        []string
	xref           []int
	offset         int
	contentObjects []string // pour stocker AddText
}

func NewPDFWriter() *PDFWriter {
	return &PDFWriter{
		objects:        []string{},
		xref:           []int{},
		offset:         0,
		contentObjects: []string{},
	}
}

func (p *PDFWriter) addRawObject(content string) {
	p.xref = append(p.xref, p.offset)
	formatted := fmt.Sprintf("%d 0 obj\n%s\nendobj\n", len(p.xref), content)
	p.objects = append(p.objects, formatted)
	p.offset += len(formatted)
}

func (p *PDFWriter) AddObject(content string) {
	p.contentObjects = append(p.contentObjects, content)
}

// Nettoyage des caractères spéciaux pour éviter les bugs d'encodage
func sanitizeText(s string) string {
	r := strings.NewReplacer(
		"é", "e",
		"è", "e",
		"ê", "e",
		"ë", "e",
		"à", "a",
		"â", "a",
		"î", "i",
		"ï", "i",
		"ô", "o",
		"ù", "u",
		"û", "u",
		"ç", "c",
		"É", "E",
		"À", "A",
		"€", "EUR",
		"😊", ":)",
		"🚀", "",
		"’", "'",
		"œ", "oe",
	)
	return r.Replace(s)
}

// Fonction publique pour écrire du texte
func (p *PDFWriter) AddText(x, y, fontSize int, text string) {
	cleanText := sanitizeText(text)
	stream := fmt.Sprintf(`<< /Length %d >>
stream
BT /F1 %d Tf %d %d Td (%s) Tj ET
endstream`,
		len(cleanText)+40, fontSize, x, y, cleanText)

	p.AddObject(stream)
}

func (p *PDFWriter) BuildPDF() string {
	header := "%PDF-1.4\n"

	// 1. Catalog
	p.addRawObject("<< /Type /Catalog /Pages 2 0 R >>")

	// 2. Pages
	p.addRawObject("<< /Type /Pages /Kids [3 0 R] /Count 1 >>")

	// 3. Page (avec /Resources et /Font intégré)
	pageObj := `<< /Type /Page
/Parent 2 0 R
/MediaBox [0 0 595 842]
/Contents 4 0 R
/Resources <<
    /Font <<
      /F1 <<
        /Type /Font
        /Subtype /Type1
        /BaseFont /Helvetica
        /Encoding /WinAnsiEncoding
      >>
    >>
  >>
>>`
	p.addRawObject(pageObj)

	// 4. Contents
	contentStream := ""
	for _, content := range p.contentObjects {
		contentStream += content + "\n"
	}
	stream := fmt.Sprintf("<< /Length %d >>\nstream\n%s\nendstream", len(contentStream), contentStream)
	p.addRawObject(stream)

	// assembler
	body := ""
	for _, obj := range p.objects {
		body += obj
	}

	xrefStart := len(header + body)
	xref := fmt.Sprintf("xref\n0 %d\n0000000000 65535 f \n", len(p.xref)+1)
	for _, pos := range p.xref {
		xref += fmt.Sprintf("%010d 00000 n \n", pos)
	}

	trailer := fmt.Sprintf(`trailer
<< /Size %d
   /Root 1 0 R >>
startxref
%d
%%%%EOF`, len(p.xref)+1, xrefStart)

	return header + body + xref + trailer
}

func (p *PDFWriter) Save(path string) error {
	data := p.BuildPDF()
	return os.WriteFile(path, []byte(data), 0644)
}
