package models

type Account struct {
	Id      string  `json:"id"`
	Name    string  `json:"name" validate:"required"`
	Bank    string  `json:"bank" validate:"required"`
	Balance float64 `json:"balance" validate:"required"`
	User_id string  `json:"user_id"`
}
