package models

type Income struct {
	Id               string `json:"id"`
	IncomeName       string `json:"income_name"`
	Incomeescription string `json:"income_description"`
	Amount           string `json:"amount"`
	Account_id       string `json:"account_id"`
}
