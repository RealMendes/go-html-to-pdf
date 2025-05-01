package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

var gotenbergURL = ""

func SendToGotenberg(html string) ([]byte, error) {
	if gotenbergURL == "" {
		gotenbergURL = os.Getenv("GOTENBERG_URL")
		if gotenbergURL == "" {
			gotenbergURL = "http://gotenberg:3000"
		}
	}

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	part, err := writer.CreateFormFile("files", "document.html")
	if err != nil {
		return nil, err
	}
	_, err = part.Write([]byte(html))
	if err != nil {
		return nil, err
	}

	writer.Close()

	resp, err := http.Post(gotenbergURL+"/forms/libreoffice/convert", writer.FormDataContentType(), &buf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pdf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return pdf, nil
} 