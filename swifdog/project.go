package swifdog

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Project struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	CreationDateTime string `json:"creationDateTime"`
}

type CreateOrPatchProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
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

func (c *Client) CreateProject(body *CreateOrPatchProjectRequest) (*Project, error) {
	if body == nil {
		body = &CreateOrPatchProjectRequest{}
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", c.Endpoint+"/projects", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject Project
	json.Unmarshal(*responseData, &responseObject)

	return &responseObject, nil
}

func (c *Client) PatchProject(projectId string, body *CreateOrPatchProjectRequest) (*Project, error) {
	if body == nil {
		body = &CreateOrPatchProjectRequest{}
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("PATCH", c.Endpoint+"/projects/"+projectId, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject Project
	json.Unmarshal(*responseData, &responseObject)

	return &responseObject, nil
}

func (c *Client) GetProject(projectId string) (*Project, error) {
	request, err := http.NewRequest("GET", c.Endpoint+"/projects/"+projectId, nil)
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject Project
	json.Unmarshal(*responseData, &responseObject)

	return &responseObject, nil
}

func (c *Client) DeleteProjectById(projectId string) error {
	request, err := http.NewRequest("DELETE", c.Endpoint+"/projects/"+projectId, nil)
	if err != nil {
		return err
	}

	_, err = c.ExecuteRequest(request)
	if err != nil {
		return err
	}

	return nil
}
