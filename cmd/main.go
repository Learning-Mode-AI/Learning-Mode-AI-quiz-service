package main

import (
	"log"
	"net/http"
	"os"

	"Learning-Mode-AI-quiz-service/pkg/router" 
	"Learning-Mode-AI-quiz-service/pkg/services" 
)

func main() {
	// Load environment variables (if any)
	port := os.Getenv("QUIZ_SERVICE_PORT")
	if port == "" {
		port = "8084" // Default port if not set
	}

	// Initialize the router from router.go
	r := router.NewRouter()

	services.InitRedis("localhost:6379", "", 0) 

	// Start the HTTP server
	log.Printf("Quiz Service is running on port %s...", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
