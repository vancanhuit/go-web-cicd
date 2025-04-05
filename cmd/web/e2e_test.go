package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHelloAPI(t *testing.T) {
	app := &application{
		Logger: slog.New(slog.DiscardHandler),
	}

	server := httptest.NewServer(app.loggerMiddleware(app.router()))
	client := &http.Client{}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, server.URL+"/api/hello", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("GET /api/hello failed: %v", err)
	}
	defer res.Body.Close() //nolint:errcheck

	if res.StatusCode != http.StatusOK {
		t.Fatalf("want %d, got %d", http.StatusOK, res.StatusCode)
	}

	want := "application/json"
	got := res.Header.Get("Content-Type")
	if got != want {
		t.Fatalf("want %q, got %q", want, got)
	}

	var body map[string]string
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		t.Fatalf("Failed to decode JSON body: %q", err)
	}
	want = "Hello from Go!"
	got = body["message"]
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
