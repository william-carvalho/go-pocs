package queue

import (
	"context"
	"testing"
	"time"
)

func TestDispatcherFlushesBeforeClose(t *testing.T) {
	t.Parallel()

	count := 0
	dispatcher := New(2, func(context.Context, any) error {
		count++
		return nil
	})

	if err := dispatcher.Enqueue(context.Background(), "a"); err != nil {
		t.Fatalf("enqueue: %v", err)
	}
	if err := dispatcher.Enqueue(context.Background(), "b"); err != nil {
		t.Fatalf("enqueue: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := dispatcher.Close(ctx); err != nil {
		t.Fatalf("close: %v", err)
	}
	if count != 2 {
		t.Fatalf("expected 2 items to flush, got %d", count)
	}
}
