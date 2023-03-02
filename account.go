package main

import (
	"encoding/json"
	"net/http"
)

type Account struct {
	Id               string `json:"id"`
	Email            string `json:"email"`
	CreationDateTime string `json:"creationDateTime"`
}

func (c *Client) GetAccount() (*Account, error) {
	request, err := http.NewRequest("GET", c.Endpoint+"/account", nil)
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject Account
	json.Unmarshal(*responseData, &responseObject)

	return &responseObject, nil
}
