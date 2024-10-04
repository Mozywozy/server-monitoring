package main

import (
	"log"
	"server-monitor/internal/monitor"
	"server-monitor/internal/storage"
	"server-monitor/web"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	storage.InitDB()

	// Start monitoring
	m := monitor.NewMonitor()
	go m.StartMonitoring()

	// Start the API server
	web.StartAPIServer()
}
