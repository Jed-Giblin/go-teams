package goteams

import (
	"github.com/joho/godotenv"
	"log"
)

type teamsConfig struct {
	sparkSecret string
	accessToken string
	username string
	botId string
}


func NewConfig(conf map[string]string ) teamsConfig {
	config := teamsConfig{}
	config.sparkSecret = conf["SparkSecret"]
	config.accessToken = conf["AccessToken"]
	config.username = conf["Username"]
	config.botId = conf["BotId"]
	return config
}


func loadEnv() map[string]string {
	var conf map[string]string
	conf, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return conf
}

