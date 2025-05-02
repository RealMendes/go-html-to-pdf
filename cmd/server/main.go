package main

import (
	"go-html-to-pdf/internal/api"
	"go-html-to-pdf/internal/worker"
	"log"
	"net/http"
)

func main() {
	jobQueue := make(chan worker.Job, 100)

	worker.StartWorker(jobQueue)

	handler := api.NewHandler(jobQueue)

	http.HandleFunc("/generate-pdf", handler.HandleGeneratePDF)
	http.HandleFunc("/pdf/", handler.HandleDownloadPDF)

	log.Println("API running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
