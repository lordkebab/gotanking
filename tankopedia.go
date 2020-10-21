package gotanking

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
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

// MapInput holds display filters
type MapInput struct {
	// Fields you want displayed. Valid fields are:
	//
	//	* name_i18n
	//	* description
	//	* camouflage_type
	Fields []string

	// Language in which you want the results. Valid values are:
	//
	// "en" — English (by default)
	//	"ru" — Русский
	//	"pl" — Polski
	//	"de" — Deutsch
	//	"fr" — Français
	//	"es" — Español
	//	"zh-cn" — 简体中文
	//	"zh-tw" — 繁體中文
	//	"tr" — Türkçe
	//	"cs" — Čeština
	//	"th" — ไทย
	//	"vi" — Tiếng Việt
	//	"ko" — 한국어

	Language string
}

// ListMaps queries the encyclopedia/arenas endpoint
func (c *WOTClient) ListMaps(input *MapInput) (*Arena, error) {
	endpoint := "/encyclopedia/arenas/"
	arenas := Arena{}

	v := url.Values{}
	v.Set("application_id", c.ApplicationID)

	if input != nil {
		// encode URL
		v.Set("language", input.Language)

		var fields string
		for _, i := range input.Fields {
			fields = fields + "," + i
		}

		v.Set("fields", fields)
	}

	resp, err := http.Get(c.baseURL + endpoint + "?" + v.Encode())
	if err != nil {
		return &arenas, err
	}

	body := new(bytes.Buffer)
	body.ReadFrom(resp.Body)

	b := body.Bytes()

	err = json.Unmarshal(b, &arenas)
	if err != nil {
		return &arenas, err
	}

	return &arenas, nil
}
