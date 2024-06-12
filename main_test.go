package main

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// Test function for GetAPIkey
func TestGetAPIKey(t *testing.T) {
	result := GetAPIKey()
	expected := "my-secret-api-key"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestGetPort(t *testing.T) {
	godotenv.Load()
	result := os.Getenv("PORT")
	expected := "8081"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}
