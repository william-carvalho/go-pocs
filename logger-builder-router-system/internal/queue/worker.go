package queue

import (
	"context"
	"errors"
	"sync"
)

var ErrClosed = errors.New("queue is closed")

type Handler func(context.Context, any) error

type job struct {
	ctx   context.Context
	value any
}

type Dispatcher struct {
	handler Handler
	ch      chan job

	mu     sync.RWMutex
	closed bool
	wg     sync.WaitGroup
	once   sync.Once
}

func New(bufferSize int, handler Handler) *Dispatcher {
	if bufferSize <= 0 {
		bufferSize = 256
	}
	d := &Dispatcher{
		handler: handler,
		ch:      make(chan job, bufferSize),
	}
	d.wg.Add(1)
	go d.run()
	return d
}

func (d *Dispatcher) Enqueue(ctx context.Context, value any) error {
	d.mu.RLock()
	closed := d.closed
	d.mu.RUnlock()
	if closed {
		return ErrClosed
	}

	select {
	case d.ch <- job{ctx: ctx, value: value}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (d *Dispatcher) Close(ctx context.Context) error {
	d.once.Do(func() {
		d.mu.Lock()
		d.closed = true
		close(d.ch)
		d.mu.Unlock()
	})

	done := make(chan struct{})
	go func() {
		d.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (d *Dispatcher) run() {
	defer d.wg.Done()
	for item := range d.ch {
		_ = d.handler(item.ctx, item.value)
	}
}
