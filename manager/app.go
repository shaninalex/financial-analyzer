package main

import (
	"context"
	"fmt"
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
			URL: "http://127.0.0.1:4434", // Kratos Admin API
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
	connection, _ := upgrader.Upgrade(w, r, nil)
	defer connection.Close()

	if !app.validateSession(r) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: get user id
	// TODO: send message about user presence ( online )
	log.Println(r.Cookies())
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

// ory_kratos_session
func (app *App) validateSession(r *http.Request) bool {
	cookie, err := r.Cookie("ory_kratos_session")
	if err != nil {
		log.Println(err)
		return false
	}
	if cookie == nil {
		log.Println("no session found in cookie")
		return false
	}
	session, _, err := app.OryClient.FrontendApi.ToSession(context.Background()).Cookie(cookie.String()).Execute()
	if err != nil {
		log.Println(err)
		return false
	}
	if !*session.Active {
		log.Println("Session is not active")
		return false
	}

	// go app.UserIsConnected(session.GetId())

	return true
}

func (app *App) UserIsConnected(userId string) {

	log.Println(fmt.Sprintf("%s is connected", userId))
}

func (app *App) UserIsDisconnected(userId string) {
	log.Println(fmt.Sprintf("%s is disconnected", userId))
}

func (app *App) messageHandler(message []byte) {
	log.Println(string(message))
}
