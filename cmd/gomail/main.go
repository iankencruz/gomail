package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"time"

	"github.com/iankencruz/gomail/mailer"

	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {

	// using the function
	_, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	err = godotenv.Load("sendgrid.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	currentTime := time.Now()

	recipients := mailer.ReadNewFile("target_list.txt")

	t := template.New("template.html")

	t, err = t.ParseFiles("template.html")
	if err != nil {
		log.Println(err)
	}

	// Read through recipients in external file
	addresses := make([]mail.Email, len(recipients))

	var body bytes.Buffer

	t.Execute(&body, struct {
		DateTime string
	}{
		DateTime: currentTime.Format("02/01/2006"),
	})

	htmlBody := body.String()

	// sendgrid functions

	from := mail.NewEmail("Paysorted Admin Team", "noreply@ynotsoft.com")
	subject := "Paysorted Feedback"

	plainTextContent := "and easy to do anywhere, even with Go"
	for i, recipient := range recipients {

		addresses[i].Address = recipient
		addresses[i].Name = ""

		message := mail.NewSingleEmail(from, subject, &addresses[i], plainTextContent, htmlBody)
		client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
		response, err := client.Send(message)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Printf("#################################### \n")
			fmt.Printf("Sending Email To: %v \n\n", &addresses[i])
			fmt.Printf("Email Status: %v \n\n", response.StatusCode)
			fmt.Printf("Completed! Email Succussfully Sent \n\n\n")
		}
	}
	fmt.Printf("\nRecieved by:\n%v \n\n", addresses)

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}
