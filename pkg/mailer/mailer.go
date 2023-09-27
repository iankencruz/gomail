package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type PayrollTemplate struct {
	Firstname      string
	StartDate      string
	EndDate        string
	DeadlineDate   string
	ProcessingDate string
}

// NewEmail ...
func NewPayrollTemplate(fname string, start time.Time) *PayrollTemplate {
	return &PayrollTemplate{
		Firstname:      fname,
		StartDate:      start.Format("02/01/2006"),
		EndDate:        start.AddDate(0, 0, 13).Format("02/01/2006"),
		DeadlineDate:   start.AddDate(0, 0, 16).Format("02/01/2006"),
		ProcessingDate: start.AddDate(0, 0, 20).Format("02/01/2006"),
	}
}

func ParseTemplate(r *http.Request, file string, c *mail.Email) (s string) {

	t := template.New(file) // Try without dir path

	t, err := t.ParseFiles("./configs/templates/" + file) // Try without dir path
	if err != nil {
		return err.Error()
	}

	err = r.ParseForm()
	if err != nil {
		fmt.Println("Error Parsing Form")
		return err.Error()
	}

	// * Get Date Input Value
	startDateForm := r.PostFormValue("startdate")

	startDate, err := time.Parse("2006-01-02", startDateForm)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	mbody := NewPayrollTemplate(c.Name, startDate)

	var body bytes.Buffer

	// Template Struct
	t.Execute(&body, mbody)
	if err != nil {
		fmt.Printf("Error: func t.Execute: %v", err)
	}

	html := body.String()

	return html

}

func SendMail(from *mail.Email, to *mail.Email, subject string, plainText string, html string, sendgridKey string) {

	message := mail.NewSingleEmail(from, subject, to, plainText, html)
	client := sendgrid.NewSendClient(sendgridKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("#################################### \n")
		fmt.Printf("Sending Email To: %v \n\n", to)
		fmt.Printf("Email Status: %v \n\n", response.StatusCode)
		if response.StatusCode == 401 {
			fmt.Printf("Error: Requires Authentication! Please Try Again... \n\n\n")
		} else if response.StatusCode == 202 {
			fmt.Printf("Completed! Email Succussfully Sent \n\n\n")

		}
	}

}
