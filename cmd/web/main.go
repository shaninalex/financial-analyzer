package main

import (
	"log"
	"os"
	"strconv"

	"github.com/shaninalex/financial-analyzer/cmd/web/websocket"
	"github.com/shaninalex/financial-analyzer/internal/rabbitmq"
	"github.com/shaninalex/financial-analyzer/internal/report"
)

var (
	DEBUG          = os.Getenv("DEBUG") // "0" or "1"
	GURU_API_KEY   = os.Getenv("GURU_API_KEY")
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
	go websocket.Websocket(port, connection, channel)

	go func() {
		// initialize report manager
		log.Println("initialize report manager")
		err = report.InitReportModule(connection, channel)
		if err != nil {
			panic(err)
		}
	}()

	defer func() {
		log.Println("Close rabbitmq connections")
		connection.Close()
		channel.Close()
	}()

}
