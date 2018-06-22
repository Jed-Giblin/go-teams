package goteams

import "encoding/json"

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
