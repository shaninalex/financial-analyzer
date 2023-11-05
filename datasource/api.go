package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Api struct {
	Context      context.Context
	Datasource   *Datasource
	Router       *gin.Engine
	MQConnection *amqp.Connection
	MQChannel    *amqp.Channel
}

func InitializeAPI(gfApiKey, alphApiKey string) (*Api, error) {

	api := &Api{
		Context:    context.TODO(),
		Datasource: InitializeDatasource(gfApiKey, alphApiKey),
		Router:     gin.Default(),
	}

	api.Router.Use(UserRequestCounter())
	api.InitRoutes()

	mq_connection, err := connectToRabbitMQ(RABBITMQ_URL)
	if err != nil {
		return nil, err
	}

	ch, err := mq_connection.Channel()
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	api.MQConnection = mq_connection
	api.MQChannel = ch

	return api, nil
}

func (api *Api) InitRoutes() {
	api.Router.GET("alphavantage/overview", nil)
	api.Router.GET("alphavantage/earnings", nil)
	api.Router.GET("alphavantage/cashflow", nil)
	api.Router.GET("gurufocus/summary", nil)
}

func (api *Api) Run(port int) {
	api.Router.Run(fmt.Sprintf(":%d", port))
}

func (api *Api) ConsumeRabbitMessages() {
	msgs, err := api.MQChannel.Consume(
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

	go func() {
		for d := range msgs {
			log.Printf("Message Body: %s", d.Body)

			if d.RoutingKey == "q.datasource" {
				var action TickerAction
				err := json.Unmarshal(d.Body, &action)
				if err != nil {
					log.Printf("Unable to unmarshal action: %s. Error: %v", d.Body, err)
				}
				key := "client_id"
				if client_id, ok := d.Headers[key]; ok {
					go api.GatheringInformation(action, client_id.(string))
				} else {
					fmt.Printf("Key not found: %s\n", key)
				}
			}
		}
	}()
}

func (api *Api) PublishResults(message []byte, client_id string) {
	log.Printf("Search for %s\n", string(message))
	err := api.MQChannel.PublishWithContext(api.Context,
		"ex.datasource",                     // exchange
		fmt.Sprintf("client_%s", client_id), // routing key
		false,                               // mandatory
		false,                               // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("message published")
	}
}

func (api *Api) GatheringInformation(action TickerAction, client_id string) {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(3)
	go func(ticker string) {
		defer wg.Done()
		data, err := api.Datasource.Alphavantage.Overview(ticker)
		if err != nil {
			log.Printf("Unable to get Alphavantage.Overview for \"%s\". Error: %v", ticker, err)
		}
		bData, _ := json.Marshal(data)
		api.PublishResults(bData, client_id)
	}(action.Ticker)

	go func(ticker string) {
		defer wg.Done()
		data, err := api.Datasource.Alphavantage.CashFlow(ticker)
		if err != nil {
			log.Printf("Unable to get Alphavantage.CashFlow for \"%s\". Error: %v", ticker, err)
		}
		bData, _ := json.Marshal(data)
		api.PublishResults(bData, client_id)
	}(action.Ticker)

	go func(ticker string) {
		defer wg.Done()
		data, err := api.Datasource.Alphavantage.Earnings(ticker)
		if err != nil {
			log.Printf("Unable to get Alphavantage.Earnings for \"%s\". Error: %v", ticker, err)
		}
		bData, _ := json.Marshal(data)
		api.PublishResults(bData, client_id)
	}(action.Ticker)

	wg.Wait()
	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("Gathering information for \"%s\" is done. Time spent: %s", action.Ticker, elapsed)
}
