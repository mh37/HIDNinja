package main

import "log"

// scan code table is a modified version of https://github.com/gsora/hid-compiler/blob/master/compiler/Scancodes.go
var scanCodes = map[rune]byte{
	// LCTRL removed as it cannot be a single rune
	// LSHIFT removed as it cannot be a single rune
	// LALT removed as it cannot be a single rune
	// LMETA removed as it cannot be a single rune
	// RCTRL removed as it cannot be a single rune
	// RSHIFT removed as it cannot be a single rune
	// RALT removed as it cannot be a single rune
	// RMETA removed as it cannot be a single rune
	// NONE removed as it cannot be a single rune
	// ERR_OVF removed as it cannot be a single rune
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
	'-':    0x2d, // MINUS
	'=':    0x2e, // EQUAL
	'[':    0x2f, // LEFTBRACE
	']':    0x30, // RIGHTBRACE
	'\\':   0x31, // BACKSLASH
	// HASHTILDE removed as it cannot be a single rune
	';':  0x33,
	'\'': 0x34,
	'`':  0x35, // GRAVE
	',': 0x36,
	'.': 0x37,
	'/': 0x38, // SLASH
	// CAPSLOCK removed as it cannot be a single rune
	// F1 removed as it cannot be a single rune
	// F2 removed as it cannot be a single rune
	// F3 removed as it cannot be a single rune
	// F4 removed as it cannot be a single rune
	// F5 removed as it cannot be a single rune
	// F6 removed as it cannot be a single rune
	// F7 removed as it cannot be a single rune
	// F8 removed as it cannot be a single rune
	// F9 removed as it cannot be a single rune
	// F10 removed as it cannot be a single rune
	// F11 removed as it cannot be a single rune
	// F12 removed as it cannot be a single rune
	// SYSRQ removed as it cannot be a single rune
	// SCROLLLOCK removed as it cannot be a single rune
	// PAUSE removed as it cannot be a single rune
	// INSERT removed as it cannot be a single rune
	// HOME removed as it cannot be a single rune
	// PAGEUP removed as it cannot be a single rune
	// DELETE removed as it cannot be a single rune
	// END removed as it cannot be a single rune
	// PAGEDOWN removed as it cannot be a single rune
	// RIGHT removed as it cannot be a single rune
	// LEFT removed as it cannot be a single rune
	// DOWN removed as it cannot be a single rune
	// UP removed as it cannot be a single rune
	// NUMLOCK removed as it cannot be a single rune
	// KPSLASH removed as it cannot be a single rune
	// KPASTERISK removed as it cannot be a single rune
	// KPMINUS removed as it cannot be a single rune
	// KPPLUS removed as it cannot be a single rune
	// KPENTER removed as it cannot be a single rune
	// KP1 removed as it cannot be a single rune
	// KP2 removed as it cannot be a single rune
	// KP3 removed as it cannot be a single rune
	// KP4 removed as it cannot be a single rune
	// KP5 removed as it cannot be a single rune
	// KP6 removed as it cannot be a single rune
	// KP7 removed as it cannot be a single rune
	// KP8 removed as it cannot be a single rune
	// KP9 removed as it cannot be a single rune
	// KP0 removed as it cannot be a single rune
	// KPDOT removed as it cannot be a single rune
	// 102ND removed as it cannot be a single rune
	// COMPOSE removed as it cannot be a single rune
	// POWER removed as it cannot be a single rune
	// KPEQUAL removed as it cannot be a single rune
	// F13 removed as it cannot be a single rune
	// F14 removed as it cannot be a single rune
	// F15 removed as it cannot be a single rune
	// F16 removed as it cannot be a single rune
	// F17 removed as it cannot be a single rune
	// F18 removed as it cannot be a single rune
	// F19 removed as it cannot be a single rune
	// F20 removed as it cannot be a single rune
	// F21 removed as it cannot be a single rune
	// F22 removed as it cannot be a single rune
	// F23 removed as it cannot be a single rune
	// F24 removed as it cannot be a single rune
	// OPEN removed as it cannot be a single rune
	// HELP removed as it cannot be a single rune
	// PROPS removed as it cannot be a single rune
	// FRONT removed as it cannot be a single rune
	// STOP removed as it cannot be a single rune
	// AGAIN removed as it cannot be a single rune
	// UNDO removed as it cannot be a single rune
	// CUT removed as it cannot be a single rune
	// COPY removed as it cannot be a single rune
	// PASTE removed as it cannot be a single rune
	// FIND removed as it cannot be a single rune
	// MUTE removed as it cannot be a single rune
	// VOLUMEUP removed as it cannot be a single rune
	// VOLUMEDOWN removed as it cannot be a single rune
	// KPCOMMA removed as it cannot be a single rune
	// RO removed as it cannot be a single rune
	// KATAKANAHIRAGANA removed as it cannot be a single rune
	// YEN removed as it cannot be a single rune
	// HENKAN removed as it cannot be a single rune
	// MUHENKAN removed as it cannot be a single rune
	// KPJPCOMMA removed as it cannot be a single rune
	// HANGEUL removed as it cannot be a single rune
	// HANJA removed as it cannot be a single rune
	// KATAKANA removed as it cannot be a single rune
	// HIRAGANA removed as it cannot be a single rune
	// ZENKAKUHANKAKU removed as it cannot be a single rune
	// KPLEFTPAREN removed as it cannot be a single rune
	// KPRIGHTPAREN removed as it cannot be a single rune
	// LEFTCTRL removed as it cannot be a single rune
	// LEFTSHIFT removed as it cannot be a single rune
	// LEFTALT removed as it cannot be a single rune
	// LEFTMETA removed as it cannot be a single rune
	// RIGHTCTRL removed as it cannot be a single rune
	// RIGHTSHIFT removed as it cannot be a single rune
	// RIGHTALT removed as it cannot be a single rune
	// RIGHTMETA removed as it cannot be a single rune
	// MEDIA_PLAYPAUSE removed as it cannot be a single rune
	// MEDIA_STOPCD removed as it cannot be a single rune
	// MEDIA_PREVIOUSSONG removed as it cannot be a single rune
	// MEDIA_NEXTSONG removed as it cannot be a single rune
	// MEDIA_EJECTCD removed as it cannot be a single rune
	// MEDIA_VOLUMEUP removed as it cannot be a single rune
	// MEDIA_VOLUMEDOWN removed as it cannot be a single rune
	// MEDIA_MUTE removed as it cannot be a single rune
	// MEDIA_WWW removed as it cannot be a single rune
	// MEDIA_BACK removed as it cannot be a single rune
	// MEDIA_FORWARD removed as it cannot be a single rune
	// MEDIA_STOP removed as it cannot be a single rune
	// MEDIA_FIND removed as it cannot be a single rune
	// MEDIA_SCROLLUP removed as it cannot be a single rune
	// MEDIA_SCROLLDOWN removed as it cannot be a single rune
	// MEDIA_EDIT removed as it cannot be a single rune
	// MEDIA_SLEEP removed as it cannot be a single rune
	// MEDIA_COFFEE removed as it cannot be a single rune
	// MEDIA_REFRESH removed as it cannot be a single rune
	// MEDIA_CALC removed as it cannot be a single rune
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
