package main

import (
	"log"
	"os"
)

var (
	DEBUG        = os.Getenv("DEBUG") // "0" or "1"
	GURU_API_KEY = os.Getenv("GURU_API_KEY")
	ALPH_API_KEY = os.Getenv("ALPH_API_KEY")
	RABBITMQ_URL = os.Getenv("RABBITMQ_URL")
)

func main() {

	connection, err := connectToRabbitMQ(RABBITMQ_URL)
	if err != nil {
		panic(err)
	}

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	api, err := InitializeApplication(GURU_API_KEY, ALPH_API_KEY, connection, channel)
	if err != nil {
		panic(err)
	}

	api.ConsumeRabbitMessages()

	defer func() {
		log.Println("close channel and connection")
		connection.Close()
		channel.Close()
	}()
}
