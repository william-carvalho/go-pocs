package builder

import (
	"fmt"

	"github.com/example/logger-builder-router-system/logger"
	"github.com/example/logger-builder-router-system/provider"
	"github.com/example/logger-builder-router-system/provider/console"
	"github.com/example/logger-builder-router-system/provider/elk"
	"github.com/example/logger-builder-router-system/provider/file"
	"github.com/example/logger-builder-router-system/router"
)

type providerFactory func(logger.Config) (provider.Provider, error)

type LoggerBuilder struct {
	cfg       logger.Config
	providers []provider.Provider
	factories []providerFactory
}

func NewLoggerBuilder() *LoggerBuilder {
	return &LoggerBuilder{
		cfg: logger.Config{
			Level: logger.InfoLevel,
		},
	}
}

func (b *LoggerBuilder) WithLevel(level logger.Level) *LoggerBuilder {
	b.cfg.Level = level
	return b
}

func (b *LoggerBuilder) WithAsync(async bool) *LoggerBuilder {
	b.cfg.Async = async
	return b
}

func (b *LoggerBuilder) WithBufferSize(size int) *LoggerBuilder {
	b.cfg.BufferSize = size
	return b
}

func (b *LoggerBuilder) WithTimeFormat(format string) *LoggerBuilder {
	b.cfg.TimeFormat = format
	return b
}

func (b *LoggerBuilder) WithErrorHandler(handler logger.ErrorHandler) *LoggerBuilder {
	b.cfg.ErrorHandler = handler
	return b
}

func (b *LoggerBuilder) AddProvider(p provider.Provider) *LoggerBuilder {
	if p != nil {
		b.providers = append(b.providers, p)
	}
	return b
}

func (b *LoggerBuilder) AddConsole() *LoggerBuilder {
	b.factories = append(b.factories, func(cfg logger.Config) (provider.Provider, error) {
		return console.New(cfg.TimeFormat), nil
	})
	return b
}

func (b *LoggerBuilder) AddFile(path string) *LoggerBuilder {
	b.factories = append(b.factories, func(cfg logger.Config) (provider.Provider, error) {
		return file.New(file.Config{
			Path:       path,
			TimeFormat: cfg.TimeFormat,
		})
	})
	return b
}

func (b *LoggerBuilder) AddELK(cfg elk.Config) *LoggerBuilder {
	b.factories = append(b.factories, func(logger.Config) (provider.Provider, error) {
		return elk.New(cfg)
	})
	return b
}

func (b *LoggerBuilder) Build() (*logger.Logger, error) {
	if err := b.cfg.Validate(); err != nil {
		return nil, err
	}

	allProviders := make([]provider.Provider, 0, len(b.providers)+len(b.factories))
	allProviders = append(allProviders, b.providers...)
	for _, factory := range b.factories {
		p, err := factory(b.cfg)
		if err != nil {
			return nil, err
		}
		allProviders = append(allProviders, p)
	}
	if len(allProviders) == 0 {
		return nil, fmt.Errorf("at least one provider is required")
	}

	return logger.New(b.cfg, router.New(allProviders...)), nil
}
