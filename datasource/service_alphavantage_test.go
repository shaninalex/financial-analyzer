package main

import (
	"testing"
)

// IMPORTANT:
// alphavantage api uses API_KEY "demo" to work only with "IBM" ticker as an
// example, and tests will use this api key. To get real one, visit this page:
// https://www.alphavantage.co/support/#api-key

func TestAlphavantageServiceInitialization(t *testing.T) {
	alphavantage := InitAlphavantage("demo", false)
	if alphavantage.API_URL != "https://www.alphavantage.co" {
		t.Errorf("Expected alphavantage.API_URL to be 'https://alphavantage.co', but got %s", alphavantage.API_URL)
	}

	if alphavantage.API_KEY == "" {
		t.Errorf("Expected alphavantage.API_KEY to be not empty")
	}

	if alphavantage.API_KEY != "demo" {
		t.Errorf("Expected alphavantage.API_KEY to be \"demo\", bot %v", alphavantage.API_KEY)
	}

}

func TestAlphavantageOverview(t *testing.T) {
	alphavantage := InitAlphavantage("demo", false)
	result, err := alphavantage.Overview("IBM")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if result == nil {
		t.Error("Expected a result, but got nil")
	} else if result.Symbol != "IBM" {
		t.Errorf("Expected Symbol to be 'IBM', but got %s", result.Symbol)
	}
}

func TestAlphavantageEarnings(t *testing.T) {
	alphavantage := InitAlphavantage("demo", false)
	result, err := alphavantage.Earnings("IBM")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if result == nil {
		t.Error("Expected a result, but got nil")
	} else if result.Symbol != "IBM" {
		t.Errorf("Expected Symbol to be 'IBM', but got %s", result.Symbol)
	}
}

func TestAlphavantageCashFlow(t *testing.T) {
	alphavantage := InitAlphavantage("demo", false)
	result, err := alphavantage.CashFlow("IBM")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if result == nil {
		t.Error("Expected a result, but got nil")
	} else if result.Symbol != "IBM" {
		t.Errorf("Expected Symbol to be 'IBM', but got %s", result.Symbol)
	}
}
