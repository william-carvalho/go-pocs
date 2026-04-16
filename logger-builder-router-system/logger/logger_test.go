package logger

import (
	"context"
	"sync"
	"testing"
	"time"
)

type stubDispatcher struct {
	mu      sync.Mutex
	entries []Entry
}

func (d *stubDispatcher) Route(_ context.Context, entry Entry) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.entries = append(d.entries, entry)
	return nil
}

func (d *stubDispatcher) Close() error { return nil }

func TestLoggerFiltersLevels(t *testing.T) {
	t.Parallel()

	dispatcher := &stubDispatcher{}
	log := New(Config{Level: WarnLevel}, dispatcher)

	log.Info("ignored", nil)
	log.Error("kept", nil)

	if len(dispatcher.entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(dispatcher.entries))
	}
	if dispatcher.entries[0].Message != "kept" {
		t.Fatalf("unexpected entry: %+v", dispatcher.entries[0])
	}
}

func TestLoggerAsyncFlushesOnClose(t *testing.T) {
	t.Parallel()

	dispatcher := &stubDispatcher{}
	log := New(Config{
		Level:      DebugLevel,
		Async:      true,
		BufferSize: 4,
	}, dispatcher)

	log.Info("user created", Fields{"user_id": 123})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := log.CloseContext(ctx); err != nil {
		t.Fatalf("close logger: %v", err)
	}

	if len(dispatcher.entries) != 1 {
		t.Fatalf("expected 1 flushed entry, got %d", len(dispatcher.entries))
	}
}
