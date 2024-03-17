/*
This package used for getting information from financials APIs such as
GuruFocus, etc.
*/
package main

import (
	"github.com/shaninalex/financial-analyzer/cmd/datasource/data"
	"github.com/shaninalex/financial-analyzer/internal/redis"
)

type Datasource struct {
	Gurufocus *data.GuruFocus
}

func InitializeDatasource(gfApiKey string, debug bool, redisClient *redis.RedisClient) *Datasource {
	return &Datasource{
		Gurufocus: data.InitGurufocus(gfApiKey, debug, redisClient),
	}
}
