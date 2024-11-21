package models

type League struct {
	TotalRosters     int                    `json:"total_rosters"`
	Status           string                 `json:"status"`
	Sport            string                 `json:"sport"`
	Settings         map[string]interface{} `json:"settings"` // Replace with specific type if known
	SeasonType       string                 `json:"season_type"`
	Season           string                 `json:"season"`
	ScoringSettings  map[string]interface{} `json:"scoring_settings"` // Replace with specific type if known
	RosterPositions  []string               `json:"roster_positions"` // Replace with specific type if array elements are objects
	PreviousLeagueID string                 `json:"previous_league_id"`
	Name             string                 `json:"name"`
	LeagueID         string                 `json:"league_id"`
	DraftID          string                 `json:"draft_id"`
	Avatar           string                 `json:"avatar"`
}
