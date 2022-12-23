package models

type Bill struct {
	Id              string `json:"id"`
	BillName        string `json:"bill_name"`
	BillDescription string `json:"bill_description"`
	Amount          string `json:"amount"`
	Account_id      string `json:"account_id"`
}
