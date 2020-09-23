package client

import (
	"testing"
)

const (
	appID      = "dummy"
	realm      = "na"
	playerName = "dummy"
)

func TestClientSetup(t *testing.T) {
	// application id should not be nil

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
			got, _ := NewClient(SetRealm(tt.realm))
			if got.baseURL != tt.want {
				t.Errorf("got %q want %q", got.baseURL, tt.want)
			}
		}
	})

	// default realm is NA
}
