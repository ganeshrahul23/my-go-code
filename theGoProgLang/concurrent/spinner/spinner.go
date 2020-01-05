package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(time.Millisecond)
	const n = 100000011100111111
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func fib(x int) int {
	d, _ := time.ParseDuration("0.5s")
	time.Sleep(d)
	if x > 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
