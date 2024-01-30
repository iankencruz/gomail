package main

import (
	"path/filepath"
	"text/template"
	"time"

	"github.com/iankencruz/gomail/internal/models"
)

type templateData struct {
	CurrentYear int
	Contacts    []models.Contact
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
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

		// Create sice containing the filepaths for our templates
		files := []string{
			"./ui/html/base.tmpl",
			"./ui/html/partials/navigation.tmpl",
			"./ui/html/partials/footer.tmpl",
			page,
		}

		// Parse the files into a template set
		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		// Add the template set to the map, using the name of the page
		// (like 'home.tmpl') as the key
		cache[name] = ts

	}

	return cache, nil

}
