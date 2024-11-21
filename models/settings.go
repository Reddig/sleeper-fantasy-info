package models

type Settings struct {
	Wins              int     `json:"wins"`
	WaiverPosition    int     `json:"waiver_position"`
	WaiverBudgetUsed  int     `json:"waiver_budget_used"`
	TotalMoves        int     `json:"total_moves"`
	Ties              int     `json:"ties"`
	Losses            int     `json:"losses"`
	FptsDecimal       float64 `json:"fpts_decimal"`
	FptsAgainstDecimal float64 `json:"fpts_against_decimal"`
	FptsAgainst       int     `json:"fpts_against"`
	Fpts              int     `json:"fpts"`
}