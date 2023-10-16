package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/bishal7679/Booking-app/internal/config"
	"github.com/bishal7679/Booking-app/internal/driver"
	handler "github.com/bishal7679/Booking-app/internal/handlers"
	"github.com/bishal7679/Booking-app/internal/helpers"
	"github.com/bishal7679/Booking-app/internal/models"
	"github.com/bishal7679/Booking-app/internal/render"
)

var PortNumber = ":8000"
var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	fmt.Println(fmt.Sprintf("Starting application on port : %s", PortNumber))
	// _ = http.ListenAndServe(PortNumber, nil)
	srv := &http.Server{
		Addr:    PortNumber,
		Handler: Routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}

func run() (*driver.DB, error) {
	// what am i going to put in the session
	gob.Register(models.BookingDetails{})
	// change this to true when you are in production environment
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database
	log.Println("Connecting to database")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=postgres password=")
	if err != nil {
		log.Fatal("cannot connect to the database! Dying...")
	}
	log.Println("Connected to database!")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache!")
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handler.NewRepo(&app,db)
	handler.Newhandlers(repo)

	render.NewTemplate(&app)
	helpers.NewHelpers(&app)

	return db, nil
}
