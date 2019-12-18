package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		param := params.Get("lname")
		fmt.Fprintf(w, "LName is : %s\n", param)
	})

	http.ListenAndServe(":8080", nil)
}
