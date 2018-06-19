package src

import (
	"github.com/joho/godotenv"
	"log"
	"godma/src/teams_client"
)

func main() {
	config := loadEnv();
	client := teams_client.NewClient(config)
}

func loadEnv() map[string]string {
	var conf map[string]string
	conf, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return conf
}