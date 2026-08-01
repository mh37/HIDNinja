package main

import (
	"log"
	"net/http"
	"os"
)

const (
	ModNone       byte = 0x00
	ModLeftCtrl   byte = 0x01
	ModLeftShift  byte = 0x02
	ModLeftAlt    byte = 0x04
	ModLeftGUI    byte = 0x08
	ModRightCtrl  byte = 0x10
	ModRightShift byte = 0x20
	ModRightAlt   byte = 0x40
	ModRightGUI   byte = 0x80
)

// Send the byte sequence of keystrokes to the virtual HID (keyboard) where it will be sent to the target host over USB
func sendKey(f *os.File, code []byte) error {
	_, err := f.Write(code)
	if err != nil {
		return err
	}
	return nil
}

// charToKeystroke takes a rune and returns the required modifier byte and string mapping for translation
type keystroke struct {
	modifier byte
	keyStr   string
}

var symbolMap = map[rune]keystroke{
	'!':  {ModLeftShift, "1"},
	'@':  {ModLeftShift, "2"},
	'#':  {ModLeftShift, "3"},
	'$':  {ModLeftShift, "4"},
	'%':  {ModLeftShift, "5"},
	'^':  {ModLeftShift, "6"},
	'&':  {ModLeftShift, "7"},
	'*':  {ModLeftShift, "8"},
	'(':  {ModLeftShift, "9"},
	')':  {ModLeftShift, "0"},
	'-':  {ModNone, "MINUS"},
	'_':  {ModLeftShift, "MINUS"},
	'=':  {ModNone, "EQUAL"},
	'+':  {ModLeftShift, "EQUAL"},
	'[':  {ModNone, "LEFTBRACE"},
	'{':  {ModLeftShift, "LEFTBRACE"},
	']':  {ModNone, "RIGHTBRACE"},
	'}':  {ModLeftShift, "RIGHTBRACE"},
	'\\': {ModNone, "BACKSLASH"},
	'|':  {ModLeftShift, "BACKSLASH"},
	';':  {ModNone, ";"},
	':':  {ModLeftShift, ";"},
	'\'': {ModNone, "'"},
	'"':  {ModLeftShift, "'"},
	'`':  {ModNone, "GRAVE"},
	'~':  {ModLeftShift, "GRAVE"},
	',':  {ModNone, ","},
	'<':  {ModLeftShift, ","},
	'.':  {ModNone, "."},
	'>':  {ModLeftShift, "."},
	'/':  {ModNone, "SLASH"},
	'?':  {ModLeftShift, "SLASH"},
	' ':  {ModNone, " "},
	'\n': {ModNone, "ENTER"},
	'\t': {ModNone, "TAB"},
}

func charToKeystroke(ch rune) (byte, string) {
	var modifier byte = ModNone
	var keyStr string

	switch {
	case ch >= 'A' && ch <= 'Z':
		modifier = ModLeftShift // LSHIFT
		keyStr = string(ch)
	case ch >= 'a' && ch <= 'z':
		keyStr = string(ch - 32)
	case ch >= '0' && ch <= '9':
		keyStr = string(ch)
	default:
		if ks, ok := symbolMap[ch]; ok {
			modifier = ks.modifier
			keyStr = ks.keyStr
		} else {
			keyStr = string(ch) // Fallback
		}
	}
	return modifier, keyStr
}

// executePayloadWithFile processes a payload string and writes it to the provided file descriptor.
func executePayloadWithFile(f *os.File, payloadString string) bool {
	payloadBuf := make([]byte, 8)
	releaseBuf := make([]byte, 8)

	//run through each character/rune in the payload string, translate it to a scancode and send it to the virtual HID
	for _, ch := range payloadString {
		modifier, keyStr := charToKeystroke(ch)

		var runeKey rune
		if len(keyStr) > 0 {
			runeKey = rune(keyStr[0])
		}

		key := translationLayer(runeKey)

		payloadBuf[0] = modifier
		payloadBuf[2] = key

		if err := sendKey(f, payloadBuf); err != nil {
			log.Println("Error sending key:", err)
		}
		// release keys
		if err := sendKey(f, releaseBuf); err != nil {
			log.Println("Error releasing key:", err)
		}
	}

	return true
}

// The function takes a payload string and processes the individual characters, so that they can be correctly translated, processed, and sent to the target host.
func executePayload(payloadString string) bool {
	f, err := os.OpenFile("/dev/hidg0", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Println("Failed to open HID device:", err)
		return false
	}
	defer f.Close()

	return executePayloadWithFile(f, payloadString)
}

func main() {
	setupRoutes()

	certFile := "server.crt"
	keyFile := "server.key"

	if _, err := os.Stat(certFile); os.IsNotExist(err) {
		log.Println("Generating self-signed certificate...")
		if err := generateCert(certFile, keyFile); err != nil {
			log.Fatalf("Failed to generate certificate: %v", err)
		}
	}

	log.Println("Waiting for client connection ...")
	log.Fatal(http.ListenAndServeTLS(":3000", certFile, keyFile, nil))
}
