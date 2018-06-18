package teams_client

import (
	"net/http"

	"io/ioutil"
	"encoding/json"
	"fmt"
	"log"
)

const WS_LIST_URI = "https://api.ciscospark.com/v1/webhooks"
const WS_NEW_URI = "https://api.ciscospark.com/v1/webhooks"

type TeamsClient struct {
	SparkSecret string
	AccessToken string
	Username string
	BotId string
	TargetURL string
}

type Message struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type WebSocketsListResponse struct {
	Websockets []WebSocket `json:"items"`
}

type WebSocket struct {
	Id string `json:"id"`
	Name string `json:"name"`
	TargetURL string `json:"targetUrl"`
	Resource string `json:"resource"`
	Event string `json:"event"`
	Filter string `json:"filter"`
	Secret string `json:"secret"`
	Status string `json:"status"`
	CreatedAt string `json:"created"`
}

func NewClient(config map[string]string ) TeamsClient {
	bot := TeamsClient{}
	bot.SparkSecret = config["SparkSecret"]
	bot.AccessToken = config["AccessToken"]
	bot.Username = config["Username"]
	bot.BotId = config["BotId"]
	bot.TargetURL = config["TargetURL"]
	return bot
}

func (b TeamsClient) Start() {
	http.HandleFunc("/", b.webSocketListenerCallBack )
	b.setupWebSocket()
	go http.ListenAndServe(":8080", nil)
	fmt.Println("Listening for Events from Cisco Teams")
}

func (b TeamsClient) setupWebSocket() {
	websockets := b.getWebsocketsList()

	if len(websockets) == 0 || ! b.webSocketExists(websockets) {
		b.createWebSocket()
	}
}

func (b TeamsClient) webSocketExists( sockets []WebSocket ) bool {
	for i := 0; i < len(sockets); i++ {
		if sockets[i].TargetURL == b.TargetURL {
			return true
		}
	}
	return false
}

func (b TeamsClient) getWebsocketsList() []WebSocket {
	client := http.Client{}
	req, _ := http.NewRequest("GET", WS_LIST_URI, nil )
	req.Header.Set("Authorization", "Bearer " + b.AccessToken)
	var list WebSocketsListResponse
	res, _ := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)

	Croak(err)

	err = json.Unmarshal(body, &list)

	Croak(err)

	return list.Websockets
}

func (b TeamsClient) createWebSocket() {

}


func (b TeamsClient) webSocketListenerCallBack( w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("X-Spark-Signature")
	if auth == b.SparkSecret  {
		w.Write([]byte("200 - Authenticated"))
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		var msg Message
		err = json.Unmarshal(body, &msg)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		b.processTeamsMessage(msg)
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("404 - Not Authenticated"))
	}
}

func (b TeamsClient) processTeamsMessage(msg Message) {

}

func Croak( e error ) {
	if e != nil {
		log.Fatal(e)
	}
}