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

// Option is the function definition for functions overriding defaults
type Option func(*Client) error

// WOTClient is the object to interface with the API
type WOTClient struct {
	client        *http.Client
	ApplicationID string
	baseURL       string
}

// NewClient returns a pointer to a new client object
func NewClient(opts ...Option) (*WOTClient, error) {

	client := &WOTClient{
		client: &http.Client{
			Timeout: DefaultClientTimeout,
		},
		ApplicationID: "nil",
		baseURL:       BaseURL,
	}

	if err := client.parseOpts(opts...); err != nil {
		return nil, err
	}

	return client, nil

}

// parseOpts overrides instantiated defaults
func (c *WOTClient) parseOpts(opts ...Option) error {
	// range over each option (function)
	// overriding defaults in sequence
	for _, option := range opts {
		err := option(c)
		if err != nil {
			return err
		}
	}

	return nil
}

// SetRealm sets the API endpoint to other realms
func SetRealm(realm string) string {
	var url string

	switch realm {
	case "na":
		url = "https://api.worldoftanks.com/wot/"
	case "eu":
		url = "https://api.worldoftanks.eu/wot/"
	case "ru":
		url = "https://api.worldoftanks.ru/wot/"
	case "asia":
		url = "https://api.worldoftanks.asia/wot/"
	default:
		url = "https://api.worldoftanks.com/wot/"
	}

	return url
}
