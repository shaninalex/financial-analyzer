package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (c *Client) triggerSearch(message []byte) {
	log.Printf("Search for %s\n", string(message))
	err := c.MQChannel.PublishWithContext(c.Context,
		"ex.datasource", // exchange
		"q.datasource",  // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
			Headers:     amqp.Table{"client_id": c.Id},
		},
	)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("message published")
	}
}
