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
func (m *postgresDBRepo) InsertBooking(book models.BookingDetails) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int

	stmt := `insert into flightbookings (travelway, flying_from, flying_to, depart, return, travel_class, full_name, address, email, country_code, mobile_no, pincode, city, state, govtidentity, user_id, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18) returning id`

	err := m.DB.QueryRowContext(ctx, stmt,
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
	).Scan(&newID)

	if err != nil {
		return 0, err
	}
	return newID, nil
}

// InsertBookingRestriction inserts a booking restrictio into database
func (m *postgresDBRepo) InsertBookingRestriction(r models.BookingRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into bookingsrestriction (travelway, flying_from, flying_to, depart, return, country_code, mobile_no, booking_id, user_id, restriction_id, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	_, err := m.DB.ExecContext(ctx, stmt,
		r.Travelway,
		r.FlyingFrom,
		r.FlyingTo,
		r.Depart,
		r.Return,
		r.CountryCode,
		r.MobileNo,
		r.BookingID,
		r.UserID,
		r.RestrictionID,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (m *postgresDBRepo) SearchExistanceBookingByUserID(country_code, mobile_no, departing, returning, travelway string, userID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var numRows int

	query := `
		select
			count(id)
		from
			bookingsrestriction
		where
			($1 = user_id and $2 = country_code and $3 = mobile_no and $4 < return and $5 > depart and $6 = travelway) or ($1 = user_id and $4 = depart and $5 = return);`

	row := m.DB.QueryRowContext(ctx,query,userID,country_code,mobile_no,departing,returning,travelway)
	err := row.Scan(&numRows)
	if err != nil {
		return false, err
	}

	if numRows == 0 {
		return true, nil
	}
	return false, nil
}

// counting passengers for a flight to check whether its full or not
func(m *postgresDBRepo) CountPassangerForDate(departing, returning, travelway, flying_from, flying_to string) (int,error){
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var numRows int

	query := `
		select
			count(id)
		from
			flightbookings
		where
			$1 = travelway and $2 = depart and $3 = return and $4 = flying_from and $5 = flying_to;`

	row2 := m.DB.QueryRowContext(ctx,query,travelway,departing,returning,flying_from,flying_to)
	err := row2.Scan(&numRows)
	if err != nil {
		return 0, err
	}
	return numRows, nil
}

