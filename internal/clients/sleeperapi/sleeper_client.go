package sleeperapi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type SleeperClient struct {
    BaseURL    string
    HTTPClient *http.Client
}

func NewClient(baseURL string) *SleeperClient {
    return &SleeperClient{
        BaseURL:    baseURL,
        HTTPClient: &http.Client{},
    }
}

func (c *SleeperClient) MakeRequest(endpoint string, method string, payload interface{}) (*http.Response, error) {
    url := c.BaseURL + endpoint
    var body []byte
    var err error
    if payload != nil {
        body, err = json.Marshal(payload)
        if err != nil {
            return nil, err
        }
    }
	
    req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if method == "GET" {
        req.Body = nil // No body for GET requests
    }
    if err != nil {
        return nil, err
    }
    req.Header.Set("Content-Type", "application/json")

    return c.HTTPClient.Do(req)
}