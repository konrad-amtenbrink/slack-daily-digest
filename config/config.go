package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Slack struct {
	Token    string
	AppToken string
}

type Config struct {
	DatabaseUrl string
	Slack       Slack
}

func LoadConfig() Config {
	godotenv.Load(".env")

	config := Config{
		DatabaseUrl: GetEnv("DATABASE_URL"),
		Slack: Slack{
			Token:    GetEnv("SLACK_AUTH_TOKEN"),
			AppToken: GetEnv("SLACK_APP_TOKEN"),
		},
	}
	return config
}

func GetEnv(env string) string {
	value, exists := os.LookupEnv(env)
	if !exists {
		log.Fatalf("Environment variable %s not set", env)
	}
	return value
}
