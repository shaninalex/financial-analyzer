package manager

import (
	"github.com/rabbitmq/amqp091-go"
)

func (rm *ReportManager) InitGatheringInformation(
	message amqp091.Delivery,
) error {
	return rm.channel.PublishWithContext(rm.ctx,
		"ex.datasource", // exchange
		"new_report",    // routing key
		false,           // mandatory
		false,           // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        message.Body,
			Headers: amqp091.Table{
				"user_id":    message.Headers["user_id"].(string),
				"client_id":  message.Headers["client_id"].(string),
				"request_id": message.Headers["request_id"].(string),
			},
		},
	)
}
