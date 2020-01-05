package main

import (
	"fmt"
	"time"
)

func main() {
	go dummy()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Main Routine")
	//time.Sleep(30 * time.Millisecond)
}

func dummy() {
	time.Sleep(110 * time.Millisecond)
	fmt.Println("Go Routine")
}
