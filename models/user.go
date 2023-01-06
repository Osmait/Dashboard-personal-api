package models

type User struct {
	Id        string `json:"id" `
	Name      string `json:"name" validate:"required,max=50"`
	LastName  string `json:"last_name" validate:"required,max=50"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,max=50,min=6"`
	Token     string `json:"token"`
	Confirmed bool   `json:"confirmed"`
}
