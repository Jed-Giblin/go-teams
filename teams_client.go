package goteams

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"errors"
	"github.com/joho/godotenv"
	"crypto/hmac"
	"crypto/sha256"
)

type TeamsConfig struct {
	SparkSecret string
	AccessToken string
	Username string
	BotId string
	TargetURL string
}

type TeamsClient struct {
	Config TeamsConfig
	// Any type that implements the interface
	EventProcessors []TeamsMessageProcessor
	Listeners sync.WaitGroup
}

type TeamsMessageProcessor interface {
	GetEvent() string
	GetResource() string
	OnMessage( msg Message )
	OnRoom( msg Room )
	//TODO - Support Memberships
	OnMembership( )
	//TODO - Support Teams
	OnTeam()
}


func NewConfig(conf map[string]string ) TeamsConfig {
	config := TeamsConfig{}
	config.SparkSecret = conf["SparkSecret"]
	config.AccessToken = conf["AccessToken"]
	config.Username = conf["Username"]
	config.BotId = conf["BotId"]
	return config
}

func NewClient() TeamsClient {
	fmt.Println("Creating a new Client")
	config := loadEnv()
	bot := TeamsClient{}
	bot.Config = NewConfig(config)
	bot.Listeners = sync.WaitGroup{}
	return bot
}

func loadEnv() map[string]string {
	var conf map[string]string
	conf, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return conf
}


func (b *TeamsClient) RegisterNewListener(ep TeamsMessageProcessor) {
	fmt.Println("Adding a new listener for " + ep.GetEvent() + " events ")
	b.EventProcessors = append(b.EventProcessors, ep)
}

func (b *TeamsClient) Start() {
	b.Listeners.Add(1)
	go b.startServer()
	fmt.Println("Listening for Events from Cisco Teams")
	b.Listeners.Wait()
}

func (b *TeamsClient) startServer() {
	http.HandleFunc("/", b.webSocketListenerCallBack )
	err := http.ListenAndServe(":8080", nil)
	Croak(err)
}

func (b *TeamsClient) webSocketListenerCallBack( w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("X-Spark-Signature")
	fmt.Println("Message Received")
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if checkMAC( body, []byte(auth), []byte(b.Config.SparkSecret) ) {
		w.Write([]byte("200 - Authenticated"))


		var envelope WebhookMessage
		err = json.Unmarshal(body, &envelope)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		switch envelope.Resource {
		case "messages":
			proc,err := b.eventProcessorsWhere( "messages" )
			Croak(err)
			var message Message
			json.Unmarshal(envelope.Data, &message)
		    proc.OnMessage(message)
		}
	} else {
		fmt.Println("WARNING - MESSAGE RECEIVED WITH INVALID AUTH HEADER: " + auth)
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("404 - Not Authenticated"))
	}
}


// Private Functions

func (b *TeamsClient) eventProcessorsWhere( resource string ) (TeamsMessageProcessor, error) {
	for i := 0; i < len(b.EventProcessors); i++ {
		if b.EventProcessors[i].GetResource() == resource {
			return b.EventProcessors[i], nil
		}
	}
	return nil, errors.New("No event processor found for " + resource)
}

func checkMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}

//func (b TeamsClient) processTeamsMessage(msg Message) {
//
//}

func Croak( e error ) {
	if e != nil {
		log.Fatal(e)
	}
}