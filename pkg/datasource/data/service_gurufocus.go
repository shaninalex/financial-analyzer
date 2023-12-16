package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/shaninalex/financial-analyzer/internal/typedefs"
)

type GuruFocus struct {
	api_url string
	api_key string
	DEBUG   bool
}

func InitGurufocus(apikey string, debug bool) *GuruFocus {
	return &GuruFocus{
		api_key: apikey,
		api_url: "https://api.gurufocus.com/public/user/",
		DEBUG:   debug,
	}
}

func (g *GuruFocus) DoGurufocusRequest(api_function typedefs.GurufocusRequestType, symbol string) (*interface{}, error) {
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
	url := fmt.Sprintf("%s/%s/stock/%s/%s",
		g.api_url,
		g.api_key,
		symbol,
		api_function,
	)
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
