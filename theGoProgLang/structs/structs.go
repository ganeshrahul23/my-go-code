package main

import "fmt"

type myStruct struct {
	s        string
	myStruct string
}

func (m myStruct) String() string {
	return m.myStruct + " " + m.s
}
func main() {
	m := myStruct{}
	m.myStruct = "Ganesh"
	m.s = "Rahul"
	fmt.Println(m)
}
