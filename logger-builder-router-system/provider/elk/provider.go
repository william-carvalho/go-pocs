package elk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/example/logger-builder-router-system/logger"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Config struct {
	Endpoint   string
	APIKey     string
	Timeout    time.Duration
	HTTPClient HTTPClient
}

type Provider struct {
	endpoint string
	apiKey   string
	timeout  time.Duration
	client   HTTPClient
}

func New(cfg Config) (*Provider, error) {
	if cfg.Endpoint == "" {
		return nil, fmt.Errorf("endpoint is required")
	}
	if cfg.Timeout <= 0 {
		cfg.Timeout = 5 * time.Second
	}
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{Timeout: cfg.Timeout}
	}
	return &Provider{
		endpoint: cfg.Endpoint,
		apiKey:   cfg.APIKey,
		timeout:  cfg.Timeout,
		client:   cfg.HTTPClient,
	}, nil
}

func (p *Provider) Name() string { return "elk" }

func (p *Provider) Write(ctx context.Context, entry logger.Entry) error {
	body, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("marshal entry: %w", err)
	}
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.endpoint, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	if p.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+p.apiKey)
	}

	resp, err := p.client.Do(req)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("elk returned status %d", resp.StatusCode)
	}
	return nil
}

func (p *Provider) Close() error { return nil }
