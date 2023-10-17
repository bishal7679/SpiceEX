package render

import (
	"net/http"
	"testing"

	"github.com/bishal7679/SpiceEx/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")
	// previously throwing error due to not having any session in this request "r"
	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("flash value of 123 not found in session")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myWriter
	err =	Template(&ww, r, "non-existent.page.html", &models.TemplateData{})
	if err == nil {
		t.Error("error writing template to browser")
	}
}

func getSession() (*http.Request, error) {
	// creating a new request
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	// creating a context
	ctx := r.Context()
	// loading the session into that context to access all season variable
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	// adding that sessional context with that request
	r = r.WithContext(ctx)

	return r, nil
}

func TestNewRenderer(t *testing.T) {
	NewRenderer(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}
