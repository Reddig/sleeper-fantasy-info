package models

// Waiver represents a waiver budget transfer (FAAB dollars transfer between rosters)
type WaiverBudget struct {
	Sender   int `json:"sender"`   // Roster ID sending the FAAB dollars
	Receiver int `json:"receiver"` // Roster ID receiving the FAAB dollars
	Amount   int `json:"amount"`   // The amount of FAAB dollars transferred
}
