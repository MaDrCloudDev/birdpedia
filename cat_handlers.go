package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type cat struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

var cats []cat

func getCatHandler(w http.ResponseWriter, r *http.Request) {
	//Convert the "cats" variable to json
	catListBytes, err := json.Marshal(cats)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of cats to the response
	w.Write(catListBytes)
}

func createCatHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of cat
	cat := cat{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the cat from the form info
	cat.Species = r.Form.Get("species")
	cat.Description = r.Form.Get("description")

	// Append our existing list of cats with a new entry
	cats = append(cats, cat)

	//Finally, we redirect the user to the original HTMl page
	// (located at `/assets/`), using the http libraries `Redirect` method
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
