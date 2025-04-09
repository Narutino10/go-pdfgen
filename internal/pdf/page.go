// internal/pdf/page.go
package pdf

func (p *PDFWriter) AddPage() {
	// Objet /Pages
	p.AddObject(`<< /Type /Pages /Kids [3 0 R] /Count 1 >>`)

	// Objet /Page
	p.AddObject(`<< /Type /Page /Parent 2 0 R /MediaBox [0 0 595 842] /Contents 4 0 R /Resources << >> >>`)
}
