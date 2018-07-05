package goteams

import (
	"io/ioutil"
	"bytes"
	"net/http"
	"fmt"
	"errors"
)

func (c *apiClient) get( url string, params map[string]string ) ([]byte,error) {
	req, err := http.NewRequest("GET", url, nil)
	Croak(err)
	if params != nil {
		q := req.URL.Query()
		for k, v := range params {
			q.Add( k, v)
		}
		req.URL.RawQuery = q.Encode()
	}
	res, err := c.client.Do(req)
	Croak(err)
	if res.StatusCode == 200 {
		return ioutil.ReadAll(res.Body)
	} else {
		return nil, errors.New("Unable to fetch " + url)
	}
}

func(c *apiClient) post( url string, body []byte) {
	req,err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Add("Authorization", "Bearer " + c.config.accessToken)
	req.Header.Add( "Content-Type", "application/json")
	res,err := c.client.Do(req)
	Croak(err)
	fmt.Println(res.Status)
}
