package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gitlab.com/shokoohi/mabnadp-challenge/models"
	"gitlab.com/shokoohi/mabnadp-challenge/routers"
	"gitlab.com/shokoohi/mabnadp-challenge/utils/cli"
)

func main() {
	// Load os's environments from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file:", err.Error())
	}
	// Initialize models and objects
	models.LoadModels()
	models.LoadQueries()
	// Check command line arguments for AutoCreate method
	if len(os.Args) >= 3 && os.Args[1] == "create" {
		max, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalln("Input valid max number, ", err.Error())
			return
		}
		// Create some task records as max limit number
		cli.AutoCreate(uint(max))
		return
	}
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
