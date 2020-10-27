package gotanking

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	//  "en" — English (by default)
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

// PlayerPersonalData holds the response data from the /account/info endpoint
type PlayerPersonalData struct {
	Data map[string]PlayerPersonalDataRec `json:"data"`
}

// PlayerPersonalDataRec is one player record from the /account/info endpoint
type PlayerPersonalDataRec struct {
	ClientLanguage string      `json:"client_language"`
	LastBattleTime int         `json:"last_battle_time"`
	AccountID      int         `json:"account_id"`
	CreatedAt      int         `json:"created_at"`
	UpdatedAt      int         `json:"updated_at"`
	Private        string      `json:"private"`
	GlobalRating   int         `json:"global_rating"`
	ClanID         int         `json:"clan_id"`
	Statistics     PlayerStats `json:"statistics"`
	Nickname       string      `json:"nickname"`
	LogoutAt       int         `json:"logout_at"`
}

// PlayerStats contains the stats for a single player
type PlayerStats struct {
	Clan               map[string]PlayerStatsRec `json:"clan"`
	All                map[string]PlayerStatsRec `json:"all"`
	RegularTeam        map[string]PlayerStatsRec `json:"regular_team"`
	TreesCut           int                       `json:"trees_cut"`
	Company            map[string]PlayerStatsRec `json:"company"`
	StrongholdSkirmish map[string]PlayerStatsRec `json:"stronghold_skirmish"`
	StrongholdDefense  map[string]PlayerStatsRec `json:"stronghold_defense"`
	Historical         map[string]PlayerStatsRec `json:"historical"`
	Team               map[string]PlayerStatsRec `json:"team"`
	Frags              int                       `json:"frags"`
}

// PlayerStatsRec holds the statistics for a single player and single category, returned from the /account/info endpoint
type PlayerStatsRec struct {
	Spotted                    int     `json:"spotted"`
	BattlesOnStunningVehicles  int     `json:"battles_on_stunning_vehicles"`
	AvgDamageBlocked           float32 `json:"average_damage_blocked"`
	DirectHitsReceived         int     `json:"direct_hits_received"`
	ExplosionHits              int     `json:"explosion_hits"`
	PiercingsReceived          int     `json:"piercings_received"`
	Piercings                  int     `json:"piercings"`
	XP                         int     `json:"xp"`
	SurvivedBattles            int     `json:"survived_battles"`
	DroppedCapturePoints       int     `json:"dropped_capture_points"`
	HitsPercents               int     `json:"hits_percents"`
	Draws                      int     `json:"draws"`
	Battles                    int     `json:"battles"`
	DamageReceived             int     `json:"damage_received"`
	AvgDamageAssisted          float32 `json:"avg_damage_assisted"`
	AvgDamageAssistedTrack     float32 `json:"avg_damage_assisted_track"`
	Frags                      int     `json:"frags"`
	StunNumber                 int     `json:"stun_number"`
	AvgDamageAssistedRadio     float32 `json:"avg_damage_assisted_radio"`
	CapturePoints              int     `json:"capture_points"`
	StunAssistedDamage         int     `json:"stun_assisted_damage"`
	Hits                       int     `json:"hits"`
	BattleAvgXP                int     `json:"battle_avg_xp"`
	Wins                       int     `json:"wins"`
	Losses                     int     `json:"losses"`
	DamageDealt                int     `json:"damage_dealt"`
	NoDamageDirectHitsReceived int     `json:"no_damage_direct_hits_received"`
	Shots                      int     `json:"shots"`
	ExplosionHitsReceived      int     `json:"explosion_hits_received"`
	TankingFactor              float32 `json:"tanking_factor"`
}

// PlayerPersonalDataInput holds filters and query parameters for player details
type PlayerPersonalDataInput struct {
	// Access token for private data.
	AccessToken string

	// Extra fields that are added to the response. See https://developers.wargaming.net/reference/all/wot/account/info for full list.
	Extra []string

	// Fields you want displayed. See https://developers.wargaming.net/reference/all/wot/account/info for full list.
	Fields []string
}

// GetAccount fetches a player's account record
func (c *WOTClient) GetAccount(search string, input *AccountInput) (*Account, error) {
	endpoint := "/account/list/"
	var account Account

	v := url.Values{}
	v.Set("application_id", c.ApplicationID)

	v.Set("search", search)
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

// GetAccountID is a convenience method to retrieve a player's account_id
func (c *WOTClient) GetAccountID(search string) int {
	resp, err := c.GetAccount(search, nil)
	if err != nil {
		panic(err)
	}

	return resp.Data[0].AccountID
}

// GetPlayerPersonalData returns player details
func (c *WOTClient) GetPlayerPersonalData(accountID int, input *PlayerPersonalDataInput) (*PlayerPersonalData, error) {
	endpoint := "/account/info/"
	var playerData PlayerPersonalData

	v := url.Values{}
	v.Set("application_id", c.ApplicationID)

	v.Set("account_id", fmt.Sprint(accountID))
	if input != nil {
		v.Set("access_token", input.AccessToken)

		var extra string
		for _, i := range input.Extra {
			extra = extra + "," + i
		}
		v.Set("extra", extra)

		var fields string
		for _, i := range input.Fields {
			fields = fields + "," + i
		}
		v.Set("fields", fields)
	}

	resp, err := http.Get(c.baseURL + endpoint + "?" + v.Encode())
	if err != nil {
		return &playerData, err
	}

	body := new(bytes.Buffer)
	body.ReadFrom(resp.Body)

	b := body.Bytes()

	err = json.Unmarshal(b, &playerData)
	if err != nil {
		return &playerData, err
	}

	return &playerData, nil
}
