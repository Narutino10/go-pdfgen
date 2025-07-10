package pdf

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

type PDFReader struct {
	content string
	objects map[int]string
}

func NewPDFReader() *PDFReader {
	return &PDFReader{
		objects: make(map[int]string),
	}
}

// LoadPDF charge un fichier PDF et parse son contenu
func (r *PDFReader) LoadPDF(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture du fichier: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content strings.Builder
	for scanner.Scan() {
		content.WriteString(scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("erreur lors de la lecture du fichier: %v", err)
	}

	r.content = content.String()
	return r.parseObjects()
}

// parseObjects extrait tous les objets PDF du contenu
func (r *PDFReader) parseObjects() error {
	objectRegex := regexp.MustCompile(`(?s)(\d+)\s+0\s+obj\s+(.*?)\s+endobj`)
	matches := objectRegex.FindAllStringSubmatch(r.content, -1)

	for _, match := range matches {
		if len(match) >= 3 {
			objNum, err := strconv.Atoi(match[1])
			if err != nil {
				continue
			}
			r.objects[objNum] = match[2]
		}
	}

	// Analyser aussi les objets imbriqués dans le contenu
	for objNum, content := range r.objects {
		if strings.Contains(content, "stream") {
			streamRegex := regexp.MustCompile(`(?s)(<< /Length \d+ >>\s*stream\s+.*?\s+endstream)`)
			streamMatches := streamRegex.FindAllString(content, -1)
			
			if len(streamMatches) > 1 {
				for i, streamContent := range streamMatches {
					fakeObjNum := objNum*100 + i 
					r.objects[fakeObjNum] = streamContent
				}
			}
		}
	}

	return nil
}

// ExtractText extrait tout le texte visible du PDF
func (r *PDFReader) ExtractText() (string, error) {
	var allTexts []string

	for objNum := 1; objNum <= 50; objNum++ { 
		if objectContent, exists := r.objects[objNum]; exists {
			// Chercher les objets contenant du texte
			if strings.Contains(objectContent, "stream") && strings.Contains(objectContent, "Tj") {
				text := r.extractTextFromStream(objectContent)
				if text != "" {
					allTexts = append(allTexts, text)
				}
			}
		}
	}

	// Chercher aussi dans tous les objets au cas où ils ne seraient pas consécutifs
	for _, objectContent := range r.objects {
		if strings.Contains(objectContent, "stream") && strings.Contains(objectContent, "Tj") {
			text := r.extractTextFromStream(objectContent)
			if text != "" {
				found := false
				for _, existing := range allTexts {
					if existing == text {
						found = true
						break
					}
				}
				if !found {
					allTexts = append(allTexts, text)
				}
			}
		}
	}

	return strings.Join(allTexts, "\n"), nil
}

// extractTextFromStream extrait le texte d'un objet stream
func (r *PDFReader) extractTextFromStream(streamContent string) string {
	streamRegex := regexp.MustCompile(`(?s)stream\s+(.*?)\s+endstream`)
	matches := streamRegex.FindStringSubmatch(streamContent)
	
	if len(matches) < 2 {
		return ""
	}

	streamData := matches[1]
	return r.parseTextCommands(streamData)
}

// parseTextCommands analyse les commandes de texte PDF
func (r *PDFReader) parseTextCommands(streamData string) string {
	textRegex := regexp.MustCompile(`\((.*?)\)\s*Tj`)
	matches := textRegex.FindAllStringSubmatch(streamData, -1)
	
	var results []string
	for _, match := range matches {
		if len(match) >= 2 {
			text := r.decodePDFString(match[1])
			if text != "" {
				results = append(results, text)
			}
		}
	}
	
	return strings.Join(results, " ")
}

// decodePDFString décode une chaîne PDF basique
func (r *PDFReader) decodePDFString(pdfString string) string {
	decoded := strings.ReplaceAll(pdfString, `\n`, "\n")
	decoded = strings.ReplaceAll(decoded, `\r`, "\r")
	decoded = strings.ReplaceAll(decoded, `\t`, "\t")
	decoded = strings.ReplaceAll(decoded, `\\`, "\\")
	decoded = strings.ReplaceAll(decoded, `\(`, "(")
	decoded = strings.ReplaceAll(decoded, `\)`, ")")
	
	decoded = r.fixUTF8Encoding(decoded)
	
	return decoded
}

// fixUTF8Encoding corrige l'encodage des caractères UTF-8
func (r *PDFReader) fixUTF8Encoding(text string) string {
	if utf8.ValidString(text) {
		text = strings.ReplaceAll(text, "Ã©", "é")
		text = strings.ReplaceAll(text, "Ã¨", "è")
		text = strings.ReplaceAll(text, "Ã ", "à")
		text = strings.ReplaceAll(text, "Ã´", "ô")
		text = strings.ReplaceAll(text, "Ã¢", "â")
		text = strings.ReplaceAll(text, "Ã®", "î")
		text = strings.ReplaceAll(text, "Ã¯", "ï")
		text = strings.ReplaceAll(text, "Ã¼", "ü")
		text = strings.ReplaceAll(text, "Ã¹", "ù")
		text = strings.ReplaceAll(text, "Ã§", "ç")
		text = strings.ReplaceAll(text, "Ã‰", "É")
		text = strings.ReplaceAll(text, "Ã€", "À")
		text = strings.ReplaceAll(text, "Ã‡", "Ç")
		text = strings.ReplaceAll(text, "Ã‹", "Ë")
		text = strings.ReplaceAll(text, "Ã'", "Ñ")
		text = strings.ReplaceAll(text, "Ã±", "ñ")
		text = strings.ReplaceAll(text, "Ã«", "ë")
		text = strings.ReplaceAll(text, "Ãª", "ê")
		text = strings.ReplaceAll(text, "Ã­", "í")
		text = strings.ReplaceAll(text, "Ã³", "ó")
		text = strings.ReplaceAll(text, "Ãº", "ú")
		text = strings.ReplaceAll(text, "Ã½", "ý")
		text = strings.ReplaceAll(text, "Ã¿", "ÿ")
	}
	
	return text
}

// GetDocumentInfo retourne des informations sur le document
func (r *PDFReader) GetDocumentInfo() map[string]interface{} {
	info := make(map[string]interface{})
	
	info["nombre_objets"] = len(r.objects)
	
	for _, objectContent := range r.objects {
		if strings.Contains(objectContent, "/Type /Catalog") {
			info["type"] = "Document PDF"
		}
		if strings.Contains(objectContent, "/Type /Pages") {
			pageCountRegex := regexp.MustCompile(`/Count\s+(\d+)`)
			matches := pageCountRegex.FindStringSubmatch(objectContent)
			if len(matches) >= 2 {
				if count, err := strconv.Atoi(matches[1]); err == nil {
					info["nombre_pages"] = count
				}
			}
		}
	}
	
	return info
}
