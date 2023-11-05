package main

import (
    "log"
	"go-sample-store/routes"
    "github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

	routes.HandleRequest()
}
