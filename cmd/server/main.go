package main

import (
	"go-personal-web/internal/config"
	"go-personal-web/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration from config.json
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Create a Gin router
	router := gin.Default()

	// Set trusted proxies (example: trust only localhost)
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// Set up HTTP handlers
	router.GET("/", func(c *gin.Context) {
		handlers.HomeHandler(c.Writer, c.Request)
	})

	// Start the server using the loaded port from config
	log.Printf("Starting server on :%s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
