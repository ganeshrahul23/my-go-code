package main

import "fmt"

func one(ptr *[32]byte) {
	for i := range ptr {
		(*ptr)[i] = 1
		//ptr[i] = 1
	}
}

func main() {
	var ptr [32]byte
	one(&ptr)
	fmt.Println(ptr)
}
