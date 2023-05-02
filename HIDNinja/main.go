package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Fatal(err)
		}

		for {

			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../PayloadInterface/index.html")
	})

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
