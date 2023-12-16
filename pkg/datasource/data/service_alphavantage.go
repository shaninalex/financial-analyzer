package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Alphavantage struct {
	DEBUG   bool
	API_KEY string
	API_URL string
}

func InitAlphavantage(api_key string, debug bool) *Alphavantage {
	return &Alphavantage{
		DEBUG:   debug,
		API_KEY: api_key,
		API_URL: "https://www.alphavantage.co",
	}
}

func (a *Alphavantage) req(function, symbol string) (*http.Response, error) {
	url := fmt.Sprintf("%s/query?function=%s&symbol=%s&apikey=%s",
		a.API_URL,
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

// Request response example: https://www.alphavantage.co/query?function=OVERVIEW&symbol=IBM&apikey=demo
// Request documentation: https://www.alphavantage.co/documentation/#company-overview
func (a *Alphavantage) Overview(ticker string) (*interface{}, error) {
	var results interface{}
	if a.DEBUG {
		fileBytes, _ := os.ReadFile("/demo_data/alphavantage_overview.json")
		err := json.Unmarshal(fileBytes, &results)
		if err != nil {
			return nil, err
		}
		return &results, nil
	}
	resp, err := a.req("OVERVIEW", ticker)
	if err != nil {
		return nil, err
	}
	if err = a.unpack(resp, &results); err != nil {
		return nil, err
	}

	return &results, nil
}

// Request documentation: https://www.alphavantage.co/documentation/#earnings
// Request example: https://www.alphavantage.co/query?function=EARNINGS&symbol=IBM&apikey=demo
func (a *Alphavantage) Earnings(ticker string) (*interface{}, error) {
	var results interface{}
	if a.DEBUG {
		fileBytes, _ := os.ReadFile("/demo_data/alphavantage_earnings.json")
		err := json.Unmarshal(fileBytes, &results)
		if err != nil {
			return nil, err
		}
		return &results, nil
	}
	resp, err := a.req("EARNINGS", ticker)
	if err != nil {
		return nil, err
	}
	if err = a.unpack(resp, &results); err != nil {
		return nil, err
	}
	return &results, nil
}

// Request documentation: https://www.alphavantage.co/documentation/#cash-flow
// Request example: https://www.alphavantage.co/query?function=CASH_FLOW&symbol=IBM&apikey=demo
func (a *Alphavantage) CashFlow(ticker string) (*interface{}, error) {
	var results interface{}
	if a.DEBUG {
		fileBytes, _ := os.ReadFile("/demo_data/alphavantage_cash_flow.json")
		err := json.Unmarshal(fileBytes, &results)
		if err != nil {
			return nil, err
		}
		return &results, nil
	}

	resp, err := a.req("CASH_FLOW", ticker)
	if err != nil {
		return nil, err
	}
	if err = a.unpack(resp, &results); err != nil {
		return nil, err
	}
	return &results, nil
}
