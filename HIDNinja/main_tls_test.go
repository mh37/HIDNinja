package main

import (
	"crypto/tls"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestMainServerTLS(t *testing.T) {
	// Setup mux like tests do to prevent double registration panics
	http.DefaultServeMux = new(http.ServeMux)

	certFile := "server.crt"
	keyFile := "server.key"
	defer os.Remove(certFile)
	defer os.Remove(keyFile)

	go main()

	// Give the server time to start up
	time.Sleep(1 * time.Second)

	// Verify server is listening with TLS
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://localhost:3000/")
	if err != nil {
		t.Fatalf("Failed to connect to TLS server: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}
}
