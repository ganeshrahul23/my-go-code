package main

import "fmt"

type dummy struct {
	name string
}

type dummyI interface {
	method() string
}

func (d *dummy) method() string {
	return "dummy"
}

func main() {
	d := dummy{name: "Ganesh"}
	var dI dummyI
	dI = &d
	fmt.Println(dI.method())
}
