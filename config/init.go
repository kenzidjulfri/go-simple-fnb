package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	ConnectDatabase()
	AutoMigrate(DB)
}
