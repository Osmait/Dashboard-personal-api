package helpers

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendMail(token string, email string) error {

	USER_EMAIL := os.Getenv("USER_EMAIL")
	PASS_EMAIL := os.Getenv("PASS_EMAIL")

	from := USER_EMAIL
	password := PASS_EMAIL

	toEmailAddress := email
	to := []string{toEmailAddress}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	auth := smtp.PlainAuth("", from, password, host)

	body := fmt.Sprintf("To: %s\r\n"+
		"Subject: Email Confirmed\r\n"+
		"\r\n"+
		"http://127.0.0.1:3000/user/confirmed/%s\r\n", to[0], token)

	message := []byte(body)

	err := smtp.SendMail(address, auth, from, to, message)

	return err

}
