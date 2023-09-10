package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// The newRouter function creates and returns the router.
// This function allows us to create and test the router independently of the main function.
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	// Define the static file directory and set it to the recently created directory.
	staticFileDirectory := http.Dir("./assets/")
	// Declare the handler, which routes requests to their respective filenames.
	// The file server is enclosed within the `stripPrefix` method to remove the "/assets/" prefix
	// when searching for files. For instance, if we enter "/assets/index.html" in the browser,
	// the file server will only look for "index.html" within the specified directory.
	// Without stripping the prefix, the file server would search for "./assets/assets/index.html," causing an error.
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method serves as a matcher, matching all routes that start with "/assets/"
	// rather than the exact route itself.
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")
	return r
}

func main() {
	// The router is now created by invoking the `newRouter` constructor function
	// defined above. The rest of the code remains unchanged.
	r := newRouter()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err.Error())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

// go build
// ./birdpedia
