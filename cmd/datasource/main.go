package main

import (
	"log"
	"os"

	"github.com/shaninalex/financial-analyzer/cmd/datasource/app"
	"github.com/shaninalex/financial-analyzer/internal/rabbitmq"
	"github.com/shaninalex/financial-analyzer/internal/redis"
)

var (
	DEBUG             = os.Getenv("DEBUG") // "0" or "1"
	GURUFOCUS_API_KEY = os.Getenv("GURUFOCUS_API_KEY")
	RABBITMQ_URL      = os.Getenv("RABBITMQ_URL")
	REDIS_URL         = os.Getenv("REDIS_URL")
)

func main() {
	connection, channel, err := rabbitmq.ConnectToRabbitMQ(RABBITMQ_URL)
	if err != nil {
		panic(err)
	}

	redisClient, err := redis.InitRedis(REDIS_URL)
	if err != nil {
		panic(err)
	}

	// initialize datasource
	log.Println("initialize datasource")

	source, err := app.Init(GURUFOCUS_API_KEY, connection, channel, redisClient)
	if err != nil {
		panic(err)
	}

	source.ConsumeRabbitMessages()

	defer func() {
		log.Println("Close rabbitmq connections")
		connection.Close()
		channel.Close()
	}()
}
