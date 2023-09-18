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

func PrepareTemplates(file string, dt any) string {
	t := template.New("./configs/templates/" + file)

	t, err := t.ParseFiles("./configs/templates/" + file)
	if err != nil {
		log.Println(err)
	}

	var body bytes.Buffer

	// struct {
	// 	DateTime string
	//  values...
	// }{
	// 	DateTime: currentTime.Format("02/01/2006"),
	//  values...
	// })

	t.Execute(&body, dt)

	htmlBody := body.String()
	fmt.Println(htmlBody)

	return htmlBody
}

func SendMail(fromSender *mail.Email, toTargets []goexcel.Contact, subject string, plainText string, html string) {

	// from := mail.NewEmail("Paysorted Admin Team", "noreply@ynotsoft.com")

	// Create sendgrid Contacts & send
	addresses := make([]mail.Email, len(toTargets))

	for i, contact := range toTargets {

		addresses[i].Address = contact.Email
		addresses[i].Name = contact.Firstname

		message := mail.NewSingleEmail(fromSender, subject, &addresses[i], plainText, html)
		client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
		response, err := client.Send(message)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Printf("#################################### \n")
			fmt.Printf("Sending Email To: %v \n\n", addresses[i])
			fmt.Printf("Email Status: %v \n\n", response.StatusCode)
			if response.StatusCode == 401 {
				fmt.Printf("Error: Requires Authentication! Please Try Again... \n\n\n")
			} else if response.StatusCode == 200 {
				fmt.Printf("Completed! Email Succussfully Sent \n\n\n")

			}
		}
	}

}
