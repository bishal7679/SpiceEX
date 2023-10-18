package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/bishal7679/SpiceEx/internal/config"
	"github.com/bishal7679/SpiceEx/internal/driver"
	"github.com/bishal7679/SpiceEx/internal/forms"
	"github.com/bishal7679/SpiceEx/internal/helpers"
	"github.com/bishal7679/SpiceEx/internal/models"
	"github.com/bishal7679/SpiceEx/internal/render"
	"github.com/bishal7679/SpiceEx/internal/repository"
	"github.com/bishal7679/SpiceEx/internal/repository/dbrepo"
)

var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewpostgresRepo(db.SQL, a),
	}
}

// Newhandlers sets the repository for the handlers
func Newhandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// remoteIP := r.RemoteAddr
	// m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.Template(w, r, "home.page.html", &models.TemplateData{})

}

// Plans is the plans page handler
func (m *Repository) Plans(w http.ResponseWriter, r *http.Request) {

	// stringmap := make(map[string]string)
	// stringmap["test"] = "hello again"

	// remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	// stringmap["remote_ip"] = remoteIP
	render.Template(w, r, "plans.page.html", &models.TemplateData{
		// StringMap: stringmap,
	})
}

// Reservation renders the make a reservation page and displays form
func (m *Repository) Plansignup(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "plansignup.page.html", &models.TemplateData{})

}

func (m *Repository) Bookflight(w http.ResponseWriter, r *http.Request) {

	// This is the emptydata when the page render for the very first time
	var emptyBookingDetails models.BookingDetails
	data := make(map[string]interface{})
	data["bookingDetails"] = emptyBookingDetails

	render.Template(w, r, "book-flight.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})

}

func (m *Repository) PostBookflight(w http.ResponseWriter, r *http.Request) {

	// comment out line 86 - 98 and uncomment 79 - 83 for testing handlers
	// err := r.ParseForm()
	// if err != nil {
	// 	helpers.ServerError(w, err)
	// 	return
	// }

	// Parse the form data, including the uploaded file
	err := r.ParseMultipartForm(10 << 20) // 10MB maximum file size
	if err != nil {
		log.Println("error in 1")
		helpers.ServerError(w, err)
		return
	}

	// Retrieve the uploaded file
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		log.Println("error in 2")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read the file data
	data := make([]byte, handler.Size)
	_, err = file.Read(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Validate file type, for example, only allow image files
	// if handler.Header.Get("Content-Type") != "image/jpeg" && handler.Header.Get("Content-Type") != "image/png" && handler.Header.Get("Content-Type") != "image/pdf" {
	// 	http.Error(w, "Only JPEG and PNG files are allowed", http.StatusBadRequest)
	// 	return
	// }

	// Create a new file on the server to save the uploaded file
	// outputFile, err := os.Create(handler.Filename)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// defer outputFile.Close()

	// departDate := r.Form.Get("depart")
	// returnDate := r.Form.Get("return")

	// // 2023-01-01 --- 01/02 03:04:05PM '06 -0700

	// layout := "2006-01-02"
	// Depart_Date, err := time.Parse(layout, departDate)
	// if err != nil {
	// 	helpers.ServerError(w,err)
	// }
	// Return_Date, err := time.Parse(layout, returnDate)
	// if err != nil {
	// 	helpers.ServerError(w,err)
	// }

	userID, err := strconv.Atoi(r.Form.Get("user_id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	bookingDetails := models.BookingDetails{
		Travelway:         r.Form.Get("check"),
		Flying_From:       r.Form.Get("flying_from"),
		Flying_To:         r.Form.Get("flying_to"),
		Departing_Date:    r.Form.Get("depart"),
		Returning_Date:    r.Form.Get("return"),
		Travel_Class:      r.Form.Get("travel_class"),
		Full_Name:         r.Form.Get("full_name"),
		Address:           r.Form.Get("address"),
		Email:             r.Form.Get("email"),
		Country_Code:      r.Form.Get("country_code"),
		Mobile_No:         r.Form.Get("mobile_no"),
		Pincode:           r.Form.Get("pincode"),
		City_Name:         r.Form.Get("city"),
		State_Name:        r.Form.Get("state"),
		UserID:            userID,
		Upload_File_As_ID: data,
	}
	form := forms.New(r.PostForm)

	form.Required("flying_from", "flying_to", "depart", "return", "full_name", "address", "email")

	form.MinLength("full_name", 6)

	form.IsEmail("email")
	if !form.Valid() {
		data := make(map[string]interface{})
		data["bookingDetails"] = bookingDetails

		render.Template(w, r, "book-flight.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return

	}

	// inserting booking details to the database
	newBookingID, err := m.DB.InsertBooking(bookingDetails)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	restriction := models.BookingRestriction{
		Travelway:     r.Form.Get("check"),
		UserID:        userID,
		BookingID:     newBookingID,
		RestrictionID: 1,
		FlyingFrom:    r.Form.Get("flying_from"),
		FlyingTo:      r.Form.Get("flying_to"),
		Depart:        r.Form.Get("depart"),
		Return:        r.Form.Get("return"),
		CountryCode:  r.Form.Get("country_code"),
		MobileNo:     r.Form.Get("mobile_no"),
	}

	err = m.DB.InsertBookingRestriction(restriction)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	m.App.Session.Put(r.Context(), "bookingDetails", bookingDetails)
	http.Redirect(w, r, "/booking-summary", http.StatusSeeOther)

}

func (m *Repository) Indonesia(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "indonesia.page.html", &models.TemplateData{})
}

func (m *Repository) Japan(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "japan.page.html", &models.TemplateData{})
}

func (m *Repository) Thailand(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "thailand.page.html", &models.TemplateData{})
}

func (m *Repository) Southkorea(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "southKorea.page.html", &models.TemplateData{})
}

// Majors renders the room page
func (m *Repository) Payment(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "payment.page.html", &models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repository) Chooseplan(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "chooseplan.page.html", &models.TemplateData{})
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	// This is the emptydata when the page render for the very first time
	var emptyContactDetails models.ContactDetails
	data := make(map[string]interface{})
	data["contactDetails"] = emptyContactDetails
	render.Template(w, r, "contact.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

func (m *Repository) PostContact(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("posted contact data"))
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	contactDetails := models.ContactDetails{
		Name:    r.Form.Get("name"),
		Email:   r.Form.Get("email"),
		Message: r.Form.Get("message"),
	}

	form := forms.New(r.PostForm)

	form.Required("name", "email", "message")
	form.IsEmail("email")
	if !form.Valid() {
		data := make(map[string]interface{})
		data["contactDetails"] = contactDetails

		render.Template(w, r, "contact.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return

	} else {
		w.Write([]byte("posted contact data successfully!"))
	}

}

func (m *Repository) BookingSummary(w http.ResponseWriter, r *http.Request) {
	bookingDetails, ok := m.App.Session.Get(r.Context(), "bookingDetails").(models.BookingDetails)
	if !ok {
		// m.App.ErrorLog.Println("cannot get item from session")
		m.App.Session.Put(r.Context(), "error", "Can't get booking-details from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	m.App.Session.Remove(r.Context(), "bookingDetails")
	data := make(map[string]interface{})
	data["bookingDetails"] = bookingDetails
	render.Template(w, r, "booking-summary.page.html", &models.TemplateData{
		Data: data,
	})
}
