package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// Form a new HTTP request for our handler.
	// We specify the method, route, and leave the request body nil in this case.
	req, err := http.NewRequest("GET", "", nil)

	// If there's an error in forming the request, fail and stop the test.
	if err != nil {
		t.Fatal(err)
	}

	// Create an HTTP recorder, which acts as the target of our HTTP request.
	// Think of it as a mini-browser that accepts the result of the HTTP request.
	recorder := httptest.NewRecorder()

	// Create an HTTP handler from our handler function.
	hf := http.HandlerFunc(handler)

	// Serve the HTTP request to our recorder. This actually executes the handler we want to test.
	hf.ServeHTTP(recorder, req)

	// Check if the status code is what we expect.
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check if the response body is what we expect.
	expected := `Hello World!`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestRouter(t *testing.T) {
	// Instantiate the router using the constructor function defined earlier.
	r := newRouter()

	// Create a new server using the "httptest" library's `NewServer` method.
	mockServer := httptest.NewServer(r)

	// The mock server we created runs a server and exposes its location in the URL attribute.
	// We make a GET request to the "hello" route we defined in the router.
	resp, err := http.Get(mockServer.URL + "/hello")

	// Handle any unexpected error.
	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 200 (ok).
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	// Read the response body and convert it to a string.
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := "Hello World!"

	// We want our response to match the one defined in our handler.
	// If it does happen to be "Hello world!", then it confirms that the
	// route is correct.
	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}

}

func TestRouterForNonExistentRoute(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	// Make a request to a route we know we didn't define, like the `PUT /hello` route.
	resp, err := http.Post(mockServer.URL+"/hello", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 405 (method not allowed).
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, got %d", resp.StatusCode)
	}

	// Read the response body and convert it to a string.
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := ""

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}

}

func TestStaticFileServer(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	// Make a GET request to the `GET /assets/` route to get the index.html file response.
	resp, err := http.Get(mockServer.URL + "/assets/")

	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 200 (ok).
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be 200, got %d", resp.StatusCode)
	}

	// Test that the content-type header is "text/html; charset=utf-8".
	// This indicates that an HTML file has been served.
	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
	}

}
