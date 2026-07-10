package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomePage(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.FileServer(http.Dir("../PayloadInterface"))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedSubstring := "<title>HIDNinja</title>"
	if !strings.Contains(rr.Body.String(), expectedSubstring) {
		t.Errorf("handler returned unexpected body: missing %q", expectedSubstring)
	}

	contentType := rr.Header().Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, "text/html")
	}
}
