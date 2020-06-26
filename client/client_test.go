package client

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

const (
	appID      = "dummy"
	realm      = "na"
	playerName = "dummy"
)

func TestClientSetup(t *testing.T) {
	t.Run("Returns a pointer to a new client object", func(t *testing.T) {
		t.Helper()

		got := NewClient("dummy", "na")
		want := &WOTClient{
			client: &http.Client{
				Timeout: DefaultClientTimeout,
			},
			ApplicationID: "dummy",
			baseURL:       "https://api.worldoftanks.com/wot/",
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %+v\n want %+v\n", got, want)
		}
	})

	t.Run("Set a different timeout than the default", func(t *testing.T) {
		t.Helper()

		client := NewClient("dummy", "na")
		client.SetTimeout(10 * time.Second)

		got := client.client.Timeout
		want := 10 * time.Second

		if got != want {
			t.Errorf("Got %d want %d", got, want)
		}
	})

	t.Run("Set a different realm", func(t *testing.T) {
		t.Helper()

		var realms = []struct {
			realm string
			want  string
		}{
			{"na", "https://api.worldoftanks.com/wot/"},
			{"eu", "https://api.worldoftanks.eu/wot/"},
			{"ru", "https://api.worldoftanks.ru/wot/"},
			{"asia", "https://api.worldoftanks.asia/wot/"},
		}

		for _, tt := range realms {
			t.Run(tt.realm, func(t *testing.T) {
				client := NewClient("dummy", tt.realm)
				got := client.baseURL

				if got != tt.want {
					t.Errorf("got %s want %s", got, tt.want)
				}
			})
		}
	})

	t.Run("Override the API URL", func(t *testing.T) {
		t.Helper()

		client := NewClient("dummy", "na")
		want := "https://example.com"
		client.OverrideURL(want)
		got := client.baseURL

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

func TestAccountEndpoints(t *testing.T) {

	t.Run("should POST data", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("Expected `POST` got `%s`", r.Method)
			}
		}))

		client := NewClient(appID, realm)
		client.OverrideURL(ts.URL)
		client.GetPlayer(playerName)
	})

	t.Run("should hit the correct URL", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			want := "/account/list/?application_id=dummy&search=dummy"
			if r.URL.String() != want {
				t.Errorf("Expected %s got %s", want, r.URL.String())
			}
		}))

		client := NewClient(appID, realm)
		client.OverrideURL(ts.URL)
		client.GetPlayer(playerName)
	})
}
