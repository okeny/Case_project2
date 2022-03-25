package config

import (
	"log"

	"github.com/joho/godotenv"
)

//This gets the env variables for setting up db connect

func Init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
