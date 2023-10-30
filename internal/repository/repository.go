package repository

import "github.com/bishal7679/SpiceEx/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertBooking(book models.BookingDetails) (int,error)

	InsertBookingRestriction(r models.BookingRestriction) error

	SearchExistanceBookingByUserID(country_code, mobile_no, departing, returning, travelway string, userID int) (bool, error)
}