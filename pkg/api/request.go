package api

import (
	"fmt"
	"net/http"
	"net/url"
)

type Request struct {
	Config     *Config
	Action     string
	Parameters map[string]string
}

func (r *Request) buildParams() (url.Values, error) {
	if r.Action == "" {
		return nil, fmt.Errorf("no action given")
	}
	if r.Config == nil {
		return nil, fmt.Errorf("config is nil")
	}

	params := url.Values{}

	// Set the base required parameters
	params.Set("Action", r.Action)
	params.Set("Format", "json")
	params.Set("Version", r.Config.Version)

	// Add any additional parameters
	for key, value := range r.Parameters {
		if value != "" {
			params.Set(key, value)
		}
	}

	return params, nil
}

func (r *Request) Build() (*http.Request, error) {
	params, err := r.buildParams()
	if err != nil {
		return nil, err
	}

	// Validate config.URL
	if r.Config.URL == "" {
		return nil, fmt.Errorf("Config.URL is empty")
	}

	// Construct the full URL with query parameters
	newURL := fmt.Sprintf("%s?%s", r.Config.URL, params.Encode())

	// Build the HTTP Request
	req, err := http.NewRequest("GET", newURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
