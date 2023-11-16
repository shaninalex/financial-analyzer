package main

import (
	"context"
	"fmt"
	"log"

	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Client struct {
	ID           string
	ClientId     string
	MQConnection *amqp.Connection
	MQChannel    *amqp.Channel
	MQQueue      *amqp.Queue
	WSConnection *websocket.Conn
	Context      context.Context
}

func InitClient(user_id string, mq *amqp.Connection, ch *amqp.Channel, ws *websocket.Conn) (*Client, error) {
	client := &Client{
		ID:           user_id,
		ClientId:     uuid.New().String(),
		MQConnection: mq,
		MQChannel:    ch,
		WSConnection: ws,
		Context:      context.TODO(),
	}

	err := ch.ExchangeDeclare(
		fmt.Sprintf("ex.client.%s", client.ID), // name
		"direct",                               // type
		true,                                   // durable
		false,                                  // auto-deleted
		false,                                  // internal
		false,                                  // no-wait
		nil,                                    // arguments
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	q, err := client.MQChannel.QueueDeclare(
		"",    // name
		false, // durable
		true,  // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	client.MQQueue = &q
	client.MQChannel.QueueBind(
		q.Name, // name
		fmt.Sprintf("client.%s__dev.%s", user_id, client.ClientId), // routing key
		fmt.Sprintf("ex.client.%s", client.ID),                     // exchange name
		false,
		nil,
	)

	return client, nil
}

func (c *Client) ConsumeMQ() {
	msgs, err := c.MQChannel.Consume(
		c.MQQueue.Name, // queue
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	if err != nil {
		log.Println("Failed to register a consumer:")
		log.Println(err)
	}

	go func() {
		for d := range msgs {
			c.WSConnection.WriteMessage(1, d.Body)
		}
	}()
}

func (c *Client) ConsumeFrontend() {
	for {
		_, message, err := c.WSConnection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var action ITickerAction
		if err := json.Unmarshal(message, &action); err != nil {
			log.Printf("error: %v", err)
			break
		}

		switch action.Action {
		case TickerActionTypeSearch:
			err := c.MQChannel.PublishWithContext(c.Context,
				"ex.datasource", // exchange
				"new_report",    // routing key
				false,           // mandatory
				false,           // immediate
				amqp.Publishing{
					ContentType: "application/json",
					Body:        message,
					Headers: amqp.Table{
						"user_id":   c.ID,
						"client_id": c.ClientId,
					},
				},
			)
			if err != nil {
				log.Println(err)
			} else {
				log.Println("message published")
			}
		}
	}
}
