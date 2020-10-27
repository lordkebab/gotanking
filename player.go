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

// PlayerVehicles contains all tanks owned by a player
type PlayerVehicles struct {
	Data map[string][]PlayerVehiclesRec
}

// PlayerVehiclesRec contains a record of a particular vehicle
type PlayerVehiclesRec struct {
	Statistics struct {
		Wins    int `json:"wins"`
		Battles int `json:"battles"`
	}
	MarkOfMastery int `json:"mark_of_mastery"`
	TankID        int `json:"tank_id"`
}

// PlayerVehiclesInput holds filters and query parameters for vehicle details
type PlayerVehiclesInput struct {
	// Access token for private data.
	AccessToken string

	// Fields you want displayed. Valid values are:
	//
	// * mark_of_mastery (0: none, 1: 3rd class, 2: 2nd class, 3: 1st class, 4: ace tanker)
	// * tank_id
	// * statistics
	// * statistics.battles
	// * statistics.wins
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

	// Tank or Tanks ID
	TankID []int
}

// PlayerAchievements holds a player's achievements
type PlayerAchievements struct {
	Data map[string]PlayerAchievementsData `json:"data"`
}

// PlayerAchievementsData contains all the achievement categories
type PlayerAchievementsData struct {
	Achievements PlayerAchievementsRec `json:"achievements"`
	Frags        PlayerFragsRec        `json:"frags"`
	MaxSeries    PlayerMaxSeriesRec    `json:"max_series"`
}

// PlayerAchievementsRec contains all earned achievements
type PlayerAchievementsRec struct {
	CariusMedal                   int `json:"medalCarius"`
	Aimer                         int `json:"aimer"`
	Invader                       int `json:"invader"`
	ShootToKill                   int `json:"shootToKill"`
	June20                        int `json:"june20"`
	MedalEkins                    int `json:"medalEkins"`
	WtHunterWins                  int `json:"wtHunterWins"`
	Duelist                       int `json:"duelist"`
	TitleSniper                   int `json:"titleSniper"`
	LeClercMedal                  int `json:"medalLeClerc"`
	Demolition                    int `json:"demolition"`
	Supporter                     int `json:"supporter"`
	SteelWall                     int `json:"steelwall"`
	LehvaslaihoMedal              int `json:"medalLehvaslaiho"`
	AbramsMedal                   int `json:"medalAbrams"`
	September20                   int `json:"september20"`
	PoppelMedal                   int `json:"medalPoppel"`
	PascucciMedal                 int `json:"medalPascucci"`
	ReliableComrade               int `json:"reliableComrade"`
	MainGun                       int `json:"mainGun"`
	Kamikaze                      int `json:"kamikaze"`
	TenYearsCoundownParticipation int `json:"TenYearsCountdownParticipation"`
	TenYearsCountdownStageMedal   int `json:"TenYearsCountdownStageMedal"`
	FirstMerit                    int `json:"firstMerit"`
	Warrior                       int `json:"warrior"`
	ArmorPiercer                  int `json:"armorPiercer"`
	Even                          int `json:"even"`
	WtBossWins                    int `json:"wtBossWins"`
	Arsonist                      int `json:"arsonist"`
	BrothersInArmMedal            int `json:"medalBrothersInArms"`
	IronMan                       int `json:"ironMan"`
	ReadyForBattleMT              int `json:"readyForBattleMT"`
	Sniper2                       int `json:"sniper2"`
	OskinMedal                    int `json:"medalOskin"`
	Charmed                       int `json:"charmed"`
	October20                     int `json:"october20"`
	Fighter                       int `json:"fighter"`
	LavrinenkoMedal               int `json:"medalLavrinenko"`
	Impenetrable                  int `json:"impenetrable"`
	Sturdy                        int `json:"sturdy"`
	KayMedal                      int `json:"medalKay"`
	FiveYearsOfService            int `json:"05YearsOfService"`
	HandOfDeath                   int `json:"handOfDeath"`
	BoneCrusher                   int `json:"bonecrusher"`
	WFC2014                       int `json:"WFC2014"`
	DumitruMedal                  int `json:"medalDumitru"`
	EvilEye                       int `json:"evileye"`
	KnispelMedal                  int `json:"medalKnispel"`
}

