package goteams

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

type newMessage struct {
	RoomID 				string 		`json:"roomId,omitempty"`
	ToPersonID 			string 		`json:"toPersonId,omitempty"`
	ToPersonEmail 		string 		`json:"toPersonEmail,omitempty"`
	Text 		  		string 		`json:"text"`
	Markdown	  		string 		`json:"markdown"`
	Files		  		[]string 	`json:"files,omitempty"`
}

