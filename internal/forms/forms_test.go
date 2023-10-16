package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("forms shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}

}

func TestForm_Has(t *testing.T) {
	// r := httptest.NewRequest("POST", "/whatever", nil)        we dont need this request anymore becoz while we writing has = form.Has("a",r) then it is searching the value of "a" in that request "r" which will be not found. So "r" is unnecessary here

	postedData := url.Values{}
	form := New(postedData)

	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it does not")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("shows form does not have filed when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	// r := httptest.NewRequest("POST", "/whatever", nil) // we dont need this request anymore
	// form := New(r.PostForm)
	postedData := url.Values{}
	form := New(postedData)

	form.MinLength("whatever", 3)
	if form.Valid() {
		t.Error("form shows minlength for non-existent field")
	}

	// checking Get function of form.Errors.Get if there is any error or not
	isError := form.Errors.Get("whatever")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}

	postedData = url.Values{}
	postedData.Add("some_field", "some value")
	form = New(postedData)

	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("shows minlength of 100 met when data is shorter")
	}

	postedData = url.Values{}
	postedData.Add("another_field", "abc123")
	form = New(postedData)

	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("shows minlength of 1 is not met when it is")
	}

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("should not have an error, but got one")
	}

}

func TestForm_IsEmail(t *testing.T) {
	// r := httptest.NewRequest("POST", "/whatever", nil)
	// form := New(r.PostForm)
	postedValues := url.Values{}
	form := New(postedValues)

	form.IsEmail("a")
	if form.Valid() {
		t.Error("form shows valid email for non-extstent field")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "bishal@gmail.com")
	form = New(postedValues)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email when we should not have")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "bishal")
	form = New(postedValues)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("got a valid email for invalid email address")
	}

}
