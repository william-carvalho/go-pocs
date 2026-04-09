package main

import (
	"log"
	"net/http"

	"tax-system/handlers"
	"tax-system/repository"
	"tax-system/service"
)

func main() {
	repo := repository.NewInMemoryTaxRuleRepository()
	taxService := service.NewTaxService(repo)
	handler := handlers.NewTaxHandler(taxService)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("tax-system listening on :8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
