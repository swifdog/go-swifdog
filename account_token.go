package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type AccountToken struct {
	ID               string `json:"id"`
	Value            string `json:"value"`
	Description      string `json:"description"`
	CreationDateTime string `json:"creationDateTime"`
}

func (c *Client) ListAccountTokens() ([]AccountToken, error) {
	request, err := http.NewRequest("GET", c.Endpoint+"/account/tokens", nil)
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject []AccountToken
	json.Unmarshal(*responseData, &responseObject)

	return responseObject, nil
}

func (c* Client) CreateAccountToken(description *string) (*AccountToken, error) {
	body := map[string]string{}

	if description != nil {
		body["description"] = *description
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", c.Endpoint+"/account/tokens", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject AccountToken
	json.Unmarshal(*responseData, &responseObject)

	return &responseObject, nil
}

func (c* Client) DeleteAccountTokenById(accountTokenId string) error {
	request, err := http.NewRequest("DELETE", c.Endpoint+"/account/tokens/"+accountTokenId, nil)
	if err != nil {
		return err
	}

	_, err = c.ExecuteRequest(request)
	if err != nil {
		return err
	}

	return nil
}