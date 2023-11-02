package main

import (
	"log"

	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
)

type Client struct {
	Hub      *Hub
	Conn     *websocket.Conn
	Send     chan []byte
	Id       string
	CSearch  chan []byte
	CProcess chan []byte
	CReport  chan []byte
}

func InitClient(hub *Hub, connection *websocket.Conn, user_id string) (*Client, error) {

	wsclient := &Client{
		Hub:      hub,
		Conn:     connection,
		Send:     make(chan []byte),
		Id:       user_id,
		CSearch:  make(chan []byte),
		CProcess: make(chan []byte),
		CReport:  make(chan []byte),
	}

	return wsclient, nil
}

func (c *Client) ListenChannels() {
	for {
		select {
		case message := <-c.Send:
			log.Printf("message: %s", string(message))
		case message := <-c.CSearch:
			log.Printf("Search for %s\n", string(message))
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

		if action.Action == TickerActionTypeSearch {
			c.CSearch <- []byte(action.Ticker)
		}

		// message = bytes.TrimSpace(message)
		// c.Send <- message
	}
}

func (c *Client) WriteMessages() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		select {
		case message := <-c.Send:
			c.Conn.WriteMessage(1, message)
		}
	}
}
