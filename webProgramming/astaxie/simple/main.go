package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for key, value := range r.Form {
		fmt.Println("Key:", key)
		fmt.Println("Value:", strings.Join(value, ""))
	}
	fmt.Fprintf(w, "Hello, World !\n")
}
func main() {
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Listen And Serve :", err)
	}
	fmt.Println("Listening on Port 8080")
}
