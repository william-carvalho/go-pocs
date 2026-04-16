package tests

import (
	"bytes"
	"testing"

	"github.com/example/logger-builder-router-system/builder"
	"github.com/example/logger-builder-router-system/logger"
	"github.com/example/logger-builder-router-system/provider/console"
)

func TestBuilderWithInjectedProvider(t *testing.T) {
	t.Parallel()

	var out bytes.Buffer
	appLogger, err := builder.NewLoggerBuilder().
		WithLevel(logger.DebugLevel).
		AddProvider(console.NewWithWriter(&out, logger.DefaultTimeFormat)).
		Build()
	if err != nil {
		t.Fatalf("build: %v", err)
	}
	defer appLogger.Close()

	appLogger.Info("integration", logger.Fields{"module": "tests"})
	if out.Len() == 0 {
		t.Fatal("expected console output")
	}
}
