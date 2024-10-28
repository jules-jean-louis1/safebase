package test

import (
	"net/http"
	"testing"
)

func TestAPIResponseHealthCheck(t *testing.T) {
	url := "http://localhost:8080/health-check"
	resp, err := http.Get(url)

	if err != nil {
		t.Errorf("Error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code error: %v", resp.StatusCode)
	}
}

func TestAPIResponsePing(t *testing.T) {
	url := "http://localhost:8080/ping"
	resp, err := http.Get(url)

	if err != nil {
		t.Errorf("Error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code error: %v", resp.StatusCode)
	}
}
