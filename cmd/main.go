package main

import (
	"Youtube-Learning-Mode-Quiz-Service/pkg/router"
	"log"
	"net/http"
	"os"
)

func main() {
	r := router.SetupRouter() // Set up the router

	port := os.Getenv("PORT")
	if port == "" {
		port = "8083" // Default port for quiz service
	}

	log.Printf("Quiz service running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r)) // Start server
}
