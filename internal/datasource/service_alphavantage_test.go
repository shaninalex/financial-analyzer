package datasource

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAlphavantageOverview(t *testing.T) {
	// Create a test server to mock HTTP responses
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mock the response for the "OVERVIEW" function
		if r.URL.Query().Get("function") == "OVERVIEW" && r.URL.Query().Get("symbol") == "AAPL" {
			// Simulate a successful JSON response
			w.Write([]byte(`{"some_field": "some_value"}`))
		} else {
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	}))
	defer server.Close()

	// Create a new Alphavantage instance for testing
	alphavantage := &Alphavantage{
		API_KEY: server.URL, // Use the test server URL as the API endpoint
	}

	// Call the Overview function with a test ticker symbol
	result, err := alphavantage.Overview("AAPL")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Check if the result fields are as expected
	if result == nil {
		t.Error("Expected a result, but got nil")
	} else if result.Ticker != "AAPL" {
		t.Errorf("Expected SomeField to be 'some_value', but got %s", result.Ticker)
	}
}
