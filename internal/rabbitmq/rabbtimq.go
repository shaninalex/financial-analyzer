package rabbitmq

import (
	"log"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectToRabbitMQ(connectionString string) (*amqp.Connection, *amqp.Channel, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	for {
		c, err := amqp.Dial(connectionString)
		if err != nil {
			log.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			log.Println("Connected to RabbitMQ")
			connection = c
			break
		}

		if counts > 5 {
			log.Println(err)
			return nil, nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("waiting for rabbitmq...")
		time.Sleep(backOff)
		continue
	}

	ch, err := connection.Channel()
	if err != nil {
		connection.Close()
		return nil, nil, err
	}

	return connection, ch, nil
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
