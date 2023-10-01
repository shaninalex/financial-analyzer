package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	ory "github.com/ory/kratos-client-go"
)

type App struct {
	OryClient *ory.APIClient
	Messages  chan string
	Context   context.Context
}

func InitializeApplication(kratosUrl string) *App {
	configuration := ory.NewConfiguration()
	configuration.Servers = []ory.ServerConfiguration{
		{
			URL: "http://kratos:4434",
		},
	}
	client := ory.NewAPIClient(configuration)

	return &App{
		OryClient: client,
		Messages:  make(chan string),
		Context:   context.TODO(),
	}
}

func (app *App) HandleConnection(w http.ResponseWriter, r *http.Request) {
	user_id := r.Header.Get("X-User")
	connection, _ := upgrader.Upgrade(w, r, nil)
	defer connection.Close()

	go app.UserIsConnected(user_id)
	for {
		mt, message, err := connection.ReadMessage()

		if err != nil || mt == websocket.CloseMessage {
			// go app.UserIsDisconnected("session.GetId()")
			break
		}

		connection.WriteMessage(websocket.TextMessage, message)
		go app.messageHandler(message)
	}
}

func (app *App) UserIsConnected(userId string) {
	// send message about user presence ( online )
	log.Printf("\n%s is connected\n", userId)
}

// func (app *App) UserIsDisconnected(userId string) {
// 	log.Println(fmt.Sprintf("%s is disconnected", userId))
// }

func (app *App) messageHandler(message []byte) {
	log.Println(string(message))
}
