package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	PORT = ":8080"
)

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/pages/{id:[0-9]+}", pageHandler)
	http.Handle("/", rtr)
	http.ListenAndServe(PORT, nil)
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fileName := "files/" + id + ".html"
	fmt.Println(fileName)
	http.ServeFile(w, r, fileName)
}
