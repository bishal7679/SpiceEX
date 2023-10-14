package handler

import (
	"net/http"

	"github.com/bishal7679/Booking-app/internal/config"
	"github.com/bishal7679/Booking-app/internal/forms"
	"github.com/bishal7679/Booking-app/internal/helpers"

	// "github.com/bishal7679/Booking-app/internal/forms"
	// "github.com/bishal7679/Booking-app/internal/helpers"
	"github.com/bishal7679/Booking-app/internal/models"
	"github.com/bishal7679/Booking-app/internal/render"
)

var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
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
	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})

}

// Plans is the plans page handler
func (m *Repository) Plans(w http.ResponseWriter, r *http.Request) {

	// stringmap := make(map[string]string)
	// stringmap["test"] = "hello again"

	// remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	// stringmap["remote_ip"] = remoteIP
	render.RenderTemplate(w, r, "plans.page.html", &models.TemplateData{
		// StringMap: stringmap,
	})
}

// Reservation renders the make a reservation page and displays form
func (m *Repository) Plansignup(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "plansignup.page.html", &models.TemplateData{})

}

func (m *Repository) Bookflight(w http.ResponseWriter, r *http.Request) {

	// This is the emptydata when the page render for the very first time
	var emptyBookingDetails models.BookingDetails
	data := make(map[string]interface{})
	data["bookingDetails"] = emptyBookingDetails

	render.RenderTemplate(w, r, "book-flight.page.html", &models.TemplateData{
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
		helpers.ServerError(w, err)
		return
	}

	// Retrieve the uploaded file
	file, fileheader, err := r.FormFile("uploadfile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

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
		Upload_File_As_ID: fileheader.Filename,
	}
	form := forms.New(r.PostForm)

	form.Required("flying_from", "flying_to", "depart", "return", "full_name", "address", "email")

	form.MinLength("full_name", 6)

	form.IsEmail("email")
	if !form.Valid() {
		data := make(map[string]interface{})
		data["bookingDetails"] = bookingDetails

		render.RenderTemplate(w, r, "book-flight.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return

	}
	m.App.Session.Put(r.Context(), "bookingDetails", bookingDetails)
	http.Redirect(w, r, "/booking-summary", http.StatusSeeOther)

}

func (m *Repository) Indonesia(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "indonesia.page.html", &models.TemplateData{})
}

func (m *Repository) Japan(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "japan.page.html", &models.TemplateData{})
}

func (m *Repository) Thailand(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "thailand.page.html", &models.TemplateData{})
}

func (m *Repository) Southkorea(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "southKorea.page.html", &models.TemplateData{})
}

// Majors renders the room page
func (m *Repository) Payment(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "payment.page.html", &models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repository) Chooseplan(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "chooseplan.page.html", &models.TemplateData{})
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	// This is the emptydata when the page render for the very first time
	var emptyContactDetails models.ContactDetails
	data := make(map[string]interface{})
	data["contactDetails"] = emptyContactDetails
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{
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

		render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{
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
	render.RenderTemplate(w, r, "booking-summary.page.html", &models.TemplateData{
		Data: data,
	})
}
