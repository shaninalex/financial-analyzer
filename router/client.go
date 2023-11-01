package main

import (
	"bytes"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
	Id   string
}

func InitClient(hub *Hub, connection *websocket.Conn, user_id string) (*Client, error) {

	wsclient := &Client{
		Hub:  hub,
		Conn: connection,
		Send: make(chan []byte),
		Id:   user_id,
	}

	return wsclient, nil
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
		message = bytes.TrimSpace(message)
		c.Send <- message
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
