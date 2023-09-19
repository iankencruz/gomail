package mailer

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"

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

func PrepareTemplates(file string, contacts []goexcel.Contact) string {
	t := template.New(file) // Try without dir path

	t, err := t.ParseFiles("./configs/templates/" + file) // Try without dir path
	if err != nil {
		log.Println(err)
	}

	var body bytes.Buffer

	// data := struct {
	// 	Name string
	// }{
	// 	Name: ,
	// }

	// struct {
	// 	DateTime string
	//  values...
	// }{
	// 	DateTime: currentTime.Format("02/01/2006"),
	//  values...
	// })

	if err := t.Execute(&body, contacts); err != nil {
		fmt.Printf("Error: func t.Execute: %v", err)
	}

	htmlBody := body.String()

	// fmt.Printf("HTML Body: %v", htmlBody)

	return htmlBody
}

func SendMail(fromSender *mail.Email, toTargets []goexcel.Contact, subject string, plainText string, html string, sendgridKey string) {

	// from := mail.NewEmail("Paysorted Admin Team", "noreply@ynotsoft.com")

	// Create sendgrid Contacts & send
	addresses := make([]mail.Email, len(toTargets))

	for i, contact := range toTargets {

		addresses[i].Address = contact.Email
		addresses[i].Name = contact.Firstname

		message := mail.NewSingleEmail(fromSender, subject, &addresses[i], plainText, html)
		client := sendgrid.NewSendClient(sendgridKey)
		response, err := client.Send(message)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Printf("#################################### \n")
			fmt.Printf("Sending Email To: %v \n\n", addresses[i])
			fmt.Printf("Email Status: %v \n\n", response.StatusCode)
			if response.StatusCode == 401 {
				fmt.Printf("Error: Requires Authentication! Please Try Again... \n\n\n")
			} else if response.StatusCode == 202 {
				fmt.Printf("Completed! Email Succussfully Sent \n\n\n")

			}
		}
	}

}
