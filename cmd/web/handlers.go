package main

import (
	// "fmt"
	"fmt"
	"log"
	"net/http"
	"os"
	// "os"
	// "github.com/iankencruz/gomail/pkg/mailer"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	contacts, err := app.contacts.GetContacts()
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, "./ui/html/index.html")
	fmt.Printf("FirstName: %s ", contacts[0].FirstName)

}

// func confirmationHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-Type", "text/html")
// 	http.ServeFile(w, r, "./configs/web/confirmation.html")
// }
//
// func uploadHandler(w http.ResponseWriter, r *http.Request) {
// 	// truncated for brevity
//
// 	// * Handle Uploading File
// 	// * Creates a ./uploads folder and
// 	// * saves uploaded file in directory to read
// 	file, fileHeader, err := mailer.UploadFile(w, r)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	// =========================
// 	// TODO Do Email Stuff Here
// 	// =========================
//
// 	fmt.Println("\nReading Files & Creating Email\n")
//
// 	// * =================================
// 	// *  Set Mailer Authentication and send data
// 	// * =================================
//
// 	//= "plainTextContent" -> Plain Text that shows up for browsers
// 	//= that do not support MIME Types and HTML Emails
// 	plainTextContent := "This is the plain text content"
//
// 	//= Read Data from an excel file
// 	contacts := mailer.ReadExcelFile(fmt.Sprintf("./configs/uploads/%s", fileHeader.Filename))
//
// 	// //=  prepare html templates by passing in the location of the template file
// 	// //=  that is stored in /config/templates/ & then passing the struct data
// 	// //=  Created earlier to pass dynamic content to the template file.
//
// 	// //=  Setup Sender details, must be authenticated with the correct domain
// 	// //=  otherwise sendgrid throws authentication error
// 	// from := mail.NewEmail("Paysorted Admin Team", "noreply@ynotsoft.com")
//
// 	// sbjct := "Timesheet Reminder || TEST "
//
// 	// for _, contact := range contacts {
//
// 	// 	mContact := mailer.CreateEmailContact(contact)
// 	// 	htmlBody := mailer.ParseTemplate(r, "template.html", &mContact)
//
// 	// 	// Sendgrid Sendmail Function
// 	// 	mailer.SendMail(from, &mContact, sbjct, plainTextContent, htmlBody, os.Getenv("SENDGRID_API_KEY"))
// 	// 	fmt.Println("\nFinished Parsing Templates && Email Contacts\n")
//
// 	// 	// ==================================
// 	// 	// = Handles deleting uploaded file uploaded at ./uploads folderex
// 	// 	// ==================================
//
// 	// 	w.Header().Add("Content-Type", "text/html")
// 	// 	http.ServeFile(w, r, "./configs/web/success.html")
// 	// }
// 	// mailer.DeleteFile(w, r, file, fileHeader)
//
// }
//
// func closeHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-Type", "text/html")
// 	http.ServeFile(w, r, "exit.html")
// 	fmt.Println("Server Closed")
// 	os.Exit(1)
// }
