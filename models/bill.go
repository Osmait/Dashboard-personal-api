package models

import "time"

type Bill struct {
	Id              string    `json:"id"`
	BillName        string    `json:"bill_name"`
	BillDescription string    `json:"bill_description"`
	Amount          int       `json:"amount"`
	Account_id      string    `json:"account_id"`
	Created_at      time.Time `json:"created_at"`
}
