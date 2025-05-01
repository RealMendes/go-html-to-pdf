package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendToGotenberg(t *testing.T) {
	// Cria um servidor HTTP mock
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/pdf")
		w.Write([]byte("PDF MOCK"))
	}))
	defer ts.Close()

	html := "<html><body>Test</body></html>"

	// Substitui a URL do Gotenberg pela do mock
	oldURL := gotenbergURL
	gotenbergURL = ts.URL
	defer func() { gotenbergURL = oldURL }()

	pdf, err := SendToGotenberg(html)
	if err != nil {
		t.Fatalf("Error sending to Gotenberg mock: %v", err)
	}

	if string(pdf) != "PDF MOCK" {
		t.Errorf("Expected 'PDF MOCK', got: %s", string(pdf))
	}
} 