package main

import (
	"log"
	"youtube-fetcher/api"
	"youtube-fetcher/config"
	"youtube-fetcher/db"
	"youtube-fetcher/youtube"
)

func main() {
	// Initialize configuration
	config.InitConfig()

	// Initialize database
	db.InitDB()
	db.MigrateModels()

	// Start background fetcher
	go youtube.StartFetcher("football") // Replace with desired query

	// Start API server
	r := api.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
