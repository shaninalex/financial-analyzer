package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

	mq, err := connectToRabbitMQ(RABBITMQ_URL)
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

	router := gin.Default()
	router.GET("/ws", func(c *gin.Context) {
		user_id := c.Request.Header.Get("X-User")
		if user_id == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "user id is empty"})
			return
		}
		ServeWebsocket(user_id, c.Writer, c.Request)
	})

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		panic(err)
	}
	router.Run(fmt.Sprintf(":%d", port))
}
