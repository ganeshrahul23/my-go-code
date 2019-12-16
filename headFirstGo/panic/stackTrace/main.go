package main

import "fmt"

func main() {
	one()
}

func one() {
	two()
}

func two() {
	three()
}

func three() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	panic("Panic in Three")

}
