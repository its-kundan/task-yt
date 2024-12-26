package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl   string
	APIKeys []string
}

var AppConfig Config

func InitConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	AppConfig = Config{
		DBUrl: os.Getenv("DB_URL"),
		APIKeys: []string{
			os.Getenv("YOUTUBE_API_KEY_1"),
			os.Getenv("YOUTUBE_API_KEY_2"), // Add more as needed
		},
	}
}
