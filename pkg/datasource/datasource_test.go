package datasource

import (
	"os"
	"testing"
)

func TestDataSourceInitialization(t *testing.T) {
	datasource := InitializeDatasource(
		os.Getenv("GURUFOCUS_API_KEY"),
		"demo",
		true,
	)

	if datasource.Gurufocus == nil {
		t.Error("Failed to initialize GuruFocus subservice")
	}

	if datasource.Alphavantage == nil {
		t.Error("Failed to initialize Alphavantage subservice")
	}
}
