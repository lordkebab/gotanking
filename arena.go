package gotanking

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
