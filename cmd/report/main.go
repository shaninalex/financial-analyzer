package main

import (
	"os"

	"github.com/shaninalex/financial-analyzer/internal/rabbitmq"
)

var (
	RABBITMQ_URL = os.Getenv("RABBITMQ_URL")
)

func main() {

	connection, channel, err := rabbitmq.ConnectToRabbitMQ(RABBITMQ_URL)
	if err != nil {
		panic(err)
	}

	manager, err := InitReportManager(connection, channel)
	if err != nil {
		panic(err)
	}

	manager.ConsumeMessages()

	defer func() {
		connection.Close()
		channel.Close()
	}()

}
