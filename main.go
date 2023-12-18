package main

import (
	"log"
	"os"
	"strconv"

	"github.com/shaninalex/financial-analyzer/internal/rabbitmq"
	"github.com/shaninalex/financial-analyzer/pkg/datasource"
	"github.com/shaninalex/financial-analyzer/pkg/report"
	"github.com/shaninalex/financial-analyzer/web"
)

var (
	DEBUG        = os.Getenv("DEBUG") // "0" or "1"
	GURU_API_KEY = os.Getenv("GURU_API_KEY")
	ALPH_API_KEY = os.Getenv("ALPH_API_KEY")
	RABBITMQ_URL = os.Getenv("RABBITMQ_URL")
	APP_PORT     = os.Getenv("APP_PORT")
)

func main() {

	port, err := strconv.Atoi(APP_PORT)
	if err != nil {
		panic(err)
	}

	connection, channel, err := rabbitmq.ConnectToRabbitMQ(RABBITMQ_URL)
	if err != nil {
		panic(err)
	}

	// initialize websocket connection
	go web.Websocket(port, connection, channel)

	// initialize datasource
	log.Println("initialize datasource")
	err = datasource.Init(connection, channel, GURU_API_KEY, ALPH_API_KEY)
	if err != nil {
		panic(err)
	}

	// initialize report manager
	log.Println("initialize report manager")
	err = report.InitReportModule(connection, channel)
	if err != nil {
		panic(err)
	}

	defer func() {
		log.Println("Close rabbitmq connections")
		connection.Close()
		channel.Close()
	}()
}
