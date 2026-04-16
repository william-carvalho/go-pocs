package file

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/example/logger-builder-router-system/logger"
	"github.com/example/logger-builder-router-system/provider/console"
)

type Config struct {
	Path       string
	TimeFormat string
	Perm       os.FileMode
	OpenFile   func(string, int, os.FileMode) (*os.File, error)
}

type Provider struct {
	file       io.WriteCloser
	timeFormat string
	mu         sync.Mutex
}

func New(cfg Config) (*Provider, error) {
	if cfg.Path == "" {
		return nil, fmt.Errorf("file path is required")
	}
	if cfg.Perm == 0 {
		cfg.Perm = 0o644
	}
	if cfg.OpenFile == nil {
		cfg.OpenFile = os.OpenFile
	}
	if err := os.MkdirAll(filepath.Dir(cfg.Path), 0o755); err != nil {
		return nil, fmt.Errorf("create log dir: %w", err)
	}
	f, err := cfg.OpenFile(cfg.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, cfg.Perm)
	if err != nil {
		return nil, fmt.Errorf("open log file: %w", err)
	}
	return &Provider{file: f, timeFormat: cfg.TimeFormat}, nil
}

func NewWithWriteCloser(w io.WriteCloser, timeFormat string) *Provider {
	return &Provider{file: w, timeFormat: timeFormat}
}

func (p *Provider) Name() string { return "file" }

func (p *Provider) Write(_ context.Context, entry logger.Entry) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	_, err := fmt.Fprintln(p.file, console.Format(entry, p.timeFormat))
	return err
}

func (p *Provider) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.file == nil {
		return nil
	}
	return p.file.Close()
}
