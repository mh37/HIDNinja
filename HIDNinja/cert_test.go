package main

import (
	"os"
	"testing"
)

func TestGenerateCert(t *testing.T) {
	certFile := "test_server.crt"
	keyFile := "test_server.key"

	defer os.Remove(certFile)
	defer os.Remove(keyFile)

	err := generateCert(certFile, keyFile)
	if err != nil {
		t.Fatalf("generateCert failed: %v", err)
	}

	if _, err := os.Stat(certFile); os.IsNotExist(err) {
		t.Errorf("cert file %s was not created", certFile)
	}

	if _, err := os.Stat(keyFile); os.IsNotExist(err) {
		t.Errorf("key file %s was not created", keyFile)
	}
}
