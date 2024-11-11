package main

import (
	"log"

	"github.com/Kennedy-lsd/ExpenseTracker/cmd/api"
)

func main() {
	err := api.Api()
	if err != nil {
		log.Fatalf("Error starting API: %v", err)
		return
	}

	log.Println("API started successfully!")
}
