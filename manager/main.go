package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// connect only through trusted servers
		return true
	},
}

var (
	KRATOS_URL = os.Getenv("KRATOS_URL")
)

func main() {
	app := InitializeApplication(KRATOS_URL)
	http.HandleFunc("/", app.HandleConnection)
	log.Println("Server is running on port 8003")
	http.ListenAndServe(":8003", nil)
}
