package main

import (
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	rabbitmq "github.com/shaninalex/financial-analyzer/internal/rabbitmq"
)

var (
	RABBITMQ_URL = os.Getenv("RABBITMQ_URL")
	APP_PORT     = os.Getenv("APP_PORT")
)

// func main() {
// 	app, err := InitApplication()
// 	if err != nil {
// 		panic(err)
// 	}

// 	port, err := strconv.Atoi(APP_PORT)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// app.Run(port)
// }

type App struct {
	Conn *amqp.Connection
}

func InitApplication() (*App, error) {
	conn, err := rabbitmq.ConnectToRabbitMQ(RABBITMQ_URL)
	failOnError(err)

	app := &App{
		Conn: conn,
	}

	app.SetupRoutes()
	return app, nil
}

func (app *App) SetupRoutes() {
	// app.Router.GET("")
	// api := app.Router.Group("/api/v2/notifications")
	// api.Use(UserIDMiddleware())
	// {
	// 	api.GET("/", app.GetNotifications)
	// 	api.PATCH("/", app.PatchNotifications)
	// 	api.DELETE("/", app.DeleteNotifications)
	// }
}

// func (app *App) Run(port int) {
// 	addr := fmt.Sprintf(":%d", port)
// 	log.Printf("App is running on %s", addr)
// 	app.Router.Run(addr)
// }
