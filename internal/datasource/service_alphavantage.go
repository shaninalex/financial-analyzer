package datasource

import (
	"encoding/json"
	"fmt"
	"io"
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

func (a *Alphavantage) unpack(resp *http.Response, resultStruct interface{}) error {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, resultStruct); err != nil {
		return err
	}
	return nil
}

type AlphavantageOverview struct {
	Ticker string
}

func (a *Alphavantage) Overview(ticker string) (*AlphavantageOverview, error) {
	resp, err := a.req("OVERVIEW", ticker)
	if err != nil {
		return nil, err
	}
	var results AlphavantageOverview
	if err = a.unpack(resp, &results); err != nil {
		return nil, err
	}
	return &results, nil
}

type AlphavantageEarnings struct {
}

func (a *Alphavantage) Earnings(ticker string) (*AlphavantageEarnings, error) {
	resp, err := a.req("EARNINGS", ticker)
	if err != nil {
		return nil, err
	}
	var results AlphavantageEarnings
	if err = a.unpack(resp, &results); err != nil {
		return nil, err
	}
	return &results, nil
}

type AlphavantageCashFlow struct {
}

func (a *Alphavantage) CashFlow(ticker string) (*AlphavantageCashFlow, error) {
	resp, err := a.req("CASH_FLOW", ticker)
	if err != nil {
		return nil, err
	}
	var results AlphavantageCashFlow
	if err = a.unpack(resp, &results); err != nil {
		return nil, err
	}
	return &results, nil
}

// type AlphavantageSeries struct {
// }
//
// func (a *Alphavantage) Series(ticker string) (*AlphavantageSeries, error) {
// 	resp, err := a.req("TIME_SERIES_MONTHLY", ticker)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var results AlphavantageSeries
// 	if err = a.unpack(resp, &results); err != nil {
// 		return nil, err
// 	}
// 	return &results, nil
// }
