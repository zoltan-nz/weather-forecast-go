package main

import (
	"github.com/zoltan-nz/weather-forecast-go/server"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "weather-forecast-go: ", log.LstdFlags|log.Lshortfile)

	srv := server.NewServer(logger)
	srv.SetupRoutes()

	logger.Printf("Starting server on :8080")
	if err := srv.Run(":8080"); err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}

}
