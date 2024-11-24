package main

import (
	"Lab4/config"
	"Lab4/router"
	"Lab4/utils"
	"log"
	"net/http"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}
	utils.InitLogger()
	defer utils.Logger.Sync()

	log.Printf("Server starting on port: %s", config.ServerPort)
	log.Printf("Log Level: %s", config.LogLevel)
	log.Printf("TLS Certificate: %s", config.TLSCertPath)

	// Set up routes and start the server
	r := router.SetupRoutes()
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
