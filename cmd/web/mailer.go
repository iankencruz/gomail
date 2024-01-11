package main

import (
// "bytes"
// "fmt"
// "html/template"
// "io"
// "log"
// "mime/multipart"
// "net/http"
// "os"
// "time"
//
// // "github.com/sendgrid/sendgrid-go"
// "github.com/sendgrid/sendgrid-go/helpers/mail"
// // "github.com/xuri/excelize/v2"
)

type PayrollTemplate struct {
	Firstname      string
	StartDate      string
	EndDate        string
	DeadlineDate   string
	ProcessingDate string
}

//
// // NewEmail ...
// func NewPayrollTemplate(fname string, start time.Time) *PayrollTemplate {
// 	return &PayrollTemplate{
// 		Firstname:      fname,
// 		StartDate:      start.Format("02/01/2006"),
// 		EndDate:        start.AddDate(0, 0, 13).Format("02/01/2006"),
// 		DeadlineDate:   start.AddDate(0, 0, 16).Format("02/01/2006"),
// 		ProcessingDate: start.AddDate(0, 0, 20).Format("02/01/2006"),
// 	}
// }
//
// func ParseTemplate(r *http.Request, file string, c *mail.Email) (s string) {
//
// 	t := template.New(file) // Try without dir path
//
// 	t, err := t.ParseFiles("./configs/templates/" + file) // Try without dir path
// 	if err != nil {
// 		return err.Error()
// 	}
//
// 	err = r.ParseForm()
// 	if err != nil {
// 		fmt.Println("Error Parsing Form")
// 		return err.Error()
// 	}
//
// 	// * Get Date Input Value
// 	startDateForm := r.PostFormValue("startdate")
//
// 	startDate, err := time.Parse("2006-01-02", startDateForm)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err.Error()
// 	}
//
// 	mbody := NewPayrollTemplate(c.Name, startDate)
//
// 	var body bytes.Buffer
//
// 	// Template Struct
// 	t.Execute(&body, mbody)
// 	if err != nil {
// 		fmt.Printf("Error: func t.Execute: %v", err)
// 	}
//
// 	html := body.String()
//
// 	return html
//
// }
//
// func SendMail(from *mail.Email, to *mail.Email, subject string, plainText string, html string, sendgridKey string) {
//
// 	message := mail.NewSingleEmail(from, subject, to, plainText, html)
// 	client := sendgrid.NewSendClient(sendgridKey)
// 	response, err := client.Send(message)
// 	if err != nil {
// 		log.Println(err)
// 	} else {
// 		fmt.Printf("#################################### \n")
// 		fmt.Printf("Sending Email To: %v \n\n", to)
// 		fmt.Printf("Email Status: %v \n\n", response.StatusCode)
// 		if response.StatusCode == 401 {
// 			fmt.Printf("Error: Requires Authentication! Please Try Again... \n\n\n")
// 		} else if response.StatusCode == 202 {
// 			fmt.Printf("Completed! Email Succussfully Sent \n\n\n")
//
// 		}
// 	}
//
// }
//
// type Contact struct {
// 	Firstname string
// 	Lastname  string
// 	Email     string
// 	Phone     string
// }
//
// // Returns a contact struct
// func ReadExcelFile(input string) []Contact {
//
// 	var contacts []Contact
// 	var cPersons Contact
//
// 	f, err := excelize.OpenFile(input)
// 	if err != nil {
// 		fmt.Printf("Error: %v", err)
// 		return contacts
// 	}
//
// 	rows, err := f.GetRows("Sheet1")
// 	if err != nil {
// 		fmt.Printf("Error: %v", err)
// 		return contacts
// 	}
//
// 	for _, row := range rows[1:] {
// 		// append column values to slice
// 		cPersons.Firstname = row[1]
// 		cPersons.Lastname = row[2]
// 		cPersons.Email = row[3]
// 		cPersons.Phone = row[4]
//
// 		contacts = append(contacts, cPersons)
// 	}
//
// 	return contacts
// }
//
// func CreateEmailContact(c Contact) mail.Email {
// 	var m mail.Email
// 	m.Address = c.Email
// 	m.Name = c.Firstname
//
// 	return m
// }
//
// // Files
//
// // * Upload File and store in projects uploads directory
// func UploadFile(w http.ResponseWriter, r *http.Request) (multipart.File, *multipart.FileHeader, error) {
// 	file, fileHeader, err := r.FormFile("file")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return file, fileHeader, err
// 	}
//
// 	defer file.Close()
//
// 	//  Create the uploads folder if it doesn't already exist
// 	err = os.MkdirAll("./configs/uploads", os.ModePerm)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return file, fileHeader, err
// 	}
//
// 	//  Create a new file in the uploads directory
// 	dst, err := os.Create(fmt.Sprintf("./configs/uploads/%s", fileHeader.Filename))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return file, fileHeader, err
// 	}
//
// 	defer dst.Close()
//
// 	// Copy the uploaded file to the filesystem
// 	// at the specified destination
// 	_, err = io.Copy(dst, file)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return file, fileHeader, err
// 	}
//
// 	fmt.Printf("\n\nUpload successful\n\n")
//
// 	return file, fileHeader, err
// }
//
// // * Delete uploaded file from projects uploads directory
// func DeleteFile(w http.ResponseWriter, r *http.Request, file multipart.File, fileHeader *multipart.FileHeader) {
// 	fmt.Printf("Deleting uploaded file %s \n", fileHeader.Filename)
// 	err := os.Remove(fmt.Sprintf("./configs/uploads/%s", fileHeader.Filename))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("Successfully deleted:  %s \n", fileHeader.Filename)
// }
