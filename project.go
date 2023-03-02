package main

import (
	"encoding/json"
	"net/http"
)

type Project struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	CreationDateTime string `json:"creationDateTime"`
}

func (c *Client) ListProjects() ([]Project, error) {
	request, err := http.NewRequest("GET", c.Endpoint+"/projects", nil)
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject []Project
	json.Unmarshal(*responseData, &responseObject)

	return responseObject, nil
}
