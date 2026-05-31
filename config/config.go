package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PushURL	string
	PlayURL	string
	Port        string
}

var App Config

func InitConfig() {

	_ = godotenv.Load()

	App = Config{
		PushURL:	os.Getenv("PushURL"),
		PlayURL:	os.Getenv("PlayURL"),
		Port:		os.Getenv("PORT"),
	}
}	