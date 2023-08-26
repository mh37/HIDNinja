package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

//TODO: Provide Modifier Key Handling
// Declare a struct type for the payload handling
/*
type hidPayload struct {
	Modifier   modifier
	Character0 string
}*/

// Send the byte sequence of keystrokes to the virtual HID (keyboard) where it will be sent to the target host over USB
func sendKey(code []byte) {
	f, err := os.OpenFile("/dev/hidg0", os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	_, err = f.Write(code)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// The function takes a payload string and processes the individual characters, so that they can be correctly translated, processed, and sent to the target host.
func executePayload(payloadString string) bool {

	//convert to upper case for standardized mapping
	payloadString = strings.ToUpper(payloadString)

	// TODO insert shift key scancodes for genuine uppercase representation
	// TODO special treatment for modifiers needed

	//run through each character/rune in the payload string, translate it to a scancode and send it to the virtual HID
	for _, ch := range payloadString {
		key := translationLayer(string(ch))
		fmt.Println(ch) // for testing purposes, remove later
		sendKey([]byte{0x00, 0x00, key, 0x00, 0x00, 0x00, 0x00, 0x00})
		// release keys
		sendKey([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	}

	return true
}

func main() {
	setupRoutes()
	log.Println("Waiting for client connection ...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
