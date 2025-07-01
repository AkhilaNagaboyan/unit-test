package main // This must match the package where the code to be tested resides

import (
	"net/http"          // Used for HTTP status codes and request/response types
	"net/http/httptest" // Used for creating test HTTP requests and response recorders
	"testing"           // Go’s built-in testing framework
)

func TestLogin(t *testing.T) {
	// Create a new HTTP GET request to the /login endpoint
	// The third parameter is the request body, which is nil for GET requests
	req := httptest.NewRequest("GET", "/login", nil)

	// Create a ResponseRecorder to record the response from the handler
	w := httptest.NewRecorder()

	// Pass the request to the default HTTP multiplexer (DefaultServeMux)
	// This will call the handler registered for /login, if one exists
	http.DefaultServeMux.ServeHTTP(w, req)

	// Check if the status code returned is 200 OK
	// If not, mark the test as failed and print the actual status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", w.Code)
	}
}
