package main

import (
	"log"
	"bytes"
	"text/template"
	"os"
	"fmt"
)

func renderHTML(data map[string]interface{}) (string, error) {
	tmpl, err := template.ParseFiles("templates/template.html")
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func StartWorker() {
	go func() {
		for job := range jobQueue {
			log.Printf("Processing job: %v", job)

			html, err := renderHTML(job.Data)
			if err != nil {
				log.Printf("Error rendering HTML: %v", err)
				continue
			}

			pdf, err := SendToGotenberg(html)
			if err != nil {
				log.Printf("Error generating PDF: %v", err)
				continue
			}

			log.Printf("PDF generated for job %s, size: %d bytes", job.ID, len(pdf))

			filePath := fmt.Sprintf("pdfs/%s.pdf", job.ID)
			err = os.WriteFile(filePath, pdf, 0644)
			if err != nil {
				log.Printf("Error saving PDF: %v", err)
				continue
			}
			log.Printf("PDF saved at %s", filePath)
		}
	}()
} 