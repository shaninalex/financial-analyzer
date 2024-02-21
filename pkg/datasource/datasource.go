/*
This package used for getting information from financials APIs such as
GuruFocus, Alphavantage etc.
*/
package datasource

import (
	"github.com/shaninalex/financial-analyzer/internal/redis"
	"github.com/shaninalex/financial-analyzer/pkg/datasource/data"
)

type Datasource struct {
	Gurufocus    *data.GuruFocus
	Alphavantage *data.Alphavantage
}

func InitializeDatasource(gfApiKey, alphApiKey string, debug bool, redisClient *redis.RedisClient) *Datasource {
	return &Datasource{
		Gurufocus:    data.InitGurufocus(gfApiKey, debug, redisClient),
		Alphavantage: data.InitAlphavantage(alphApiKey, debug),
	}
}
