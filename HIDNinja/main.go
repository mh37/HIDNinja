package main

import (
	"log"
	"net/http"
	"os"
	"unicode"
)

// Send the byte sequence of keystrokes to the virtual HID (keyboard) where it will be sent to the target host over USB
func sendKey(f *os.File, code []byte) error {
	_, err := f.Write(code)
	if err != nil {
		return err
	}
	return nil
}

// charToKeystroke takes a rune and returns the required modifier byte and mapping for translation
func charToKeystroke(ch rune) (byte, rune) {
	var modifier byte = 0x00
	var keyRune rune

	switch {
	case ch >= 'A' && ch <= 'Z':
		modifier = 0x02 // LSHIFT
		keyRune = ch
	case ch >= 'a' && ch <= 'z':
		keyRune = unicode.ToUpper(ch)
	case ch >= '0' && ch <= '9':
		keyRune = ch
	default:
		switch ch {
		case '!': modifier = 0x02; keyRune = '1'
		case '@': modifier = 0x02; keyRune = '2'
		case '#': modifier = 0x02; keyRune = '3'
		case '$': modifier = 0x02; keyRune = '4'
		case '%': modifier = 0x02; keyRune = '5'
		case '^': modifier = 0x02; keyRune = '6'
		case '&': modifier = 0x02; keyRune = '7'
		case '*': modifier = 0x02; keyRune = '8'
		case '(': modifier = 0x02; keyRune = '9'
		case ')': modifier = 0x02; keyRune = '0'
		case '-': keyRune = '-'
		case '_': modifier = 0x02; keyRune = '-'
		case '=': keyRune = '='
		case '+': modifier = 0x02; keyRune = '='
		case '[': keyRune = '['
		case '{': modifier = 0x02; keyRune = '['
		case ']': keyRune = ']'
		case '}': modifier = 0x02; keyRune = ']'
		case '\\': keyRune = '\\'
		case '|': modifier = 0x02; keyRune = '\\'
		case ';': keyRune = ';'
		case ':': modifier = 0x02; keyRune = ';'
		case '\'': keyRune = '\''
		case '"': modifier = 0x02; keyRune = '\''
		case '`': keyRune = '`'
		case '~': modifier = 0x02; keyRune = '`'
		case ',': keyRune = ','
		case '<': modifier = 0x02; keyRune = ','
		case '.': keyRune = '.'
		case '>': modifier = 0x02; keyRune = '.'
		case '/': keyRune = '/'
		case '?': modifier = 0x02; keyRune = '/'
		case ' ': keyRune = ' '
		case '\n': keyRune = '\n'
		case '\t': keyRune = '\t'
		default:
			keyRune = ch // Fallback
		}
	}
	return modifier, keyRune
}

// The function takes a payload string and processes the individual characters, so that they can be correctly translated, processed, and sent to the target host.
func executePayload(payloadString string) bool {

	f, err := os.OpenFile("/dev/hidg0", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Println("Error opening /dev/hidg0:", err)
		return false
	}
	defer f.Close()

	//run through each character/rune in the payload string, translate it to a scancode and send it to the virtual HID
	for _, ch := range payloadString {
		modifier, keyRune := charToKeystroke(ch)
		key := translationLayer(keyRune)

		if err := sendKey(f, []byte{modifier, 0x00, key, 0x00, 0x00, 0x00, 0x00, 0x00}); err != nil {
			log.Println("Error sending key:", err)
		}
		// release keys
		if err := sendKey(f, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}); err != nil {
			log.Println("Error releasing key:", err)
		}
	}

	return true
}

func main() {
	setupRoutes()
	log.Println("Waiting for client connection ...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
