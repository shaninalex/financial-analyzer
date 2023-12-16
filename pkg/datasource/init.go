package datasource

import (
	"github.com/rabbitmq/amqp091-go"
	rabbitmq "github.com/shaninalex/financial-analyzer/internal/rabbitmq"
)

func Init(GURU_API_KEY, ALPH_API_KEY, RABBITMQ_URL string) (*amqp091.Connection, *amqp091.Channel, error) {
	connection, err := rabbitmq.ConnectToRabbitMQ(RABBITMQ_URL)
	if err != nil {
		return nil, nil, err
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, nil, err
	}

	if err != nil {
		return nil, nil, err
	}

	api, err := InitializeApplication(GURU_API_KEY, ALPH_API_KEY, connection, channel, RABBITMQ_URL)
	if err != nil {
		return nil, nil, err
	}

	api.ConsumeRabbitMessages()

	return connection, channel, nil
}
