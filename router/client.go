package main

import (
	"context"
	"fmt"
	"log"

	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Client struct {
	Conn         *websocket.Conn
	Send         chan []byte
	Id           string
	Context      context.Context
	CSearch      chan []byte
	CProcess     chan []byte
	CReport      chan []byte
	MQConnection *amqp.Connection
	MQChannel    *amqp.Channel
	MQQueue      *amqp.Queue
}

func InitClient(connection *websocket.Conn, user_id string) (*Client, error) {

	client := &Client{
		Conn:    connection,
		Id:      user_id,
		Context: context.TODO(),
	}

	mq_connection, err := connectToRabbitMQ(RABBITMQ_URL)
	if err != nil {
		return nil, err
	}

	ch, err := mq_connection.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		fmt.Sprintf("q_client-%s", client.Id), // name
		true,                                  // durable
		false,                                 // delete when unused
		false,                                 // exclusive
		false,                                 // no-wait
		nil,                                   // arguments
	)
	if err != nil {
		return nil, err
	}
	routing_key := fmt.Sprintf("client_%s", client.Id)
	ch.QueueBind(q.Name, routing_key, "ex.datasource", false, nil)
	ch.QueueBind(q.Name, routing_key, "ex.email", false, nil)
	ch.QueueBind(q.Name, routing_key, "ex.notifications", false, nil)
	ch.QueueBind(q.Name, routing_key, "ex.report", false, nil)
	ch.QueueBind(q.Name, "", "ex.global.notifications", false, nil)

	client.MQConnection = mq_connection
	client.MQChannel = ch
	client.MQQueue = &q
	return client, nil
}

func (c *Client) ConsumeRMQMessages() {
	msgs, err := c.MQChannel.Consume(
		c.MQQueue.Name, // queue
		"",             // consumer
		false,          // auto-ack
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
			c.Conn.WriteMessage(1, d.Body)
		}
	}()
}

func (c *Client) ReadMessages() {
	for {
		_, message, err := c.Conn.ReadMessage()
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

		if action.Action == TickerActionTypeSearch {
			c.triggerSearch(message)
		}
	}
}

func (c *Client) Disconnect() {
	if c.MQConnection != nil {
		err := c.MQConnection.Close()
		if err != nil {
			log.Printf("Error closing RabbitMQ connection: %v", err)
		}
		fmt.Println("RabbitMQ connection closed")
		close(c.Send)
		// close(c.CProcess)
		// close(c.CReport)
	}
}
