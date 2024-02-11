package models

import (
	"database/sql"
	"fmt"
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

func (e *EmailModel) GetAllEmails() ([]Email, error) {
	var emails []Email
	fmt.Printf("Email GetAllContacts Method")
	return emails, nil
}
