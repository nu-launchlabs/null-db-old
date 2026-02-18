package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nu-launchlabs/null-db/api"
)

func main() {
	// load .env file
	if err := godotenv.Load(); err != nil {
		slog.Warn("No .env file found")
	}

	h := api.NewHandler()
	router := http.NewServeMux()

	// routes
	router.HandleFunc("/health", h.HealthCheckHandler)

	// load env
	port := getEnv("PORT", "8080")

	// start server
	slog.Info("Starting server on :" + port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		slog.Error("ListenAndServe failed", "error", err)
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
