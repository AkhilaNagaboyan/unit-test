package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Login Successful!")
}

func TestLogin(t *testing.T) {
	// Register the route explicitly for testing
	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginHandler)

	req := httptest.NewRequest("GET", "/login", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %d", w.Code)
	}
}
