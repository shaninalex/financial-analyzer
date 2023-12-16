package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/shaninalex/financial-analyzer/internal/typedefs"
)

type GuruFocus struct {
	API_URL string
	API_KEY string
	DEBUG   bool
}

func InitGurufocus(apikey string, debug bool) *GuruFocus {
	return &GuruFocus{
		API_KEY: apikey,
		API_URL: "https://api.gurufocus.com/public/user/",
		DEBUG:   debug,
	}
}

func (g *GuruFocus) DoGurufocusRequest(api_function typedefs.GurufocusRequestType, symbol string, args ...string) (*interface{}, error) {
	var result interface{}
	if g.DEBUG {
		fileBytes, _ := os.ReadFile(
			fmt.Sprintf("/demo_data/gurufocus_%s.json", api_function),
		)
		err := json.Unmarshal(fileBytes, &result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	}

	client := http.Client{}

	var url string
	if len(args) > 0 {
		url = fmt.Sprintf("%s/%s/stock/%s/%s%s",
			g.API_URL,
			g.API_KEY,
			symbol,
			api_function,
			args,
		)
	} else {
		url = fmt.Sprintf("%s/%s/stock/%s/%s",
			g.API_URL,
			g.API_KEY,
			symbol,
			api_function,
		)
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (g *GuruFocus) Summary(ticker string) (*interface{}, error) {
	data, err := g.DoGurufocusRequest(typedefs.GurufocusRequestSummary, ticker)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (g *GuruFocus) Financials(ticker string) (*interface{}, error) {
	data, err := g.DoGurufocusRequest(typedefs.GurufocusRequestFinancials, ticker)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (g *GuruFocus) Dividends(ticker string) (*interface{}, error) {
	data, err := g.DoGurufocusRequest(typedefs.GurufocusRequestDividend, ticker)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// api_function method: "price"
// years: current date - 10 years (?start_date=20131113&end_date=20231113)
// https://api.gurufocus.com/public/user/<api_key>/stock/AAPL/price?start_date=20131113&end_date=20231113
func (g *GuruFocus) Price(ticker string) (*interface{}, error) {
	endDate := time.Now()
	startDate := endDate.AddDate(-10, 0, 0)
	endDateString := endDate.Format("20060102")
	startDateString := startDate.Format("20060102")
	queryString := fmt.Sprintf("?start_date=%s&end_date=%s", startDateString, endDateString)
	data, err := g.DoGurufocusRequest(typedefs.GurufocusRequestPrice, ticker, queryString)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (g *GuruFocus) Keyratios(ticker string) (*interface{}, error) {
	// TODO: additional args
	data, err := g.DoGurufocusRequest(typedefs.GurufocusRequestKeyratios, ticker)
	if err != nil {
		return nil, err
	}
	return data, nil
}
