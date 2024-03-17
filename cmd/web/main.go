package main

import (
	"log"
	"os"
	"strconv"

	"github.com/shaninalex/financial-analyzer/cmd/web/websocket"
	"github.com/shaninalex/financial-analyzer/internal/rabbitmq"
)

var (
	RABBITMQ_URL   = os.Getenv("RABBITMQ_URL")
	WEBSOCKET_PORT = os.Getenv("WEBSOCKET_PORT")
)

func main() {

	port, err := strconv.Atoi(WEBSOCKET_PORT)
	if err != nil {
		panic(err)
	}

	connection, channel, err := rabbitmq.ConnectToRabbitMQ(RABBITMQ_URL)
	if err != nil {
		panic(err)
	}

	// initialize websocket connection
	websocket.Websocket(port, connection, channel)

	defer func() {
		log.Println("Close rabbitmq connections")
		connection.Close()
		channel.Close()
	}()

}
