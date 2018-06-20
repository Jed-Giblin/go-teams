package goteams

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"bytes"
)

const NEW_MESSAGE_URL = "https://api.ciscospark.com/v1/messages"

func SendMessageToTeams(msg NewMessage) Message {
	client := http.Client{}
	body, err := json.Marshal(&msg)
	Croak(err)
	req,err := http.NewRequest("POST", NEW_MESSAGE_URL, bytes.NewBuffer(body))
	Croak(err)
	res,err := client.Do(req)
	msgBody, err := ioutil.ReadAll(res.Body)
	Croak(err)
	defer res.Body.Close()
	var outputMessage Message
	json.Unmarshal(msgBody, &outputMessage)
	return outputMessage
}