package teams_client

import (
	"net/http"
	"godma/src/teams_config"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"log"
	"godma/src/teams_message_processor"
	. "godma/src/teams_messages"
)

const WS_LIST_URI = "https://api.ciscospark.com/v1/webhooks"
const WS_NEW_URI = "https://api.ciscospark.com/v1/webhooks"

type ResourceType int32
const (
	MESSAGES 	ResourceType = 0
	MEMBERSHIPS ResourceType = 1
	ROOMS       ResourceType = 2
	ALL_RES			ResourceType = 3
)

type TeamsClient struct {
	Config teams_config.TeamsConfig
	// Any type that implements the interface
	EventProcessor teams_message_processor.TeamsMessageProcessor
}

type WebSocketsListResponse struct {
	Websockets []ExistingWebSocket `json:"items"`
}

type callback func( w http.ResponseWriter, r *http.Request )

func NewClient(config map[string]string ) TeamsClient {
	bot := TeamsClient{}
	bot.Config = teams_config.NewConfig(config)
	return bot
}

func (b TeamsClient) RegisterNewListener(socket NewWebSocket, cb callback) {
	fmt.Println("Checking if the passed websocket already exists")
	socketsList := b.getWebsocketsList()
	for i := 0; i < len(socketsList); i ++ {
		if socket.isEqual(socketsList[i]) {

		}
	}
}

func (b TeamsClient) Start() {
	http.HandleFunc("/", b.webSocketListenerCallBack )
	//b.setupWebSocket()
	go http.ListenAndServe(":8080", nil)
	fmt.Println("Listening for Events from Cisco Teams")
}

//func (b TeamsClient) setupWebSocket() {
//	websockets := b.getWebsocketsList()
//
//	if len(websockets) == 0 || ! b.webSocketExists(websockets) {
//		b.createWebSocket()
//	}
//}

//
//func (b TeamsClient) webSocketExists( sockets []WebSocket ) bool {
//	for i := 0; i < len(sockets); i++ {
//		if sockets[i].TargetURL == b.Config.TargetURL {
//			return true
//		}
//	}
//	return false
//}

func (b TeamsClient) getWebsocketsList() []ExistingWebSocket {
	client := http.Client{}
	req, _ := http.NewRequest("GET", WS_LIST_URI, nil )
	req.Header.Set("Authorization", "Bearer " + b.Config.AccessToken)
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
	if auth == b.Config.SparkSecret  {
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