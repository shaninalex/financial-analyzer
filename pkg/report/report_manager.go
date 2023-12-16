package report

import "github.com/rabbitmq/amqp091-go"

type ReportManager struct {
	connection *amqp091.Connection
	channel *amqp091.Channel
}

func InitReportManager(connection *amqp091.Connection, channel *amqp091.Channel) (*ReportManager, error) {

	return &ReportManager{
		connection: connection,
		channel: channel,
	}, nil
}

func (rm *ReportManager) ConsumeMessages() {
}

