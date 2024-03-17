// Note: Probably also goes into separate service.
package report

import (
	"github.com/rabbitmq/amqp091-go"
)

func InitReportModule(connection *amqp091.Connection, channel *amqp091.Channel) error {

	manager, err := InitReportManager(connection, channel)
	if err != nil {
		return err
	}

	manager.ConsumeMessages()

	return nil
}
