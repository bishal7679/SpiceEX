package dbrepo

import (
	"context"
	"time"

	"github.com/bishal7679/SpiceEx/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// InsertBooking inserts a booking into the database
func (m *postgresDBRepo) InsertBooking(book models.BookingDetails) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into flightbookings (check, flying_from, flying_to, depart, return, travel_class, full_name, address, email, country_code, mobile_no, pincode, city, state, govtidentity, user_id, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)`

	_, err := m.DB.ExecContext(ctx,stmt,
		book.Travelway,
		book.Flying_From,
		book.Flying_To,
		book.Departing_Date,
		book.Returning_Date,
		book.Travel_Class,
		book.Full_Name,
		book.Address,
		book.Email,
		book.Country_Code,
		book.Mobile_No,
		book.Pincode,
		book.City_Name,
		book.State_Name,
		book.Upload_File_As_ID,
		book.UserID,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}
	return nil
}
