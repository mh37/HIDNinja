package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"unicode"
)

//TODO: Provide Modifier Key Handling

// Send the byte sequence of keystrokes to the virtual HID (keyboard) where it will be sent to the target host over USB
func sendKey(f *os.File, code []byte) error {
	_, err := f.Write(code)
	if err != nil {
		return err
	}
	return nil
}

// charToKeystroke takes a rune and returns the required modifier byte and string mapping for translation
func charToKeystroke(ch rune) (byte, string) {
	var modifier byte = 0x00
	var keyStr string

	switch {
	case ch >= 'A' && ch <= 'Z':
		modifier = 0x02 // LSHIFT
		keyStr = string(ch)
	case ch >= 'a' && ch <= 'z':
		keyStr = strings.ToUpper(string(ch))
	case ch >= '0' && ch <= '9':
		keyStr = string(ch)
	default:
		switch ch {
		case '!': modifier = 0x02; keyStr = "1"
		case '@': modifier = 0x02; keyStr = "2"
		case '#': modifier = 0x02; keyStr = "3"
		case '$': modifier = 0x02; keyStr = "4"
		case '%': modifier = 0x02; keyStr = "5"
		case '^': modifier = 0x02; keyStr = "6"
		case '&': modifier = 0x02; keyStr = "7"
		case '*': modifier = 0x02; keyStr = "8"
		case '(': modifier = 0x02; keyStr = "9"
		case ')': modifier = 0x02; keyStr = "0"
		case '-': keyStr = "MINUS"
		case '_': modifier = 0x02; keyStr = "MINUS"
		case '=': keyStr = "EQUAL"
		case '+': modifier = 0x02; keyStr = "EQUAL"
		case '[': keyStr = "LEFTBRACE"
		case '{': modifier = 0x02; keyStr = "LEFTBRACE"
		case ']': keyStr = "RIGHTBRACE"
		case '}': modifier = 0x02; keyStr = "RIGHTBRACE"
		case '\\': keyStr = "BACKSLASH"
		case '|': modifier = 0x02; keyStr = "BACKSLASH"
		case ';': keyStr = ";"
		case ':': modifier = 0x02; keyStr = ";"
		case '\'': keyStr = "'"
		case '"': modifier = 0x02; keyStr = "'"
		case '`': keyStr = "GRAVE"
		case '~': modifier = 0x02; keyStr = "GRAVE"
		case ',': keyStr = ","
		case '<': modifier = 0x02; keyStr = ","
		case '.': keyStr = "."
		case '>': modifier = 0x02; keyStr = "."
		case '/': keyStr = "SLASH"
		case '?': modifier = 0x02; keyStr = "SLASH"
		case ' ': keyStr = " "
		case '\n': keyStr = "ENTER"
		case '\t': keyStr = "TAB"
		default:
			keyStr = string(ch) // Fallback
		}
	}
	return modifier, keyStr
}

// The function takes a payload string and processes the individual characters, so that they can be correctly translated, processed, and sent to the target host.
func executePayload(payloadString string) bool {

	//run through each character/rune in the payload string, translate it to a scancode and send it to the virtual HID
	for _, ch := range payloadString {
		modifier, keyStr := charToKeystroke(ch)
		key := translationLayer(keyStr)

		if err := sendKey([]byte{modifier, 0x00, key, 0x00, 0x00, 0x00, 0x00, 0x00}); err != nil {
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
