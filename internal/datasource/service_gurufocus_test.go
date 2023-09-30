package datasource

import (
	"os"
	"testing"
)

// IMPORTANT!
// GuruFocuse does not provide demo api key. It should be real one
func TestGurufocusInitialization(t *testing.T) {
	gurufocus := InitGurufocus(os.Getenv("GURUFOCUS_API_KEY"))
	if gurufocus.api_url != "https://api.gurufocus.com/public/user/" {
		t.Errorf("Expected gurufocus.api_url to be \"https://api.gurufocus.com/public/user/\", but got %s", gurufocus.api_key)
	}

	if gurufocus.api_url == "" {
		t.Errorf("api_key should not be empty. Check if you provide api_key in environment variables")
	}
}

func TestGurufocusSummary(t *testing.T) {
	gurufocus := InitGurufocus(os.Getenv("GURUFOCUS_API_KEY"))
	test_symbol := "IBM"
	results, err := gurufocus.Summary(test_symbol)
	if err != nil {
		t.Errorf("Summary request faild with: %s", err.Error())
	}

	if results.Summary.CompanyData.Symbol != test_symbol {
		t.Errorf("Expected to be results.Summary.CompanyData.Symbol to be equal \"%s\", but got %s",
			test_symbol, results.Summary.CompanyData.Symbol)
	}
}
