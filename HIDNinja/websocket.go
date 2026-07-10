package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

// Initialize a websocket Upgrader and configure its read and write buffer size to 1024 bytes.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		if origin == "" {
			return true
		}
		u, err := url.Parse(origin)
		if err != nil {
			return false
		}
		return u.Host == r.Host
	},
}

// Configure the WebSocket endpoint
func wsEndpoint(w http.ResponseWriter, r *http.Request) {

	// upgrade the connection to a websocket
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}

	log.Println("Client connected successfully")

	reader(conn)
}

// reader function that listens for incoming payloads
func reader(conn *websocket.Conn) {
	defer conn.Close()
	// keep listening for incoming payloads
	for {
		//read incoming payload
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}

		//print received message to console
		log.Printf("%q sent: %q\n", conn.RemoteAddr().String(), string(msg))

		if err = conn.WriteMessage(msgType, msg); err != nil {
			log.Println("Error writing message:", err)
			return
		}

		//execute payload
		executePayload(string(msg))
	}
}

// Defines the routes, such as the WebSocket Endpoint and Homepage
func setupRoutes() {
	http.HandleFunc("/echo", wsEndpoint)
	fs := http.FileServer(http.Dir("../PayloadInterface"))
	http.Handle("/", fs)
}
