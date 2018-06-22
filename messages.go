package goteams

import (
	"encoding/json"
	"bytes"
	"net/http"
)

const NEW_MESSAGE_URL = "https://api.ciscospark.com/v1/messages"

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
	Files		  		[]string 	`json:"files"`
}



func (m *Message ) Respond( text string, markdown string, files []string) {
	newMsg := newMessage{}
	newMsg.Files = files
	newMsg.Markdown = markdown
	newMsg.Text = text
	newMsg.ToPersonEmail = m.ToPersonEmail
	newMsg.ToPersonID = m.ToPersonID
	newMsg.RoomID = m.RoomID
	newMsg.send()
}

func( m *newMessage ) send() {
	client := http.Client{}
	body, err := json.Marshal(&m)
	Croak(err)
	req,err := http.NewRequest("POST", NEW_MESSAGE_URL, bytes.NewBuffer(body))
	Croak(err)
	client.Do(req)
}