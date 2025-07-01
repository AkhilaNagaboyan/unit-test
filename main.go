package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Login Successful!")
	})
	http.ListenAndServe(":8080", nil)
}
