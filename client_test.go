package gotanking

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	appID      = "dummy"
	realm      = "na"
	playerName = "dummy"
)

var (
	mux    *http.ServeMux
	client *WOTClient
	server *httptest.Server
)

func TestClientSetup(t *testing.T) {
	// application id should not be nil
	t.Run("application id should not be nil", func(t *testing.T) {
		_, err := NewClient()

		assertError(t, err, ErrNilApplicationID)
	})

	// client sets realm appropriately
	t.Run("client sets realm appropriately", func(t *testing.T) {

		realmTests := []struct {
			realm string
			want  string
		}{
			{realm: "na", want: "https://api.worldoftanks.com/wot/"},
			{realm: "eu", want: "https://api.worldoftanks.eu/wot/"},
			{realm: "ru", want: "https://api.worldoftanks.ru/wot/"},
			{realm: "asia", want: "https://api.worldoftanks.asia/wot/"},
			{realm: "moon", want: "https://api.worldoftanks.com/wot/"},
		}

		for _, tt := range realmTests {
			got, _ := NewClient(SetRealm(tt.realm), SetAppID("dummy"))
			if got.baseURL != tt.want {
				t.Errorf("got %q want %q", got.baseURL, tt.want)
			}
		}
	})

	// default realm is NA
	t.Run("default realm is NA", func(t *testing.T) {
		got, _ := NewClient(SetAppID("dummy"))
		want := "https://api.worldoftanks.com/wot/"

		if got.baseURL != want {
			t.Errorf("got %q want %q", got.baseURL, want)
		}
	})

	// base URL can be changed
	t.Run("base URL can be changed", func(t *testing.T) {
		url := "http://localhost:8080/api/"
		got, _ := NewClient(SetAppID("dummy"), SetBaseURL("http://localhost:8080/api/"))

		if got.baseURL != url {
			t.Errorf("got %q want %q", got.baseURL, url)
		}

	})
}

// when we expect an error
func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected an error here")
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func serverSetup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client, _ = NewClient(SetAppID("dummy"), SetBaseURL(server.URL))

	return func() {
		server.Close()
	}
}

// fixture returns the test data for the path endpoint
func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/" + path)
	if err != nil {
		panic(err)
	}

	return string(b)
}
