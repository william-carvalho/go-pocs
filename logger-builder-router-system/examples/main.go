package main

import (
	"context"
	"log"
	"time"

	"github.com/example/logger-builder-router-system/builder"
	"github.com/example/logger-builder-router-system/logger"
	"github.com/example/logger-builder-router-system/provider/elk"
)

func main() {
	appLogger, err := builder.NewLoggerBuilder().
		WithLevel(logger.DebugLevel).
		WithAsync(true).
		AddConsole().
		AddFile("tmp/app.log").
		AddELK(elk.Config{Endpoint: "http://localhost:9200/logs/_doc"}).
		Build()
	if err != nil {
		log.Fatal(err)
	}

	appLogger.Info("user created", logger.Fields{
		"user_id": 123,
		"module":  "auth",
	})

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := appLogger.CloseContext(shutdownCtx); err != nil {
		log.Printf("logger close error: %v", err)
	}
}
