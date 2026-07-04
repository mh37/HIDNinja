package main

import (
	"testing"
)

func TestTranslationLayer(t *testing.T) {
	tests := []struct {
		name     string
		input    rune
		expected byte
	}{
		{
			name:     "Valid key A",
			input:    'A',
			expected: 0x04,
		},
		{
			name:     "Zero rune",
			input:    0,
			expected: 0x00,
		},
		{
			name:     "Unknown rune",
			input:    '?',
			expected: 0x00,
		},
		{
			name:     "Lowercase character",
			input:    'a',
			expected: 0x00,
		},
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
