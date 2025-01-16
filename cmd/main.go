package main

import (
	"log"
	"net/http"
	"os"

	"Learning-Mode-AI-quiz-service/pkg/config" // Import the config package
	"Learning-Mode-AI-quiz-service/pkg/router"
	"Learning-Mode-AI-quiz-service/pkg/services"
)

func main() {
	// Load configuration
	config.InitConfig()

	// Set up port
	port := os.Getenv("QUIZ_SERVICE_PORT")
	if port == "" {
		port = "8084" // Default port if not set
	}

	// Initialize Redis using the RedisHost from config
	services.InitRedis(config.RedisHost, "", 0)

	// Initialize the router
	r := router.NewRouter()

	// Start the HTTP server
	log.Printf("Quiz Service is running on port %s...", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
