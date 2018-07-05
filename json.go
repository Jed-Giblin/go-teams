package goteams

import "encoding/json"

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

type newMessage struct {
	RoomID 				string 		`json:"roomId,omitempty"`
	ToPersonID 			string 		`json:"toPersonId,omitempty"`
	ToPersonEmail 		string 		`json:"toPersonEmail,omitempty"`
	Text 		  		string 		`json:"text"`
	Markdown	  		string 		`json:"markdown"`
	Files		  		[]string 	`json:"files,omitempty"`
}

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


type Room struct {
	ID 				string 		`json:"id"`
	Title			string		`json:"title"`
	Type			string		`json:"type"`
	Locked			bool		`json:"isLocked"`
	TeamID			string		`json:"teamId"`
	LastActivity	string		`json:"lastActivity"`
}

type Person struct {
	ID				string	`json:"id"`
	Emails			[]string	`json:"emails"`
	DisplayName 	string	`json:"displayName"`
	NickName		string	`json:"nickName"`
	FirstName		string	`json:"firstName"`
	LastName		string	`json:"lastName"`
	Avatar			string	`json:"avatar"`
	OrgID			string	`json:"orgId"`
	Roles			[]string `json:"roles"`
	Licenses		[]string	`json:"licenses"`
	Created			string    `json:"created"`
	Timezone		string	`json:"timezone"`
	LastActivity	string	`json:"lastActivity"`
	Status			string	`json:"status"`
	InvitePending	bool	`json:"invitePending"`
	LoginEnabled	bool	`json:"loginEnabled"`
	Type			string		`json:"type"`
}