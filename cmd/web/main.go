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
	err := run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("Starting application on port : %s", PortNumber))
	// _ = http.ListenAndServe(PortNumber, nil)
	srv := &http.Server{
		Addr:    PortNumber,
		Handler: Routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}

func run() error {
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
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache!")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handler.NewRepo(&app)
	handler.Newhandlers(repo)

	render.NewTemplate(&app)
	helpers.NewHelpers(&app)

	return nil
}
