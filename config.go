package goteams

import (
	"github.com/joho/godotenv"
	"log"
)

type TeamsConfig struct {
	SparkSecret string
	AccessToken string
	Username string
	BotId string
	TargetURL string
}


func loadEnv() map[string]string {
	var conf map[string]string
	conf, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return conf
}

