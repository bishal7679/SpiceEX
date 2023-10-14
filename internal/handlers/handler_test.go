package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"chooseplan", "/chooseplan", "GET", []postData{}, http.StatusOK},
	{"plans", "/plans", "GET", []postData{}, http.StatusOK},
	{"plansignup", "/plansignup", "GET", []postData{}, http.StatusOK},
	{"payment", "/payment", "GET", []postData{}, http.StatusOK},
	{"book-flight", "/book-flight", "GET", []postData{}, http.StatusOK},
	{"indonesia", "/indonesia", "GET", []postData{}, http.StatusOK},
	{"japan", "/japan", "GET", []postData{}, http.StatusOK},
	{"southkorea", "/southkorea", "GET", []postData{}, http.StatusOK},
	{"thailand", "/thailand", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},

	// {"post-search-avail", "/search-availability", "POST", []postData{
	// 	{key: "start", value: "2022-01-01"},
	// 	{key: "end", value: "2022-02-02"},
	// }, http.StatusOK},
	{"book-flight post", "/book-flight", "POST", []postData{
		{key: "check", value: "Roundtrip"},
		{key: "flying_from", value: "usa"},
		{key: "flying_to", value: "india"},
		{key: "depart", value: "2023-09-02"},
		{key: "return", value: "2023-09-04"},
		{key: "adult_no", value: "0"},
		{key: "child_no", value: "1"},
		{key: "travel_class", value: "Economy"},

		{key: "full_name", value: "john doe"},
		{key: "address", value: "Benno Mohr Street 95 Blumenau/SC - London"},
		{key: "email", value: "john12@gmail.com"},
		{key: "country_code", value: "+81"},
		{key: "mobile_no", value: "87900055433"},
		{key: "pincode", value: "814500"},
		{key: "state", value: "something"},
		{key: "city", value: "something"},
		{key: "uploadfile", value: "nature1.jpg"},
	}, http.StatusOK},

	{"contact post", "/contact", "POST", []postData{
		{key: "name", value: "john doe"},
		{key: "email", value: "john12@gmail.com"},
		{key: "message", value: "hello spiceEX!"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	// getting the handler
	routes := getRoutes()
	// creating one server where a webclient will request
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				// t.Log(err)
				// t.Fatal(err)
				fmt.Println("error in get")
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			// url.Values{} is basically used to store all the post data by the form
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				// t.Log(err)
				// t.Fatal(err)
				fmt.Println("error here")
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}

		}
	}
}
