package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

  func GetDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
					GetEnvironmentVariable("DB_HOST"),
					GetEnvironmentVariable("DB_USER"),
					GetEnvironmentVariable("DB_PW"),
					GetEnvironmentVariable("DB_NAME"),
					GetEnvironmentVariable("DB_PORT"),
					GetEnvironmentVariable("DB_SSL_MODE"),
					GetEnvironmentVariable("DB_TIME_ZONE"))
  	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if (err != nil) {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	return db, err
  }