package datasource

import (
	"log"

	rabbitmq "github.com/shaninalex/financial-analyzer/internal/rabbitmq"
)

func Init(GURU_API_KEY, ALPH_API_KEY, RABBITMQ_URL string) {
	connection, err := rabbitmq.ConnectToRabbitMQ(RABBITMQ_URL)
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

	api, err := InitializeApplication(GURU_API_KEY, ALPH_API_KEY, connection, channel, RABBITMQ_URL)
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
