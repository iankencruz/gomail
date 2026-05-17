package models

import (
	"database/sql"
	"fmt"
	"time"

	"gopkg.in/gomail.v2"
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
func (e *Email) SendEmail(target []string, subject string, body string) {

	// Receiver address
	// to := target

	// subj := []byte(subject)
	// message := []byte(to + "Subject:" + body)

	// Sender Data
	// from := "ian@from.com"
	// pass := "password"

	// // smtp settings
	host := "127.0.0.1"
	port := 1025

	m := gomail.NewMessage()

	m.SetHeader("From", "ian.cruz@live.com")
	m.SetHeader("To", "bob@example.com", "cora@example.com")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer(host, port, "user", "123456")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	// Authentication.
	// auth := smtp.PlainAuth("", from, pass, host)

	// Sending email.
	// err := smtp.SendMail(host+":"+port, auth, from, to, message)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// // }
	// fmt.Println("Email Sent Successfully!")
}
