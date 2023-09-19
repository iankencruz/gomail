package main

import (
	"log"

	"github.com/iankencruz/gomail/pkg/routes"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("sendgrid.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	routes.InitializeRoutes()

	// currentTime := time.Now()

}
