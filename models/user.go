package models

type User struct {
	Id        string `json:"id" `
	Name      string `json:"name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Token     string `json:"token"`
	Confirmed bool   `json:"confirmed"`
}
