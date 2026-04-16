package logger

import (
	"context"
	"errors"
	"time"

	"github.com/example/logger-builder-router-system/internal/queue"
)

const DefaultTimeFormat = time.RFC3339Nano

type ErrorHandler func(error)

type Config struct {
	Level        Level
	Async        bool
	BufferSize   int
	TimeFormat   string
	ErrorHandler ErrorHandler
}

type dispatcher interface {
	Route(context.Context, Entry) error
	Close() error
}

type Logger struct {
	cfg        Config
	dispatcher dispatcher
	queue      *queue.Dispatcher
	closed     bool
}

func New(cfg Config, d dispatcher) *Logger {
	cfg = cfg.normalize()

	l := &Logger{
		cfg:        cfg,
		dispatcher: d,
	}
	if cfg.Async {
		l.queue = queue.New(cfg.BufferSize, func(ctx context.Context, value any) error {
			entry := value.(Entry)
			if err := d.Route(ctx, entry); err != nil {
				cfg.ErrorHandler(err)
			}
			return nil
		})
	}
	return l
}

func (l *Logger) Debug(message string, fields Fields) {
	l.log(context.Background(), DebugLevel, message, fields)
}
func (l *Logger) Info(message string, fields Fields) {
	l.log(context.Background(), InfoLevel, message, fields)
}
func (l *Logger) Warn(message string, fields Fields) {
	l.log(context.Background(), WarnLevel, message, fields)
}
func (l *Logger) Error(message string, fields Fields) {
	l.log(context.Background(), ErrorLevel, message, fields)
}

func (l *Logger) DebugContext(ctx context.Context, message string, fields Fields) {
	l.log(ctx, DebugLevel, message, fields)
}

func (l *Logger) InfoContext(ctx context.Context, message string, fields Fields) {
	l.log(ctx, InfoLevel, message, fields)
}

func (l *Logger) WarnContext(ctx context.Context, message string, fields Fields) {
	l.log(ctx, WarnLevel, message, fields)
}

func (l *Logger) ErrorContext(ctx context.Context, message string, fields Fields) {
	l.log(ctx, ErrorLevel, message, fields)
}

func (l *Logger) Close() error {
	return l.CloseContext(context.Background())
}

func (l *Logger) CloseContext(ctx context.Context) error {
	if l.closed {
		return nil
	}
	l.closed = true

	var errs []error
	if l.queue != nil {
		if err := l.queue.Close(ctx); err != nil {
			errs = append(errs, err)
		}
	}
	if err := l.dispatcher.Close(); err != nil {
		errs = append(errs, err)
	}
	return errors.Join(errs...)
}

func (l *Logger) log(ctx context.Context, level Level, message string, fields Fields) {
	if l.closed || level < l.cfg.Level {
		return
	}
	if ctx == nil {
		ctx = context.Background()
	}

	entry := Entry{
		Timestamp: time.Now().UTC(),
		Level:     level,
		Message:   message,
		Fields:    fields.Clone(),
	}

	if l.cfg.Async {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err := l.queue.Enqueue(ctx, entry); err != nil {
			l.cfg.ErrorHandler(err)
		}
		return
	}

	if err := l.dispatcher.Route(ctx, entry); err != nil {
		l.cfg.ErrorHandler(err)
	}
}

func (c Config) Validate() error {
	if !c.Level.IsValid() {
		return errors.New("invalid level")
	}
	if c.BufferSize < 0 {
		return errors.New("buffer size must be >= 0")
	}
	return nil
}

func (c Config) normalize() Config {
	if c.BufferSize <= 0 {
		c.BufferSize = 256
	}
	if c.TimeFormat == "" {
		c.TimeFormat = DefaultTimeFormat
	}
	if c.ErrorHandler == nil {
		c.ErrorHandler = func(error) {}
	}
	return c
}
