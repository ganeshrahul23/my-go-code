package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go dummy(&wg)
	wg.Wait()
	fmt.Println("Main Routine")
}

func dummy(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Dummy")
}
