package main

import (
	"path/filepath"
	"text/template"
	"time"

	"github.com/iankencruz/gomail/internal/models"
)

type templateData struct {
	CurrentYear int
	Contact     models.Contact
	Contacts    []models.Contact
}

func humanDate(t time.Time) string {
	return t.Format("02/01/2006")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {

	// Init a new Cache
	cache := map[string]*template.Template{}

	// Use filepath.Glob() to get a slice of all the filepaths of "./ui/html/pages/*.tmpl
	// This will essentially give us a slice of all the filepaths for our application 'page' templates
	// e.g: [ui/html/pages/home.tmpl | ui/html/pages/view.tmpl]
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	// Loop through the page filepath one-by-one
	for _, page := range pages {
		// Extract the filename 'home.tmpl' from the full filepaths
		// and assign it to the name variable
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}

		// Parse the files into a template set
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Add the template set to the map, using the name of the page
		// (like 'home.tmpl') as the key
		cache[name] = ts

	}

	return cache, nil

}
