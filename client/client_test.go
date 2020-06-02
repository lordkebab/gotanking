package client

import (
	"net/http"
	"reflect"
	"testing"
)

func TestClient(t *testing.T) {
	t.Run("Returns a pointer to a new client object", func(t *testing.T) {
		t.Helper()

		got := NewClient("dummy")
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
}
