package report

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type ReportManager struct {
	connection *amqp091.Connection
	channel    *amqp091.Channel
}

func InitReportManager(connection *amqp091.Connection, channel *amqp091.Channel) (*ReportManager, error) {

	return &ReportManager{
		connection: connection,
		channel:    channel,
	}, nil
}

func (rm *ReportManager) ConsumeMessages() {

	messages, err := rm.channel.Consume("q.report", "", true, false, false, false, nil)
	if err != nil {
		log.Println("Failed to register consumer:")
		log.Println(err)
	}

	log.Println("ReportManager: start consume messages...")
	go func() {
		for m := range messages {
			log.Println("ReportManager: get message:")
			log.Println(m.Body)
			log.Println(m.Type)
		}
	}()
}
