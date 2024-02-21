package redis

import (
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func InitRedis(addr string) (*RedisClient, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &RedisClient{
		client: rdb,
	}, nil
}

func (r *RedisClient) Set() {}
func (r *RedisClient) Get() {}
