package main

import "log"

// scan code table is a modified version of https://github.com/gsora/hid-compiler/blob/master/compiler/Scancodes.go
var scanCodes = map[rune]byte{
	'A':    0x04,
	'B':    0x05,
	'C':    0x06,
	'D':    0x07,
	'E':    0x08,
	'F':    0x09,
	'G':    0x0a,
	'H':    0x0b,
	'I':    0x0c,
	'J':    0x0d,
	'K':    0x0e,
	'L':    0x0f,
	'M':    0x10,
	'N':    0x11,
	'O':    0x12,
	'P':    0x13,
	'Q':    0x14,
	'R':    0x15,
	'S':    0x16,
	'T':    0x17,
	'U':    0x18,
	'V':    0x19,
	'W':    0x1a,
	'X':    0x1b,
	'Y':    0x1c,
	'Z':    0x1d,
	'1':    0x1e,
	'2':    0x1f,
	'3':    0x20,
	'4':    0x21,
	'5':    0x22,
	'6':    0x23,
	'7':    0x24,
	'8':    0x25,
	'9':    0x26,
	'0':    0x27,
	'\n':   0x28, // ENTER
	'\x1b': 0x29, // ESC
	'\b':   0x2a, // BACKSPACE
	'\t':   0x2b, // TAB
	' ':    0x2c,
	';':  0x33,
	'\'': 0x34,
	',': 0x36,
	'.': 0x37,
}

// Translates char/key into HID compatible code
func translationLayer(s rune) byte {
	if val, ok := scanCodes[s]; ok {
		return val
	} else {
		log.Printf("NOT FOUND: %q", string(s))
		return 0x00
	}
}
