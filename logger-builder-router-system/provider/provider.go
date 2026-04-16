package provider

import (
	"context"

	"github.com/example/logger-builder-router-system/logger"
)

type Provider interface {
	Name() string
	Write(context.Context, logger.Entry) error
	Close() error
}
