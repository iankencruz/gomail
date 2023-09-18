package main

import (
	"log"
	"time"

	xc "github.com/iankencruz/gomail/pkg/goexcel"
	"github.com/iankencruz/gomail/pkg/mailer"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("sendgrid.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// routes.InitializeRoutes()

	// currentTime := time.Now()

	// // rec := area.ReadNewFile("sdadwa"

	data := struct {
		//TODO  Read date input and format to template
		DateTime string
	}{
		DateTime: time.Now().Format("02/01/2006"),
	}

	//*** Read Data from an excel file
	htmlContent := mailer.PrepareTemplates("template.html", data)

	contacts := xc.ReadExcelFile("sample.xlsx")

	mailer.SendMail(contacts, "subject", "Plain", htmlContent)

	// // Read through recipients in external file

	// var body bytes.Buffer

	// t.Execute(&body, struct {
	// 	DateTime string
	// }{
	// 	DateTime: currentTime.Format("02/01/2006"),
	// })

}
