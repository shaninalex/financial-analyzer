package tests

import (
	"testing"

	"github.com/shaninalex/financial-analyzer/internal/redis"
)

const (
	redisUrl = "localhost:6379"
)

func TestCreateRedisClient(t *testing.T) {
	redisClient, err := redis.InitRedis(redisUrl)
	if err != nil {
		t.Errorf("Unable to initialize redis client. Error: %v", err)
	}
	if redisClient == nil {
		t.Error("Initialized redis client should not be nil")
	}
}

func TestGetSetKeys(t *testing.T) {
	redisClient, _ := redis.InitRedis(redisUrl)
	err := redisClient.Set("key", "value")
	if err != nil {
		t.Errorf("Unable to set key. Error: %v", err)
	}

	result, err := redisClient.Get("key")
	if err != nil {
		t.Errorf("Unable to get key. Error: %v", err)
	}

	if result != "value" {
		t.Errorf("Wrong value. Want \"value\", got: \"%s\"", result)
	}
}
