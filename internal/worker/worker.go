package worker

import (
	"bytes"
	"fmt"
	"go-html-to-pdf/internal/gotenberg"
	"log"
	"os"
	"text/template"
)

type Job struct {
	ID   string                 `json:"id"`
	Data map[string]interface{} `json:"data"`
}

func RenderHTML(data map[string]interface{}) (string, error) {
	tmpl, err := template.ParseFiles("../../templates/template.html")
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

func StartWorker(jobQueue <-chan Job) {
	go func() {
		for job := range jobQueue {
			log.Printf("Processing job: %v", job)

			html, err := RenderHTML(job.Data)
			if err != nil {
				log.Printf("Error rendering HTML: %v", err)
				continue
			}

			pdf, err := gotenberg.SendToGotenberg(html)
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
