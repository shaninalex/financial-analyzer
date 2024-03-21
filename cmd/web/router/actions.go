package router

import (
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (c *Client) CreateReport(message []byte) error {
	return c.MQChannel.PublishWithContext(c.Context,
		// "ex.datasource", // exchange
		"ex.report",  // exchange
		"new_report", // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
			Headers: amqp.Table{
				"user_id":    c.ID,
				"client_id":  c.ClientId,
				"request_id": uuid.NewString(),
			},
		},
	)
}