// PlayerFragsRec contains achievement progress
type PlayerFragsRec struct {
	CrucialShotMedal         int `json:"crucialShotMedal"`
	InfiltratorMedal         int `json:"infiltratorMedal"`
	SentinelMedal            int `json:"sentinelMedal"`
	PrematureDetonationMedal int `json:"prematureDetonationMedal"`
	FightingReconMedal       int `json:"fightingReconnaissanceMedal"`
	FireAndSteelMedal        int `json:"fireAndSteelMedal"`
	RangerMedal              int `json:"rangerMedal"`
	ReliableComrade          int `json:"reliableComrade"`
	WolfAmongSheepMedal      int `json:"wolfAmongSheepMedal"`
	HeavyFireMedal           int `json:"heavyFireMedal"`
	BruteForceMedal          int `json:"bruteForceMedal"`
	GuerrillaMedal           int `json:"guerrilaMedal"`
	PromisingFighterMedal    int `json:"promisingFighterMedal"`
	PyromaniacMedal          int `json:"pyromaniacMedal"`
	GeniusForWarMedal        int `json:"geniusForWarMedal"`
	Sinai                    int `json:"sinai"`
	BeastHunter              int `json:"beastHunter"`
	PattonValley             int `json:"pattonValley"`
}

// PlayerAchievementsInput holds filters and query parameters for player achievements
type PlayerAchievementsInput struct {
	// Fields you want displayed. Valid values are:
	//
	// * achievements
	// * frags
	// * max_series
	Fields []string

	// Language for the results
	Language string
}

// PlayerMaxSeriesRec contains the maximum achievement series values
type PlayerMaxSeriesRec struct {
	ArmorPiercer         int `json:"armorPiercer"`
	Aimer                int `json:"aimer"`
	TitleSniper          int `json:"titleSniper"`
	TacticalBreakthrough int `json:"tacticalBreakthrough"`
	Invincible           int `json:"invincible"`
	VictoryMarch         int `json:"victoryMarch"`
	DeathTrack           int `json:"deathTrack"`
	EFC2016              int `json:"EFC2016"`
	DieHard              int `json:"dieHard"`
	WFC2014              int `json:"WFC2014"`
	HandOfDeath          int `json:"handOfDeath"`
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

// GetPlayerVehicles returns all tanks owned by a player
func (c *WOTClient) GetPlayerVehicles(accountID int, input *PlayerVehiclesInput) (*PlayerVehicles, error) {
	endpoint := "/account/tanks/"
	var playerVehicles PlayerVehicles

	v := url.Values{}
	v.Set("application_id", c.ApplicationID)

	v.Set("account_id", fmt.Sprint(accountID))
	if input != nil {
		v.Set("access_token", input.AccessToken)

		var tanks string
		for _, i := range input.TankID {
			tanks = tanks + "," + fmt.Sprint(i)
		}
		v.Set("tank_id", tanks)

		var fields string
		for _, i := range input.Fields {
			fields = fields + "," + i
		}
		v.Set("fields", fields)

		v.Set("language", input.Language)
	}

	resp, err := http.Get(c.baseURL + endpoint + "?" + v.Encode())
	if err != nil {
		return &playerVehicles, err
	}

	body := new(bytes.Buffer)
	body.ReadFrom(resp.Body)

	b := body.Bytes()

	err = json.Unmarshal(b, &playerVehicles)
	if err != nil {
		return &playerVehicles, err
	}

	return &playerVehicles, nil
}

// GetPlayerAchievements returns player achievements
func (c *WOTClient) GetPlayerAchievements(accountID int, input *PlayerAchievementsInput) (*PlayerAchievements, error) {
	endpoint := "/account/achievements/"
	var playerAchievements PlayerAchievements

	v := url.Values{}
	v.Set("application_id", c.ApplicationID)

	v.Set("account_id", fmt.Sprint(accountID))
	if input != nil {

		var fields string
		for _, i := range input.Fields {
			fields = fields + "," + i
		}
		v.Set("fields", fields)

		v.Set("language", input.Language)
	}

	resp, err := http.Get(c.baseURL + endpoint + "?" + v.Encode())
	if err != nil {
		return &playerAchievements, err
	}

	body := new(bytes.Buffer)
	body.ReadFrom(resp.Body)

	b := body.Bytes()

	err = json.Unmarshal(b, &playerAchievements)
	if err != nil {
		return &playerAchievements, err
	}

	return &playerAchievements, nil
}
