package models

type Roster struct {
	Starters   []string  `json:"starters"`
	Settings   Settings  `json:"settings"`
	RosterID   int       `json:"roster_id"`
	Reserve    []string  `json:"reserve"`
	Players    []string  `json:"players"`
	OwnerID    string    `json:"owner_id"`
	LeagueID   string    `json:"league_id"`
}
