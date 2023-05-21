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
	3:   `\0\0\x48\0\0\0\0\0`, // Pause / Break
	8:   `\0\0\x2a\0\0\0\0\0`, // Backspace / Delete
	9:   `\0\0\x2b\0\0\0\0\0`, // Tab
	12:  `\0\0\x53\0\0\0\0\0`, // Clear
	13:  `\0\0\x28\0\0\0\0\0`, // Enter
	16:  `\0\0\xe1\0\0\0\0\0`, // Shift (Left)
	17:  `\0\0\xe0\0\0\0\0\0`, // Ctrl (left)
	18:  `\0\0\xe1\0\0\0\0\0`, // Alt (left)
	19:  `\0\0\x48\0\0\0\0\0`, // Pause / Break
	20:  `\0\0\x39\0\0\0\0\0`, // Caps Lock
	21:  `\0\0\x90\0\0\0\0\0`, // Hangeul
	25:  `\0\0\x91\0\0\0\0\0`, // Hanja
	27:  `\0\0\x29\0\0\0\0\0`, // Escape
	32:  `\0\0\x2c\0\0\0\0\0`, // Spacebar
	33:  `\0\0\x4b\0\0\0\0\0`, // Page Up
	34:  `\0\0\x4e\0\0\0\0\0`, // Page Down
	35:  `\0\0\x4d\0\0\0\0\0`, // End
	36:  `\0\0\x4a\0\0\0\0\0`, // Home
	37:  `\0\0\x50\0\0\0\0\0`, // Left Arrow
	38:  `\0\0\x52\0\0\0\0\0`, // Up Arrow
	39:  `\0\0\x4f\0\0\0\0\0`, // Right Arrow
	40:  `\0\0\x51\0\0\0\0\0`, // Down Arrow
	41:  `\0\0\x77\0\0\0\0\0`, // Select
	43:  `\0\0\x74\0\0\0\0\0`, // Execute
	44:  `\0\0\x46\0\0\0\0\0`, // Print Screen
	45:  `\0\0\x49\0\0\0\0\0`, // Insert
	46:  `\0\0\x4c\0\0\0\0\0`, // Delete
	47:  `\0\0\x75\0\0\0\0\0`, // Help
	48:  `\0\0\x27\0\0\0\0\0`, // 0
	49:  `\0\0\x1e\0\0\0\0\0`, // 1
	50:  `\0\0\x1f\0\0\0\0\0`, // 2
	51:  `\0\0\x20\0\0\0\0\0`, // 3
	52:  `\0\0\x21\0\0\0\0\0`, // 4
	53:  `\0\0\x22\0\0\0\0\0`, // 5
	54:  `\0\0\x23\0\0\0\0\0`, // 6
	55:  `\0\0\x24\0\0\0\0\0`, // 7
	56:  `\0\0\x25\0\0\0\0\0`, // 8
	57:  `\0\0\x26\0\0\0\0\0`, // 9
	59:  `\0\0\x33\0\0\0\0\0`, // Semicolon
	60:  `\0\0\xc5\0\0\0\0\0`, // <
	61:  `\0\0\x2e\0\0\0\0\0`, // Equal sign
	65:  `\0\0\x04\0\0\0\0\0`, // a
	66:  `\0\0\x05\0\0\0\0\0`, // b
	67:  `\0\0\x06\0\0\0\0\0`, // c
	68:  `\0\0\x07\0\0\0\0\0`, // d
	69:  `\0\0\x08\0\0\0\0\0`, // e
	70:  `\0\0\x09\0\0\0\0\0`, // f
	71:  `\0\0\x0a\0\0\0\0\0`, // g
	72:  `\0\0\x0b\0\0\0\0\0`, // h
	73:  `\0\0\x0c\0\0\0\0\0`, // i
	74:  `\0\0\x0d\0\0\0\0\0`, // j
	75:  `\0\0\x0e\0\0\0\0\0`, // k
	76:  `\0\0\x0f\0\0\0\0\0`, // l
	77:  `\0\0\x10\0\0\0\0\0`, // m
	78:  `\0\0\x11\0\0\0\0\0`, // n
	79:  `\0\0\x12\0\0\0\0\0`, // o
	80:  `\0\0\x13\0\0\0\0\0`, // p
	81:  `\0\0\x14\0\0\0\0\0`, // q
	82:  `\0\0\x15\0\0\0\0\0`, // r
	83:  `\0\0\x16\0\0\0\0\0`, // s
	84:  `\0\0\x17\0\0\0\0\0`, // t
	85:  `\0\0\x18\0\0\0\0\0`, // u
	86:  `\0\0\x19\0\0\0\0\0`, // v
	87:  `\0\0\x1a\0\0\0\0\0`, // w
	88:  `\0\0\x1b\0\0\0\0\0`, // x
	89:  `\0\0\x1c\0\0\0\0\0`, // y
	90:  `\0\0\x1d\0\0\0\0\0`, // z
	91:  `\0\0\xe3\0\0\0\0\0`, // Windows key / Meta Key (Left)
	96:  `\0\0\x62\0\0\0\0\0`, // Numpad 0
	97:  `\0\0\x59\0\0\0\0\0`, // Numpad 1
	98:  `\0\0\x5a\0\0\0\0\0`, // Numpad 2
	99:  `\0\0\x5b\0\0\0\0\0`, // Numpad 3
	100: `\0\0\x5c\0\0\0\0\0`, // Numpad 4
	101: `\0\0\x5d\0\0\0\0\0`, // Numpad 5
	102: `\0\0\x5e\0\0\0\0\0`, // Numpad 6
	103: `\0\0\x5f\0\0\0\0\0`, // Numpad 7
	104: `\0\0\x60\0\0\0\0\0`, // Numpad 8
	105: `\0\0\x61\0\0\0\0\0`, // Numpad 9
	112: `\0\0\x3b\0\0\0\0\0`, // F1
	113: `\0\0\x3c\0\0\0\0\0`, // F2
	114: `\0\0\x3d\0\0\0\0\0`, // F3
	115: `\0\0\x3e\0\0\0\0\0`, // F4
	116: `\0\0\x3f\0\0\0\0\0`, // F5
	117: `\0\0\x40\0\0\0\0\0`, // F6
	118: `\0\0\x41\0\0\0\0\0`, // F7
	119: `\0\0\x42\0\0\0\0\0`, // F8
	120: `\0\0\x43\0\0\0\0\0`, // F9
	121: `\0\0\x44\0\0\0\0\0`, // F10
	122: `\0\0\x45\0\0\0\0\0`, // F11
	123: `\0\0\x46\0\0\0\0\0`, // F12
	124: `\0\0\x68\0\0\0\0\0`, // F13
	125: `\0\0\x69\0\0\0\0\0`, // F14
	126: `\0\0\x6a\0\0\0\0\0`, // F15
	127: `\0\0\x6b\0\0\0\0\0`, // F16
	128: `\0\0\x6c\0\0\0\0\0`, // F17
	129: `\0\0\x6d\0\0\0\0\0`, // F18
	130: `\0\0\x6e\0\0\0\0\0`, // F19
	131: `\0\0\x6f\0\0\0\0\0`, // F20
	132: `\0\0\x70\0\0\0\0\0`, // F21
	133: `\0\0\x71\0\0\0\0\0`, // F22
	134: `\0\0\x72\0\0\0\0\0`, // F23
	144: `\0\0\x53\0\0\0\0\0`, // Num Lock
	145: `\0\0\x47\0\0\0\0\0`, // Scroll Lock
	161: `\0\0\x1e\0\0\0\0\0`, // !
	163: `\0\0\x32\0\0\0\0\0`, // Hash
	173: `\0\0\x2d\0\0\0\0\0`, // Minus
	179: `\0\0\xe8\0\0\0\0\0`, // Media play/pause
	168: `\0\0\xfa\0\0\0\0\0`, // Refresh
	186: `\0\0\x33\0\0\0\0\0`, // Semicolon
	187: `\0\0\x2e\0\0\0\0\0`, // Equal sign
	188: `\0\0\x36\0\0\0\0\0`, // Comma
	189: `\0\0\x2d\0\0\0\0\0`, // Minus sign
	190: `\0\0\x37\0\0\0\0\0`, // Period
	191: `\0\0\x38\0\0\0\0\0`, // Forward slash
	192: `\0\0\x35\0\0\0\0\0`, // Accent grave
	219: `\0\0\x2f\0\0\0\0\0`, // Left bracket ([, {])
	220: `\0\0\x31\0\0\0\0\0`, // Back slash
	221: `\0\0\x30\0\0\0\0\0`, // Right bracket (], })
	222: `\0\0\x34\0\0\0\0\0`, // Single quote
	223: `\0\0\x35\0\0\0\0\0`, // Accent grave (`)
}

// Defining the read and write buffer size for the upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//Translates char/key into HID compatible code
func translationLayer(c rune) string {
	// for example: letter 'a' --> `\0\0\x4\0\0\0\0\0`
	hidCode := keycodeMap[int(c)]
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
