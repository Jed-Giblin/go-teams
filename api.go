package goteams

import (
	"net/http"
	"encoding/json"
)

type apiClient struct {
	client http.Client
	config teamsConfig
}

func newApiClient(con teamsConfig) apiClient {
	c := apiClient{}
	c.client = http.Client{}
	c.config = con
	return c
}

func(c *apiClient) sendMessage(msg newMessage) {
	body, err := json.Marshal(&msg)
	Croak(err)
	c.post(NEW_MESSAGE_URL, body)
}
func( c *apiClient) getUserDetails( uId string ) Person {
	msgBody, err := c.get(NEW_PERSON_URL + "/" + uId, nil)
	Croak(err)
	var person Person
	Croak(err)
	json.Unmarshal(msgBody, &person)
	return person
}

func( c *apiClient ) getFullMessage( msgId string ) Message {
	msgBody, err := c.get( NEW_MESSAGE_URL + "/" + msgId , nil )
	var message Message
	Croak(err)
	json.Unmarshal(msgBody, &message)
	return message
}
