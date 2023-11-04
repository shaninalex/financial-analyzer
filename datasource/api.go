package main

import (
	"fmt"

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
