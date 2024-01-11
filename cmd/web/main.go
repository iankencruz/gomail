package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"

	// "github.com/iankencruz/gomail/pkg/routes"
	"github.com/iankencruz/gomail/internal/models"
	"github.com/joho/godotenv"
)

type application struct {
	// TODO: add models
	title    string
	contacts *models.Contact
}

func main() {

	// Specify Web Address Port with custom flags
	addr := flag.String("addr", ":8080", "HTTP Network Address")

	// TODO: Add DSN when implementing SQL

	err := godotenv.Load("sendgrid.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := &application{
		contacts: &models.Contact{
			ID:        2,
			FirstName: "Ian",
			LastName:  "Dela Cruz",
			Email:     "iacruz@test.com",
			Phone:     "034923872",
		},
		title: "GoMail",
	}

	mux := http.NewServeMux()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger.Info("Starting server on", slog.Any("addr:", ":8080"))

	mux.HandleFunc("/", app.home)
	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)

	// routes.InitializeRoutes()
	// currentTime := time.Now()

}
