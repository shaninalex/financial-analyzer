package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectToRabbitMQ(connectionString string) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	return conn, ch, nil
}

func ReconnectToRabbitMQ(connection *amqp.Connection, channel *amqp.Channel, connectionString string) (*amqp.Connection, *amqp.Channel, error) {
	if connection == nil || connection.IsClosed() || channel == nil || channel.IsClosed() {
		log.Println("Reconnect to RabbitMQ")

		newConnection, newChannel, err := ConnectToRabbitMQ(connectionString)
		if err != nil {
			return nil, nil, err
		}

		log.Println("Recreate RabbitMQ channel")
		return newConnection, newChannel, nil
	}

	return connection, channel, nil
}
