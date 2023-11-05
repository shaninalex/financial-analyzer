package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Api struct {
	datasource   *Datasource
	router       *gin.Engine
	MQConnection *amqp.Connection
	MQChannel    *amqp.Channel
	MQQueue      *amqp.Queue
}

func InitializeAPI(gfApiKey, alphApiKey string) (*Api, error) {

	api := &Api{
		datasource: InitializeDatasource(gfApiKey, alphApiKey),
		router:     gin.Default(),
	}

	api.router.Use(UserRequestCounter())
	api.InitRoutes()

	mq_connection, err := connectToRabbitMQ(RABBITMQ_URL)
	if err != nil {
		return nil, err
	}

	ch, err := mq_connection.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"q_datasource", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		"ex_datasource", // name
		"direct",        // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		return nil, err
	}
	err = ch.QueueBind(
		q.Name,
		"",
		"ex_datasource",
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}
	api.MQConnection = mq_connection
	api.MQChannel = ch
	api.MQQueue = &q

	return api, nil
}

func (api *Api) InitRoutes() {
	api.router.GET("alphavantage/overview", nil)
	api.router.GET("alphavantage/earnings", nil)
	api.router.GET("alphavantage/cashflow", nil)
	api.router.GET("gurufocus/summary", nil)
}

func (api *Api) Run(port int) {
	api.router.Run(fmt.Sprintf(":%d", port))
}

func (api *Api) ConsumeRabbitMessages() {
	msgs, err := api.MQChannel.Consume(
		"q_datasource", // queue
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	if err != nil {
		log.Println("Failed to register a consumer:")
		log.Println(err)
	}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
}
