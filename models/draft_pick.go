package models

// DraftPick represents the structure of a draft pick involved in a trade
type DraftPick struct {
	Season          string `json:"season"`           // The season this draft pick belongs to
	Round           int    `json:"round"`            // The round this draft pick is from
	RosterID        int    `json:"roster_id"`        // Roster ID of the original owner
	PreviousOwnerID int    `json:"previous_owner_id"` // Previous owner's roster ID
	OwnerID         int    `json:"owner_id"`         // The new owner of the draft pick
}
