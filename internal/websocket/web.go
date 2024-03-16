package websocket

import (
	"fmt"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/shaninalex/financial-analyzer/internal/router"
)

func Websocket(port int, connection *amqp.Connection, channel *amqp.Channel) {

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("X-User")
		if userID == "" {
			http.Error(w, "user id is empty", http.StatusUnauthorized)
			return
		}
		router.ServeWebsocket(userID, connection, channel, w, r)
	})

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server is running on http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}
