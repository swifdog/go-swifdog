package swifdog

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type RegistryCredential struct {
	ID               string `json:"id"`
	RegistryURL      string `json:"registryUrl"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	CreationDateTime string `json:"creationDateTime"`
}

type CreateOrPatchRegistryCredentialRequest struct {
	RegistryURL string `json:"registryUrl"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

func (c *Client) ListRegistryCredential(projectId string) ([]RegistryCredential, error) {
	request, err := http.NewRequest("GET", c.Endpoint+"/projects/"+projectId+"/registrycredentials", nil)
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject []RegistryCredential
	json.Unmarshal(*responseData, &responseObject)

	return responseObject, nil
}

func (c *Client) CreateRegistryCredential(projectId string, body *CreateOrPatchRegistryCredentialRequest) (*RegistryCredential, error) {
	if body == nil {
		body = &CreateOrPatchRegistryCredentialRequest{}
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", c.Endpoint+"/projects/"+projectId+"/registrycredentials", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject RegistryCredential
	json.Unmarshal(*responseData, &responseObject)

	return &responseObject, nil
}

func (c *Client) PatchRegistryCredential(projectId string, registryCredentialId string, body *CreateOrPatchRegistryCredentialRequest) (*RegistryCredential, error) {
	if body == nil {
		body = &CreateOrPatchRegistryCredentialRequest{}
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("PATCH", c.Endpoint+"/projects/"+projectId+"/registrycredentials/"+registryCredentialId, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject RegistryCredential
	json.Unmarshal(*responseData, &responseObject)

	return &responseObject, nil
}

func (c *Client) GetRegistryCredential(projectId string, registryCredentialId string) (*RegistryCredential, error) {
	request, err := http.NewRequest("GET", c.Endpoint+"/projects/"+projectId+"/registrycredentials/"+registryCredentialId, nil)
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject RegistryCredential
	json.Unmarshal(*responseData, &responseObject)

	return &responseObject, nil
}

func (c *Client) DeleteRegistryCredentialById(projectId string, registryCredentialId string) error {
	request, err := http.NewRequest("DELETE", c.Endpoint+"/projects/"+projectId+"/registrycredentials/"+registryCredentialId, nil)
	if err != nil {
		return err
	}

	_, err = c.ExecuteRequest(request)
	if err != nil {
		return err
	}

	return nil
}
