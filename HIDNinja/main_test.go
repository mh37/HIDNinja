package main

import (
	"bytes"
	"log"
	"os"
	"strings"
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

func TestExecutePayloadWithFile(t *testing.T) {
	// Create a temporary file to mock the HID device
	tmpFile, err := os.CreateTemp("", "hid-mock")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Test a simple payload
	payload := "aA"
	success := executePayloadWithFile(tmpFile, payload)
	if !success {
		t.Errorf("Expected success, got false")
	}

	// Read back and verify length or content
	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to read temp file: %v", err)
	}

	// 8 bytes per key press, 8 bytes per release.
	// 2 characters -> 4 writes -> 32 bytes total.
	if len(content) != 32 {
		t.Errorf("Expected 32 bytes written, got %d", len(content))
	}
	tmpFile.Close()
}

func TestExecutePayloadWithFile_WriteError(t *testing.T) {
	// Create a temporary file to mock the HID device
	tmpFile, err := os.CreateTemp("", "hid-mock-err")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Close it to force a write error
	tmpFile.Close()

	// Capture log output to check for error logs
	var logBuf bytes.Buffer
	log.SetOutput(&logBuf)
	defer log.SetOutput(os.Stderr)

	payload := "a"
	success := executePayloadWithFile(tmpFile, payload)

	// Function returns true even if write fails because it doesn't interrupt execution
	if !success {
		t.Errorf("Expected success even when writing fails, got false")
	}

	logOutput := logBuf.String()
	if !strings.Contains(logOutput, "Error sending key:") {
		t.Errorf("Expected error log for sending key, got %q", logOutput)
	}
}

func TestExecutePayload_OpenError(t *testing.T) {
	var logBuf bytes.Buffer
	log.SetOutput(&logBuf)
	defer log.SetOutput(os.Stderr)

	success := executePayload("test")

	// Verify the result depending on if /dev/hidg0 exists
	if _, err := os.Stat("/dev/hidg0"); os.IsNotExist(err) {
		if success {
			t.Errorf("Expected false when /dev/hidg0 doesn't exist")
		}
		if !strings.Contains(logBuf.String(), "Failed to open HID device:") {
			t.Errorf("Expected error log about failing to open HID device")
		}
	}
}
