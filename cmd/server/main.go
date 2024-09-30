package main

import (
	"go-personal-web/internal/config"
	"go-personal-web/internal/handlers"
	"log"
	"net/http"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Set up HTTP handlers
	http.HandleFunc("/", handlers.HomeHandler)

	// Start the server
	log.Printf("Starting server on :%s", cfg.Port)
	if err := http.ListenAndServe("0.0.0.0:"+cfg.Port, nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
