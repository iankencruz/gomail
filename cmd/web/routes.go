package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// "log"
// "net/http"

func (app *application) routes() http.Handler {

	// Initialize router
	r := chi.NewRouter()

	// Use MiddleWare logger
	r.Use(middleware.Logger)

	// Serve Static Files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// Create Routes and assign handlers
	r.Get("/", app.home)

	// Contacts Routes
	r.Get("/contacts", app.listcontacts)
	r.Get("/contacts/view/{id}", app.contactView)
	r.Delete("/contacts/view/{id}", app.deleteContact)
	r.Get("/contacts/create", app.contactCreate)
	r.Post("/contacts/create", app.contactCreatePost)

	return r
}
