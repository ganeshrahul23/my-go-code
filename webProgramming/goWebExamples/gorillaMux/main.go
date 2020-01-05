package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title, ok := vars["title"]
		if ok {
			fmt.Fprintf(w, "Title is %s\n", title)
		}
		page, ok := vars["page"]
		if ok {
			fmt.Fprintf(w, "Page is %s\n", page)
		}
	})

	http.ListenAndServe(":8080", r)

}
