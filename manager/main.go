package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// connect only through trusted servers
		return true
	},
}

func main() {
	http.HandleFunc("/", echo)
	log.Println("Server is running on port 8003")
	http.ListenAndServe(":8003", nil)
}

func echo(w http.ResponseWriter, r *http.Request) {
	connection, _ := upgrader.Upgrade(w, r, nil)
	defer connection.Close()

	// TODO: validate Ory cookie
	// TODO: get user id
	// TODO: send message about user presence ( online )
	log.Println(r.Cookies())
	for {
		mt, message, err := connection.ReadMessage()

		if err != nil || mt == websocket.CloseMessage {
			break
		}

		connection.WriteMessage(websocket.TextMessage, message)
		go messageHandler(message)
	}
}

func messageHandler(message []byte) {
	log.Println(string(message))
}
