package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/websocket"
)

var keycodeMap = map[int]string{
	3:   "x48", // Pause / Break
	8:   "x2a", // Backspace / Delete
	9:   "x2b", // Tab
	12:  "x53", // Clear
	13:  "x28", // Enter
	16:  "xe1", // Shift (Left)
	17:  "xe0", // Ctrl (left)
	18:  "xe1", // Alt (left)
	19:  "x48", // Pause / Break
	20:  "x39", // Caps Lock
	21:  "x90", // Hangeul
	25:  "x91", // Hanja
	27:  "x29", // Escape
	32:  "x2c", // Spacebar
	33:  "x4b", // Page Up
	34:  "x4e", // Page Down
	35:  "x4d", // End
	36:  "x4a", // Home
	37:  "x50", // Left Arrow
	38:  "x52", // Up Arrow
	39:  "x4f", // Right Arrow
	40:  "x51", // Down Arrow
	41:  "x77", // Select
	43:  "x74", // Execute
	44:  "x46", // Print Screen
	45:  "x49", // Insert
	46:  "x4c", // Delete
	47:  "x75", // Help
	48:  "x27", // 0
	49:  "x1e", // 1
	50:  "x1f", // 2
	51:  "x20", // 3
	52:  "x21", // 4
	53:  "x22", // 5
	54:  "x23", // 6
	55:  "x24", // 7
	56:  "x25", // 8
	57:  "x26", // 9
	59:  "x33", // Semicolon
	60:  "xc5", // <
	61:  "x2e", // Equal sign
	65:  "x04", // a
	66:  "x05", // b
	67:  "x06", // c
	68:  "x07", // d
	69:  "x08", // e
	70:  "x09", // f
	71:  "x0a", // g
	72:  "x0b", // h
	73:  "x0c", // i
	74:  "x0d", // j
	75:  "x0e", // k
	76:  "x0f", // l
	77:  "x10", // m
	78:  "x11", // n
	79:  "x12", // o
	80:  "x13", // p
	81:  "x14", // q
	82:  "x15", // r
	83:  "x16", // s
	84:  "x17", // t
	85:  "x18", // u
	86:  "x19", // v
	87:  "x1a", // w
	88:  "x1b", // x
	89:  "x1c", // y
	90:  "x1d", // z
	91:  "xe3", // Windows key / Meta Key (Left)
	96:  "x62", // Numpad 0
	97:  "x59", // Numpad 1
	98:  "x5a", // Numpad 2
	99:  "x5b", // Numpad 3
	100: "x5c", // Numpad 4
	101: "x5d", // Numpad 5
	102: "x5e", // Numpad 6
	103: "x5f", // Numpad 7
	104: "x60", // Numpad 8
	105: "x61", // Numpad 9
	112: "x3b", // F1
	113: "x3c", // F2
	114: "x3d", // F3
	115: "x3e", // F4
	116: "x3f", // F5
	117: "x40", // F6
	118: "x41", // F7
	119: "x42", // F8
	120: "x43", // F9
	121: "x44", // F10
	122: "x45", // F11
	123: "x46", // F12
	124: "x68", // F13
	125: "x69", // F14
	126: "x6a", // F15
	127: "x6b", // F16
	128: "x6c", // F17
	129: "x6d", // F18
	130: "x6e", // F19
	131: "x6f", // F20
	132: "x70", // F21
	133: "x71", // F22
	134: "x72", // F23
	144: "x53", // Num Lock
	145: "x47", // Scroll Lock
	161: "x1e", // !
	163: "x32", // Hash
	173: "x2d", // Minus
	179: "xe8", // Media play/pause
	168: "xfa", // Refresh
	186: "x33", // Semicolon
	187: "x2e", // Equal sign
	188: "x36", // Comma
	189: "x2d", // Minus sign
	190: "x37", // Period
	191: "x38", // Forward slash
	192: "x35", // Accent grave
	219: "x2f", // Left bracket ([, {])
	220: "x31", // Back slash
	221: "x30", // Right bracket (], })
	222: "x34", // Single quote
	223: "x35", // Accent grave (`)
}

// Defining the read and write buffer size for the upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//Translates char/key into HID compatible code
func translationLayer(c rune) string {
	// for example: letter 'a' --> `\0\0\x4\0\0\0\0\0`
	keycode := keycodeMap[int(c)]
	hidCode := `\0\0\` + keycode + `\0\0\0\0\0`
	return hidCode
}

//execute bash command
func execCmd(key string) (err error) {
	//example cmd: echo -ne "\0\0\x4\0\0\0\0\0" > /dev/hidg0
	gadget := "/dev/hidg0"
	cmdObj := exec.Command("bash", "-c", `echo -ne "`+key+`" > `+gadget)
	cmdObj.Stdout = os.Stdout
	cmdObj.Stderr = os.Stderr
	err = cmdObj.Run()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// TODO: This works for single character payloads but not more. Needs investigating.
func executePayload(payloadString string) bool {

	//run through each character/rune in the payload string
	for _, ch := range payloadString {
		key := translationLayer(ch)
		fmt.Println("HID KEY: " + key)
		execCmd(key)
		// release key
		execCmd(`\0\0\0\0\0\0\0\0`)
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
