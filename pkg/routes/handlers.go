package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/iankencruz/gomail/pkg/goexcel"
	"github.com/iankencruz/gomail/pkg/mailer"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "./configs/web/index.html")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// truncated for brevity

	// * Handle Uploading File
	// * Creates a ./uploads folder and
	// * saves uploaded file in directory to read
	file, fileHeader, err := uploadFile(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// =========================
	// TODO Do Email Stuff Here
	// =========================

	fmt.Println("\nReading Files & Creating Email\n")

	// * =================================
	// *  Set Mailer Authentication and send data
	// * =================================

	//= "plainTextContent" -> Plain Text that shows up for browsers
	//= that do not support MIME Types and HTML Emails
	plainTextContent := "This is the plain text content"

	//= Read Data from an excel file
	contacts := goexcel.ReadExcelFile(fmt.Sprintf("./configs/uploads/%s", fileHeader.Filename))

	//=  prepare html templates by passing in the location of the template file
	//=  that is stored in /config/templates/ & then passing the struct data
	//=  Created earlier to pass dynamic content to the template file.

	//=  Setup Sender details, must be authenticated with the correct domain
	//=  otherwise sendgrid throws authentication error
	from := mail.NewEmail("Paysorted Admin Team", "noreply@ynotsoft.com")

	// for _, target := range contacts {
	// 	fmt.Printf("Target Fname: %v \n", target.Firstname)
	// 	fmt.Printf("Target Lname: %v \n", target.Lastname)
	// 	fmt.Printf("Target Email: %v \n", target.Email)
	// 	fmt.Printf("Target Phone: %v \n", target.Phone)
	// }

	sbjct := "Timesheet Reminder || TEST "

	// Sendgrid Sendmail Function
	mailer.SendMail(r, "template.html", from, contacts, sbjct, plainTextContent, os.Getenv("SENDGRID_API_KEY"))

	fmt.Println("\nFinished Parsing Templates && Email Contacts\n")

	// ==================================
	// = Handles deleting uploaded file uploaded at ./uploads folderex
	// ==================================
	deleteFile(w, r, file, fileHeader)

	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "./configs/web/success.html")

}

func closeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "exit.html")
	fmt.Println("Server Closed")
	os.Exit(1)
}
