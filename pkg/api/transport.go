package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// AnetTransport wraps an http.RoundTripper and injects Atlantic.Net HMAC auth
type AnetTransport struct {
	Transport http.RoundTripper
	Config    *Config
}

// signatureRequest generates the HMAC-SHA256 signature for a given string
func (at *AnetTransport) signatureRequest(stringToSign string) (string, error) {
	if at == nil {
		return "", fmt.Errorf("transport is nil")
	}
	if at.Config == nil {
		return "", fmt.Errorf("transport config is nil")
	}
	if at.Config.Secret == "" {
		return "", fmt.Errorf("transport config secret is empty")
	}

	h := hmac.New(sha256.New, []byte(at.Config.Secret))
	_, _ = h.Write([]byte(stringToSign)) // Write never errors
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return strings.TrimRight(signature, "\n"), nil
}

// RoundTrip executes a single HTTP transaction and adds custom authentication
func (at *AnetTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if at.Transport == nil {
		at.Transport = http.DefaultTransport
	}

	if at.Config == nil {
		return nil, fmt.Errorf("AnetTransport.Config is nil")
	}

	// Generate the timestamp and random UUID
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	randomGUID := uuid.New().String()

	// Build the string to sign and compute signature
	stringToSign := timestamp + randomGUID
	signature, err := at.signatureRequest(stringToSign)
	if err != nil {
		return nil, err
	}

	// Clone the request to avoid mutating caller's request
	reqCopy := req.Clone(req.Context())

	// Parse existing query params
	urlParams, err := url.ParseQuery(reqCopy.URL.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to parse query params: %w", err)
	}

	// Add authentication params
	urlParams.Set("ACSAccessKeyId", at.Config.Key)
	urlParams.Set("Timestamp", timestamp)
	urlParams.Set("Rndguid", randomGUID)
	urlParams.Set("Signature", signature)

	// Rebuild query string
	reqCopy.URL.RawQuery = urlParams.Encode()

	// Pass request down the chain
	return at.Transport.RoundTrip(reqCopy)
}
