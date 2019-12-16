package main

import "fmt"

type Interface1 interface {
	Dummy1()
}

type Interface2 interface {
	Dummy2()
}

type Interface1_2 interface {
	Interface1
	Interface2
}

type type1 float64
type type2 float64

func (t type1) Dummy1() {
	fmt.Println("Dummy 1")
}

func (t type2) Dummy1() {
	fmt.Println("Dummy 1")
}

func (t type2) Dummy2() {
	fmt.Println("Dummy 2")
}

func main() {
	var i1 Interface1
	var i2 Interface2
	var i1_2 Interface1_2

	i1 = type1(2.0)
	i1.Dummy1()

	i2 = type2(3.0)
	i2.Dummy2()

	i1_2 = type2(4.0)
	i1_2.Dummy1()
	i1_2.Dummy2()
}
