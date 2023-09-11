package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func readbyWord(file string) []string {

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

func main() {

	// using the function
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	currentTime := time.Now()

	recipients := readbyWord(dir + "/recipients_list.txt")

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

	from := mail.NewEmail("Paysorted Admin", "noreply@ynotsoft.com")
	subject := "Paysorted Feedback"

	// to := mail.NewEmail("ian.cruz@ynotconsulting.com.au", "ian.cruz@ynotconsulting.com.au")
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
