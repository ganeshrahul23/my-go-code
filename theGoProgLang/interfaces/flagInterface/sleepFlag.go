package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	period := flag.Duration("period", time.Second, "sleep period")
	flag.Parse()
	fmt.Printf("%T\n", period)

	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
