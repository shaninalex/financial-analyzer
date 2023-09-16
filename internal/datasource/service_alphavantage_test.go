package datasource

import (
	"os"
	"testing"
)

// IMPORTANT:
// https://www.alphavantage.co/support/#api-key

func TestAlphavantageServiceInitialization(t *testing.T) {
	alphavantage := InitAlphavantage(os.Getenv("ALPHAVANTAGE_API_KEY"))
	if alphavantage.API_URL != "https://www.alphavantage.co" {
		t.Errorf("Expected alphavantage.API_URL to be 'https://alphavantage.co', but got %s", alphavantage.API_URL)
	}

	if alphavantage.API_KEY == "" {
		t.Errorf("Expected alphavantage.API_KEY to be not empty")
	}


}

func TestAlphavantageOverview(t *testing.T) {
	alphavantage := InitAlphavantage(os.Getenv("ALPHAVANTAGE_API_KEY"))
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
	alphavantage := InitAlphavantage(os.Getenv("ALPHAVANTAGE_API_KEY"))
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
	alphavantage := InitAlphavantage(os.Getenv("ALPHAVANTAGE_API_KEY"))
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
