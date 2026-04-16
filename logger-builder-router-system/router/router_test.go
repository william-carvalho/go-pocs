package router

import (
	"context"
	"errors"
	"testing"

	"github.com/example/logger-builder-router-system/logger"
)

type providerStub struct {
	writeErr error
	closeErr error
	entries  []logger.Entry
}

func (p *providerStub) Name() string { return "stub" }

func (p *providerStub) Write(_ context.Context, entry logger.Entry) error {
	p.entries = append(p.entries, entry)
	return p.writeErr
}

func (p *providerStub) Close() error { return p.closeErr }

func TestRouterWritesToAllProviders(t *testing.T) {
	t.Parallel()

	first := &providerStub{}
	second := &providerStub{}
	rt := New(first, second)

	err := rt.Route(context.Background(), logger.Entry{Level: logger.InfoLevel, Message: "ok"})
	if err != nil {
		t.Fatalf("route: %v", err)
	}
	if len(first.entries) != 1 || len(second.entries) != 1 {
		t.Fatalf("expected both providers to receive one entry")
	}
}

func TestRouterJoinsErrors(t *testing.T) {
	t.Parallel()

	writeErr := errors.New("write error")
	closeErr := errors.New("close error")
	rt := New(&providerStub{writeErr: writeErr}, &providerStub{closeErr: closeErr})

	if err := rt.Route(context.Background(), logger.Entry{Level: logger.InfoLevel}); !errors.Is(err, writeErr) {
		t.Fatalf("expected joined write error, got %v", err)
	}
	if err := rt.Close(); !errors.Is(err, closeErr) {
		t.Fatalf("expected joined close error, got %v", err)
	}
}
