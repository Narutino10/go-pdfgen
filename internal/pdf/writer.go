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

func (p *PDFWriter) AddObject(content string) {
	p.xref = append(p.xref, p.offset)
	formatted := fmt.Sprintf("%d 0 obj\n%s\nendobj\n", len(p.xref), content)
	p.objects = append(p.objects, formatted)
	p.offset += len(formatted)
}

func (p *PDFWriter) BuildPDF() string {
	header := "%PDF-1.4\n"
	body := ""
	for _, obj := range p.objects {
		body += obj
	}

	xrefStart := len(header + body)
	xref := "xref\n0 " + fmt.Sprint(len(p.xref)+1) + "\n0000000000 65535 f \n"
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
