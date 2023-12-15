package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

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
		log.Println(r.Host)
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

func main() {

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("X-User")
		if userID == "" {
			http.Error(w, "user id is empty", http.StatusUnauthorized)
			return
		}
		ServeWebsocket(userID, w, r)
	})

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		panic(err)
	}

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server is running on http://localhost%s\n", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}
