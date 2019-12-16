package main

import (
	"fmt"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/counter", counter)
	http.ListenAndServe("localhost:8888", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	fmt.Println(count)
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Counter = %d\n", count)
	mu.Unlock()
}
