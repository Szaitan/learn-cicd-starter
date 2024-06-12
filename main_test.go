package main

import (
	"testing"
)

// Test function for GetAPIkey
func TestGetAPIKey(t *testing.T) {
	result := GetAPIKey()
	expected := "my-secret-api-key"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}
