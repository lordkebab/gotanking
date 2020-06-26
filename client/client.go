package client

import (
	"fmt"
	"net/http"
	"net/url"
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
func NewClient(applicationID string, realm string) *WOTClient {
	return &WOTClient{
		client: &http.Client{
			Timeout: DefaultClientTimeout,
		},
		ApplicationID: applicationID,
		baseURL:       SetRealm(realm),
	}
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

// SetTimeout allows overriding of the default 5 second client timeout
func (w *WOTClient) SetTimeout(d time.Duration) {
	w.client.Timeout = d
}

// OverrideURL overrides the default URL in the client
func (w *WOTClient) OverrideURL(url string) {
	w.baseURL = url
}

// GetPlayer retrieves a player record
func (w *WOTClient) GetPlayer(playerName string) {
	endpoint := "/account/list/?"

	url := url.Values{}
	url.Add("application_id", w.ApplicationID)
	url.Add("search", playerName)
	endpoint = w.baseURL + endpoint + url.Encode()

	request, _ := http.NewRequest(http.MethodPost, endpoint, nil)
	_, err := w.client.Do(request)

	if err != nil {
		fmt.Println(err)
	}
}
