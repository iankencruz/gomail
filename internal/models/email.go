package models

import (
	"database/sql"
	"fmt"
	"net/smtp"
	"time"
)

type Email struct {
	ID      int
	Title   string
	Message string
	Created time.Time
	Sent    time.Time
}

type EmailModel struct {
	DB *sql.DB
}

func (e *EmailModel) Insert() (int, error) {
	fmt.Printf("Email Model Insert DB Method")
	return 1, nil
}

func (e *EmailModel) Delete(id int) {
	fmt.Printf("Email Delete Method")
}

func (e *EmailModel) Get(id int) (Email, error) {
	var email Email
	fmt.Printf("Email Get Method")
	return email, nil
}

func (e *EmailModel) GetAllEmails(target []string) ([]Email, error) {
	var emails []Email
	fmt.Printf("Email GetAllContacts Method")
	return emails, nil
}
func (e *Email) SendEmail(target []string) {
	message := []byte("This is a test email")

	// Sender Data
	from := "ian@from.com"
	pass := "password"

	// Receiver address
	to := target

	// smtp settings
	host := "127.0.0.1"
	port := "1025"

	// Authentication.
	auth := smtp.PlainAuth("", from, pass, host)

	// Sending email.
	err := smtp.SendMail(host+":"+port, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
