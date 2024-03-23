package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/shaninalex/financial-analyzer/internal/redis"
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
	Methods      []providerMethod
}

func Init(gfAppKey string, connection *amqp.Connection, channel *amqp.Channel, redisClient *redis.RedisClient) (*App, error) {
	app := &App{
		Context:      context.TODO(),
		Datasource:   InitializeDatasource(gfAppKey, true, redisClient),
		MQConnection: connection,
		MQChannel:    channel,
	}

	app.Methods = []providerMethod{
		{f: app.Datasource.Gurufocus.Summary, dataType: string(typedefs.GurufocusRequestSummary)},
		{f: app.Datasource.Gurufocus.Dividends, dataType: string(typedefs.GurufocusRequestFinancials)},
		{f: app.Datasource.Gurufocus.Financials, dataType: string(typedefs.GurufocusRequestDividend)},
		{f: app.Datasource.Gurufocus.Price, dataType: string(typedefs.GurufocusRequestPrice)},
		{f: app.Datasource.Gurufocus.Keyratios, dataType: string(typedefs.GurufocusRequestKeyratios)},
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
			var action typedefs.Action
			err := json.Unmarshal(d.Body, &action)
			if err != nil {
				log.Printf("Unable to unmarshal action: %s. Error: %v", d.Body, err)
				continue
			}
			app.GatheringInformation(
				action,
				d.Headers["user_id"].(string),
				d.Headers["client_id"].(string),
				d.Headers["request_id"].(string),
			)
		}
	}
}

func (app *App) PublishResults(message any, user_id, client_id, message_type, ticker, report_id string) {
	comp_message, _ := json.Marshal(map[string]interface{}{
		"action": "data_result",
		"payload": map[string]interface{}{
			"ticker":    ticker,
			"report_id": report_id,
			"type":      message_type,
			"data":      message,
		},
	})

	// client device
	routing_key := fmt.Sprintf("client.%s__dev.%s", user_id, client_id)

	err := app.MQChannel.PublishWithContext(app.Context,
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

	update_message, _ := json.Marshal(&typedefs.Action{
		Action: typedefs.ActionTypeUpdateReport,
		Payload: map[string]interface{}{
			"type":      message_type,
			"report_id": report_id,
		},
	})

	err = app.MQChannel.PublishWithContext(app.Context,
		"ex.report",     // exchange
		"update_report", // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        update_message,
			Headers: amqp.Table{
				"user_id":   user_id,
				"client_id": client_id,
			},
		},
	)
}

func (app *App) GatheringInformation(action typedefs.Action, user_id, client_id, report_id string) {
	var wg sync.WaitGroup

	ticker, ok := action.Payload["ticker"].(string)
	if !ok {
		log.Printf("unable to create report. Payload is corrapted: %v", action.Payload)
		return
	}

	for _, p := range app.Methods {
		wg.Add(1)
		go func(ticker, datatype, user, client string, method func(t string) (*interface{}, error)) {
			defer wg.Done()
			data, err := method(ticker)
			if err != nil {
				log.Printf("Unable to get %s for \"%s\". Error: %v", datatype, ticker, err)
			} else {
				app.PublishResults(data, user, client, datatype, ticker, report_id)
			}
		}(ticker, p.dataType, user_id, client_id, p.f)
	}
	wg.Wait()

	// TODO: If there no errors for this report - mark report as "Success=true"
	log.Println("End gathering information for report")
}
