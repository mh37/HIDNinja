package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/websocket"
)

// Defining the read and write buffer size for the upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func translationLayer(payloadString string) string {
	// TBD

	return ""
}

func executePayload(payloadString string) bool {

	//For testing purposes at the moment
	app := "echo"
	arg0 := "-ne"
	arg1 := "\\0\\0\\x4\\0\\0\\0\\0\\0"
	arg2 := ">"
	arg3 := "/dev/hidg0"

	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	// Print the output
	fmt.Println(string(stdout))

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
		executePayload("")
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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
