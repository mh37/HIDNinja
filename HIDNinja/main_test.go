package main

import (
	"testing"
)

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
