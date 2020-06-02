package client

import (
	"net/http"
	"time"
)

const (
	// BaseURL of the API
	BaseURL string = "https://api.worldoftanks.com/wot/"
	// DefaultClientTimeout defines the client timeout value
	DefaultClientTimeout time.Duration = 5 * time.Second
)

// WOTClient is the object to interface with the API
type WOTClient struct {
	client        *http.Client
	ApplicationID string
	baseURL       string
}

// NewClient returns a pointer to a new client object
func NewClient(applicationID string) *WOTClient {
	return &WOTClient{
		client: &http.Client{
			Timeout: DefaultClientTimeout,
		},
		ApplicationID: applicationID,
		baseURL:       BaseURL,
	}
}
