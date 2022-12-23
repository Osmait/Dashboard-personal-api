package models

type SignUpLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
