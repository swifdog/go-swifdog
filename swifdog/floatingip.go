package swifdog

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type FloatingIP struct {
	ID string `json:"id"`
	IP string `json:"ip"`
	Version int `json:"version"`
	Endpoints []FloatingIPEndpoint `json:"endpoints"`
}

type FloatingIPEndpoint struct {
	PacketId string `json:"packetId"`
	Protocol string `json:"protocol"`
	ContainerPort int `json:"containerPort"`
	TargetPort int `json:"targetPort"`
}

func (c *Client) ListFloatingIPs(projectId string) ([]FloatingIP, error) {
	request, err := http.NewRequest("GET", c.Endpoint+"/projects/"+projectId+"/floatingips", nil)
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject []FloatingIP
	json.Unmarshal(*responseData, &responseObject)

	return responseObject, nil
}

func (c *Client) CreateFloatingIP(projectId string, body *FloatingIP) (*FloatingIP, error) {
	if body == nil {
		body = &FloatingIP{}
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", c.Endpoint+"/projects/"+projectId+"/floatingips", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject FloatingIP
	json.Unmarshal(*responseData, &responseObject)

	return &responseObject, nil
}

func (c *Client) PatchFloatingIP(projectId string, floatingIpId string, body *FloatingIP) (*FloatingIP, error) {
	if body == nil {
		body = &FloatingIP{}
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("PATCH", c.Endpoint+"/projects/"+projectId+"/floatingips/"+floatingIpId, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject FloatingIP
	json.Unmarshal(*responseData, &responseObject)

	return &responseObject, nil
}

func (c *Client) GetFloatingIP(projectId string, floatingIpId string) (*FloatingIP, error) {
	request, err := http.NewRequest("GET", c.Endpoint+"/projects/"+projectId+"/floatingips/"+floatingIpId, nil)
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject FloatingIP
	json.Unmarshal(*responseData, &responseObject)

	return &responseObject, nil
}

func (c *Client) DeleteFloatingIPById(projectId string, floatingIpId string) error {
	request, err := http.NewRequest("DELETE", c.Endpoint+"/projects/"+projectId+"/floatingips/"+floatingIpId, nil)
	if err != nil {
		return err
	}

	_, err = c.ExecuteRequest(request)
	if err != nil {
		return err
	}

	return nil
}
