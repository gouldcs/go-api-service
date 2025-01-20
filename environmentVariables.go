package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvironmentVariable(key string) string {
    envLoadErr := godotenv.Load(".env")
    if envLoadErr != nil { log.Fatalf("error loading .env file.")}

	return os.Getenv(key)
}