package goteams

import "encoding/json"

type WebhookResponse interface {}

// This is a Message object from the Teams API
type Message struct {
	ID 					string 		`json:"id"`
	RoomID 				string 		`json:"roomId"`
	RoomType 			string 		`json:"roomtType"`
	ToPersonID 			string 		`json:"toPersonId"`
	ToPersonEmail 		string 		`json:"toPersonEmail"`
	Text 		  		string 		`json:"text"`
	Markdown	  		string 		`json:"markdown"`
	Files		  		[]string 	`json:"files"`
	SenderID	  		string   	`json:"personId"`
	SenderEmail	  		string   	`json:"personEmail"`
	MessageTime			string   	`json:"created"`
	MentionedPeople 	[]string 	`json:"mentionedPeople"`
	MentionedGroups 	[]string 	`json:"mentionedGroups"`
}

type NewMessage struct {
	RoomID 				string 		`json:"roomId"`
	ToPersonID 			string 		`json:"toPersonId"`
	ToPersonEmail 		string 		`json:"toPersonEmail"`
	Text 		  		string 		`json:"text"`
	Markdown	  		string 		`json:"markdown"`
	Files		  		[]string 	`json:"files"`
}


// This ia Room object
type Room struct {
	ID 				string 		`json:"id"`
	Title			string		`json:"title"`
	Type			string		`json:"type"`
	Locked			bool		`json:"isLocked"`
	TeamID			string		`json:"teamId"`
	LastActivity	string		`json:"lastActivity"`
}

// This is the message taht comes in from the webhook
type WebhookMessage struct {
	ID 				string 		`json:"id"`
	Name 			string		`json:"name"`
	Resource   		string 		`json:"resource"`
	Event 			string 		`json:"event"`
	Filter 			string  	`json:"filter"`
	OrgID			string 		`json:"orgId"`
	CreatedBy 		string 		`json:"createdBy"`
	AppID 			string 		`json:"appId"`
	OwnedBy			string 		`json:"ownedBy"`
	Status 			string 		`json:"status"`
	ActorID			string		`json:"actorId"`
	Data 			json.RawMessage
}

type NewWebSocket struct {
	Name string `json:"name"`
	TargetURL string `json:"targetUrl"`
	Resource string `json:"resource"`
	Event string `json:"event"`
	Filter string `json:"filter"`
	Secret string `json:"secret"`
}

type ExistingWebSocket struct {
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


type WebSocketsListResponse struct {
	Websockets []ExistingWebSocket `json:"items"`
}