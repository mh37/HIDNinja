package main

import (
	"testing"
)

func TestTranslationLayer(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected byte
	}{
		{"Valid Key A", "A", 0x04},
		{"Valid Key ENTER", "ENTER", 0x28},
		{"Valid Key 1", "1", 0x1e},
		{"Valid Key Space", " ", 0x2c},
		{"Valid Key LCTRL", "LCTRL", 0x01},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := translationLayer(tt.input)
			if result != tt.expected {
				t.Errorf("translationLayer(%q) = 0x%02x; want 0x%02x", tt.input, result, tt.expected)
			}
		})
	}
}
