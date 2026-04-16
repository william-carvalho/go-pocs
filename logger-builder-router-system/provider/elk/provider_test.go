package elk

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/example/logger-builder-router-system/logger"
)

type clientFunc func(*http.Request) (*http.Response, error)

func (f clientFunc) Do(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestProviderWritesJSON(t *testing.T) {
	t.Parallel()

	var body string
	p, err := New(Config{
		Endpoint: "http://elk.example/logs",
		Timeout:  time.Second,
		APIKey:   "secret",
		HTTPClient: clientFunc(func(req *http.Request) (*http.Response, error) {
			data, readErr := io.ReadAll(req.Body)
			if readErr != nil {
				t.Fatalf("read body: %v", readErr)
			}
			body = string(data)
			if req.Header.Get("Authorization") != "Bearer secret" {
				t.Fatalf("unexpected auth header: %s", req.Header.Get("Authorization"))
			}
			return &http.Response{StatusCode: http.StatusCreated, Body: http.NoBody}, nil
		}),
	})
	if err != nil {
		t.Fatalf("new provider: %v", err)
	}

	err = p.Write(context.Background(), logger.Entry{
		Level:   logger.InfoLevel,
		Message: "hello",
		Fields:  logger.Fields{"service": "api"},
	})
	if err != nil {
		t.Fatalf("write: %v", err)
	}
	if !strings.Contains(body, "\"message\":\"hello\"") {
		t.Fatalf("expected JSON body, got %s", body)
	}
}
