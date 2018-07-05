package goteams

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"github.com/hashicorp/golang-lru"
)

const NEW_MESSAGE_URL = "https://api.ciscospark.com/v1/messages"
const NEW_PERSON_URL = "https://api.ciscospark.com/v1/people"

type TeamsClient struct {
	// Any type that implements the interface
	EventProcessors map[string][]TeamsMessageProcessor
	Listeners sync.WaitGroup
	client apiClient
	lruCache *lru.Cache
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

func NewClient() TeamsClient {
	fmt.Println("Creating a new Client")
	config := loadEnv()
	bot := TeamsClient{}
	bot.lruCache, _ = lru.New(50)
	bot.Listeners = sync.WaitGroup{}
	bot.client = newApiClient(NewConfig(config))
	return bot
}

func (b *TeamsClient) RegisterNewListener(ep TeamsMessageProcessor) {
	fmt.Println("Adding a new listener for " + ep.GetResource() + " events ")
	b.EventProcessors[ep.GetResource()] = append(b.EventProcessors[ep.GetResource()], ep);
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
	b.client.sendMessage(newMsg)
}

func (b *TeamsClient) GetUserDetails( userId string ) Person {
	return b.client.getUserDetails(userId)
}


// Private Functions

func (b *TeamsClient) webSocketListenerCallBack( w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("X-Spark-Signature")
	fmt.Println("Message Received")
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if checkMAC( body, []byte(auth), []byte(b.client.config.sparkSecret) ) {
		w.Write([]byte("200 - Authenticated"))

		var envelope WebhookMessage
		err = json.Unmarshal(body, &envelope)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		switch envelope.Resource {
		case "messages":
			b.actOnMessage(envelope)
		}
	} else {
		fmt.Println("WARNING - MESSAGE RECEIVED WITH INVALID AUTH HEADER: " + auth)
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("404 - Not Authenticated"))
	}
}
/**
 Private
 Act on an incoming message by scanning for event processors
 */
func( b *TeamsClient) actOnMessage(  envelope WebhookMessage ) {
	var message Message
	json.Unmarshal(envelope.Data, &message)
	// Skip messages from myself
	if message.SenderEmail == b.client.config.username {
		return
	}
	decryptedMessage := b.client.getFullMessage( message.ID )
	for i := 0; i < len(b.EventProcessors["messages"]); i++ {
		proc := b.EventProcessors["message"][i]
		proc.OnMessage(*b, decryptedMessage)
	}
}

func (b *TeamsClient) startServer() {
	http.HandleFunc("/", b.webSocketListenerCallBack )
	err := http.ListenAndServe(":8080", nil)
	Croak(err)
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