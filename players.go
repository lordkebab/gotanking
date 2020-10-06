package gotanking

import "fmt"

// PlayerResponse holds the API response to the /players endpoint
type PlayerResponse struct {
	Status string      `json:"status"`
	Data   []PlayerRec `json:"data"`
}

// PlayerRec is a record for a single player
type PlayerRec struct {
	Nickname  string `json:"nickname"`
	AccountID int    `json:"account_id"`
}

// FormattedPlayer pretty prints a PlayerRec
func (p PlayerRec) FormattedPlayer() string {
	return fmt.Sprintf("Player:  %s\nAccount#:  %d", p.Nickname, p.AccountID)
}
