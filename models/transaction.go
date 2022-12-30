package models

import "time"

type Transaction struct {
	Id             string    `json:"id"`
	Name           string    `json:"transaction_name"`
	Description    string    `json:"transaction_description"`
	Amount         float64   `json:"amount"`
	TypeTransation string    `json:"type_transation"`
	UserId         string    `json:"user_id"`
	Account_id     string    `json:"account_id"`
	Created_at     time.Time `json:"created_at"`
}
