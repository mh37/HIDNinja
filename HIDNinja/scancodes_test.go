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
		{
			name:     "Valid key A",
			input:    "A",
			expected: 0x04,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: 0x00,
		},
		{
			name:     "Unknown character",
			input:    "?",
			expected: 0x00,
		},
		{
			name:     "Lowercase character",
			input:    "a",
			expected: 0x00,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var runeInput rune
			if len(tt.input) > 0 {
				runeInput = rune(tt.input[0])
			}
			result := translationLayer(runeInput)
			if result != tt.expected {
				t.Errorf("translationLayer(%q) = 0x%02x; want 0x%02x", tt.input, result, tt.expected)
			}
		})
	}
}
