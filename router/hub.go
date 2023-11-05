package main

type Hub struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
}

func InitHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {

		// when new connection established - save client to Hub
		case client := <-h.Register:
			h.Clients[client] = true

		// when connection closed - remove client and close all client channels
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
				close(client.CSearch)
				close(client.CProcess)
				close(client.CReport)
				// client.MQConnection.Close()
				// client.MQChannel.Close()
			}
		}
	}
}
