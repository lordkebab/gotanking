package gotanking

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

// Account represents data from the account/list endpoint
type Account struct {
	Data []AccountRecord `json:"data"`
}

// AccountRecord holds one search result from the account/list endpoint
type AccountRecord struct {
	Nickname  string `json:"nickname"`
	AccountID int    `json:"account_id"`
}

// AccountInput holds display filters
type AccountInput struct {
	// Fields you want displayed. Valid values are:
	//
	// * nickname
	// * account_id
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

	// Limit the number of returned entries, up to 100
	Limit string

	// Type of search matching. Valid values are:
	//
	// * startswith (by default)
	// * exact
	SearchType string
}

// GetAccount fetches a player's account record
func (c *WOTClient) GetAccount(search string, input *AccountInput) (*Account, error) {
	endpoint := "/account/list"
	var account Account

	v := url.Values{}
	if input != nil {
		v.Set("language", input.Language)
		v.Set("limit", input.Limit)
		v.Set("type", input.SearchType)

		var fields string
		for _, i := range input.Fields {
			fields = fields + "," + i
		}

		v.Set("fields", fields)
	}

	resp, err := http.Get(c.baseURL + endpoint + "?" + v.Encode())
	if err != nil {
		return &account, err
	}

	body := new(bytes.Buffer)
	body.ReadFrom(resp.Body)

	b := body.Bytes()

	err = json.Unmarshal(b, &account)
	if err != nil {
		return &account, err
	}

	return &account, nil

}
