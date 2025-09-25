package api

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	// The configurations to bind to the client
	config *Config
}

func NewClient(cfg *Config) *Client {
	return &Client{config: cfg}
}

func (c *Client) send(request Request) (*http.Response, error) {
	// Build the Request
	req, err := request.Build()
	if err != nil {
		return nil, err
	}

	// Send the HTTP request
	httpClient := &http.Client{
		Timeout:   time.Duration(c.config.Timeout) * time.Second,
		Transport: &AnetTransport{Config: c.config},
	}
	return httpClient.Do(req)
}

func (c *Client) RequestRaw(action string, params map[string]string) ([]byte, error) {
	req := Request{
		Action:     action,
		Parameters: params,
		Config:     c.config,
	}

	r, err := c.send(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	byteResp, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return byteResp, err
}

func (c *Client) Request(action string, params map[string]string) (string, error) {
	byteResp, err := c.RequestRaw(action, params)
	if err != nil {
		return "", err
	}

	resp := fmt.Sprintf("%s\n", string(byteResp))

	return resp, nil
}
