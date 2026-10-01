package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealth(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	handler().ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("code = %d", w.Code)
	}
	if strings.TrimSpace(w.Body.String()) != `{"status":"ok"}` {
		t.Fatalf("body = %s", w.Body.String())
	}
}
