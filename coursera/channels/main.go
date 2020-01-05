package main

import "fmt"

func main() {
	c := make(chan int)
	go a(c)
	go b(c)
	<-c
	<-c
}

func prod(a int, b int, c chan int) {
	c <- a * b
}

func a(c chan int) {
	fmt.Println("Inside a")
	c <- 0
	fmt.Println("End of a")
}

func b(c chan int) {
	fmt.Println("Inside b")
	<-c
	fmt.Println("End of b")
	c <- 1
	c <- 1
}
