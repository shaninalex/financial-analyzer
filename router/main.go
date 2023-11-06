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

	client, err := InitClient(conn, user_id)
	if err != nil {
		log.Println(err)
		return
	}

	go client.ReadMessages()
	go client.ListenChannels()
	go client.ConsumeRMQMessages()

	// close(client.Send)
	// close(client.CSearch)
	// close(client.CProcess)
	// close(client.CReport)
	// defer client.MQConnection.Close()
	// defer client.MQChannel.Close()
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
