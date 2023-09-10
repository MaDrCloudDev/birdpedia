package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestGetBirdsHandler(t *testing.T) {

	// Set up a sample bird for testing
	birds = []Bird{
		{"sparrow", "A small harmless bird"},
	}

	// Create a new HTTP request for GET method
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	// Create an HTTP handler function for getting birds
	hf := http.HandlerFunc(getBirdHandler)

	// Serve the HTTP request
	hf.ServeHTTP(recorder, req)

	// Check if the status code is OK
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check if the response body matches the expected bird
	expected := Bird{"sparrow", "A small harmless bird"}
	b := []Bird{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b[0]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestCreateBirdsHandler(t *testing.T) {

	// Set up a sample bird for testing
	birds = []Bird{
		{"sparrow", "A small harmless bird"},
	}

	// Create a new form for creating a bird
	form := newCreateBirdForm()
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	// Create an HTTP handler function for creating birds
	hf := http.HandlerFunc(createBirdHandler)

	// Serve the HTTP request
	hf.ServeHTTP(recorder, req)

	// Check if the status code is Found (302)
	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check if the new bird was added to the list
	expected := Bird{"eagle", "A bird of prey"}

	if err != nil {
		t.Fatal(err)
	}

	actual := birds[1]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func newCreateBirdForm() *url.Values {
	form := url.Values{}
	form.Set("species", "eagle")
	form.Set("description", "A bird of prey")
	return &form
}
