package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

var scanCodes = map[string]byte{
	"LCTRL":              0x01,
	"LSHIFT":             0x02,
	"LALT":               0x04,
	"LMETA":              0x08,
	"RCTRL":              0x10,
	"RSHIFT":             0x20,
	"RALT":               0x40,
	"RMETA":              0x80,
	"NONE":               0x00,
	"ERR_OVF":            0x01,
	"A":                  0x04,
	"B":                  0x05,
	"C":                  0x06,
	"D":                  0x07,
	"E":                  0x08,
	"F":                  0x09,
	"G":                  0x0a,
	"H":                  0x0b,
	"I":                  0x0c,
	"J":                  0x0d,
	"K":                  0x0e,
	"L":                  0x0f,
	"M":                  0x10,
	"N":                  0x11,
	"O":                  0x12,
	"P":                  0x13,
	"Q":                  0x14,
	"R":                  0x15,
	"S":                  0x16,
	"T":                  0x17,
	"U":                  0x18,
	"V":                  0x19,
	"W":                  0x1a,
	"X":                  0x1b,
	"Y":                  0x1c,
	"Z":                  0x1d,
	"1":                  0x1e,
	"2":                  0x1f,
	"3":                  0x20,
	"4":                  0x21,
	"5":                  0x22,
	"6":                  0x23,
	"7":                  0x24,
	"8":                  0x25,
	"9":                  0x26,
	"0":                  0x27,
	"ENTER":              0x28,
	"ESC":                0x29,
	"BACKSPACE":          0x2a,
	"TAB":                0x2b,
	" ":                  0x2c,
	"MINUS":              0x2d,
	"EQUAL":              0x2e,
	"LEFTBRACE":          0x2f,
	"RIGHTBRACE":         0x30,
	"BACKSLASH":          0x31,
	"HASHTILDE":          0x32,
	";":                  0x33,
	"'":                  0x34,
	"GRAVE":              0x35,
	",":                  0x36,
	".":                  0x37,
	"SLASH":              0x38,
	"CAPSLOCK":           0x39,
	"F1":                 0x3a,
	"F2":                 0x3b,
	"F3":                 0x3c,
	"F4":                 0x3d,
	"F5":                 0x3e,
	"F6":                 0x3f,
	"F7":                 0x40,
	"F8":                 0x41,
	"F9":                 0x42,
	"F10":                0x43,
	"F11":                0x44,
	"F12":                0x45,
	"SYSRQ":              0x46,
	"SCROLLLOCK":         0x47,
	"PAUSE":              0x48,
	"INSERT":             0x49,
	"HOME":               0x4a,
	"PAGEUP":             0x4b,
	"DELETE":             0x4c,
	"END":                0x4d,
	"PAGEDOWN":           0x4e,
	"RIGHT":              0x4f,
	"LEFT":               0x50,
	"DOWN":               0x51,
	"UP":                 0x52,
	"NUMLOCK":            0x53,
	"KPSLASH":            0x54,
	"KPASTERISK":         0x55,
	"KPMINUS":            0x56,
	"KPPLUS":             0x57,
	"KPENTER":            0x58,
	"KP1":                0x59,
	"KP2":                0x5a,
	"KP3":                0x5b,
	"KP4":                0x5c,
	"KP5":                0x5d,
	"KP6":                0x5e,
	"KP7":                0x5f,
	"KP8":                0x60,
	"KP9":                0x61,
	"KP0":                0x62,
	"KPDOT":              0x63,
	"102ND":              0x64,
	"COMPOSE":            0x65,
	"POWER":              0x66,
	"KPEQUAL":            0x67,
	"F13":                0x68,
	"F14":                0x69,
	"F15":                0x6a,
	"F16":                0x6b,
	"F17":                0x6c,
	"F18":                0x6d,
	"F19":                0x6e,
	"F20":                0x6f,
	"F21":                0x70,
	"F22":                0x71,
	"F23":                0x72,
	"F24":                0x73,
	"OPEN":               0x74,
	"HELP":               0x75,
	"PROPS":              0x76,
	"FRONT":              0x77,
	"STOP":               0x78,
	"AGAIN":              0x79,
	"UNDO":               0x7a,
	"CUT":                0x7b,
	"COPY":               0x7c,
	"PASTE":              0x7d,
	"FIND":               0x7e,
	"MUTE":               0x7f,
	"VOLUMEUP":           0x80,
	"VOLUMEDOWN":         0x81,
	"KPCOMMA":            0x85,
	"RO":                 0x87,
	"KATAKANAHIRAGANA":   0x88,
	"YEN":                0x89,
	"HENKAN":             0x8a,
	"MUHENKAN":           0x8b,
	"KPJPCOMMA":          0x8c,
	"HANGEUL":            0x90,
	"HANJA":              0x91,
	"KATAKANA":           0x92,
	"HIRAGANA":           0x93,
	"ZENKAKUHANKAKU":     0x94,
	"KPLEFTPAREN":        0xb6,
	"KPRIGHTPAREN":       0xb7,
	"LEFTCTRL":           0xe0,
	"LEFTSHIFT":          0xe1,
	"LEFTALT":            0xe2,
	"LEFTMETA":           0xe3,
	"RIGHTCTRL":          0xe4,
	"RIGHTSHIFT":         0xe5,
	"RIGHTALT":           0xe6,
	"RIGHTMETA":          0xe7,
	"MEDIA_PLAYPAUSE":    0xe8,
	"MEDIA_STOPCD":       0xe9,
	"MEDIA_PREVIOUSSONG": 0xea,
	"MEDIA_NEXTSONG":     0xeb,
	"MEDIA_EJECTCD":      0xec,
	"MEDIA_VOLUMEUP":     0xed,
	"MEDIA_VOLUMEDOWN":   0xee,
	"MEDIA_MUTE":         0xef,
	"MEDIA_WWW":          0xf0,
	"MEDIA_BACK":         0xf1,
	"MEDIA_FORWARD":      0xf2,
	"MEDIA_STOP":         0xf3,
	"MEDIA_FIND":         0xf4,
	"MEDIA_SCROLLUP":     0xf5,
	"MEDIA_SCROLLDOWN":   0xf6,
	"MEDIA_EDIT":         0xf7,
	"MEDIA_SLEEP":        0xf8,
	"MEDIA_COFFEE":       0xf9,
	"MEDIA_REFRESH":      0xfa,
	"MEDIA_CALC":         0xfb,
}

// Defining the read and write buffer size for the upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Translates char/key into HID compatible code
func translationLayer(s string) byte {
	if val, ok := scanCodes[s]; ok {
		return val
	} else {
		log.Fatal("NOT FOUND:", s == "\n")
		return 0x00
	}
}

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

	//convert to upper case for mapping
	payloadString = strings.ToUpper(payloadString)

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

// Serve the index.html for the payload web interface
func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../PayloadInterface/index.html")
}

// WebSocket endpoint
func wsEndpoint(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade the connection to a websocket
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Client Connected")

	reader(conn)
}

// reader function that listens for incoming payloads
func reader(conn *websocket.Conn) {
	// keep listening for incoming payloads
	for {
		//read incoming payload
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		//print received message to console
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}

		//execute payload
		executePayload(string(msg))
	}
}

// Defines the routes, such as the WebSocket Endpoint and Homepage
func setupRoutes() {
	http.HandleFunc("/echo", wsEndpoint)
	http.HandleFunc("/", homePage)
}

func main() {
	setupRoutes()
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
