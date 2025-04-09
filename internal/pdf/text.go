package pdf

import (
	"fmt"
)

func (p *PDFWriter) AddText(x, y, fontSize int, text string) {
	stream := fmt.Sprintf(`<< /Length %d >>
stream
BT /F1 %d Tf %d %d Td (%s) Tj ET
endstream`,
		len(text)+40, fontSize, x, y, text)

	p.AddObject(stream)
}

func (p *PDFWriter) AddObject(content string) {
	p.addRawObject(content)
}
