package models

type Transaction struct {
	Type          string         `json:"type"`           // Type of transaction ("trade", "free_agent", etc.)
	TransactionID string         `json:"transaction_id"` // Unique ID for the transaction
	StatusUpdated int64          `json:"status_updated"` // Timestamp of when the status was last updated
	Status        string         `json:"status"`         // Status of the transaction (e.g., "complete")
	Settings      *interface{}   `json:"settings"`       // Nullable, for trades it's null
	RosterIDs     []int          `json:"roster_ids"`     // List of roster IDs involved in the transaction
	Metadata      *interface{}   `json:"metadata"`       // Nullable, additional info (e.g., waiver notes)
	Leg           int            `json:"leg"`            // Week number in football
	Drops         map[string]int `json:"drops"`          // Players dropped (key = player_id, value = roster_id)
	DraftPicks    []DraftPick    `json:"draft_picks"`    // List of draft picks involved in the trade
	Creator       string         `json:"creator"`        // User ID who initiated the transaction
	Created       int64          `json:"created"`        // Timestamp of when the transaction was created
	ConsenterIDs  []int          `json:"consenter_ids"`  // Roster IDs who agreed to the transaction
	Adds          map[string]int `json:"adds"`           // Players added (key = player_id, value = roster_id)
	WaiverBudget  []WaiverBudget `json:"waiver_budget"`  // List of waiver budget transfers (sender, receiver, amount)
}
