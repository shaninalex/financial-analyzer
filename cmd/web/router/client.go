// Package router provides functionality for routing messages to different
// RabbitMQ queues.
//
// The router module is responsible for receiving messages and directing them
// to the appropriate RabbitMQ queue based on predefined routing rules.
// It acts as an intermediary between message producers and consumers,
// ensuring efficient message delivery.

// Usage:
// ```go
// client, err := InitClient(user_id, connection, channel, wsconnection)
//
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	````
//
// Arguments:
// user_id - user Id, appended to request through Oathkeeper
// connection - *amqp.Connection
// channel - *amqp.Channel
// wsconnection - pointer to websocket connection object
package router

import (
	"context"
	"fmt"
	"log"

	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/shaninalex/financial-analyzer/internal/typedefs"
)

type Client struct {
	ID           string
	ClientId     string
	MQConnection *amqp.Connection
	MQChannel    *amqp.Channel
	MQQueue      *amqp.Queue
	WSConnection *websocket.Conn
	Context      context.Context
	Account      typedefs.IAccount
}

func InitClient(user_id string, mq *amqp.Connection, ch *amqp.Channel, ws *websocket.Conn) (*Client, error) {
	acc, err := typedefs.InitAccount(user_id, mq, ch)
	if err != nil {
		return nil, err
	}

	client := &Client{
		ID:           user_id,
		ClientId:     uuid.New().String(),
		MQConnection: mq,
		MQChannel:    ch,
		WSConnection: ws,
		Context:      context.TODO(),
		Account:      acc,
	}

	err = ch.ExchangeDeclare(
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

	client.MQChannel.QueueBind(
		q.Name,                                 // name
		fmt.Sprintf("client.%s.all", user_id),  // routing key
		fmt.Sprintf("ex.client.%s", client.ID), // exchange name
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
			i, err := c.MQChannel.QueueDelete(c.MQQueue.Name, false, false, true)
			log.Printf("ws connection closed with status: %d\n", i)
			if err != nil {
				log.Printf("Unable to delete queue: %v\n", err)
			}
			break
		}

		var action typedefs.Action
		if err := json.Unmarshal(message, &action); err != nil {
			log.Printf("error: %v", err)
			break
		}

		switch action.Action {
		case typedefs.ActionTypeReport:
			able, msg := c.Account.AbleToReport()
			if !able {
				c.RequestDenied(*msg)
				break
			}
			err := c.CreateReport(message)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (c *Client) RequestDenied(msg string) {
	message := map[string]interface{}{
		"action": "notification",
		"payload": map[string]interface{}{
			"level":   "error",
			"message": msg,
		},
	}
	body, err := json.Marshal(message)
	if err != nil {
		// TODO: need global logger
		log.Println(err)
	}
	c.WSConnection.WriteMessage(1, body)
}
