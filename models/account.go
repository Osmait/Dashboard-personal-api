package models

type Account struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Bank    string `json:"bank"`
	Balance int    `json:"balance"`
	User_id string `json:"user_id"`
}
