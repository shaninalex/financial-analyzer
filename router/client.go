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
	Hub          *Hub
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

func InitClient(hub *Hub, connection *websocket.Conn, user_id string) (*Client, error) {

	wsclient := &Client{
		Hub:      hub,
		Conn:     connection,
		Send:     make(chan []byte),
		Id:       user_id,
		Context:  context.TODO(),
		CSearch:  make(chan []byte),
		CProcess: make(chan []byte),
		CReport:  make(chan []byte),
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
		fmt.Sprintf("q_client-%s", wsclient.Id), // name
		false,                                   // durable
		false,                                   // delete when unused
		false,                                   // exclusive
		false,                                   // no-wait
		nil,                                     // arguments
	)
	if err != nil {
		return nil, err
	}

	wsclient.MQConnection = mq_connection
	wsclient.MQChannel = ch
	wsclient.MQQueue = &q
	return wsclient, nil
}

func (c *Client) ConsumeRMQMessages() {
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
			c.Send <- d.Body
			log.Printf("Received a message: %s", d.Body)
		}
	}()
}

func (c *Client) ListenChannels() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		select {
		case message := <-c.CSearch:
			c.triggerSearch(message)
		case message := <-c.CProcess:
			log.Printf("Init Process for %s\n", string(message))
		case message := <-c.CReport:
			log.Printf("Report generating for %s\n", string(message))

		// send to frontend
		case message := <-c.Send:
			c.Conn.WriteMessage(1, message)
		}
	}
}

func (c *Client) ReadMessages() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

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

		switch action.Action {
		case TickerActionTypeSearch:
			c.CSearch <- []byte(action.Ticker)
		case TickerActionTypeProcess:
			c.CProcess <- []byte(action.Ticker)
		case TickerActionTypeReport:
			c.CReport <- []byte(action.Ticker)
		}
	}
}
