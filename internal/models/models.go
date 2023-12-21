package models

import (
	"time"
)

// BookingDetails holds bookings data for a user
type BookingDetails struct {
	ID                int
	Travelway         string
	Flying_From       string
	Flying_To         string
	Departing_Date    string
	Returning_Date    string
	Travel_Class      string
	Full_Name         string
	Address           string
	Email             string
	Country_Code      string
	Mobile_No         string
	Pincode           string
	State_Name        string
	City_Name         string
	Upload_File_As_ID []byte
	UserID            int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type ContactDetails struct {
	Name    string
	Email   string
	Message string
}

// models for database.......

// UserSignup is the user-signup model
type UserSignup struct {
	ID          int
	UserName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Restriction is the Restriction model
type Restriction struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// BookingRestriction is the BookingRestriction model
type BookingRestriction struct {
	ID            int
	Travelway     string
	FlyingFrom    string
	FlyingTo      string
	Depart        string
	Return        string
	CountryCode   string
	MobileNo      string
	BookingID     int
	UserID        int
	RestrictionID int
	Booking       BookingDetails
	Restriction   Restriction
	User          UserSignup
}

// Payment is the Payment model
type Payment struct {
	ID         int
	Email      string
	CardNumber string
	Validity   string
	CVV        int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// MailData holds an email address
type MailData struct {
	To      string
	From    string
	Subject string
	Content string
}
