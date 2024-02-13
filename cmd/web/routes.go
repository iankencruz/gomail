package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/alice"
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

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	// Create Routes and assign handlers
	r.Get("/", dynamic.ThenFunc(http.HandlerFunc(app.home)).ServeHTTP)
	r.Get("/contacts", dynamic.ThenFunc(http.HandlerFunc(app.listcontacts)).ServeHTTP)
	r.Get("/contacts/view/{id}", dynamic.ThenFunc(http.HandlerFunc(app.contactView)).ServeHTTP)
	r.Delete("/contacts/view/{id}", dynamic.ThenFunc(http.HandlerFunc(app.deleteContact)).ServeHTTP)
	r.Get("/contacts/create", dynamic.ThenFunc(http.HandlerFunc(app.contactCreate)).ServeHTTP)
	r.Post("/contacts/create", dynamic.ThenFunc(http.HandlerFunc(app.contactCreatePost)).ServeHTTP)

	// Email Routes
	r.Get("/emails/create", dynamic.ThenFunc(http.HandlerFunc(app.emailCreate)).ServeHTTP)
	r.Post("/emails/create", dynamic.ThenFunc(http.HandlerFunc(app.emailCreatePost)).ServeHTTP)

	return r
}
