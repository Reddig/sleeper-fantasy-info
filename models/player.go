package models

type Player struct {
	PlayerID            string   `json:"player_id"`            // Unique player ID
	Hashtag             string   `json:"hashtag"`              // Hashtag identifier
	DepthChartPosition  *string      `json:"depth_chart_position"` // Depth chart position
	Status             string   `json:"status"`               // Player status (Active, Inactive, etc.)
	Sport              string   `json:"sport"`                // Sport type (e.g., NFL)
	FantasyPositions   []string `json:"fantasy_positions"`    // List of fantasy positions (e.g., ["QB"])
	Number             int      `json:"number"`               // Jersey number
	SearchLastName     string   `json:"search_last_name"`     // Last name for search
	InjuryStartDate    *string  `json:"injury_start_date"`    // Nullable injury start date
	Weight             string   `json:"weight"`               // Weight (e.g., "220")
	Position           string   `json:"position"`             // Player's position (e.g., "QB")
	PracticeParticipation *string `json:"practice_participation"` // Nullable participation status
	SportradarID       string   `json:"sportradar_id"`        // Sportradar ID
	Team               string   `json:"team"`                 // Team abbreviation (e.g., "NE" for New England)
	LastName           string   `json:"last_name"`            // Last name of the player
	College            string   `json:"college"`              // College the player attended
	FantasyDataID      int      `json:"fantasy_data_id"`      // Fantasy data ID
	InjuryStatus       *string  `json:"injury_status"`        // Nullable injury status
	Height             string   `json:"height"`               // Height (e.g., "6'4\"")
	SearchFullName     string   `json:"search_full_name"`     // Full name for search
	Age                int      `json:"age"`                  // Age of the player
	StatsID            int   `json:"stats_id"`             // Stats ID (if applicable)
	BirthCountry       string   `json:"birth_country"`        // Country of birth
	ESPNID             int   `json:"espn_id"`              // ESPN ID
	SearchRank         int      `json:"search_rank"`          // Search rank
	FirstName          string   `json:"first_name"`           // First name
	DepthChartOrder    int      `json:"depth_chart_order"`    // Order in the depth chart
	YearsExperience    int      `json:"years_exp"`            // Years of experience
	RotowireID         *int  `json:"rotowire_id"`          // Nullable Rotowire ID
	RotoworldID        int      `json:"rotoworld_id"`         // Rotoworld ID
	SearchFirstName    string   `json:"search_first_name"`    // First name for search
	YahooID            *int  `json:"yahoo_id"`             // Nullable Yahoo ID
	Transactions 		map[string]Transaction
}