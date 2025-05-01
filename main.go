package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Job struct {
	ID   string                 `json:"id"`
	Data map[string]interface{} `json:"data"`
}

var jobQueue chan Job

func main() {
	jobQueue = make(chan Job, 100)

	StartWorker()

	http.HandleFunc("/generate-pdf", handleGeneratePDF)
	http.HandleFunc("/pdf/", handleDownloadPDF)

	log.Println("API running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleGeneratePDF(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var job Job
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid payload"))
		return
	}

	jobQueue <- job

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Job received and added to the queue"))
}

func handleDownloadPDF(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/pdf/"):]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID not provided"))
		return
	}
	filePath := "pdfs/" + id + ".pdf"
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+id+".pdf\"")
	http.ServeFile(w, r, filePath)
} 