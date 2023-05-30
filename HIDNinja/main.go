package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

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

func executePayload(payloadString string) bool {

	//convert to upper case for standardized mapping
	payloadString = strings.ToUpper(payloadString)

	// TODO insert shift key scancodes for genuine uppercase

	//TODO special treatment for modifiers needed

	//run through each character/rune in the payload string
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
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
