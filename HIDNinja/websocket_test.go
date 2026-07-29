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

func TestSetupRoutes(t *testing.T) {
	// Reset the DefaultServeMux to avoid panic from multiple registrations
	http.DefaultServeMux = new(http.ServeMux)
	setupRoutes()

	// Test "/"
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status OK for /, got %v", rr.Code)
	}

	// Test "/echo"
	reqEcho, err := http.NewRequest("GET", "/echo", nil)
	if err != nil {
		t.Fatal(err)
	}
	rrEcho := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rrEcho, reqEcho)
	if rrEcho.Code != http.StatusBadRequest {
		t.Errorf("Expected status Bad Request for /echo, got %v", rrEcho.Code)
	}
}
