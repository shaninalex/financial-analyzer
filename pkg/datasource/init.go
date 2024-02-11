package datasource

import (
	"github.com/rabbitmq/amqp091-go"
)

func Init(connection *amqp091.Connection, channel *amqp091.Channel, GURU_API_KEY, ALPH_API_KEY string) error {
	api, err := InitializeApplication(GURU_API_KEY, ALPH_API_KEY, connection, channel)
	if err != nil {
		return err
	}

	api.ConsumeRabbitMessages()

	return nil
}
