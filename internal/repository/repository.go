package repository

import "github.com/bishal7679/SpiceEx/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertBooking(book models.BookingDetails) error
}