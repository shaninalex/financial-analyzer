package main

import (
	"log"
	"os"

	"github.com/shaninalex/financial-analyzer/internal/rabbitmq"
	"github.com/shaninalex/financial-analyzer/internal/redis"
)

var (
	DEBUG        = os.Getenv("DEBUG") // "0" or "1"
	GURU_API_KEY = os.Getenv("GURU_API_KEY")
	RABBITMQ_URL = os.Getenv("RABBITMQ_URL")
	APP_PORT     = os.Getenv("APP_PORT")

	// for kratos proxy
	PORT       = os.Getenv("PORT")
	KRATOS_URL = os.Getenv("KRATOS_URL")
	REDIS_URL  = os.Getenv("REDIS_URL")
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

	api, err := InitializeApplication(GURU_API_KEY, connection, channel, redisClient)
	if err != nil {
		panic(err)
	}

	api.ConsumeRabbitMessages()

	// return nil
	// err = InitDatasource(connection, channel, GURU_API_KEY, redisClient)
	// if err != nil {
	// 	panic(err)
	// }

	defer func() {
		log.Println("Close rabbitmq connections")
		connection.Close()
		channel.Close()
	}()

}
