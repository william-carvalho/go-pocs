package router

import (
	"context"
	"errors"

	"github.com/example/logger-builder-router-system/logger"
	"github.com/example/logger-builder-router-system/provider"
)

type Router struct {
	providers []provider.Provider
}

func New(providers ...provider.Provider) *Router {
	copied := make([]provider.Provider, len(providers))
	copy(copied, providers)
	return &Router{providers: copied}
}

func (r *Router) Route(ctx context.Context, entry logger.Entry) error {
	var errs []error
	for _, p := range r.providers {
		if err := p.Write(ctx, entry); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func (r *Router) Close() error {
	var errs []error
	for _, p := range r.providers {
		if err := p.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}
