package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
}

var App Config

func InitConfig() {

	_ = godotenv.Load()

	App = Config{
		Port:         os.Getenv("PORT"),
	}
}	