package datasource

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Alphavantage struct {
	API_KEY string
	APIURL  string
}

func InitAlphavantage(api_key string) *Alphavantage {
	return &Alphavantage{
		API_KEY: api_key,
		APIURL:  "https://www.alphavantage.co",
	}
}

func (a *Alphavantage) req(function, symbol string) (*http.Response, error) {
	url := fmt.Sprintf("%s/query&function=%s&symbol=%s&apiKey=%s",
		a.APIURL,
		function,
		symbol,
		a.API_KEY,
	)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *Alphavantage) unpack(data *http.Response, int interface{}) (*interface{}, error) {
	if err := json.Unmarshal([]byte(data), &int); err != nil {
		return nil, err
	}
	return nil
}

func (a *Alphavantage) Overview(ticker string) (*AlphavantageOverview, error) {
	var results AlphavantageOverview
	resp, err := a.req("OVERVIEW", ticker)
	if err != nil {
		// TODO: send signal about error to "office" service via rabbitmq
		return nil, err
	}

	if err != a.unpack(res.)
	return &results, nil
}

type AlphavantageOverview struct {
}
