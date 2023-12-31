package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/shaninalex/financial-analyzer/internal/rabbitmq"
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

func ServeWebsocket(user_id string, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	mq, err := rabbitmq.ConnectToRabbitMQ(RABBITMQ_URL)
	if err != nil {
		log.Println(err)
		return
	}

	ch, err := mq.Channel()
	if err != nil {
		log.Println(err)
		return
	}

	client, err := InitClient(user_id, mq, ch, conn)
	if err != nil {
		log.Println(err)
		return
	}

	client.ConsumeMQ()
	client.ConsumeFrontend()

	defer func() {
		log.Println("close channel and connection")
		ch.Close()
		mq.Close()
	}()
}
