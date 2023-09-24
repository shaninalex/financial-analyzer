package api

import (
	"testing"
)

func TestApCreation(t *testing.T) {

	_, err := InitializeApi()
	if err != nil {
		t.Error("Failed to initialize API")
	}
}
