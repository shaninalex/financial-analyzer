package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GuruFocus struct {
	api_url string
	api_key string
}

func InitGurufocus(apikey string) *GuruFocus {
	return &GuruFocus{
		api_key: apikey,
		api_url: "https://api.gurufocus.com/public/user/",
	}
}

func (g *GuruFocus) reqStock(function, symbol string) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s/stock/%s/%s",
		g.api_url,
		g.api_key,
		symbol,
		function,
	)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GuruFocus) unpack(resp *http.Response, resultStruct interface{}) error {
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

func (g *GuruFocus) Summary(ticker string) (*GuruFocusSummary, error) {
	resp, err := g.reqStock("summary", ticker)
	if err != nil {
		return nil, err
	}
	var results GuruFocusSummary
	if err = g.unpack(resp, &results); err != nil {
		return nil, err
	}
	return &results, nil
}
