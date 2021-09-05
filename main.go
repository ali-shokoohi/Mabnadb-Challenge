package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gitlab.com/shokoohi/mabnadp-challenge/models"
	"gitlab.com/shokoohi/mabnadp-challenge/routers"
)

func main() {
	// Load os's environments from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file:", err.Error())
	}
	// Initialize models and objects
	models.LoadModels()
	// Get address and port values from os's environments
	var (
		address = os.Getenv("listen_address")
		port    = os.Getenv("listen_port")
	)
	// Use default port if It's empty
	if port == "" {
		port = "8000"
	}
	// Get router object from our router package
	router := routers.GetRouter()
	// Prepare listen address for listening
	listen := address + ":" + port
	log.Println("Listening at =>", listen)
	// Start listening
	log.Fatal(router.Run(listen))
}
