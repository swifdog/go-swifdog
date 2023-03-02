package swifdog

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Packet struct {
	ID                   string                  `json:"id"`
	Name                 string                  `json:"name"`
	Image                string                  `json:"image"`
	RegistryCredentialId string                  `json:"registryCredentialId"`
	EnvironmentVariables []EnvironmentVariable   `json:"envs"`
	VolumeMounts         []PersistentVolumeMount `json:"volumeMounts"`
	CreationDateTime     string                  `json:"creationDateTime"`
}

type EnvironmentVariable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PersistentVolumeMount struct {
	VolumeId   string `json:"volumeId"`
	VolumeName string `json:"volumeName"`
	MountPath  string `json:"mountPath"`
}

type CreateOrPatchPacketRequest struct {
	Name                 string                  `json:"name"`
	Image                string                  `json:"image"`
	RegistryCredentialId string                  `json:"registryCredentialId"`
	EnvironmentVariables []EnvironmentVariable   `json:"envs"`
	VolumeMounts         []PersistentVolumeMount `json:"volumeMounts"`

	Packet
}

func (pv *PersistentVolume) asMount(mountPath string) *PersistentVolumeMount {
	return &PersistentVolumeMount{
		VolumeId:  pv.ID,
		MountPath: mountPath,
	}
}

func (c *Client) ListPackets(projectId string) ([]Packet, error) {
	request, err := http.NewRequest("GET", c.Endpoint+"/projects/"+projectId+"/packets", nil)
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject []Packet
	json.Unmarshal(*responseData, &responseObject)

	return responseObject, nil
}

func (c *Client) CreatePacket(projectId string, body *CreateOrPatchPacketRequest) (*Packet, error) {
	if body == nil {
		body = &CreateOrPatchPacketRequest{}
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", c.Endpoint+"/projects/"+projectId+"/packets", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject Packet
	json.Unmarshal(*responseData, &responseObject)

	return &responseObject, nil
}

func (c *Client) PatchPacket(projectId string, packetId string, body *CreateOrPatchPacketRequest) (*Packet, error) {
	if body == nil {
		body = &CreateOrPatchPacketRequest{}
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("PATCH", c.Endpoint+"/projects/"+projectId+"/packets/"+packetId, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject Packet
	json.Unmarshal(*responseData, &responseObject)

	return &responseObject, nil
}

func (c *Client) GetPacket(projectId string, packetId string) (*Packet, error) {
	request, err := http.NewRequest("GET", c.Endpoint+"/projects/"+projectId+"/packets/"+packetId, nil)
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject Packet
	json.Unmarshal(*responseData, &responseObject)

	return &responseObject, nil
}

func (c *Client) DeletePacketById(projectId string, packetId string) error {
	request, err := http.NewRequest("DELETE", c.Endpoint+"/projects/"+projectId+"/packets/"+packetId, nil)
	if err != nil {
		return err
	}

	_, err = c.ExecuteRequest(request)
	if err != nil {
		return err
	}

	return nil
}
