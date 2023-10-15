package main

import (
	"net/http"

	//"github.com/bishal7679/golang-repo/pkg/config"
	"github.com/bishal7679/Booking-app/internal/config"
	handler "github.com/bishal7679/Booking-app/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(Repo.Home))
	// mux.Get("/about", http.HandlerFunc(Repo.About))
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	// mux.Use(NoSurf)
	// mux.Use(WriteToConsole)
	mux.Use(SessionLoad)
	mux.Get("/", handler.Repo.Home)
	mux.Get("/chooseplan", handler.Repo.Chooseplan)
	mux.Get("/plans", handler.Repo.Plans)
	mux.Get("/plansignup", handler.Repo.Plansignup)
	mux.Get("/payment", handler.Repo.Payment)
	mux.Get("/book-flight", handler.Repo.Bookflight)
	mux.Post("/book-flight", handler.Repo.PostBookflight)
	mux.Get("/booking-summary", handler.Repo.BookingSummary)
	mux.Get("/indonesia", handler.Repo.Indonesia)
	mux.Get("/japan", handler.Repo.Japan)
	mux.Get("/thailand", handler.Repo.Thailand)
	mux.Get("/southkorea", handler.Repo.Southkorea)

	mux.Get("/contact", handler.Repo.Contact)
	mux.Post("/contact", handler.Repo.PostContact)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
