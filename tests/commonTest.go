package tests

import (
	"log"

	"github.com/joho/godotenv"
)

func SetUpEnv() {
	err := godotenv.Load("../../../../.env.test")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
