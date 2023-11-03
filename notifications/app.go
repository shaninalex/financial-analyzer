package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

type App struct {
	Conn   *amqp.Connection
	Router *gin.Engine
}

func InitApplication() (*App, error) {
	conn, err := connectToRabbitMQ(RABBITMQ_URL)
	failOnError(err)
	app := &App{
		Conn:   conn,
		Router: gin.Default(),
	}

	app.SetupRoutes()
	return app, nil
}

func (app *App) SetupRoutes() {
	app.Router.GET("")
	api := app.Router.Group("/api/v2/notifications")
	api.Use(UserIDMiddleware())
	{
		api.Get("/", app.GetNotifications)
		api.Patch("/", app.GetNotifications)
		api.Delete("/", app.GetNotifications)
	}
}

func (app *App) Run(port int) {
	addr := fmt.Sprintf(":%d", port)
	log.Printf("App is running on %s", addr)
	app.Router.Run(addr)
}
