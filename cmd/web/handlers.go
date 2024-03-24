package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/iankencruz/gomail/internal/models"
	"github.com/iankencruz/gomail/internal/validator"
	// "github.com/iankencruz/gomail/pkg/mailer"
)

type contactCreateForm struct {
	Fname string
	Lname string
	Email string
	Phone string
	validator.Validator
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	contacts, err := app.contacts.GetAllContacts()
	if err != nil {
		fmt.Printf("Server Error: %s", err.Error())
		return
	}

	data := app.newTemplateData(r)
	data.Contacts = contacts

	app.render(w, r, http.StatusOK, "home.tmpl", data)

}

func (app *application) deleteContact(w http.ResponseWriter, r *http.Request) {
	// get id Param
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		fmt.Printf("Server Error: %v", err.Error())
		return
	}

	if err != nil {
		fmt.Printf("Server Error: %v", err.Error())
		return
	}

	app.contacts.Delete(id)
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func (app *application) listcontacts(w http.ResponseWriter, r *http.Request) {
	contacts, err := app.contacts.GetAllContacts()
	if err != nil {
		fmt.Printf("Server Error: %s", err.Error())
		return
	}

	data := app.newTemplateData(r)
	data.Contacts = contacts

	app.render(w, r, http.StatusOK, "contacts.tmpl", data)
}

func (app *application) contactView(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		fmt.Printf("Server Error: %v", err.Error())
		return
	}

	contact, err := app.contacts.Get(id)

	// Check Error
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			fmt.Printf("Server Error: %v", err.Error())
		} else {
			fmt.Printf("Server Error: %v", err.Error())
		}
		return
	}
	// Logging URL Param
	// fmt.Printf("%+v", contact)

	data := app.newTemplateData(r)
	data.Contact = contact

	app.render(w, r, http.StatusOK, "contact_view.tmpl", data)
}

func (app *application) contactCreate(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("contactCreate Form : Get Request"))
	contacts, err := app.contacts.GetAllContacts()
	if err != nil {
		fmt.Printf("Server Error: %s", err.Error())
		return
	}
	data := app.newTemplateData(r)
	data.Contacts = contacts

	app.render(w, r, http.StatusOK, "contact_create.tmpl", data)
}

func (app *application) contactCreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("Server Error: %v", err.Error())
		return
	}

	// TODO: Automatic Form Decoder/Parsing
	form := contactCreateForm{
		Fname: r.PostForm.Get("first-name"),
		Lname: r.PostForm.Get("last-name"),
		Email: r.PostForm.Get("email"),
		Phone: r.PostForm.Get("phone-number"),
	}

	// Initialize a map to hold any validation errors for the form fields.
	form.CheckField(validator.NotBlank(form.Fname), "first-name", "'First Name' cannot be blank")
	form.CheckField(validator.NotBlank(form.Lname), "last-name", "'Last Name' cannot be blank")
	form.CheckField(validator.NotBlank(form.Email), "email", "Email cannot be blank")
	form.CheckField(validator.NotBlank(form.Phone), "phone-number", "Phone number  cannot be blank")
	// check MaxChars of form
	form.CheckField(validator.MaxChars(form.Fname, 100), "first-name", "First Name can be no longer than 100 characters")
	form.CheckField(validator.MaxChars(form.Lname, 100), "last-name", "Last Name can be no longer than 100 characters")
	form.CheckField(validator.MaxChars(form.Email, 100), "email", "Email can be no longer than 100 characters")
	form.CheckField(validator.MaxChars(form.Phone, 10), "phone-number", "Phone number can be no longer than 10 characters")

	form.CheckField(validator.EmailValidate(form.Email), "email", "This field must be a valid email address")

	// TODO: Render Error to HTML
	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		// render error partial
		app.render(w, r, http.StatusUnprocessableEntity, "contact_create.tmpl", data)
		fmt.Printf("Form Validation Error: %v", form.FieldErrors)
		return
	}

	_, err = app.contacts.Insert(form.Fname, form.Lname, form.Email, form.Phone)
	if err != nil {
		fmt.Printf("Server Error: %v", err.Error())
		return
	}

	// TODO: Sessions to flash creation message

	// contact, err := app.contacts.Get(id)
	// data := app.newTemplateData(r)
	// data.Contact = contact
	// data.Form = form

	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

// EMAIL Handlers
func (app *application) emailCreate(w http.ResponseWriter, r *http.Request) {
	// emails, err := app.emails.GetAllEmails()
	// if err != nil {
	// fmt.Printf("Server Error: %s", err.Error())
	// return
	// }

	contacts, err := app.contacts.GetAllContacts()
	if err != nil {
		fmt.Printf("Server Error: %s", err.Error())
		return
	}
	data := app.newTemplateData(r)
	data.Contacts = contacts
	// data.Emails = emails

	app.render(w, r, http.StatusOK, "email_create.tmpl", data)

}

func (a *application) emailCreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("Server Error: %v", err.Error())
		return
	}
	fmt.Printf("%+v\n", r.Form)
}

// func (app *application) contactGetAll(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Get All Contacts"))
// }

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
