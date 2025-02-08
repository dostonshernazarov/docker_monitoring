package main

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadConfig loads environment variables from .env
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Warning: No .env file found. Using defaults.")
	}
}
