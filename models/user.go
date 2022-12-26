package models

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	Confirmed bool   `json:"confirmed"`
}
