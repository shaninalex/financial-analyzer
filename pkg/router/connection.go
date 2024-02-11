package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	APP_PORT     = os.Getenv("APP_PORT")
	RABBITMQ_URL = os.Getenv("RABBITMQ_URL")
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// log.Println(r.Host)
		// TODO: accept origin only from front
		return true
	},
}

func ServeWebsocket(user_id string, connection *amqp.Connection, channel *amqp.Channel, w http.ResponseWriter, r *http.Request) {
	wsconnection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer wsconnection.Close()

	client, err := InitClient(user_id, connection, channel, wsconnection)
	if err != nil {
		log.Println(err)
		return
	}

	client.ConsumeMQ()
	client.ConsumeFrontend()
}
