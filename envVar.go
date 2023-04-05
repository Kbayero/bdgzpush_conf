package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Function for load envoronment variables
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getenv(key string) string {
	value, defined := os.LookupEnv(key)
	if !defined {
		log.Fatalf("Error loading environment variable: %s: environment variable does not exist\n", key)
	}
	if (value == "") || (value == " ") {
		log.Fatalf("Error loading environment variable: %s: empty environment variable\n", key)
	}
	return value
}
