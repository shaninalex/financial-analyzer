package redis

import (
	"context"
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

	return &RedisClient{
		client: rdb,
		ctx:    context.Background(),
	}, nil
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
