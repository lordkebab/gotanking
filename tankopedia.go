package gotanking

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Arena represents data from the encyclopedia/arenas endpoint
type Arena struct {
	Data map[string]ArenaRecord `json:"data"`
}

// ArenaRecord represents a single arena record
type ArenaRecord struct {
	Name string `json:"name_i18n"`
	Camo string `json:"camouflage_type"`
	Desc string `json:"description"`
	ID   string `json:"arena_id"`
}

// ListMaps queries the encyclopedia/arenas endpoint
func (c *WOTClient) ListMaps() (Arena, error) {
	endpoint := "/encyclopedia/arenas"
	arenas := Arena{}

	resp, err := http.Get(c.baseURL + endpoint)
	if err != nil {
		return arenas, err
	}

	body := new(bytes.Buffer)
	body.ReadFrom(resp.Body)
	b := body.Bytes()

	err = json.Unmarshal(b, &arenas)
	if err != nil {
		return Arena{}, err
	}

	return arenas, nil
}
