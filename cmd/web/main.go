package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iankencruz/gomail/internal/models"
	"github.com/joho/godotenv"
)

type application struct {
	// TODO: add models
	title          string
	contact        *models.Contact
	contacts       *models.ContactModel
	email          *models.Email
	emails         *models.EmailModel
	templateCache  map[string]*template.Template
	sessionManager *scs.SessionManager
}

func (a *application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a handler"))
}

func main() {

	// Specify Web Address Port with custom flags
	addr := flag.String("addr", ":8080", "HTTP Network Address")
	dsn := flag.String("dsn", "root:password@/gomail?parseTime=true", "MySQL datasource name")
	flag.Parse()

	db, err := openDB(*dsn)
	if err != nil {
		fmt.Printf("DB Error: %s", err.Error())
		os.Exit(1)
	}
	defer db.Close()

	err = godotenv.Load("sendgrid.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	templateCache, err := newTemplateCache()
	if err != nil {
		fmt.Fprintf(os.Stdout, err.Error())
		os.Exit(1)
	}

	sessionManager := scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 3 * time.Hour

	app := &application{
		contacts:       &models.ContactModel{DB: db},
		emails:         &models.EmailModel{DB: db},
		templateCache:  templateCache,
		sessionManager: sessionManager,
	}

	fmt.Print(app.templateCache)
	srv := &http.Server{
		Addr:           *addr,
		MaxHeaderBytes: 524288,
		Handler:        app.routes(),
		ErrorLog:       slog.NewLogLogger(slog.NewTextHandler(os.Stdout, nil), slog.LevelDebug),
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger.Info("Starting server on", slog.Any("addr:", ":8080"))

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)

	// routes.InitializeRoutes()
	// currentTime := time.Now()

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err

	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
