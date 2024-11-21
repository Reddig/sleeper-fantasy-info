package models

type User struct {
	Username string `json:"username"`
	UserID string `json:"user_id"`
	DisplayName string `json:"display_name"`
	Avatar string `json:"avatar"`
	Transactions 		[]Transaction
}