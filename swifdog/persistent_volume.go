package swifdog

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type PersistentVolume struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Capacity         string `json:"capacity"`
	CreationDateTime string `json:"creationDateTime"`
}

type CreateOrPatchPersistentVolumeRequest struct {
	Name     string `json:"name"`
	Capacity string `json:"capacity"`
}

func (c *Client) ListPersistentVolume(projectId string) ([]PersistentVolume, error) {
	request, err := http.NewRequest("GET", c.Endpoint+"/projects/"+projectId+"/volumes", nil)
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject []PersistentVolume
	json.Unmarshal(*responseData, &responseObject)

	return responseObject, nil
}

func (c *Client) CreatePersistentVolume(projectId string, body *CreateOrPatchPersistentVolumeRequest) (*PersistentVolume, error) {
	if body == nil {
		body = &CreateOrPatchPersistentVolumeRequest{}
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", c.Endpoint+"/projects/"+projectId+"/volumes", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject PersistentVolume
	json.Unmarshal(*responseData, &responseObject)

	return &responseObject, nil
}

func (c *Client) PatchPersistentVolume(projectId string, persistentVolumeId string, body *CreateOrPatchPersistentVolumeRequest) (*PersistentVolume, error) {
	if body == nil {
		body = &CreateOrPatchPersistentVolumeRequest{}
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("PATCH", c.Endpoint+"/projects/"+projectId+"/volumes/"+persistentVolumeId, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject PersistentVolume
	json.Unmarshal(*responseData, &responseObject)

	return &responseObject, nil
}

func (c *Client) DeletePersistentVolumeById(projectId string, persistentVolumeId string) error {
	request, err := http.NewRequest("DELETE", c.Endpoint+"/projects/"+projectId+"/volumes/"+persistentVolumeId, nil)
	if err != nil {
		return err
	}

	_, err = c.ExecuteRequest(request)
	if err != nil {
		return err
	}

	return nil
}
