package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

/*
Dec		Hex				CharDescription
0:		0x00, //		NUL	Null
1:		0x01, //		SOH	Start of Header
2:		0x02, //		STX	Start of Text
3:		0x03, //		ETX	End of Text
4:		0x04, //		EOT	End of Transmission
5:		0x05, //		ENQ	Enquiry
6:		0x06, //		ACK	Acknowledge
7:		0x07, //		BEL	Bell
8:		0x08, //		BS	Backspace
9:		0x09, //		HT	Horizontal Tab
10:		0x0A, //		LF	Line Feed
11:		0x0B, //		VT	Vertical Tab
12:		0x0C, //		FF	Form Feed
13:		0x0D, //		CR	Carriage Return
14:		0x0E, //		SO	Shift Out
15:		0x0F, //		SI	Shift In
16:		0x10, //		DLE	Data Link Escape
17:		0x11, //		DC1	Device Control 1
18:		0x12, //		DC2	Device Control 2
19:		0x13, //		DC3	Device Control 3
20:		0x14, //		DC4	Device Control 4
21:		0x15, //		NAK	Negative Acknowledge
22:		0x16, //		SYN	Synchronize
23:		0x17, //		ETB	End of Transmission Block
24:		0x18, //		CAN	Cancel
25:		0x19, //		EM	End of Medium
26:		0x1A, //		SUB	Substitute
27:		0x1B, //		ESC	Escape
28:		0x1C, //		FS	File Separator
29:		0x1D, //		GS	Group Separator
30:		0x1E, //		RS	Record Separator
31:		0x1F, //		US	Unit Separator
32:		0x20, //		space	Space
33:		0x21, //		!	Exclamation mark
34:		0x22, //		"	Double quote
35:		0x23, //		#	Number
36:		0x24, //		$	Dollar sign
37:		0x25, //		%	Percent
38:		0x26, //		&	Ampersand
39:		0x27, //		'	Single quote
40:		0x28, //		(	Left parenthesis
41:		0x29, //		)	Right parenthesis
42:		0x2A, //		*	Asterisk
43:		0x2B, //		+	Plus
44:		0x2C, //		,	Comma
45:		0x2D, //		-	Minus
46:		0x2E, //		.	Period
47:		0x2F, //		/	Slash
48:		0x30, //		0	Zero
49:		0x31, //		1	One
50:		0x32, //		2	Two
51:		0x33, //		3	Three
52:		0x34, //		4	Four
53:		0x35, //		5	Five
54:		0x36, //		6	Six
55:		0x37, //		7	Seven
56:		0x38, //		8	Eight
57:		0x39, //		9	Nine
58:		0x3A, //		:	Colon
59:		0x3B, //		;	Semicolon
60:		0x3C, //		<	Less than
61:		0x3D, //		=	Equality sign
62:		0x3E, //		>	Greater than
63:		0x3F, //		?	Question mark
64:		0x40, //		@	At sign
65:		0x41, //		A	Capital A
66:		0x42, //		B	Capital B
67:		0x43, //		C	Capital C
68:		0x44, //		D	Capital D
69:		0x45, //		E	Capital E
70:		0x46, //		F	Capital F
71:		0x47, //		G	Capital G
72:		0x48, //		H	Capital H
73:		0x49, //		I	Capital I
74:		0x4A, //		J	Capital J
75:		0x4B, //		K	Capital K
76:		0x4C, //		L	Capital L
77:		0x4D, //		M	Capital M
78:		0x4E, //		N	Capital N
79:		0x4F, //		O	Capital O
80:		0x50, //		P	Capital P
81:		0x51, //		Q	Capital Q
82:		0x52, //		R	Capital R
83:		0x53, //		S	Capital S
84:		0x54, //		T	Capital T
85:		0x55, //		U	Capital U
86:		0x56, //		V	Capital V
87:		0x57, //		W	Capital W
88:		0x58, //		X	Capital X
89:		0x59, //		Y	Capital Y
90:		0x5A, //		Z	Capital Z
91:		0x5B, //		[	Left square bracket
92:		0x5C, //		\	Backslash
93:		0x5D, //		]	Right square bracket
94:		0x5E, //		^	Caret / circumflex
95:		0x5F, //		_	Underscore
96:		0x60, //		`	Grave / accent
97:		0x61, //		a	Small a
98:		0x62, //		b	Small b
99:		0x63, //		c	Small c
100:	0x64, //	 	d	Small d
101:	0x65, //	 	e	Small e
102:	0x66, //	 	f	Small f
103:	0x67, //	 	g	Small g
104:	0x68, //	 	h	Small h
105:	0x69, //	 	i	Small i
106:	0x6A, //	 	j	Small j
107:	0x6B, //	 	k	Small k
108:	0x6C, //	 	l	Small l
109:	0x6D, //	 	m	Small m
110:	0x6E, //	 	n	Small n
111:	0x6F, //	 	o	Small o
112:	0x70, //	 	p	Small p
113:	0x71, //	 	q	Small q
114:	0x72, //	 	r	Small r
115:	0x73, //	 	s	Small s
116:	0x74, //	 	t	Small t
117:	0x75, //	 	u	Small u
118:	0x76, //	 	v	Small v
119:	0x77, //	 	w	Small w
120:	0x78, //	 	x	Small x
121:	0x79, //	 	y	Small y
122:	0x7A, //	 	z	Small z
123:	0x7B, //	 	{	Left curly bracket
124:	0x7C, //	 	|	Vertical bar
125:	0x7D, //	 	}	Right curly bracket
126:	0x7E, //	 	~	Tilde
127:	0x7F, //	 	DEL	Delete
*/
var keycodeMap = map[int]byte{
	3:   0x48, // Pause / Break
	8:   0x2a, // Backspace / Delete
	9:   0x2b, // Tab
	12:  0x53, // Clear
	13:  0x28, // Enter
	16:  0xe1, // Shift (Left)
	17:  0xe0, // Ctrl (left)
	18:  0xe1, // Alt (left)
	19:  0x48, // Pause / Break
	20:  0x39, // Caps Lock
	21:  0x90, // Hangeul
	25:  0x91, // Hanja
	27:  0x29, // Escape
	32:  0x2c, // Spacebar
	33:  0x4b, // Page Up
	34:  0x4e, // Page Down
	35:  0x4d, // End
	36:  0x4a, // Home
	37:  0x50, // Left Arrow
	38:  0x52, // Up Arrow
	39:  0x4f, // Right Arrow
	40:  0x51, // Down Arrow
	41:  0x77, // Select
	43:  0x74, // Execute
	44:  0x46, // Print Screen
	45:  0x49, // Insert
	46:  0x4c, // Delete
	47:  0x75, // Help
	48:  0x27, // 0
	49:  0x1e, // 1
	50:  0x1f, // 2
	51:  0x20, // 3
	52:  0x21, // 4
	53:  0x22, // 5
	54:  0x23, // 6
	55:  0x24, // 7
	56:  0x25, // 8
	57:  0x26, // 9
	59:  0x33, // Semicolon
	60:  0xc5, // <
	61:  0x2e, // Equal sign
	65:  0x04, // a
	66:  0x05, // b
	67:  0x06, // c
	68:  0x07, // d
	69:  0x08, // e
	70:  0x09, // f
	71:  0x0a, // g
	72:  0x0b, // h
	73:  0x0c, // i
	74:  0x0d, // j
	75:  0x0e, // k
	76:  0x0f, // l
	77:  0x10, // m
	78:  0x11, // n
	79:  0x12, // o
	80:  0x13, // p
	81:  0x14, // q
	82:  0x15, // r
	83:  0x16, // s
	84:  0x17, // t
	85:  0x18, // u
	86:  0x19, // v
	87:  0x1a, // w
	88:  0x1b, // x
	89:  0x1c, // y
	90:  0x1d, // z
	91:  0xe3, // Windows key / Meta Key (Left)
	96:  0x62, // Numpad 0
	97:  0x59, // Numpad 1
	98:  0x5a, // Numpad 2
	99:  0x5b, // Numpad 3
	100: 0x5c, // Numpad 4
	101: 0x5d, // Numpad 5
	102: 0x5e, // Numpad 6
	103: 0x5f, // Numpad 7
	104: 0x60, // Numpad 8
	105: 0x61, // Numpad 9
	112: 0x3b, // F1
	113: 0x3c, // F2
	114: 0x3d, // F3
	115: 0x3e, // F4
	116: 0x3f, // F5
	117: 0x40, // F6
	118: 0x41, // F7
	119: 0x42, // F8
	120: 0x43, // F9
	121: 0x44, // F10
	122: 0x45, // F11
	123: 0x46, // F12
	124: 0x68, // F13
	125: 0x69, // F14
	126: 0x6a, // F15
	127: 0x6b, // F16
	128: 0x6c, // F17
	129: 0x6d, // F18
	130: 0x6e, // F19
	131: 0x6f, // F20
	132: 0x70, // F21
	133: 0x71, // F22
	134: 0x72, // F23
	144: 0x53, // Num Lock
	145: 0x47, // Scroll Lock
	161: 0x1e, // !
	163: 0x32, // Hash
	173: 0x2d, // Minus
	179: 0xe8, // Media play/pause
	168: 0xfa, // Refresh
	186: 0x33, // Semicolon
	187: 0x2e, // Equal sign
	188: 0x36, // Comma
	189: 0x2d, // Minus sign
	190: 0x37, // Period
	191: 0x38, // Forward slash
	192: 0x35, // Accent grave
	219: 0x2f, // Left bracket ([, {])
	220: 0x31, // Back slash
	221: 0x30, // Right bracket (], })
	222: 0x34, // Single quote
	223: 0x35, // Accent grave (`)
}

// Defining the read and write buffer size for the upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Translates char/key into HID compatible code
func translationLayer(c rune) byte {
	hidCode := keycodeMap[int(c)]
	return hidCode
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

// TODO: Not all keycodes are correct, why?
func executePayload(payloadString string) bool {

	//run through each character/rune in the payload string
	for _, ch := range payloadString {
		key := translationLayer(ch)
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
