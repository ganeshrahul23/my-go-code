package main

import "fmt"

func main() {
	slice := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

	newSlice1 := slice[1:3]
	fmt.Println("newSlice1 :", newSlice1, "Len :", len(newSlice1), "Cap :", cap(newSlice1))
	fmt.Println("Full newSlice1 :", newSlice1[:cap(newSlice1)])

	newSlice2 := slice[1:3:4]
	fmt.Println("newSlice2 :", newSlice2, "Len :", len(newSlice2), "Cap :", cap(newSlice2))
	fmt.Println("Full newSlice2 :", newSlice2[:cap(newSlice2)])
}
