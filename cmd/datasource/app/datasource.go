/*
This package used for getting information from financials APIs such as
GuruFocus, etc.
*/
package app

import (
	"github.com/shaninalex/financial-analyzer/cmd/datasource/app/data"
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
