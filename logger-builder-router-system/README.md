# Logger Builder Router System

A simple and clean Go logging library built around three core ideas:

- one application-facing logger API
- a builder for configuration
- a router that dispatches logs to one or many providers

## Project Structure

```text
logger-builder-router-system/
├── builder/
├── examples/
├── internal/queue/
├── logger/
├── provider/
│   ├── console/
│   ├── elk/
│   └── file/
├── router/
├── tests/
├── go.mod
└── README.md
```

## Features

- log levels: `DEBUG`, `INFO`, `WARN`, `ERROR`
- structured fields with `map[string]any`
- one or many outputs at the same time
- sync or async logging per logger instance
- graceful `Close()` for async flush
- console, file, and ELK-compatible HTTP output
- provider interface for easy extension

## Quick Start

```go
appLogger, err := builder.NewLoggerBuilder().
    WithLevel(logger.InfoLevel).
    WithAsync(true).
    AddConsole().
    AddFile("app.log").
    AddELK(elk.Config{Endpoint: "http://localhost:9200/logs/_doc"}).
    Build()
if err != nil {
    return err
}
defer appLogger.Close()

appLogger.Info("user created", logger.Fields{
    "user_id": 123,
    "module":  "auth",
})
```

## How It Works

- `builder` collects config and provider choices.
- `logger` exposes the public API used by application code.
- `router` sends each log entry to every configured provider.
- `internal/queue` is used only when async mode is enabled.
- each provider implements the same `provider.Provider` interface.

## Sync vs Async

- sync mode writes in the caller goroutine
- async mode pushes entries into a buffered channel and a worker goroutine writes them in the background
- `Close()` should always be called so pending async entries are flushed

## Adding a New Provider

Implement the provider interface:

```go
type Provider interface {
    Name() string
    Write(context.Context, logger.Entry) error
    Close() error
}
```

Then pass it to the builder with `AddProvider(...)`.

## Setup

```bash
go test ./...
go run ./examples
```

## Notes

- the ELK provider is a small HTTP adapter intended as a clean starting point
- the file provider is intentionally simple and can later be replaced with a rotating writer without changing the logger API
- logging methods do not return errors directly; provider failures are routed through the configured error handler
