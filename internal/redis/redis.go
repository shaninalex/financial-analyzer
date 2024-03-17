package redis

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

func InitRedis(addr string) (*RedisClient, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	client := &RedisClient{
		client: rdb,
		ctx:    context.Background(),
	}
	status, err := client.client.Ping(client.ctx).Result()
	if err != nil {
		return nil, err
	}

	if status != "PONG" {
		log.Printf("Wrong ping message. Want \"PONG\". Got: \"%s\"", status)
		return nil, err
	}

	return client, nil
}

func (r *RedisClient) Set(key, value string) error {
	err := r.client.Set(r.ctx, key, value, time.Duration(time.Hour*1)).Err()
	if err != nil {
		panic(err)
	}
	return nil
}

func (r *RedisClient) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}
