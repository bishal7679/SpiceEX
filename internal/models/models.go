package models

// BookingDetails holds bookings data for a user
type BookingDetails struct {
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
	Upload_File_As_ID string
}

type ContactDetails struct {
	Name    string
	Email   string
	Message string
}
