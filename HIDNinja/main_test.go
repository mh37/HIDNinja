package main

import (
	"os"
	"testing"
)

func TestSendKey(t *testing.T) {
	// Create a temporary file to mock the HID device
	tmpFile, err := os.CreateTemp("", "hid-mock")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// Test writing to the file
	testData := []byte("test data")
	err = sendKey(tmpFile, testData)
	if err != nil {
		t.Fatalf("sendKey returned an error: %v", err)
	}

	// Verify the contents written to the file
	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to read temp file: %v", err)
	}
	if string(content) != string(testData) {
		t.Errorf("Expected content %q, got %q", string(testData), string(content))
	}

	// Test writing to a closed file (error case)
	tmpFile.Close()
	err = sendKey(tmpFile, testData)
	if err == nil {
		t.Errorf("Expected an error when writing to a closed file, got nil")
	}
}

func TestCharToKeystroke(t *testing.T) {
	tests := []struct {
		char        rune
		expectedMod byte
		expectedKey string
	}{
		{'A', 0x02, "A"},
		{'a', 0x00, "A"},
		{'1', 0x00, "1"},
		{'!', 0x02, "1"},
		{'-', 0x00, "MINUS"},
		{'_', 0x02, "MINUS"},
		{'\n', 0x00, "ENTER"},
		{',', 0x00, ","},
		{'<', 0x02, ","},
		{'>', 0x02, "."},
		{' ', 0x00, " "},
		{';', 0x00, ";"},
		{':', 0x02, ";"},
		{'?', 0x02, "SLASH"},
		{'©', 0x00, "©"}, // Fallback case
	}

	for _, tt := range tests {
		t.Run(string(tt.char), func(t *testing.T) {
			mod, key := charToKeystroke(tt.char)
			if mod != tt.expectedMod || key != tt.expectedKey {
				t.Errorf("charToKeystroke(%q) = 0x%02x, %q; want 0x%02x, %q", tt.char, mod, key, tt.expectedMod, tt.expectedKey)
			}
		})
	}
}
