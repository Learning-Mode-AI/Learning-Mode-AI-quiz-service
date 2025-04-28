package main

import (
	"net/http"
	"os"

	"Learning-Mode-AI-quiz-service/pkg/config" // Import the config package
	"Learning-Mode-AI-quiz-service/pkg/router"
	"Learning-Mode-AI-quiz-service/pkg/services"
	logrus "github.com/sirupsen/logrus"
)

// Initialize the logger
var logger = logrus.New()

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

	// Log service start message
	logger.WithFields(logrus.Fields{
		"port": port,
	}).Info("Quiz Service is running on port...")

	// Start the HTTP server
	if err := http.ListenAndServe(":"+port, r); err != nil {
		logger.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("Failed to start server")
	}
}
