package main

import (
	"fmt"
	"os"

	"github.com/shaninalex/financial-analyzer/cmd/report/manager"
	"github.com/shaninalex/financial-analyzer/internal/db"
	"github.com/shaninalex/financial-analyzer/internal/rabbitmq"
)

var (
	RABBITMQ_URL           = os.Getenv("RABBITMQ_URL")
	POSTGRES_USER          = os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD      = os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DATABASE_NAME = os.Getenv("POSTGRES_DATABASE_NAME")
	POSTGRES_HOST          = os.Getenv("POSTGRES_HOST")
)

func main() {

	connection, channel, err := rabbitmq.ConnectToRabbitMQ(RABBITMQ_URL)
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		POSTGRES_HOST, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DATABASE_NAME,
	)
	database, err := db.InitDatabase(dsn, "main", "psql")
	if err != nil {
		panic(err)
	}

	manager, err := manager.InitReportManager(connection, channel, database)
	if err != nil {
		panic(err)
	}

	manager.ConsumeMessages()

	defer func() {
		connection.Close()
		channel.Close()
	}()

}
