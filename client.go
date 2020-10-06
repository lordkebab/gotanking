package gotanking

import (
	"errors"
	"net/http"
	"time"
)

const (
	// BaseURL of the API
	BaseURL string = "https://api.worldoftanks.com/wot/"
	// DefaultClientTimeout defines the client timeout value
	DefaultClientTimeout time.Duration = 5 * time.Second
)

var (
	// ErrNilApplicationID is raised when the client is instantiated without an API key
	ErrNilApplicationID error = errors.New("Application ID cannot be nil")
)

// Option is the function definition for functions overriding defaults
type Option func(*WOTClient) error

// WOTClient is the object to interface with the API
type WOTClient struct {
	client        *http.Client
	ApplicationID string
	baseURL       string
	realm         string
}

// NewClient returns a pointer to a new client object
func NewClient(appID string, opts ...Option) (*WOTClient, error) {

	c := &WOTClient{
		client: &http.Client{
			Timeout: DefaultClientTimeout,
		},
		ApplicationID: appID,
		baseURL:       BaseURL,
		realm:         "na",
	}

	if err := c.parseOpts(opts...); err != nil {
		return nil, err
	}

	if c.ApplicationID == "" {
		return nil, ErrNilApplicationID
	}

	return c, nil
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

// SetAppID sets the API key for the client
func SetAppID(key string) Option {
	return func(c *WOTClient) error {
		c.ApplicationID = key
		return nil
	}
}

// SetBaseURL sets the URL irrespective of the realm. This is used for testing against a test server.
func SetBaseURL(url string) Option {
	return func(c *WOTClient) error {
		c.baseURL = url
		return nil
	}
}

// SetRealm sets the API endpoint to other realms
func SetRealm(realm string) Option {
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

	return func(c *WOTClient) error {
		c.baseURL = url
		c.realm = realm
		return nil
	}
}
