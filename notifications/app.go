package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type App struct {
	Conn *amqp.Connection
}

func InitApplication(rabbit_mq_url string) (*App, error) {
	conn, err := connectToRabbitMQ(RABBITMQ_URL)
	failOnError(err)
	app := &App{
		Conn: conn,
	}

	return app, nil
}

func (app *App) Run() {
	log.Println("Notification service is running...")
}
