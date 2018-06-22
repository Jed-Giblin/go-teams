package goteams

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"errors"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"bytes"
)

const NEW_MESSAGE_URL = "https://api.ciscospark.com/v1/messages"

type TeamsClient struct {
	Config TeamsConfig
	// Any type that implements the interface
	EventProcessors []TeamsMessageProcessor
	Listeners sync.WaitGroup
	client http.Client
}

type TeamsMessageProcessor interface {
	GetEvent() string
	GetResource() string
	OnMessage( client TeamsClient, msg Message )
	OnRoom( client TeamsClient, msg Room )
	//TODO - Support Memberships
	OnMembership(client TeamsClient)
	//TODO - Support Teams
	OnTeam(client TeamsClient)
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
	bot.client = http.Client{}
	return bot
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

func (b *TeamsClient) Respond( text string, markdown string, files []string, oMsg Message) {
	newMsg := newMessage{}
	newMsg.Files = files
	newMsg.Markdown = markdown
	newMsg.Text = text
	newMsg.ToPersonEmail = oMsg.ToPersonEmail
	newMsg.ToPersonID = oMsg.ToPersonID
	newMsg.RoomID = oMsg.RoomID
	b.sendMessage(newMsg)
}

func(b *TeamsClient) sendMessage(msg newMessage) {
	body, err := json.Marshal(&msg)
	Croak(err)
	fmt.Println(string(body))
	Croak(err)
	req,err := http.NewRequest("POST", NEW_MESSAGE_URL, bytes.NewBuffer(body))
	req.Header.Add("Authorization", "Bearer " + b.Config.AccessToken)
	fmt.Println(req.Header.Get("Authorization"))
	Croak(err)
	res, err := b.client.Do(req)
	Croak(err)
	fmt.Println(res.Status)
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

	fmt.Println("The body of the message:")
	fmt.Printf("%x\n", body)
	fmt.Println("The supplied HMAC-SHA1 hash")
	fmt.Println(auth)
	fmt.Println("The the expected encryption")

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
		    proc.OnMessage(*b, message)
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

func checkMAC(unsignedData, receivedHMAC, key []byte) bool {
	mac := hmac.New(sha1.New, key)
	mac.Write(unsignedData)
	expectedMAC := mac.Sum(nil)
	fmt.Println(hex.EncodeToString([]byte(expectedMAC)))
	return string(receivedHMAC) == hex.EncodeToString([]byte(expectedMAC))
}

func Croak( e error ) {
	if e != nil {
		log.Fatal(e)
	}
}