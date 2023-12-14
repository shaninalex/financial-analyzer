package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

type App struct {
	Context      context.Context
	Datasource   *Datasource
	Router       *gin.Engine
	MQConnection *amqp.Connection
	MQChannel    *amqp.Channel
}

func InitializeApplication(gfAppKey, alphAppKey string, connection *amqp.Connection, channel *amqp.Channel) (*App, error) {
	app := &App{
		Context:      context.TODO(),
		Datasource:   InitializeDatasource(gfAppKey, alphAppKey, true),
		Router:       gin.Default(),
		MQConnection: connection,
		MQChannel:    channel,
	}

	return app, nil
}

func (app *App) ConsumeRabbitMessages() {
	msgs, err := app.MQChannel.Consume(
		"q.datasource",       // queue
		"datasource_service", // consumer
		true,                 // auto-ack
		false,                // exclusive
		false,                // no-local
		false,                // no-wait
		nil,                  // args
	)
	if err != nil {
		log.Println("Failed to register a consumer:")
		log.Println(err)
	}

	for d := range msgs {
		log.Printf("Message Body: %s", d.Body)

		// {
		// 	"ticker": "IBM",
		// 	"action": "new_report"
		// }
		if d.RoutingKey == "new_report" {
			var action TickerAction
			err := json.Unmarshal(d.Body, &action)
			if err != nil {
				log.Printf("Unable to unmarshal action: %s. Error: %v", d.Body, err)
				continue
			}
			log.Println(string(d.Body))
			log.Println(d.Headers)
			app.GatheringInformation(
				action,
				d.Headers["user_id"].(string),
				d.Headers["client_id"].(string),
			)
		}
	}

}

func (app *App) PublishResults(message any, user_id string, client_id string, message_type string, ticker string) {
	comp_message, _ := json.Marshal(map[string]interface{}{
		"action": "results",
		"ticker": ticker,
		"type":   message_type,
		"data":   message,
	})

	routing_key := fmt.Sprintf("client.%s__dev.%s", user_id, client_id)
	// routing_key := "client.b5aa9b5e-a1b8-43f3-9fe1-9e8ab85d8025__dev.6f4314b5-1b08-4734-86f8-9b467dd1c9ec"
	fmt.Printf("Routing key: %s", routing_key)

	err := app.reconnectToRabbitMQ()
	if err != nil {
		log.Printf("Error publishing message: %v", err)
	}

	err = app.MQChannel.PublishWithContext(app.Context,
		fmt.Sprintf("ex.client.%s", user_id), // exchange
		routing_key,                          // routing key
		false,                                // mandatory
		false,                                // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        comp_message,
		},
	)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("message \"%s\" for \"%s\" is published", message_type, ticker)
	}
}

func (app *App) reconnectToRabbitMQ() error {
	if app.MQConnection.IsClosed() {
		log.Println("Reconnect to rabbitmq")
		connection, err := connectToRabbitMQ(RABBITMQ_URL)
		if err != nil {
			return err
		}
		app.MQConnection = connection
	}

	if app.MQChannel.IsClosed() {
		log.Println("Recreate rabbitmq channel")
		channel, err := app.MQConnection.Channel()
		if err != nil {
			return err
		}

		app.MQChannel = channel
	}

	return nil
}

func (app *App) GatheringInformation(action TickerAction, user_id string, client_id string) {
	overview, err := app.Datasource.Alphavantage.Overview(action.Ticker)
	if err != nil {
		log.Printf("Unable to get Alphavantage.Overview for \"%s\". Error: %v", action.Ticker, err)
	} else {
		app.PublishResults(overview, user_id, client_id, "alph_overview", action.Ticker)
	}

	cashflow, err := app.Datasource.Alphavantage.CashFlow(action.Ticker)
	if err != nil {
		log.Printf("Unable to get Alphavantage.CashFlow for \"%s\". Error: %v", action.Ticker, err)
	} else {
		app.PublishResults(cashflow, user_id, client_id, "alph_cashflow", action.Ticker)
	}

	earnings, err := app.Datasource.Alphavantage.Earnings(action.Ticker)
	if err != nil {
		log.Printf("Unable to get Alphavantage.Earnings for \"%s\". Error: %v", action.Ticker, err)
	} else {
		app.PublishResults(earnings, user_id, client_id, "alph_earnings", action.Ticker)
	}
}
