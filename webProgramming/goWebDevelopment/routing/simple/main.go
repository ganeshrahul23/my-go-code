package main

import "net/http"

const (
	PORT = ":8080"
)

func main() {
	fileServer := http.FileServer(http.Dir("files"))
	http.ListenAndServe(PORT, fileServer)
}
