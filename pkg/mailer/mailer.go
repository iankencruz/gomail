package mailer

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/iankencruz/gomail/pkg/goexcel"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var Info = "This is package mailer"

// Comment added
func ReadNewFile(file string) []string {

	var result []string

	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {

		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return result

}

// func PrepareTemplates(file string, contacts []goexcel.Contact) string {
// 	t := template.New(file) // Try without dir path

// 	t, err := t.ParseFiles("./configs/templates/" + file) // Try without dir path
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	var body bytes.Buffer

// 	if err := t.Execute(&body, contacts); err != nil {
// 		fmt.Printf("Error: func t.Execute: %v", err)
// 	}

// 	htmlBody := body.String()

// 	// fmt.Printf("HTML Body: %v", htmlBody)

// 	return htmlBody
// }

type MailBody struct {
	Firstname      string
	StartDate      string
	EndDate        string
	ProcessingDate string
}

// NewEmail ...
func NewMailBody(fname string, start string, end string, processDate string) *MailBody {
	return &MailBody{
		Firstname:      fname,
		StartDate:      start,
		EndDate:        end,
		ProcessingDate: processDate,
	}
}

func SendMail(r *http.Request, file string, fromSender *mail.Email, toTargets []goexcel.Contact, subject string, plainText string, sendgridKey string) {

	t := template.New(file) // Try without dir path

	t, err := t.ParseFiles("./configs/templates/" + file) // Try without dir path
	if err != nil {
		log.Println(err)
	}

	err = r.ParseForm()
	if err != nil {
		fmt.Println("Error Parsing Form")
		return
	}

	// * Get Date Input Value
	inputStartDate := r.PostFormValue("startdate")
	inputEndDate := r.PostFormValue("enddate")
	inputProcessDate := r.PostFormValue("processingdate")

	startDate, err := time.Parse("2006-01-02", inputStartDate)
	if err != nil {
		fmt.Println(err)
		return
	}
	endDate, err := time.Parse("2006-01-02", inputEndDate)
	if err != nil {
		fmt.Println(err)
		return
	}
	processDate, err := time.Parse("2006-01-02", inputProcessDate)
	if err != nil {
		fmt.Println(err)
		return
	}
	// * Format Date Input
	fmtStartDate := startDate.Format("02/01/2006")
	fmtEndDate := endDate.Format("02/01/2006")
	fmtProcessDate := processDate.Format("02/01/2006")

	for _, contact := range toTargets {

		var body bytes.Buffer

		// Create sendgrid Contacts & send

		target := mail.NewEmail(contact.Firstname, contact.Email)
		mbody := NewMailBody(contact.Firstname, fmtStartDate, fmtEndDate, fmtProcessDate)

		// Template Struct

		t.Execute(&body, mbody)
		if err != nil {
			fmt.Printf("Error: func t.Execute: %v", err)
		}

		htmlBody := body.String()

		message := mail.NewSingleEmail(fromSender, subject, target, plainText, htmlBody)
		client := sendgrid.NewSendClient(sendgridKey)
		response, err := client.Send(message)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Printf("#################################### \n")
			fmt.Printf("Sending Email To: %v \n\n", target)
			fmt.Printf("Email Status: %v \n\n", response.StatusCode)
			if response.StatusCode == 401 {
				fmt.Printf("Error: Requires Authentication! Please Try Again... \n\n\n")
			} else if response.StatusCode == 202 {
				fmt.Printf("Completed! Email Succussfully Sent \n\n\n")

			}
		}
	}

}
