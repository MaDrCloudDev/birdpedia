package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

var birds []Bird

func getBirdHandler(w http.ResponseWriter, r *http.Request) {
	// Convert the "birds" variable to JSON format
	birdListBytes, err := json.Marshal(birds)

	// In case of an error, log it and send a server error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// If everything goes well, write the JSON list of birds to the response
	w.Write(birdListBytes)
}

func createBirdHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of Bird
	bird := Bird{}

	// We send all our data as HTML form data.
	// The `ParseForm` method of the request parses the form values.
	err := r.ParseForm()

	// If there's an error, respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Retrieve information about the bird from the form data
	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	// Add the new bird entry to our existing list of birds
	birds = append(birds, bird)

	// Finally, redirect the user back to the original HTML page (located at `/assets/`)
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
