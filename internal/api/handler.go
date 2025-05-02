package api

import (
	"encoding/json"
	"go-html-to-pdf/internal/worker"
	"net/http"
)

type Handler struct {
	JobQueue chan<- worker.Job
}

func NewHandler(jobQueue chan<- worker.Job) *Handler {
	return &Handler{JobQueue: jobQueue}
}

func (h *Handler) HandleGeneratePDF(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var job worker.Job
	if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid payload"))
		return
	}

	h.JobQueue <- job

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Job received and added to the queue"))
}

func (h *Handler) HandleDownloadPDF(w http.ResponseWriter, r *http.Request) {
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
