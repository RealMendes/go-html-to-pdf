package main

import (
	"testing"
)

func TestRenderHTML(t *testing.T) {
	data := map[string]interface{}{
		"titulo":   "Test PDF",
		"mensagem": "Test message",
		"itens": map[string]interface{}{
			"item1": "value1",
			"item2": "value2",
		},
	}

	html, err := renderHTML(data)
	if err != nil {
		t.Fatalf("Error rendering HTML: %v", err)
	}

	if len(html) == 0 {
		t.Error("Generated HTML is empty")
	}

	if !contains(html, "Test PDF") || !contains(html, "Test message") {
		t.Error("Expected content not found in generated HTML")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && (contains(s[1:], substr) || contains(s[:len(s)-1], substr)))
} 