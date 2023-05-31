package swifdog

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type IngressRule struct {
	ID            string            `json:"id"`
	Hostname      string            `json:"hostname"`
	AutoManageSSL bool              `json:"autoManageSSL"`
	PathRules     []IngressRulePath `json:"pathRules"`
}

type IngressRulePath struct {
	Path          string `json:"path"`
	PacketId      string `json:"packetId"`
	ContainerPort int    `json:"containerPort"`
}

func (c *Client) ListIngressRule(projectId string) ([]IngressRule, error) {
	request, err := http.NewRequest("GET", c.Endpoint+"/projects/"+projectId+"/ingressrules", nil)
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject []IngressRule
	json.Unmarshal(*responseData, &responseObject)

	return responseObject, nil
}

func (c *Client) CreateIngressRule(projectId string, body *IngressRule) (*IngressRule, error) {
	if body == nil {
		body = &IngressRule{}
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", c.Endpoint+"/projects/"+projectId+"/ingressrules", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject IngressRule
	json.Unmarshal(*responseData, &responseObject)

	return &responseObject, nil
}

func (c *Client) PatchIngressRule(projectId string, ingressRuleId string, body *IngressRule) (*IngressRule, error) {
	if body == nil {
		body = &IngressRule{}
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("PATCH", c.Endpoint+"/projects/"+projectId+"/ingressrules/"+ingressRuleId, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject IngressRule
	json.Unmarshal(*responseData, &responseObject)

	return &responseObject, nil
}

func (c *Client) GetIngressRule(projectId string, ingressRuleId string) (*IngressRule, error) {
	request, err := http.NewRequest("GET", c.Endpoint+"/projects/"+projectId+"/ingressrules/"+ingressRuleId, nil)
	if err != nil {
		return nil, err
	}

	responseData, err := c.ExecuteRequest(request)
	if err != nil {
		return nil, err
	}

	var responseObject IngressRule
	json.Unmarshal(*responseData, &responseObject)

	return &responseObject, nil
}

func (c *Client) DeleteIngressRuleById(projectId string, ingressRuleId string) error {
	request, err := http.NewRequest("DELETE", c.Endpoint+"/projects/"+projectId+"/ingressrules/"+ingressRuleId, nil)
	if err != nil {
		return err
	}

	_, err = c.ExecuteRequest(request)
	if err != nil {
		return err
	}

	return nil
}
