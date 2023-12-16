package datasource

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"

	rabbitmq "github.com/shaninalex/financial-analyzer/internal/rabbitmq"
	"github.com/shaninalex/financial-analyzer/internal/typedefs"
)

type providerMethod struct {
	f        func(ticker string) (*interface{}, error)
	dataType string
}

type App struct {
	Context      context.Context
	Datasource   *Datasource
	MQConnection *amqp.Connection
	MQChannel    *amqp.Channel
	RabbitmqUrl  string
	Methods      []providerMethod
}

func InitializeApplication(gfAppKey, alphAppKey string, connection *amqp.Connection, channel *amqp.Channel, RABBITMQ_URL string) (*App, error) {
	app := &App{
		Context:      context.TODO(),
		Datasource:   InitializeDatasource(gfAppKey, alphAppKey, true),
		MQConnection: connection,
		MQChannel:    channel,
		RabbitmqUrl:  RABBITMQ_URL,
	}

	app.Methods = []providerMethod{
		{f: app.Datasource.Gurufocus.Summary, dataType: string(typedefs.GurufocusRequestSummary)},
		{f: app.Datasource.Gurufocus.Dividends, dataType: string(typedefs.GurufocusRequestFinancials)},
		{f: app.Datasource.Gurufocus.Financials, dataType: string(typedefs.GurufocusRequestDividend)},
		{f: app.Datasource.Gurufocus.Price, dataType: string(typedefs.GurufocusRequestPrice)},
		{f: app.Datasource.Gurufocus.Keyratios, dataType: string(typedefs.GurufocusRequestKeyratios)},
		// {f: app.Datasource.Alphavantage.CashFlow, dataType: string(typedefs.AlphavantageRequestCashFlow)},
		// {f: app.Datasource.Alphavantage.Earnings, dataType: string(typedefs.AlphavantageRequestEarnings)},
		// {f: app.Datasource.Alphavantage.Overview, dataType: string(typedefs.AlphavantageRequestOverview)},
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
		log.Println("Failed to register a consumer.", err)
	}

	for d := range msgs {
		if d.RoutingKey == "new_report" {
			var action typedefs.ITickerAction
			err := json.Unmarshal(d.Body, &action)
			if err != nil {
				log.Printf("Unable to unmarshal action: %s. Error: %v", d.Body, err)
				continue
			}
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
		"action": "data_result",
		"payload": map[string]interface{}{
			"ticker": ticker,
			"type":   message_type,
			"data":   message,
		},
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
		connection, err := rabbitmq.ConnectToRabbitMQ(app.RabbitmqUrl)
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

func (app *App) GatheringInformation(action typedefs.ITickerAction, user_id string, client_id string) {
	var wg sync.WaitGroup
	for _, p := range app.Methods {
		wg.Add(1)
		go func(ticker, datatype, user, client string, method func(t string) (*interface{}, error)) {
			defer wg.Done()
			data, err := method(ticker)
			if err != nil {
				log.Printf("Unable to get %s for \"%s\". Error: %v", datatype, ticker, err)
			} else {
				app.PublishResults(data, user, client, datatype, ticker)
			}
		}(action.Ticker, p.dataType, user_id, client_id, p.f)
	}
	wg.Wait()
	log.Println("End gathering information for report")
}
