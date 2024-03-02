package datasource

import (
	"github.com/rabbitmq/amqp091-go"
	"github.com/shaninalex/financial-analyzer/internal/redis"
)

func Init(connection *amqp091.Connection, channel *amqp091.Channel, GURU_API_KEY string, redisClient *redis.RedisClient) error {
	api, err := InitializeApplication(GURU_API_KEY, connection, channel, redisClient)
	if err != nil {
		return err
	}

	api.ConsumeRabbitMessages()

	return nil
}
