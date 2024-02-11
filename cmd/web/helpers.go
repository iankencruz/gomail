package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		fmt.Printf("Server Error: %s", err.Error())
		return
	}
	// Write out the provided HTTP status code ('200 OK', '400 Bad Request' etc).
	w.WriteHeader(status)
	// Execute the template set and write the response body. Again, if there
	// is any error we call the the serverError() helper.
	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		fmt.Printf("Server Error: %s", err.Error())
	}

}

func (app *application) renderPartial(w http.ResponseWriter, file string, partial string, data templateData) {
	tmpl := template.Must(template.ParseFiles(file))

	err := tmpl.ExecuteTemplate(w, partial, data)
	if err != nil {
		fmt.Printf("Server Error %s", err.Error())
		return
	}
}

func (app *application) newTemplateData(r *http.Request) templateData {
	return templateData{
		CurrentYear: time.Now().Year(),
	}
}
