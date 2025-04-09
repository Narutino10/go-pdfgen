// internal/pdf/writer.go
package pdf

import (
	"fmt"
	"os"
)

type PDFWriter struct {
	objects []string
	xref    []int
	offset  int
}

func NewPDFWriter() *PDFWriter {
	return &PDFWriter{
		objects: []string{},
		xref:    []int{},
		offset:  0,
	}
}

func (p *PDFWriter) addRawObject(content string) {
	p.xref = append(p.xref, p.offset)
	formatted := fmt.Sprintf("%d 0 obj\n%s\nendobj\n", len(p.xref), content)
	p.objects = append(p.objects, formatted)
	p.offset += len(formatted)
}

func (p *PDFWriter) BuildPDF(text string) string {
	p.objects = []string{} 
	p.xref = []int{}
	p.offset = 0

	header := "%PDF-1.4\n"

	// 1. Catalog
	p.addRawObject("<< /Type /Catalog /Pages 2 0 R >>")

	// 2. Pages
	p.addRawObject("<< /Type /Pages /Kids [3 0 R] /Count 1 >>")

	// 3. Page
	p.addRawObject("<< /Type /Page /Parent 2 0 R /MediaBox [0 0 595 842] /Contents 4 0 R /Resources << >> >>")

	// 4. Contents (le texte)
	stream := fmt.Sprintf("<< /Length %d >>\nstream\nBT /F1 24 Tf 100 700 Td (%s) Tj ET\nendstream", len(text)+40, text)
	p.addRawObject(stream)

	// Build body
	body := ""
	for _, obj := range p.objects {
		body += obj
	}

	// Build xref
	xrefStart := len(header + body)
	xref := fmt.Sprintf("xref\n0 %d\n0000000000 65535 f \n", len(p.xref)+1)
	for _, pos := range p.xref {
		xref += fmt.Sprintf("%010d 00000 n \n", pos)
	}

	// Build trailer
	trailer := fmt.Sprintf(`trailer
<< /Size %d
   /Root 1 0 R >>
startxref
%d
%%%%EOF`, len(p.xref)+1, xrefStart)

	return header + body + xref + trailer
}

func (p *PDFWriter) Save(path string, text string) error {
	data := p.BuildPDF(text)
	return os.WriteFile(path, []byte(data), 0644)
}
