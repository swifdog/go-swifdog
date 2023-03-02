package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type ResponseError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type Client struct {
	Endpoint            string
	AuthorizationHeader *string
	HttpClient          *http.Client
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func NewBasicClient(email string, password string) (*Client, error) {
	header := "Basic " + basicAuth(email, password)

	return &Client{
		Endpoint:            "http://localhost:9090",
		AuthorizationHeader: &header,
		HttpClient:          &http.Client{},
	}, nil
}

func NewBearerTokenClient(token string) (*Client, error) {
	header := "Bearer " + token

	return &Client{
		Endpoint:            "http://localhost:9090",
		AuthorizationHeader: &header,
		HttpClient:          &http.Client{},
	}, nil
}

func (c *Client) ExecuteRequest(request *http.Request) (*[]byte, error) {
	// only add authorization header if header is not nil!
	if c.AuthorizationHeader != nil {
		request.Header.Add("Authorization", *c.AuthorizationHeader)
	}
	request.Header.Add("Content-Type", "application/json")

	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Status 2xx shows a success! Anyhting else will be an error.
	if response.StatusCode >= 200 || response.StatusCode < 299 {
		return &responseData, nil
	} else {
		var responseObject ResponseError
		json.Unmarshal(responseData, &responseObject)

		return nil, errors.New(responseObject.Message)
	}
}
