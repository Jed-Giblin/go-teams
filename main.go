package main

import (
	"github.com/joho/godotenv"
	"log"
	"godma/teams_client"
)

func main() {
	loadEnv();
}

func loadEnv() {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot := teams_client.NewClient(myEnv)
	bot.Start()
}