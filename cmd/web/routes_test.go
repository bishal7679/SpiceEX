package main

import (
	"fmt"
	"testing"

	"github.com/bishal7679/Booking-app/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := Routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing; test passed
	default:
		t.Errorf(fmt.Sprintf("type is not *chi.Mux, type is %T", v))
	}
}
